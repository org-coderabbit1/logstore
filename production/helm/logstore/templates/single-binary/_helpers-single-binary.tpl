{{/*
singleBinary common labels
*/}}
{{- define "logstore.singleBinaryLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: single-binary
{{- end }}


{{/* singleBinary selector labels */}}
{{- define "logstore.singleBinarySelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: single-binary
{{- end }}

{{/*
singleBinary priority class name
*/}}
{{- define "logstore.singleBinaryPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.singleBinary.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}

{{/* singleBinary replicas calculation */}}
{{- define "logstore.singleBinaryReplicas" -}}
{{- $replicas := 1 }}
{{- $usingObjectStorage := eq (include "logstore.isUsingObjectStorage" .) "true" }}
{{- if and $usingObjectStorage (gt (int .Values.singleBinary.replicas) 1)}}
{{- $replicas = int .Values.singleBinary.replicas -}}
{{- end }}
{{- printf "%d" $replicas }}
{{- end }}
