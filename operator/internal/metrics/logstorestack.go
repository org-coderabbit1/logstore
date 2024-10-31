package metrics

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

const (
	metricsPrefix = "logstorestack_"
)

var (
	metricsCommonLabels = []string{
		"stack_namespace",
		"stack_name",
		"size",
	}

	logstoreStackInfoDesc = prometheus.NewDesc(
		metricsPrefix+"info",
		"Information about deployed LogstoreStack instances. Value is always 1.",
		metricsCommonLabels, nil,
	)

	logstoreStackConditionsCountDesc = prometheus.NewDesc(
		metricsPrefix+"status_condition",
		"Counts the current status conditions of the LogstoreStack.",
		append(metricsCommonLabels, "condition", "reason", "status"), nil,
	)
)

func RegisterLogstoreStackCollector(log logr.Logger, k8sClient client.Client, registry prometheus.Registerer) error {
	metrics := &logstoreStackCollector{
		log:       log,
		k8sClient: k8sClient,
	}

	return registry.Register(metrics)
}

type logstoreStackCollector struct {
	log       logr.Logger
	k8sClient client.Client
}

func (l *logstoreStackCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- logstoreStackInfoDesc
	ch <- logstoreStackConditionsCountDesc
}

func (l *logstoreStackCollector) Collect(m chan<- prometheus.Metric) {
	ctx := context.TODO()

	stackList := &logstorev1.LogstoreStackList{}
	err := l.k8sClient.List(ctx, stackList)
	if err != nil {
		l.log.Error(err, "failed to get list of LogstoreStacks for metrics")
		return
	}

	for _, stack := range stackList.Items {
		labels := []string{
			stack.Namespace,
			stack.Name,
			string(stack.Spec.Size),
		}

		m <- prometheus.MustNewConstMetric(logstoreStackInfoDesc, prometheus.GaugeValue, 1.0, labels...)

		for _, c := range stack.Status.Conditions {
			activeValue := 0.0
			if c.Status == metav1.ConditionTrue {
				activeValue = 1.0
			}

			// This mirrors the behavior of kube_state_metrics, which creates two metrics for each condition,
			// one for each status (true/false).
			m <- prometheus.MustNewConstMetric(
				logstoreStackConditionsCountDesc,
				prometheus.GaugeValue, activeValue,
				append(labels, c.Type, c.Reason, "true")...,
			)
			m <- prometheus.MustNewConstMetric(
				logstoreStackConditionsCountDesc,
				prometheus.GaugeValue, 1.0-activeValue,
				append(labels, c.Type, c.Reason, "false")...,
			)
		}
	}
}
