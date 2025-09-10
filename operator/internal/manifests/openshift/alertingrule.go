package openshift

import (
	"maps"
	"strings"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

func AlertingRuleTenantLabels(ar *logstorev1.AlertingRule) {
	switch ar.Spec.TenantID {
	case tenantApplication:
		labels := map[string]string{
			ocpMonitoringGroupByLabel: ar.Namespace,
		}
		labelMatchers := strings.Split(opaDefaultLabelMatchers, ",")
		for _, label := range labelMatchers {
			labels[label] = ar.Namespace
		}
		appendAlertingRuleLabels(ar, labels)
	case tenantInfrastructure, tenantAudit, tenantNetwork:
		appendAlertingRuleLabels(ar, map[string]string{
			ocpMonitoringGroupByLabel: ar.Namespace,
		})
	default:
		// Do nothing
	}
}

func appendAlertingRuleLabels(ar *logstorev1.AlertingRule, labels map[string]string) {
	for groupIdx, group := range ar.Spec.Groups {
		for ruleIdx, rule := range group.Rules {
			if rule.Labels == nil {
				rule.Labels = map[string]string{}
			}

			maps.Copy(rule.Labels, labels)

			group.Rules[ruleIdx] = rule
		}
		ar.Spec.Groups[groupIdx] = group
	}
}
