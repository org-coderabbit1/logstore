{{/*
bloom-builder fullname
*/}}
{{- define "logstore.bloomBuilderFullname" -}}
{{ include "logstore.fullname" . }}-bloom-builder
{{- end }}

{{/*
bloom-builder common labels
*/}}
{{- define "logstore.bloomBuilderLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: bloom-builder
{{- end }}

{{/*
bloom-builder selector labels
*/}}
{{- define "logstore.bloomBuilderSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: bloom-builder
{{- end }}

{{/*
bloom-builder priority class name
*/}}
{{- define "logstore.bloomBuilderPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.bloomBuilder.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
