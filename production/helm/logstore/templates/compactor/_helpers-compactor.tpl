{{/*
compactor fullname
*/}}
{{- define "logstore.compactorFullname" -}}
{{ include "logstore.fullname" . }}-compactor
{{- end }}

{{/*
compactor common labels
*/}}
{{- define "logstore.compactorLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: compactor
{{- end }}

{{/*
compactor selector labels
*/}}
{{- define "logstore.compactorSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: compactor
{{- end }}

{{/*
compactor image
*/}}
{{- define "logstore.compactorImage" -}}
{{- $dict := dict "logstore" .Values.logstore.image "service" .Values.compactor.image "global" .Values.global.image "defaultVersion" .Chart.AppVersion -}}
{{- include "logstore.logstoreImage" $dict -}}
{{- end }}

{{/*
compactor readinessProbe
*/}}
{{- define "logstore.compactor.readinessProbe" -}}
{{- with .Values.compactor.readinessProbe }}
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
compactor livenessProbe
*/}}
{{- define "logstore.compactor.livenessProbe" -}}
{{- with .Values.compactor.livenessProbe }}
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
compactor priority class name
*/}}
{{- define "logstore.compactorPriorityClassName" }}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.compactor.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}

{{/*
Create the name of the compactor service account
*/}}
{{- define "logstore.compactorServiceAccountName" -}}
{{- if .Values.compactor.serviceAccount.create -}}
    {{ default (print (include "logstore.serviceAccountName" .) "-compactor") .Values.compactor.serviceAccount.name }}
{{- else -}}
    {{ default (include "logstore.serviceAccountName" .) .Values.compactor.serviceAccount.name }}
{{- end -}}
{{- end -}}
