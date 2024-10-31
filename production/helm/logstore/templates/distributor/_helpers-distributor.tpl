{{/*
distributor fullname
*/}}
{{- define "logstore.distributorFullname" -}}
{{ include "logstore.fullname" . }}-distributor
{{- end }}

{{/*
distributor common labels
*/}}
{{- define "logstore.distributorLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: distributor
{{- end }}

{{/*
distributor selector labels
*/}}
{{- define "logstore.distributorSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: distributor
{{- end }}

{{/*
distributor priority class name
*/}}
{{- define "logstore.distributorPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.distributor.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
