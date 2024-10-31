{{/*
memcached Service
Params:
  ctx = . context
  valuesSection = name of the section in values.yaml
  component = name of the component
valuesSection and component are specified separately because helm prefers camelcase for naming convetion and k8s components are named with snake case.
*/}}
{{- define "logstore.memcached.service" -}}
{{ with (index $.ctx.Values $.valuesSection) }}
{{- if .enabled -}}
apiVersion: v1
kind: Service
metadata:
  name: {{ include "logstore.resourceName" (dict "ctx" $.ctx "component" $.component) }}
  labels:
    {{- include "logstore.labels" $.ctx | nindent 4 }}
    app.kubernetes.io/component: "memcached-{{ $.component }}"
    {{- with .service.labels }}
    {{- toYaml . | nindent 4 }}
    {{- end }}
  annotations:
    {{- toYaml .service.annotations | nindent 4 }}
  namespace: {{ $.ctx.Release.Namespace | quote }}
spec:
  type: ClusterIP
  clusterIP: None
  ports:
    - name: memcached-client
      port: {{ .port }}
      targetPort: {{ .port }}
    {{ if $.ctx.Values.memcachedExporter.enabled -}}
    - name: http-metrics
      port: 9150
      targetPort: 9150
    {{ end }}
  selector:
    {{- include "logstore.selectorLabels" $.ctx | nindent 4 }}
    app.kubernetes.io/component: "memcached-{{ $.component }}"
{{- end -}}
{{- end -}}
{{- end -}}
