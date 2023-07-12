{{/*
Docker image name for logstore helm test
*/}}
{{- define "logstore.helmTestImage" -}}
{{- $dict := dict "service" .Values.test.image "global" .Values.global.image "defaultVersion" "latest" -}}
{{- include "logstore.baseImage" $dict -}}
{{- end -}}


{{/*
test common labels
*/}}
{{- define "logstore.helmTestLabels" -}}
{{ include "logstore.labels" . }}
app.kubernetes.io/component: helm-test
{{- end }}
