{{/*
query-scheduler fullname
*/}}
{{- define "logstore.querySchedulerFullname" -}}
{{ include "logstore.fullname" . }}-query-scheduler
{{- end }}

{{/*
query-scheduler common labels
*/}}
{{- define "logstore.querySchedulerLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: query-scheduler
{{- end }}

{{/*
query-scheduler selector labels
*/}}
{{- define "logstore.querySchedulerSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: query-scheduler
{{- end }}

{{/*
query-scheduler image
*/}}
{{- define "logstore.querySchedulerImage" -}}
{{- $dict := dict "logstore" .Values.logstore.image "service" .Values.queryScheduler.image "global" .Values.global.image "defaultVersion" .Chart.AppVersion -}}
{{- include "logstore.logstoreImage" $dict -}}
{{- end }}

{{/*
query-scheduler priority class name
*/}}
{{- define "logstore.querySchedulerPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.queryScheduler.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
