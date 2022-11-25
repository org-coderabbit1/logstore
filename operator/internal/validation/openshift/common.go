package openshift

import (
	"strings"

	logstorev1beta1 "example.com/acme/logstore/operator/apis/logstore/v1beta1"

	"example.com/acme/logstore/pkg/logql/syntax"
	"github.com/prometheus/prometheus/model/labels"
)

const (
	namespaceLabelName        = "kubernetes_namespace_name"
	namespaceOpenshiftLogging = "openshift-logging"

	tenantAudit          = "audit"
	tenantApplication    = "application"
	tenantInfrastructure = "infrastructure"
)

func validateRuleExpression(namespace, tenantID, rawExpr string) error {
	// Check if the LogQL parser can parse the rule expression
	expr, err := syntax.ParseExpr(rawExpr)
	if err != nil {
		return logstorev1beta1.ErrParseLogQLExpression
	}

	sampleExpr, ok := expr.(syntax.SampleExpr)
	if !ok {
		return logstorev1beta1.ErrParseLogQLNotSample
	}

	matchers := sampleExpr.Selector().Matchers()
	if tenantID != tenantAudit && !validateIncludesNamespace(namespace, matchers) {
		return logstorev1beta1.ErrRuleMustMatchNamespace
	}

	return nil
}

func validateIncludesNamespace(namespace string, matchers []*labels.Matcher) bool {
	for _, m := range matchers {
		if m.Name == namespaceLabelName && m.Type == labels.MatchEqual && m.Value == namespace {
			return true
		}
	}

	return false
}

func tenantForNamespace(namespace string) []string {
	if strings.HasPrefix(namespace, "openshift") ||
		strings.HasPrefix(namespace, "kube-") ||
		namespace == "default" {
		if namespace == namespaceOpenshiftLogging {
			return []string{tenantAudit, tenantInfrastructure}
		}

		return []string{tenantInfrastructure}
	}

	return []string{tenantApplication}
}
