{{/*
query-frontend fullname
*/}}
{{- define "logstore.queryFrontendFullname" -}}
{{ include "logstore.fullname" . }}-query-frontend
{{- end }}

{{/*
query-frontend common labels
*/}}
{{- define "logstore.queryFrontendLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: query-frontend
{{- end }}

{{/*
query-frontend selector labels
*/}}
{{- define "logstore.queryFrontendSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: query-frontend
{{- end }}

{{/*
query-frontend priority class name
*/}}
{{- define "logstore.queryFrontendPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.queryFrontend.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
