{{/*
pattern ingester fullname
*/}}
{{- define "logstore.patternIngesterFullname" -}}
{{ include "logstore.fullname" . }}-pattern-ingester
{{- end }}

{{/*
pattern ingester common labels
*/}}
{{- define "logstore.patternIngesterLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: pattern-ingester
{{- end }}

{{/*
pattern ingester selector labels
*/}}
{{- define "logstore.patternIngesterSelectorLabels" -}}
{{ include "logstore.selectorLabels" . }}
app.kubernetes.io/component: pattern-ingester
{{- end }}

{{/*
pattern ingester readinessProbe
*/}}
{{- define "logstore.patternIngester.readinessProbe" }}
{{- with .Values.patternIngester.readinessProbe | default .Values.logstore.readinessProbe }}
readinessProbe:
  {{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{/*
pattern ingester priority class name
*/}}
{{- define "logstore.patternIngesterPriorityClassName" }}
{{- $pcn := coalesce .Values.global.priorityClassName .Values.patternIngester.priorityClassName -}}
{{- if $pcn }}
priorityClassName: {{ $pcn }}
{{- end }}
{{- end }}

{{/*
Create the name of the pattern ingester service account
*/}}
{{- define "logstore.patternIngesterServiceAccountName" -}}
{{- if .Values.patternIngester.serviceAccount.create -}}
    {{ default (print (include "logstore.serviceAccountName" .) "-pattern-ingester") .Values.patternIngester.serviceAccount.name }}
{{- else -}}
    {{ default (include "logstore.serviceAccountName" .) .Values.patternIngester.serviceAccount.name }}
{{- end -}}
{{- end -}}
