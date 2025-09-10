{{/*
bloom gateway fullname
*/}}
{{- define "logstore.bloomGatewayFullname" -}}
{{ include "logstore.fullname" . }}-bloom-gateway
{{- end }}

{{/*
bloom gateway common labels
*/}}
{{- define "logstore.bloomGatewayLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: bloom-gateway
{{- end }}

{{/*
bloom gateway selector labels
*/}}
{{- define "logstore.bloomGatewaySelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: bloom-gateway
{{- end }}

{{/*
bloom gateway livenessProbe
*/}}
{{- define "logstore.bloomGateway.livenessProbe" }}
{{- with .Values.bloomGateway.livenessProbe | default .Values.logstore.livenessProbe }}
livenessProbe:
  {{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{/*
bloom gateway readinessProbe
*/}}
{{- define "logstore.bloomGateway.readinessProbe" }}
{{- with .Values.bloomGateway.readinessProbe | default .Values.logstore.readinessProbe }}
readinessProbe:
  {{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{/*
bloom gateway startupProbe
*/}}
{{- define "logstore.bloomGateway.startupProbe" }}
{{- with .Values.bloomGateway.startupProbe | default .Values.logstore.startupProbe }}
startupProbe:
  {{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{/*
bloom gateway priority class name
*/}}
{{- define "logstore.bloomGatewayPriorityClassName" }}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.bloomGateway.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}

{{/*
Create the name of the bloom gateway service account
*/}}
{{- define "logstore.bloomGatewayServiceAccountName" -}}
{{- if .Values.bloomGateway.serviceAccount.create -}}
    {{ default (print (include "logstore.serviceAccountName" .) "-bloom-gateway") .Values.bloomGateway.serviceAccount.name }}
{{- else -}}
    {{ default (include "logstore.serviceAccountName" .) .Values.bloomGateway.serviceAccount.name }}
{{- end -}}
{{- end -}}
