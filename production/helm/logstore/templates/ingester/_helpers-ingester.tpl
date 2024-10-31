{{/*
ingester fullname
*/}}
{{- define "logstore.ingesterFullname" -}}
{{ include "logstore.fullname" . }}-ingester
{{- end }}

{{/*
ingester common labels
*/}}
{{- define "logstore.ingesterLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: ingester
{{- end }}

{{/*
ingester selector labels
*/}}
{{- define "logstore.ingesterSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: ingester
{{- end }}

{{/*
ingester priority class name
*/}}
{{- define "logstore.ingesterPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.ingester.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}

{{- define "logstore.ingester.readinessProbe" -}}
{{- with .Values.ingester.readinessProbe }}  
readinessProbe:
  {{- toYaml . | nindent 2 }}
{{- else }}
{{- with .Values.logstore.readinessProbe }}
readinessProbe:
  {{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}
{{- end -}}

{{- define "logstore.ingester.livenessProbe" -}}
{{- with .Values.ingester.livenessProbe }}
livenessProbe:
  {{- toYaml . | nindent 2 }}
{{- else }}
{{- with .Values.logstore.livenessProbe }}
livenessProbe:
  {{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}
{{- end -}}

{{/*
expects global context
*/}}
{{- define "logstore.ingester.replicaCount" -}}
{{- ceil (divf .Values.ingester.replicas 3) -}}
{{- end -}}

{{/*
expects a dict
{
  "replicas": replicas in a zone,
  "ctx": global context
}
*/}}
{{- define "logstore.ingester.maxUnavailable" -}}
{{- ceil (mulf .replicas (divf (int .ctx.Values.ingester.zoneAwareReplication.maxUnavailablePct) 100)) -}}
{{- end -}}