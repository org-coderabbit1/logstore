{{/*
Enforce valid label value.
See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
*/}}
{{- define "logstore.validLabelValue" -}}
{{- (regexReplaceAllLiteral "[^a-zA-Z0-9._-]" . "-") | trunc 63 | trimSuffix "-" | trimSuffix "_" | trimSuffix "." }}
{{- end }}

{{/*
Expand the name of the chart.
*/}}
{{- define "logstore.name" -}}
{{- $default := ternary "enterprise-logs" "logstore" .Values.enterprise.enabled }}
{{- coalesce .Values.nameOverride $default | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
singleBinary fullname
*/}}
{{- define "logstore.singleBinaryFullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Resource name template
Params:
  ctx = . context
  component = component name (optional)
  rolloutZoneName = rollout zone name (optional)
*/}}
{{- define "logstore.resourceName" -}}
{{- $resourceName := include "logstore.fullname" .ctx -}}
{{- if .component -}}{{- $resourceName = printf "%s-%s" $resourceName .component -}}{{- end -}}
{{- if and (not .component) .rolloutZoneName -}}{{- printf "Component name cannot be empty if rolloutZoneName (%s) is set" .rolloutZoneName | fail -}}{{- end -}}
{{- if .rolloutZoneName -}}{{- $resourceName = printf "%s-%s" $resourceName .rolloutZoneName -}}{{- end -}}
{{- if gt (len $resourceName) 253 -}}{{- printf "Resource name (%s) exceeds kubernetes limit of 253 character. To fix: shorten release name if this will be a fresh install or shorten zone names (e.g. \"a\" instead of \"zone-a\") if using zone-awareness." $resourceName | fail -}}{{- end -}}
{{- $resourceName -}}
{{- end -}}

{{/*
Return if deployment mode is simple scalable
*/}}
{{- define "logstore.deployment.isScalable" -}}
  {{- and (eq (include "logstore.isUsingObjectStorage" . ) "true") (or (eq .Values.deploymentMode "SingleBinary<->SimpleScalable") (eq .Values.deploymentMode "SimpleScalable") (eq .Values.deploymentMode "SimpleScalable<->Distributed")) }}
{{- end -}}

{{/*
Return if deployment mode is single binary
*/}}
{{- define "logstore.deployment.isSingleBinary" -}}
  {{- or (eq .Values.deploymentMode "SingleBinary") (eq .Values.deploymentMode "SingleBinary<->SimpleScalable") }}
{{- end -}}

{{/*
Return if deployment mode is distributed
*/}}
{{- define "logstore.deployment.isDistributed" -}}
  {{- and (eq (include "logstore.isUsingObjectStorage" . ) "true") (or (eq .Values.deploymentMode "Distributed") (eq .Values.deploymentMode "SimpleScalable<->Distributed")) }}
{{- end -}}


{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "logstore.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := include "logstore.name" . }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Cluster label for rules and alerts.
*/}}
{{- define "logstore.clusterLabel" -}}
{{- if .Values.clusterLabelOverride }}
{{- .Values.clusterLabelOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := include "logstore.name" . }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

{{/* Create a default storage config that uses filesystem storage
This is required for CI, but Logstore will not be queryable with this default
applied, thus it is encouraged that users override this.
*/}}
{{- define "logstore.storageConfig" -}}
{{- if .Values.logstore.storageConfig -}}
{{- .Values.logstore.storageConfig | toYaml | nindent 4 -}}
{{- else }}
{{- .Values.logstore.defaultStorageConfig | toYaml | nindent 4 }}
{{- end}}
{{- end}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "logstore.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "logstore.labels" -}}
helm.sh/chart: {{ include "logstore.chart" . }}
{{ include "logstore.selectorLabels" . }}
{{- if or (.Chart.AppVersion) (.Values.logstore.image.tag) }}
app.kubernetes.io/version: {{ include "logstore.validLabelValue" (.Values.logstore.image.tag | default .Chart.AppVersion) | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "logstore.selectorLabels" -}}
app.kubernetes.io/name: {{ include "logstore.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "logstore.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "logstore.name" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Base template for building docker image reference
*/}}
{{- define "logstore.baseImage" }}
{{- $registry := .global.registry | default .service.registry | default "" -}}
{{- $repository := .service.repository | default "" -}}
{{- $ref := ternary (printf ":%s" (.service.tag | default .defaultVersion | toString)) (printf "@%s" .service.digest) (empty .service.digest) -}}
{{- if and $registry $repository -}}
  {{- printf "%s/%s%s" $registry $repository $ref -}}
{{- else -}}
  {{- printf "%s%s%s" $registry $repository $ref -}}
{{- end -}}
{{- end -}}

{{/*
Docker image name for Logstore
*/}}
{{- define "logstore.logstoreImage" -}}
{{- $dict := dict "service" .Values.logstore.image "global" .Values.global.image "defaultVersion" .Chart.AppVersion -}}
{{- include "logstore.baseImage" $dict -}}
{{- end -}}

{{/*
Docker image name for enterprise logs
*/}}
{{- define "logstore.enterpriseImage" -}}
{{- $dict := dict "service" .Values.enterprise.image "global" .Values.global.image "defaultVersion" .Values.enterprise.version -}}
{{- include "logstore.baseImage" $dict -}}
{{- end -}}

{{/*
Docker image name
*/}}
{{- define "logstore.image" -}}
{{- if .Values.enterprise.enabled -}}{{- include "logstore.enterpriseImage" . -}}{{- else -}}{{- include "logstore.logstoreImage" . -}}{{- end -}}
{{- end -}}

{{/*
Docker image name for kubectl container
*/}}
{{- define "logstore.kubectlImage" -}}
{{- $dict := dict "service" .Values.kubectlImage "global" .Values.global.image "defaultVersion" "latest" -}}
{{- include "logstore.baseImage" $dict -}}
{{- end -}}

{{/*
Generated storage config for logstore common config
*/}}
{{- define "logstore.commonStorageConfig" -}}
{{- if .Values.minio.enabled -}}
s3:
  endpoint: {{ include "logstore.minio" $ }}
  bucketnames: chunks
  secret_access_key: {{ $.Values.minio.rootPassword }}
  access_key_id: {{ $.Values.minio.rootUser }}
  s3forcepathstyle: true
  insecure: true
{{- else if eq .Values.logstore.storage.type "s3" -}}
{{- with .Values.logstore.storage.s3 }}
s3:
  {{- with .s3 }}
  s3: {{ . }}
  {{- end }}
  {{- with .endpoint }}
  endpoint: {{ . }}
  {{- end }}
  {{- with .region }}
  region: {{ . }}
  {{- end}}
  bucketnames: {{ $.Values.logstore.storage.bucketNames.chunks }}
  {{- with .secretAccessKey }}
  secret_access_key: {{ . }}
  {{- end }}
  {{- with .accessKeyId }}
  access_key_id: {{ . }}
  {{- end }}
  {{- with .signatureVersion }}
  signature_version: {{ . }}
  {{- end }}
  s3forcepathstyle: {{ .s3ForcePathStyle }}
  insecure: {{ .insecure }}
  {{- with .disable_dualstack }}
  disable_dualstack: {{ . }}
  {{- end }}
  {{- with .http_config}}
  http_config:
{{ toYaml . | indent 4 }}
  {{- end }}
  {{- with .backoff_config}}
  backoff_config:
{{ toYaml . | indent 4 }}
  {{- end }}
  {{- with .sse }}
  sse:
{{ toYaml . | indent 4 }}
  {{- end }}
{{- end -}}

{{- else if eq .Values.logstore.storage.type "gcs" -}}
{{- with .Values.logstore.storage.gcs }}
gcs:
  bucket_name: {{ $.Values.logstore.storage.bucketNames.chunks }}
  chunk_buffer_size: {{ .chunkBufferSize }}
  request_timeout: {{ .requestTimeout }}
  enable_http2: {{ .enableHttp2 }}
{{- end -}}
{{- else if eq .Values.logstore.storage.type "azure" -}}
{{- with .Values.logstore.storage.azure }}
azure:
  account_name: {{ .accountName }}
  {{- with .accountKey }}
  account_key: {{ . }}
  {{- end }}
  {{- with .connectionString }}
  connection_string: {{ . }}
  {{- end }}
  container_name: {{ $.Values.logstore.storage.bucketNames.chunks }}
  use_managed_identity: {{ .useManagedIdentity }}
  use_federated_token: {{ .useFederatedToken }}
  {{- with .userAssignedId }}
  user_assigned_id: {{ . }}
  {{- end }}
  {{- with .requestTimeout }}
  request_timeout: {{ . }}
  {{- end }}
  {{- with .endpointSuffix }}
  endpoint_suffix: {{ . }}
  {{- end }}
  {{- with .chunkDelimiter }}
  chunk_delimiter: {{ . }}
  {{- end }}
{{- end -}}
{{- else if eq .Values.logstore.storage.type "alibabacloud" -}}
{{- with .Values.logstore.storage.alibabacloud }}
alibabacloud:
  bucket: {{ $.Values.logstore.storage.bucketNames.chunks }}
  endpoint: {{ .endpoint }}
  access_key_id: {{ .accessKeyId }}
  secret_access_key: {{ .secretAccessKey }}
{{- end -}}
{{- else if eq .Values.logstore.storage.type "swift" -}}
{{- with .Values.logstore.storage.swift }}
swift:
{{ toYaml . | indent 2 }}
{{- end -}}
{{- else -}}
{{- with .Values.logstore.storage.filesystem }}
filesystem:
  chunks_directory: {{ .chunks_directory }}
  rules_directory: {{ .rules_directory }}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Storage config for ruler
*/}}
{{- define "logstore.rulerStorageConfig" -}}
{{- if .Values.minio.enabled -}}
type: "s3"
s3:
  bucketnames: ruler
{{- else if eq .Values.logstore.storage.type "s3" -}}
{{- with .Values.logstore.storage.s3 }}
type: "s3"
s3:
  {{- with .s3 }}
  s3: {{ . }}
  {{- end }}
  {{- with .endpoint }}
  endpoint: {{ . }}
  {{- end }}
  {{- with .region }}
  region: {{ . }}
  {{- end}}
  bucketnames: {{ $.Values.logstore.storage.bucketNames.ruler }}
  {{- with .secretAccessKey }}
  secret_access_key: {{ . }}
  {{- end }}
  {{- with .accessKeyId }}
  access_key_id: {{ . }}
  {{- end }}
  s3forcepathstyle: {{ .s3ForcePathStyle }}
  insecure: {{ .insecure }}
  {{- with .http_config }}
  http_config: {{ toYaml . | nindent 6 }}
  {{- end }}
{{- end -}}
{{- else if eq .Values.logstore.storage.type "gcs" -}}
{{- with .Values.logstore.storage.gcs }}
type: "gcs"
gcs:
  bucket_name: {{ $.Values.logstore.storage.bucketNames.ruler }}
  chunk_buffer_size: {{ .chunkBufferSize }}
  request_timeout: {{ .requestTimeout }}
  enable_http2: {{ .enableHttp2 }}
{{- end -}}
{{- else if eq .Values.logstore.storage.type "azure" -}}
{{- with .Values.logstore.storage.azure }}
type: "azure"
azure:
  account_name: {{ .accountName }}
  {{- with .accountKey }}
  account_key: {{ . }}
  {{- end }}
  {{- with .connectionString }}
  connection_string: {{ . }}
  {{- end }}
  container_name: {{ $.Values.logstore.storage.bucketNames.ruler }}
  use_managed_identity: {{ .useManagedIdentity }}
  use_federated_token: {{ .useFederatedToken }}
  {{- with .userAssignedId }}
  user_assigned_id: {{ . }}
  {{- end }}
  {{- with .requestTimeout }}
  request_timeout: {{ . }}
  {{- end }}
  {{- with .endpointSuffix }}
  endpoint_suffix: {{ . }}
  {{- end }}
{{- end -}}
{{- else if eq .Values.logstore.storage.type "swift" -}}
{{- with .Values.logstore.storage.swift }}
swift:
  {{- with .auth_version }}
  auth_version: {{ . }}
  {{- end }}
  auth_url: {{ .auth_url }}
  {{- with .internal }}
  internal: {{ . }}
  {{- end }}
  username: {{ .username }}
  user_domain_name: {{ .user_domain_name }}
  {{- with .user_domain_id }}
  user_domain_id: {{ . }}
  {{- end }}
  {{- with .user_id }}
  user_id: {{ . }}
  {{- end }}
  password: {{ .password }}
  {{- with .domain_id }}
  domain_id: {{ . }}
  {{- end }}
  domain_name: {{ .domain_name }}
  project_id: {{ .project_id }}
  project_name: {{ .project_name }}
  project_domain_id: {{ .project_domain_id }}
  project_domain_name: {{ .project_domain_name }}
  region_name: {{ .region_name }}
  container_name: {{ .container_name }}
  max_retries: {{ .max_retries | default 3 }}
  connect_timeout: {{ .connect_timeout | default "10s" }}
  request_timeout: {{ .request_timeout | default "5s" }}
{{- end -}}
{{- else }}
type: "local"
{{- end -}}
{{- end -}}

{{/* Logstore ruler config */}}
{{- define "logstore.rulerConfig" }}
ruler:
  storage:
    {{- include "logstore.rulerStorageConfig" . | nindent 4}}
{{- if (not (empty .Values.logstore.rulerConfig)) }}
{{- toYaml .Values.logstore.rulerConfig | nindent 2}}
{{- end }}
{{- end }}

{{/* Enterprise Logs Admin API storage config */}}
{{- define "enterprise-logs.adminAPIStorageConfig" }}
storage:
  {{- if .Values.minio.enabled }}
  backend: "s3"
  s3:
    bucket_name: admin
  {{- else if eq .Values.logstore.storage.type "s3" -}}
  {{- with .Values.logstore.storage.s3 }}
  backend: "s3"
  s3:
    bucket_name: {{ $.Values.logstore.storage.bucketNames.admin }}
  {{- end -}}
  {{- else if eq .Values.logstore.storage.type "gcs" -}}
  {{- with .Values.logstore.storage.gcs }}
  backend: "gcs"
  gcs:
    bucket_name: {{ $.Values.logstore.storage.bucketNames.admin }}
  {{- end -}}
  {{- else if eq .Values.logstore.storage.type "azure" -}}
  {{- with .Values.logstore.storage.azure }}
  backend: "azure"
  azure:
    account_name: {{ .accountName }}
    {{- with .accountKey }}
    account_key: {{ . }}
    {{- end }}
    {{- with .connectionString }}
    connection_string: {{ . }}
    {{- end }}
    container_name: {{ $.Values.logstore.storage.bucketNames.admin }}
    {{- with .endpointSuffix }}
    endpoint_suffix: {{ . }}
    {{- end }}
  {{- end -}}
  {{- else if eq .Values.logstore.storage.type "swift" -}}
  {{- with .Values.logstore.storage.swift }}
  backend: "swift"
  swift:
    {{- with .auth_version }}
    auth_version: {{ . }}
    {{- end }}
    auth_url: {{ .auth_url }}
    {{- with .internal }}
    internal: {{ . }}
    {{- end }}
    username: {{ .username }}
    user_domain_name: {{ .user_domain_name }}
    {{- with .user_domain_id }}
    user_domain_id: {{ . }}
    {{- end }}
    {{- with .user_id }}
    user_id: {{ . }}
    {{- end }}
    password: {{ .password }}
    {{- with .domain_id }}
    domain_id: {{ . }}
    {{- end }}
    domain_name: {{ .domain_name }}
    project_id: {{ .project_id }}
    project_name: {{ .project_name }}
    project_domain_id: {{ .project_domain_id }}
    project_domain_name: {{ .project_domain_name }}
    region_name: {{ .region_name }}
    container_name: {{ .container_name }}
    max_retries: {{ .max_retries | default 3 }}
    connect_timeout: {{ .connect_timeout | default "10s" }}
    request_timeout: {{ .request_timeout | default "5s" }}
  {{- end -}}
  {{- else }}
  backend: "filesystem"
  filesystem:
    dir: {{ .Values.logstore.storage.filesystem.admin_api_directory }}
  {{- end -}}
{{- end }}

{{/*
Calculate the config from structured and unstructured text input
*/}}
{{- define "logstore.calculatedConfig" -}}
{{ tpl (mergeOverwrite (tpl .Values.logstore.config . | fromYaml) .Values.logstore.structuredConfig | toYaml) . }}
{{- end }}

{{/*
The volume to mount for logstore configuration
*/}}
{{- define "logstore.configVolume" -}}
{{- if eq .Values.logstore.configStorageType "Secret" -}}
secret:
  secretName: {{ tpl .Values.logstore.configObjectName . }}
{{- else -}}
configMap:
  name: {{ tpl .Values.logstore.configObjectName . }}
  items:
    - key: "config.yaml"
      path: "config.yaml"
{{- end -}}
{{- end -}}

{{/*
Memcached Docker image
*/}}
{{- define "logstore.memcachedImage" -}}
{{- $dict := dict "service" .Values.memcached.image "global" .Values.global.image -}}
{{- include "logstore.image" $dict -}}
{{- end }}

{{/*
Memcached Exporter Docker image
*/}}
{{- define "logstore.memcachedExporterImage" -}}
{{- $dict := dict "service" .Values.memcachedExporter.image "global" .Values.global.image -}}
{{- include "logstore.image" $dict -}}
{{- end }}

{{/* Allow KubeVersion to be overridden. */}}
{{- define "logstore.kubeVersion" -}}
  {{- default .Capabilities.KubeVersion.Version .Values.kubeVersionOverride -}}
{{- end -}}

{{/*
Return the appropriate apiVersion for ingress.
*/}}
{{- define "logstore.ingress.apiVersion" -}}
  {{- if and (.Capabilities.APIVersions.Has "networking.k8s.io/v1") (semverCompare ">= 1.19-0" (include "logstore.kubeVersion" .)) -}}
      {{- print "networking.k8s.io/v1" -}}
  {{- else if .Capabilities.APIVersions.Has "networking.k8s.io/v1beta1" -}}
    {{- print "networking.k8s.io/v1beta1" -}}
  {{- else -}}
    {{- print "extensions/v1beta1" -}}
  {{- end -}}
{{- end -}}

{{/*
Return if ingress is stable.
*/}}
{{- define "logstore.ingress.isStable" -}}
  {{- eq (include "logstore.ingress.apiVersion" .) "networking.k8s.io/v1" -}}
{{- end -}}

{{/*
Return if ingress supports ingressClassName.
*/}}
{{- define "logstore.ingress.supportsIngressClassName" -}}
  {{- or (eq (include "logstore.ingress.isStable" .) "true") (and (eq (include "logstore.ingress.apiVersion" .) "networking.k8s.io/v1beta1") (semverCompare ">= 1.18-0" (include "logstore.kubeVersion" .))) -}}
{{- end -}}

{{/*
Return if ingress supports pathType.
*/}}
{{- define "logstore.ingress.supportsPathType" -}}
  {{- or (eq (include "logstore.ingress.isStable" .) "true") (and (eq (include "logstore.ingress.apiVersion" .) "networking.k8s.io/v1beta1") (semverCompare ">= 1.18-0" (include "logstore.kubeVersion" .))) -}}
{{- end -}}

{{/*
Generate list of ingress service paths based on deployment type
*/}}
{{- define "logstore.ingress.servicePaths" -}}
{{- if (eq (include "logstore.deployment.isSingleBinary" .) "true") -}}
{{- include "logstore.ingress.singleBinaryServicePaths" . }}
{{- else if (eq (include "logstore.deployment.isDistributed" .) "true") -}}
{{- include "logstore.ingress.distributedServicePaths" . }}
{{- else if and (eq (include "logstore.deployment.isScalable" .) "true") (not .Values.read.legacyReadTarget ) -}}
{{- include "logstore.ingress.scalableServicePaths" . }}
{{- else -}}
{{- include "logstore.ingress.legacyScalableServicePaths" . }}
{{- end -}}
{{- end -}}


{{/*
Ingress service paths for distributed deployment
*/}}
{{- define "logstore.ingress.distributedServicePaths" -}}
{{- $distributorServiceName := include "logstore.distributorFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $distributorServiceName "paths" .Values.ingress.paths.distributor )}}
{{- $queryFrontendServiceName := include "logstore.queryFrontendFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $queryFrontendServiceName "paths" .Values.ingress.paths.queryFrontend )}}
{{- $rulerServiceName := include "logstore.rulerFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $rulerServiceName "paths" .Values.ingress.paths.ruler)}}
{{- end -}}

{{/*
Ingress service paths for legacy simple scalable deployment when backend components were part of read component.
*/}}
{{- define "logstore.ingress.scalableServicePaths" -}}
{{- $readServiceName := include "logstore.readFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $readServiceName "paths" .Values.ingress.paths.queryFrontend )}}
{{- $writeServiceName := include "logstore.writeFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $writeServiceName "paths" .Values.ingress.paths.distributor )}}
{{- $backendServiceName := include "logstore.backendFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $backendServiceName "paths" .Values.ingress.paths.ruler )}}
{{- end -}}

{{/*
Ingress service paths for legacy simple scalable deployment
*/}}
{{- define "logstore.ingress.legacyScalableServicePaths" -}}
{{- $readServiceName := include "logstore.readFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $readServiceName "paths" .Values.ingress.paths.queryFrontend )}}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $readServiceName "paths" .Values.ingress.paths.ruler )}}
{{- $writeServiceName := include "logstore.writeFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $writeServiceName "paths" .Values.ingress.paths.distributor )}}
{{- end -}}

{{/*
Ingress service paths for single binary deployment
*/}}
{{- define "logstore.ingress.singleBinaryServicePaths" -}}
{{- $serviceName := include "logstore.singleBinaryFullname" . }}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $serviceName "paths" .Values.ingress.paths.distributor )}}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $serviceName "paths" .Values.ingress.paths.queryFrontend )}}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "serviceName" $serviceName "paths" .Values.ingress.paths.ruler )}}
{{- end -}}

{{/*
Ingress service path helper function
Params:
  ctx = . context
  serviceName = fully qualified k8s service name
  paths = list of url paths to allow ingress for
*/}}
{{- define "logstore.ingress.servicePath" -}}
{{- $ingressApiIsStable := eq (include "logstore.ingress.isStable" .ctx) "true" -}}
{{- $ingressSupportsPathType := eq (include "logstore.ingress.supportsPathType" .ctx) "true" -}}
{{- range .paths }}
- path: {{ . }}
  {{- if $ingressSupportsPathType }}
  pathType: Prefix
  {{- end }}
  backend:
    {{- if $ingressApiIsStable }}
    service:
      name: {{ $.serviceName }}
      port:
        number: {{ $.ctx.Values.logstore.server.http_listen_port }}
    {{- else }}
    serviceName: {{ $.serviceName }}
    servicePort: {{ $.ctx.Values.logstore.server.http_listen_port }}
    {{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create the service endpoint including port for MinIO.
*/}}
{{- define "logstore.minio" -}}
{{- if .Values.minio.enabled -}}
{{- .Values.minio.address | default (printf "%s-%s.%s.svc:%s" .Release.Name "minio" .Release.Namespace (.Values.minio.service.port | toString)) -}}
{{- end -}}
{{- end -}}

{{/* Determine if deployment is using object storage */}}
{{- define "logstore.isUsingObjectStorage" -}}
{{- or (eq .Values.logstore.storage.type "gcs") (eq .Values.logstore.storage.type "s3") (eq .Values.logstore.storage.type "azure") (eq .Values.logstore.storage.type "swift") (eq .Values.logstore.storage.type "alibabacloud") -}}
{{- end -}}

{{/* Configure the correct name for the memberlist service */}}
{{- define "logstore.memberlist" -}}
{{ include "logstore.name" . }}-memberlist
{{- end -}}

{{/* Determine the public host for the Logstore cluster */}}
{{- define "logstore.host" -}}
{{- $isSingleBinary := eq (include "logstore.deployment.isSingleBinary" .) "true" -}}
{{- $url := printf "%s.%s.svc.%s.:%s" (include "logstore.gatewayFullname" .) .Release.Namespace .Values.global.clusterDomain (.Values.gateway.service.port | toString)  }}
{{- if and $isSingleBinary (not .Values.gateway.enabled)  }}
  {{- $url = printf "%s.%s.svc.%s.:%s" (include "logstore.singleBinaryFullname" .) .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
{{- end }}
{{- printf "%s" $url -}}
{{- end -}}

{{/* Determine the public endpoint for the Logstore cluster */}}
{{- define "logstore.address" -}}
{{- printf "http://%s" (include "logstore.host" . ) -}}
{{- end -}}

{{/* Name of the cluster */}}
{{- define "logstore.clusterName" -}}
{{- $name := .Values.enterprise.cluster_name | default .Release.Name }}
{{- printf "%s" $name -}}
{{- end -}}

{{/* Name of kubernetes secret to persist GEL admin token to */}}
{{- define "enterprise-logs.adminTokenSecret" }}
{{- .Values.enterprise.adminToken.secret | default (printf "%s-admin-token" (include "logstore.name" . )) -}}
{{- end -}}

{{/* Prefix for provisioned secrets created for each provisioned tenant */}}
{{- define "enterprise-logs.provisionedSecretPrefix" }}
{{- .Values.enterprise.provisioner.provisionedSecretPrefix | default (printf "%s-provisioned" (include "logstore.name" . )) -}}
{{- end -}}

{{/* Name of kubernetes secret to persist canary credentials in */}}
{{- define "enterprise-logs.selfMonitoringTenantSecret" }}
{{- .Values.enterprise.canarySecret | default (printf "%s-%s" (include "enterprise-logs.provisionedSecretPrefix" . ) .Values.monitoring.selfMonitoring.tenant.name) -}}
{{- end -}}

{{/* Snippet for the nginx file used by gateway */}}
{{- define "logstore.nginxFile" }}
worker_processes  5;  ## Default: 1
error_log  /dev/stderr;
pid        /tmp/nginx.pid;
worker_rlimit_nofile 8192;

events {
  worker_connections  4096;  ## Default: 1024
}

http {
  client_body_temp_path /tmp/client_temp;
  proxy_temp_path       /tmp/proxy_temp_path;
  fastcgi_temp_path     /tmp/fastcgi_temp;
  uwsgi_temp_path       /tmp/uwsgi_temp;
  scgi_temp_path        /tmp/scgi_temp;

  client_max_body_size  {{ .Values.gateway.nginxConfig.clientMaxBodySize }};

  proxy_read_timeout    600; ## 10 minutes
  proxy_send_timeout    600;
  proxy_connect_timeout 600;

  proxy_http_version    1.1;

  default_type application/octet-stream;
  log_format   {{ .Values.gateway.nginxConfig.logFormat }}

  {{- if .Values.gateway.verboseLogging }}
  access_log   /dev/stderr  main;
  {{- else }}

  map $status $loggable {
    ~^[23]  0;
    default 1;
  }
  access_log   /dev/stderr  main  if=$loggable;
  {{- end }}

  sendfile     on;
  tcp_nopush   on;
  {{- if .Values.gateway.nginxConfig.resolver }}
  resolver {{ .Values.gateway.nginxConfig.resolver }};
  {{- else }}
  resolver {{ .Values.global.dnsService }}.{{ .Values.global.dnsNamespace }}.svc.{{ .Values.global.clusterDomain }}.;
  {{- end }}

  {{- with .Values.gateway.nginxConfig.httpSnippet }}
  {{- tpl . $ | nindent 2 }}
  {{- end }}

  server {
    {{- if (.Values.gateway.nginxConfig.ssl) }}
    listen             8080 ssl;
    {{- if .Values.gateway.nginxConfig.enableIPv6 }}
    listen             [::]:8080 ssl;
    {{- end }}
    {{- else }}
    listen             8080;
    {{- if .Values.gateway.nginxConfig.enableIPv6 }}
    listen             [::]:8080;
    {{- end }}
    {{- end }}

    {{- if .Values.gateway.basicAuth.enabled }}
    auth_basic           "Logstore";
    auth_basic_user_file /etc/nginx/secrets/.htpasswd;
    {{- end }}

    location = / {
      return 200 'OK';
      auth_basic off;
    }

    ########################################################
    # Configure backend targets

    {{- $backendHost := include "logstore.backendFullname" .}}
    {{- $readHost := include "logstore.readFullname" .}}
    {{- $writeHost := include "logstore.writeFullname" .}}

    {{- if .Values.read.legacyReadTarget }}
    {{- $backendHost = include "logstore.readFullname" . }}
    {{- end }}

    {{- $httpSchema := .Values.gateway.nginxConfig.schema }}

    {{- $writeUrl    := printf "%s://%s.%s.svc.%s:%s" $httpSchema $writeHost   .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
    {{- $readUrl     := printf "%s://%s.%s.svc.%s:%s" $httpSchema $readHost    .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
    {{- $backendUrl  := printf "%s://%s.%s.svc.%s:%s" $httpSchema $backendHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}

    {{- if .Values.gateway.nginxConfig.customWriteUrl }}
    {{- $writeUrl  = .Values.gateway.nginxConfig.customWriteUrl }}
    {{- end }}
    {{- if .Values.gateway.nginxConfig.customReadUrl }}
    {{- $readUrl = .Values.gateway.nginxConfig.customReadUrl }}
    {{- end }}
    {{- if .Values.gateway.nginxConfig.customBackendUrl }}
    {{- $backendUrl = .Values.gateway.nginxConfig.customBackendUrl }}
    {{- end }}

    {{- $singleBinaryHost := include "logstore.singleBinaryFullname" . }}
    {{- $singleBinaryUrl  := printf "%s://%s.%s.svc.%s:%s" $httpSchema $singleBinaryHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}

    {{- $distributorHost := include "logstore.distributorFullname" .}}
    {{- $ingesterHost := include "logstore.ingesterFullname" .}}
    {{- $queryFrontendHost := include "logstore.queryFrontendFullname" .}}
    {{- $indexGatewayHost := include "logstore.indexGatewayFullname" .}}
    {{- $rulerHost := include "logstore.rulerFullname" .}}
    {{- $compactorHost := include "logstore.compactorFullname" .}}
    {{- $schedulerHost := include "logstore.querySchedulerFullname" .}}


    {{- $distributorUrl := printf "%s://%s.%s.svc.%s:%s" $httpSchema $distributorHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) -}}
    {{- $ingesterUrl := printf "%s://%s.%s.svc.%s:%s" $httpSchema $ingesterHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
    {{- $queryFrontendUrl := printf "%s://%s.%s.svc.%s:%s" $httpSchema $queryFrontendHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
    {{- $indexGatewayUrl := printf "%s://%s.%s.svc.%s:%s" $httpSchema $indexGatewayHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
    {{- $rulerUrl := printf "%s://%s.%s.svc.%s:%s" $httpSchema $rulerHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
    {{- $compactorUrl := printf "%s://%s.%s.svc.%s:%s" $httpSchema $compactorHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}
    {{- $schedulerUrl := printf "%s://%s.%s.svc.%s:%s" $httpSchema $schedulerHost .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.http_listen_port | toString) }}

    {{- if eq (include "logstore.deployment.isSingleBinary" .) "true"}}
    {{- $distributorUrl = $singleBinaryUrl }}
    {{- $ingesterUrl = $singleBinaryUrl }}
    {{- $queryFrontendUrl = $singleBinaryUrl }}
    {{- $indexGatewayUrl = $singleBinaryUrl }}
    {{- $rulerUrl = $singleBinaryUrl }}
    {{- $compactorUrl = $singleBinaryUrl }}
    {{- $schedulerUrl = $singleBinaryUrl }}
    {{- else if eq (include "logstore.deployment.isScalable" .) "true"}}
    {{- $distributorUrl = $writeUrl }}
    {{- $ingesterUrl = $writeUrl }}
    {{- $queryFrontendUrl = $readUrl }}
    {{- $indexGatewayUrl = $backendUrl }}
    {{- $rulerUrl = $backendUrl }}
    {{- $compactorUrl = $backendUrl }}
    {{- $schedulerUrl = $backendUrl }}
    {{- end -}}

    # Distributor
    location = /api/prom/push {
      proxy_pass       {{ $distributorUrl }}$request_uri;
    }
    location = /logstore/api/v1/push {
      proxy_pass       {{ $distributorUrl }}$request_uri;
    }
    location = /distributor/ring {
      proxy_pass       {{ $distributorUrl }}$request_uri;
    }
    location = /otlp/v1/logs {
      proxy_pass       {{ $distributorUrl }}$request_uri;
    }

    # Ingester
    location = /flush {
      proxy_pass       {{ $ingesterUrl }}$request_uri;
    }
    location ^~ /ingester/ {
      proxy_pass       {{ $ingesterUrl }}$request_uri;
    }
    location = /ingester {
      internal;        # to suppress 301
    }

    # Ring
    location = /ring {
      proxy_pass       {{ $ingesterUrl }}$request_uri;
    }

    # MemberListKV
    location = /memberlist {
      proxy_pass       {{ $ingesterUrl }}$request_uri;
    }

    # Ruler
    location = /ruler/ring {
      proxy_pass       {{ $rulerUrl }}$request_uri;
    }
    location = /api/prom/rules {
      proxy_pass       {{ $rulerUrl }}$request_uri;
    }
    location ^~ /api/prom/rules/ {
      proxy_pass       {{ $rulerUrl }}$request_uri;
    }
    location = /logstore/api/v1/rules {
      proxy_pass       {{ $rulerUrl }}$request_uri;
    }
    location ^~ /logstore/api/v1/rules/ {
      proxy_pass       {{ $rulerUrl }}$request_uri;
    }
    location = /prometheus/api/v1/alerts {
      proxy_pass       {{ $rulerUrl }}$request_uri;
    }
    location = /prometheus/api/v1/rules {
      proxy_pass       {{ $rulerUrl }}$request_uri;
    }

    # Compactor
    location = /compactor/ring {
      proxy_pass       {{ $compactorUrl }}$request_uri;
    }
    location = /logstore/api/v1/delete {
      proxy_pass       {{ $compactorUrl }}$request_uri;
    }
    location = /logstore/api/v1/cache/generation_numbers {
      proxy_pass       {{ $compactorUrl }}$request_uri;
    }

    # IndexGateway
    location = /indexgateway/ring {
      proxy_pass       {{ $indexGatewayUrl }}$request_uri;
    }

    # QueryScheduler
    location = /scheduler/ring {
      proxy_pass       {{ $schedulerUrl }}$request_uri;
    }

    # Config
    location = /config {
      proxy_pass       {{ $ingesterUrl }}$request_uri;
    }

    {{- if and .Values.enterprise.enabled .Values.enterprise.adminApi.enabled }}
    # Admin API
    location ^~ /admin/api/ {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }
    location = /admin/api {
      internal;        # to suppress 301
    }
    {{- end }}


    # QueryFrontend, Querier
    location = /api/prom/tail {
      proxy_pass       {{ $queryFrontendUrl }}$request_uri;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection "upgrade";
    }
    location = /logstore/api/v1/tail {
      proxy_pass       {{ $queryFrontendUrl }}$request_uri;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection "upgrade";
    }
    location ^~ /api/prom/ {
      proxy_pass       {{ $queryFrontendUrl }}$request_uri;
    }
    location = /api/prom {
      internal;        # to suppress 301
    }
    location ^~ /logstore/api/v1/ {
      proxy_pass       {{ $queryFrontendUrl }}$request_uri;
    }
    location = /logstore/api/v1 {
      internal;        # to suppress 301
    }

    {{- with .Values.gateway.nginxConfig.serverSnippet }}
    {{ . | nindent 4 }}
    {{- end }}
  }
}
{{- end }}

{{/* Configure enableServiceLinks in pod */}}
{{- define "logstore.enableServiceLinks" -}}
{{- if semverCompare ">=1.13-0" (include "logstore.kubeVersion" .) -}}
{{- if or (.Values.logstore.enableServiceLinks) (ne .Values.logstore.enableServiceLinks false) -}}
enableServiceLinks: true
{{- else -}}
enableServiceLinks: false
{{- end -}}
{{- end -}}
{{- end -}}

{{/* Determine compactor address based on target configuration */}}
{{- define "logstore.compactorAddress" -}}
{{- $isSimpleScalable := eq (include "logstore.deployment.isScalable" .) "true" -}}
{{- $isDistributed := eq (include "logstore.deployment.isDistributed" .) "true" -}}
{{- $isSingleBinary := eq (include "logstore.deployment.isSingleBinary" .) "true" -}}
{{- $compactorAddress := include "logstore.backendFullname" . -}}
{{- if and $isSimpleScalable .Values.read.legacyReadTarget -}}
{{/* 2 target configuration */}}
{{- $compactorAddress = include "logstore.readFullname" . -}}
{{- else if $isSingleBinary -}}
{{/* single binary */}}
{{- $compactorAddress = include "logstore.singleBinaryFullname" . -}}
{{/* distributed */}}
{{- else if $isDistributed -}}
{{- $compactorAddress = include "logstore.compactorFullname" . -}}
{{- end -}}
{{- printf "http://%s:%s" $compactorAddress (.Values.logstore.server.http_listen_port | toString) }}
{{- end }}

{{/* Determine query-scheduler address */}}
{{- define "logstore.querySchedulerAddress" -}}
{{- $schedulerAddress := ""}}
{{- $isDistributed := eq (include "logstore.deployment.isDistributed" .) "true" -}}
{{- if $isDistributed -}}
{{- $schedulerAddress = printf "%s.%s.svc.%s:%s" (include "logstore.querySchedulerFullname" .) .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.grpc_listen_port | toString) -}}
{{- end -}}
{{- printf "%s" $schedulerAddress }}
{{- end }}

{{/* Determine querier address */}}
{{- define "logstore.querierAddress" -}}
{{- $querierAddress := "" }}
{{- $isDistributed := eq (include "logstore.deployment.isDistributed" .) "true" -}}
{{- if $isDistributed -}}
{{- $querierHost := include "logstore.querierFullname" .}}
{{- $querierUrl := printf "http://%s.%s.svc.%s:3100" $querierHost .Release.Namespace .Values.global.clusterDomain }}
{{- $querierAddress = $querierUrl }}
{{- end -}}
{{- printf "%s" $querierAddress }}
{{- end }}

{{/* Determine index-gateway address */}}
{{- define "logstore.indexGatewayAddress" -}}
{{- $idxGatewayAddress := ""}}
{{- $isDistributed := eq (include "logstore.deployment.isDistributed" .) "true" -}}
{{- $isScalable := eq (include "logstore.deployment.isScalable" .) "true" -}}
{{- if $isDistributed -}}
{{- $idxGatewayAddress = printf "dns+%s-headless.%s.svc.%s:%s" (include "logstore.indexGatewayFullname" .) .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.grpc_listen_port | toString) -}}
{{- end -}}
{{- if $isScalable -}}
{{- $idxGatewayAddress = printf "dns+%s-headless.%s.svc.%s:%s" (include "logstore.backendFullname" .) .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.grpc_listen_port | toString) -}}
{{- end -}}
{{- printf "%s" $idxGatewayAddress }}
{{- end }}

{{/* Determine bloom-planner address */}}
{{- define "logstore.bloomPlannerAddress" -}}
{{- $bloomPlannerAddress := ""}}
{{- $isDistributed := eq (include "logstore.deployment.isDistributed" .) "true" -}}
{{- $isScalable := eq (include "logstore.deployment.isScalable" .) "true" -}}
{{- if $isDistributed -}}
{{- $bloomPlannerAddress = printf "%s-headless.%s.svc.%s:%s" (include "logstore.bloomPlannerFullname" .) .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.grpc_listen_port | toString) -}}
{{- end -}}
{{- if $isScalable -}}
{{- $bloomPlannerAddress = printf "%s-headless.%s.svc.%s:%s" (include "logstore.backendFullname" .) .Release.Namespace .Values.global.clusterDomain (.Values.logstore.server.grpc_listen_port | toString) -}}
{{- end -}}
{{- printf "%s" $bloomPlannerAddress}}
{{- end }}

{{/* Determine bloom-gateway address */}}
{{- define "logstore.bloomGatewayAddresses" -}}
{{- $bloomGatewayAddresses := ""}}
{{- $isDistributed := eq (include "logstore.deployment.isDistributed" .) "true" -}}
{{- $isScalable := eq (include "logstore.deployment.isScalable" .) "true" -}}
{{- if $isDistributed -}}
{{- $bloomGatewayAddresses = printf "dnssrvnoa+_grpc._tcp.%s-headless.%s.svc.%s" (include "logstore.bloomGatewayFullname" .) .Release.Namespace .Values.global.clusterDomain -}}
{{- end -}}
{{- if $isScalable -}}
{{- $bloomGatewayAddresses = printf "dnssrvnoa+_grpc._tcp.%s-headless.%s.svc.%s" (include "logstore.backendFullname" .) .Release.Namespace .Values.global.clusterDomain -}}
{{- end -}}
{{- printf "%s" $bloomGatewayAddresses}}
{{- end }}

{{- define "logstore.config.checksum" -}}
checksum/config: {{ include (print .Template.BasePath "/config.yaml") . | sha256sum }}
{{- end -}}

{{/*
Return the appropriate apiVersion for PodDisruptionBudget.
*/}}
{{- define "logstore.pdb.apiVersion" -}}
  {{- if and (.Capabilities.APIVersions.Has "policy/v1") (semverCompare ">=1.21-0" (include "logstore.kubeVersion" .)) -}}
    {{- print "policy/v1" -}}
  {{- else -}}
    {{- print "policy/v1beta1" -}}
  {{- end -}}
{{- end -}}

{{/*
Return the object store type for use with the test schema.
*/}}
{{- define "logstore.testSchemaObjectStore" -}}
  {{- if .Values.minio.enabled -}}
    s3
  {{- else -}}
    filesystem
  {{- end -}}
{{- end -}}

{{/*
Return the appropriate apiVersion for HorizontalPodAutoscaler.
*/}}
{{- define "logstore.hpa.apiVersion" -}}
  {{- if and (.Capabilities.APIVersions.Has "autoscaling/v2") (semverCompare ">= 1.19-0" (include "logstore.kubeVersion" .)) -}}
      {{- print "autoscaling/v2" -}}
  {{- else if .Capabilities.APIVersions.Has "autoscaling/v2beta2" -}}
    {{- print "autoscaling/v2beta2" -}}
  {{- else -}}
    {{- print "autoscaling/v2beta1" -}}
  {{- end -}}
{{- end -}}
