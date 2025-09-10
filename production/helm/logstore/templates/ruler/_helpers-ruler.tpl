{{/*
ruler fullname
*/}}
{{- define "logstore.rulerFullname" -}}
{{ include "logstore.fullname" . }}-ruler
{{- end }}

{{/*
ruler common labels
*/}}
{{- define "logstore.rulerLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: ruler
{{- end }}

{{/*
ruler selector labels
*/}}
{{- define "logstore.rulerSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: ruler
{{- end }}

{{/*
ruler image
*/}}
{{- define "logstore.rulerImage" -}}
{{- $dict := dict "logstore" .Values.logstore.image "service" .Values.ruler.image "global" .Values.global.image "defaultVersion" .Chart.AppVersion -}}
{{- include "logstore.logstoreImage" $dict -}}
{{- end }}

{{/*
format rules dir
*/}}
{{- define "logstore.rulerRulesDirName" -}}
rules-{{ . | replace "_" "-" | trimSuffix "-" | lower }}
{{- end }}

{{/*
ruler priority class name
*/}}
{{- define "logstore.rulerPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.ruler.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
