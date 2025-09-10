{{/*
memcached Service
Params:
  ctx = . context
  valuesSection = name of the section in values.yaml
  memcacheConfig = cache config
  component = name of the component
valuesSection and component are specified separately because helm prefers camelcase for naming convetion and k8s components are named with snake case.
*/}}
{{- define "logstore.memcached.service" -}}
{{ with $.memcacheConfig }}
{{- if and .enabled ($.ctx.Values.memcached.enabled) -}}
apiVersion: v1
kind: Service
metadata:
  name: {{ include "logstore.resourceName" (dict "ctx" $.ctx "component" $.component "suffix" .suffix) }}
  labels:
    {{- include "logstore.labels" $.ctx | nindent 4 }}
    app.kubernetes.io/component: "memcached-{{ $.component }}{{ include "logstore.memcached.suffix" .suffix }}"
    {{- with .service.labels }}
    {{- toYaml . | nindent 4 }}
    {{- end }}
  annotations:
    {{- toYaml .service.annotations | nindent 4 }}
  namespace: {{ include "logstore.namespace" $.ctx | quote }}
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
    app.kubernetes.io/component: "memcached-{{ $.component }}{{ include "logstore.memcached.suffix" .suffix }}"
{{- end -}}
{{- end -}}
{{- end -}}
