{{/*
read fullname
*/}}
{{- define "logstore.readFullname" -}}
{{ include "logstore.name" . }}-read
{{- end }}

{{/*
read common labels
*/}}
{{- define "logstore.readLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: read
{{- end }}

{{/*
read selector labels
*/}}
{{- define "logstore.readSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: read
{{- end }}

{{/*
read priority class name
*/}}
{{- define "logstore.readPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.read.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
