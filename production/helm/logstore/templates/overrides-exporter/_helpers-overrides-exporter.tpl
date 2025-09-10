{{/*
overrides-exporter fullname
*/}}
{{- define "logstore.overridesExporterFullname" -}}
{{ include "logstore.fullname" . }}-overrides-exporter
{{- end }}

{{/*
overrides-exporter common labels
*/}}
{{- define "logstore.overridesExporterLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: overrides-exporter
{{- end }}

{{/*
overrides-exporter selector labels
*/}}
{{- define "logstore.overridesExporterSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: overrides-exporter
{{- end }}

{{/*
overrides-exporter priority class name
*/}}
{{- define "logstore.overridesExporterPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.overridesExporter.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
