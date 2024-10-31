{{/*
querier fullname
*/}}
{{- define "logstore.querierFullname" -}}
{{ include "logstore.fullname" . }}-querier
{{- end }}

{{/*
querier common labels
*/}}
{{- define "logstore.querierLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: querier
{{- end }}

{{/*
querier selector labels
*/}}
{{- define "logstore.querierSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: querier
{{- end }}

{{/*
querier priority class name
*/}}
{{- define "logstore.querierPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.querier.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
