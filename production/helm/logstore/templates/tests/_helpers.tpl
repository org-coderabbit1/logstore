{{/*
Docker image name for logstore helm test
*/}}
{{- define "logstore.helm-test-image" -}}
{{- $dict := dict "service" .Values.test.image "global" .Values.global.image "defaultVersion" "latest" -}}
{{- include "logstore.baseImage" $dict -}}
{{- end -}}
