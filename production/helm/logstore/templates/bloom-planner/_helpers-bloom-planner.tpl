{{/*
bloom planner fullname
*/}}
{{- define "logstore.bloomPlannerFullname" -}}
{{ include "logstore.fullname" . }}-bloom-planner
{{- end }}

{{/*
bloom planner common labels
*/}}
{{- define "logstore.bloomPlannerLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: bloom-planner
{{- end }}

{{/*
bloom planner selector labels
*/}}
{{- define "logstore.bloomPlannerSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: bloom-planner
{{- end }}

{{/*
bloom planner readinessProbe
*/}}
{{- define "logstore.bloomPlanner.readinessProbe" -}}
{{- with .Values.bloomPlanner.readinessProbe }}
readinessProbe:
  {{- toYaml . | nindent 2 }}
{{- else }}
{{- with .Values.logstore.readinessProbe }}
readinessProbe:
  {{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}
{{- end -}}

{{/*
bloom planner priority class name
*/}}
{{- define "logstore.bloomPlannerPriorityClassName" }}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.bloomPlanner.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}

{{/*
Create the name of the bloom planner service account
*/}}
{{- define "logstore.bloomPlannerServiceAccountName" -}}
{{- if .Values.bloomPlanner.serviceAccount.create -}}
    {{ default (print (include "logstore.serviceAccountName" .) "-bloom-planner") .Values.bloomPlanner.serviceAccount.name }}
{{- else -}}
    {{ default (include "logstore.serviceAccountName" .) .Values.bloomPlanner.serviceAccount.name }}
{{- end -}}
{{- end -}}
