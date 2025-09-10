{{/*
gateway fullname
*/}}
{{- define "logstore.gatewayFullname" -}}
{{ include "logstore.fullname" . }}-gateway
{{- end }}

{{/*
gateway common labels
*/}}
{{- define "logstore.gatewayLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: gateway
{{- end }}

{{/*
gateway selector labels
*/}}
{{- define "logstore.gatewaySelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: gateway
{{- end }}

{{/*
gateway auth secret name
*/}}
{{- define "logstore.gatewayAuthSecret" -}}
{{ .Values.gateway.basicAuth.existingSecret | default (include "logstore.gatewayFullname" . ) }}
{{- end }}

{{/*
gateway Docker image
*/}}
{{- define "logstore.gatewayImage" -}}
{{- $dict := dict "service" .Values.gateway.image "global" .Values.global.image -}}
{{- include "logstore.baseImage" $dict -}}
{{- end }}

{{/*
gateway priority class name
*/}}
{{- define "logstore.gatewayPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.gateway.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
