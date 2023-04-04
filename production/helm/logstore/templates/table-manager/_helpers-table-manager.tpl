{{/*
table-manager fullname
*/}}
{{- define "logstore.tableManagerFullname" -}}
{{ include "logstore.fullname" . }}-table-manager
{{- end }}

{{/*
table-manager common labels
*/}}
{{- define "logstore.tableManagerLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: table-manager
{{- end }}

{{/*
table-manager selector labels
*/}}
{{- define "logstore.tableManagerSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: table-manager
{{- end }}

{{/*
table-manager priority class name
*/}}
{{- define "logstore.tableManagerPriorityClassName" -}}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.tableManager.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}
