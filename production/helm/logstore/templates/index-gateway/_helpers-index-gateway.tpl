{{/*
index-gateway fullname
*/}}
{{- define "logstore.indexGatewayFullname" -}}
{{ include "logstore.fullname" . }}-index-gateway
{{- end }}

{{/*
index-gateway common labels
*/}}
{{- define "logstore.indexGatewayLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: index-gateway
{{- end }}

{{/*
index-gateway selector labels
*/}}
{{- define "logstore.indexGatewaySelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: index-gateway
{{- end }}

{{/*
index-gateway image
*/}}
{{- define "logstore.indexGatewayImage" -}}
{{- $dict := dict "logstore" .Values.logstore.image "service" .Values.indexGateway.image "global" .Values.global.image "defaultVersion" .Chart.AppVersion -}}
{{- include "logstore.logstoreImage" $dict -}}
{{- end }}

{{/*
index-gateway priority class name
*/}}
{{- define "logstore.indexGatewayPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.indexGateway.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
