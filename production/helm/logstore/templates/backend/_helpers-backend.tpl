{{/*
backend fullname
*/}}
{{- define "logstore.backendFullname" -}}
{{ include "logstore.name" . }}-backend
{{- end }}

{{/*
backend common labels
*/}}
{{- define "logstore.backendLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: backend
{{- end }}

{{/*
backend selector labels
*/}}
{{- define "logstore.backendSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: backend
{{- end }}

{{/*
backend priority class name
*/}}
{{- define "logstore.backendPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.backend.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
