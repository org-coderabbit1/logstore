package main

import (
	"flag"
	"os"

	"github.com/ViaQ/logerr/v2/kverrors"
	"github.com/ViaQ/logerr/v2/log"
	configv1 "github.com/openshift/api/config/v1"
	routev1 "github.com/openshift/api/route/v1"
	cloudcredentialv1 "github.com/openshift/cloud-credential-operator/pkg/apis/cloudcredential/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	runtimemetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	ctrlconfigv1 "example.com/acme/logstore/operator/api/config/v1"
	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	logstorev1beta1 "example.com/acme/logstore/operator/api/logstore/v1beta1"
	"example.com/acme/logstore/operator/internal/config"
	logstorectrl "example.com/acme/logstore/operator/internal/controller/logstore"
	"example.com/acme/logstore/operator/internal/metrics"
	"example.com/acme/logstore/operator/internal/operator"
	"example.com/acme/logstore/operator/internal/validation"
	"example.com/acme/logstore/operator/internal/validation/openshift"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(logstorev1beta1.AddToScheme(scheme))

	utilruntime.Must(logstorev1.AddToScheme(scheme))

	utilruntime.Must(ctrlconfigv1.AddToScheme(scheme))

	// +kubebuilder:scaffold:scheme
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "",
		"The controller will load its initial configuration from this file. "+
			"Omit this flag to use the default configuration values. "+
			"Command-line flags override configuration from this file.",
	)
	flag.Parse()

	logger := log.NewLogger("logstore-operator")
	ctrl.SetLogger(logger)

	var err error

	ctrlCfg, tokenCCOAuth, options, err := config.LoadConfig(scheme, configFile)
	if err != nil {
		logger.Error(err, "failed to load operator configuration")
		os.Exit(1)
	}

	if tokenCCOAuth != nil {
		logger.Info("Discovered OpenShift Cluster within a token cco authentication environment")
	}

	if ctrlCfg.Gates.LogstoreStackAlerts && !ctrlCfg.Gates.ServiceMonitors {
		logger.Error(kverrors.New("LogstoreStackAlerts flag requires ServiceMonitors"), "")
		os.Exit(1)
	}

	if ctrlCfg.Gates.ServiceMonitorTLSEndpoints && !ctrlCfg.Gates.HTTPEncryption {
		logger.Error(kverrors.New("ServiceMonitorTLSEndpoints flag requires HTTPEncryption"), "")
		os.Exit(1)
	}

	if ctrlCfg.Gates.ServiceMonitors || ctrlCfg.Gates.ServiceMonitorTLSEndpoints {
		utilruntime.Must(monitoringv1.AddToScheme(scheme))
	}

	if ctrlCfg.Gates.LogstoreStackGateway {
		utilruntime.Must(configv1.AddToScheme(scheme))

		if ctrlCfg.Gates.OpenShift.Enabled {
			utilruntime.Must(routev1.AddToScheme(scheme))
			utilruntime.Must(cloudcredentialv1.AddToScheme(scheme))
		}
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), options)
	if err != nil {
		logger.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&logstorectrl.LogstoreStackReconciler{
		Client:       mgr.GetClient(),
		Log:          logger.WithName("controllers").WithName("logstorestack"),
		Scheme:       mgr.GetScheme(),
		FeatureGates: ctrlCfg.Gates,
		AuthConfig:   tokenCCOAuth,
	}).SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", "logstorestack")
		os.Exit(1)
	}

	if ctrlCfg.Gates.ServiceMonitors && ctrlCfg.Gates.OpenShift.Enabled && ctrlCfg.Gates.OpenShift.Dashboards {
		var ns string
		ns, err = operator.GetNamespace()
		if err != nil {
			logger.Error(err, "unable to read in operator namespace")
			os.Exit(1)
		}

		if err = (&logstorectrl.DashboardsReconciler{
			Client:     mgr.GetClient(),
			Scheme:     mgr.GetScheme(),
			Log:        logger.WithName("controllers").WithName(logstorectrl.ControllerNameLogstoreDashboards),
			OperatorNs: ns,
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", logstorectrl.ControllerNameLogstoreDashboards)
			os.Exit(1)
		}
	}

	if ctrlCfg.Gates.LogstoreStackWebhook {
		v := &validation.LogstoreStackValidator{}
		if err = v.SetupWebhookWithManager(mgr); err != nil {
			logger.Error(err, "unable to create webhook", "webhook", "logstorestack")
			os.Exit(1)
		}
	}
	if err = (&logstorectrl.AlertingRuleReconciler{
		Client: mgr.GetClient(),
		Log:    logger.WithName("controllers").WithName(logstorectrl.ControllerNameAlertingRule),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", logstorectrl.ControllerNameAlertingRule)
		os.Exit(1)
	}
	if ctrlCfg.Gates.AlertingRuleWebhook {
		v := &validation.AlertingRuleValidator{}
		if ctrlCfg.Gates.OpenShift.ExtendedRuleValidation {
			v.ExtendedValidator = openshift.AlertingRuleValidator
		}

		if err = v.SetupWebhookWithManager(mgr); err != nil {
			logger.Error(err, "unable to create webhook", "webhook", "alertingrule")
			os.Exit(1)
		}
	}
	if err = (&logstorectrl.RecordingRuleReconciler{
		Client: mgr.GetClient(),
		Log:    logger.WithName("controllers").WithName(logstorectrl.ControllerNameRecordingRule),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", logstorectrl.ControllerNameRecordingRule)
		os.Exit(1)
	}
	if ctrlCfg.Gates.RecordingRuleWebhook {
		v := &validation.RecordingRuleValidator{}
		if ctrlCfg.Gates.OpenShift.ExtendedRuleValidation {
			v.ExtendedValidator = openshift.RecordingRuleValidator
		}

		if err = v.SetupWebhookWithManager(mgr); err != nil {
			logger.Error(err, "unable to create webhook", "webhook", "recordingrule")
			os.Exit(1)
		}
	}
	if err = (&logstorectrl.RulerConfigReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", logstorectrl.ControllerNameRulerConfig)
		os.Exit(1)
	}
	if ctrlCfg.Gates.RulerConfigWebhook {
		v := &validation.RulerConfigValidator{}
		if err = v.SetupWebhookWithManager(mgr); err != nil {
			logger.Error(err, "unable to create webhook", "webhook", "rulerconfig")
			os.Exit(1)
		}
	}
	if ctrlCfg.Gates.BuiltInCertManagement.Enabled {
		if err = (&logstorectrl.CertRotationReconciler{
			Client:       mgr.GetClient(),
			Log:          logger.WithName("controllers").WithName(logstorectrl.ControllerNameCertRotation),
			Scheme:       mgr.GetScheme(),
			FeatureGates: ctrlCfg.Gates,
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", logstorectrl.ControllerNameCertRotation)
			os.Exit(1)
		}
	}
	if err = (&logstorectrl.LogstoreStackZoneAwarePodReconciler{
		Client: mgr.GetClient(),
		Log:    logger.WithName("controllers").WithName(logstorectrl.ControllerNameZoneAware),
	}).SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", logstorectrl.ControllerNameZoneAware)
		os.Exit(1)
	}

	// +kubebuilder:scaffold:builder

	if err = mgr.AddHealthzCheck("health", healthz.Ping); err != nil {
		logger.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err = mgr.AddReadyzCheck("check", healthz.Ping); err != nil {
		logger.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	logger.Info("registering metrics")
	err = metrics.RegisterLogstoreStackCollector(logger, mgr.GetClient(), runtimemetrics.Registry)
	if err != nil {
		logger.Error(err, "failed to register metrics")
		os.Exit(1)
	}

	logger.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		logger.Error(err, "problem running manager")
		os.Exit(1)
	}
}
