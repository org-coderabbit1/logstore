{{/*
dashboards name
*/}}
{{- define "logstore.dashboardsName" -}}
{{ include "logstore.name" . }}-dashboards
{{- end }}
