{{/*
memcached StatefulSet
Params:
  ctx = . context
  memcacheConfig = cache config
  valuesSection = name of the section in values.yaml
  component = name of the component
valuesSection and component are specified separately because helm prefers camelcase for naming convetion and k8s components are named with snake case.
*/}}
{{- define "logstore.memcached.pdb" -}}
{{ with $.memcacheConfig }}
{{- if and .enabled -}}
{{- if gt (int .replicas) 1 }}
apiVersion: {{ include "logstore.pdb.apiVersion" $.ctx }}
kind: PodDisruptionBudget
metadata:
  name: {{ include "logstore.resourceName" (dict "ctx" $.ctx "component" $.component "suffix" .suffix) }}
  namespace: {{ include "logstore.namespace" $.ctx }}
  labels:
    {{- include "logstore.selectorLabels" $.ctx | nindent 4 }}
    app.kubernetes.io/component: "memcached-{{ $.component }}{{ include "logstore.memcached.suffix" .suffix }}"
spec:
  selector:
    matchLabels:
      {{- include "logstore.selectorLabels" $.ctx | nindent 6 }}
      app.kubernetes.io/component: "memcached-{{ $.component }}{{ include "logstore.memcached.suffix" .suffix }}"
  {{- with .maxUnavailable }}
  maxUnavailable: {{ . }}
  {{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}
