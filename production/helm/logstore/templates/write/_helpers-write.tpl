{{/*
write fullname
*/}}
{{- define "logstore.writeFullname" -}}
{{ include "logstore.name" . }}-write
{{- end }}

{{/*
write common labels
*/}}
{{- define "logstore.writeLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: write
{{- end }}

{{/*
write selector labels
*/}}
{{- define "logstore.writeSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: write
{{- end }}

{{/*
write priority class name
*/}}
{{- define "logstore.writePriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.write.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
