{{/*
canary fullname
*/}}
{{- define "logstore-canary.fullname" -}}
{{ include "logstore.name" . }}-canary
{{- end }}

{{/*
canary common labels
*/}}
{{- define "logstore-canary.labels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: canary
{{- end }}

{{/*
canary selector labels
*/}}
{{- define "logstore-canary.selectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: canary
{{- end }}

{{/*
Docker image name for logstore-canary
*/}}
{{- define "logstore-canary.image" -}}
{{- $dict := dict "service" .Values.monitoring.selfMonitoring.logstoreCanary.image "global" .Values.global.image "defaultVersion" .Chart.AppVersion -}}
{{- include "logstore.baseImage" $dict -}}
{{- end -}}

{{/*
canry priority class name
*/}}
{{- define "logstore-canary.priorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.read.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
