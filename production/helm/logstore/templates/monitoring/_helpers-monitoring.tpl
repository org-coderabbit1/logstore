{{/*
Client definition for LogsInstance
*/}}
{{- define "logstore.logsInstanceClient" -}}
{{- $isSingleBinary := eq (include "logstore.deployment.isSingleBinary" .) "true" -}}
{{- $url := printf "http://%s.%s.svc.%s:3100/logstore/api/v1/push" (include "logstore.writeFullname" .) .Release.Namespace .Values.global.clusterDomain }}
{{- if $isSingleBinary  }}
  {{- $url = printf "http://%s.%s.svc.%s:3100/logstore/api/v1/push" (include "logstore.singleBinaryFullname" .) .Release.Namespace .Values.global.clusterDomain }}
{{- else if .Values.gateway.enabled -}}
  {{- $url = printf "http://%s.%s.svc.%s/logstore/api/v1/push" (include "logstore.gatewayFullname" .) .Release.Namespace .Values.global.clusterDomain }}
{{- end -}}
- url: {{ $url }}
  externalLabels:
    cluster: {{ include "logstore.clusterLabel" . }}
  {{- if .Values.enterprise.enabled }}
  basicAuth:
    username:
      name: {{ include "enterprise-logs.selfMonitoringTenantSecret" . }}
      key: username
    password:
      name: {{ include "enterprise-logs.selfMonitoringTenantSecret" . }}
      key: password
  {{- else if .Values.logstore.auth_enabled }}
  tenantId: {{ .Values.monitoring.selfMonitoring.tenant.name | quote }}
  {{- end }}
{{- end -}}

{{/*
Convert a recording rule group to yaml
*/}}
{{- define "logstore.ruleGroupToYaml" -}}
{{- range . }}
- name: {{ .name }}
  rules:
    {{- toYaml .rules | nindent 4 }}
{{- end }}
{{- end }}

{{/*
AcmeAgent priority class name
*/}}
{{- define "acme-agent.priorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.monitoring.selfMonitoring.acmeAgent.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
