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
  {{- and (eq (include "logstore.isUsingObjectStorage" . ) "true") (eq (int .Values.singleBinary.replicas) 0) }}
{{- end -}}

{{/*
Return if deployment mode is single binary
*/}}
{{- define "logstore.deployment.isSingleBinary" -}}
  {{- $nonZeroReplicas := gt (int .Values.singleBinary.replicas) 0 }}
  {{- or (eq (include "logstore.isUsingObjectStorage" . ) "false") ($nonZeroReplicas) }}
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
{{- $tag := .service.tag | default .defaultVersion | toString -}}
{{- if and $registry $repository -}}
  {{- printf "%s/%s:%s" $registry $repository $tag -}}
{{- else -}}
  {{- printf "%s%s:%s" $registry $repository $tag -}}
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
  bucketnames: {{ $.Values.logstore.storage.bucketNames.chunks }}
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
  s3forcepathstyle: {{ .s3ForcePathStyle }}
  insecure: {{ .insecure }}
  {{- with .http_config}}
  http_config:
    {{- with .idle_conn_timeout }}
    idle_conn_timeout: {{ . }}
    {{- end}}
    {{- with .response_header_timeout }}
    response_header_timeout: {{ . }}
    {{- end}}
    {{- with .insecure_skip_verify }}
    insecure_skip_verify: {{ . }}
    {{- end}}
    {{- with .ca_file}}
    ca_file: {{ . }}
    {{- end}}
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
  container_name: {{ $.Values.logstore.storage.bucketNames.chunks }}
  use_managed_identity: {{ .useManagedIdentity }}
  {{- with .userAssignedId }}
  user_assigned_id: {{ . }}
  {{- end }}
  {{- with .requestTimeout }}
  request_timeout: {{ . }}
  {{- end }}
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
  bucketnames: {{ $.Values.logstore.storage.bucketNames.ruler }}
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
  container_name: {{ $.Values.logstore.storage.bucketNames.ruler }}
  use_managed_identity: {{ .useManagedIdentity }}
  {{- with .userAssignedId }}
  user_assigned_id: {{ . }}
  {{- end }}
  {{- with .requestTimeout }}
  request_timeout: {{ . }}
  {{- end }}
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

{{/*
Return the appropriate apiVersion for ingress.
*/}}
{{- define "logstore.ingress.apiVersion" -}}
  {{- if and (.Capabilities.APIVersions.Has "networking.k8s.io/v1") (semverCompare ">= 1.19-0" .Capabilities.KubeVersion.Version) -}}
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
  {{- or (eq (include "logstore.ingress.isStable" .) "true") (and (eq (include "logstore.ingress.apiVersion" .) "networking.k8s.io/v1beta1") (semverCompare ">= 1.18-0" .Capabilities.KubeVersion.Version)) -}}
{{- end -}}

{{/*
Return if ingress supports pathType.
*/}}
{{- define "logstore.ingress.supportsPathType" -}}
  {{- or (eq (include "logstore.ingress.isStable" .) "true") (and (eq (include "logstore.ingress.apiVersion" .) "networking.k8s.io/v1beta1") (semverCompare ">= 1.18-0" .Capabilities.KubeVersion.Version)) -}}
{{- end -}}

{{/*
Generate list of ingress service paths based on deployment type
*/}}
{{- define "logstore.ingress.servicePaths" -}}
{{- if (eq (include "logstore.deployment.isScalable" .) "true") -}}
{{- include "logstore.ingress.scalableServicePaths" . }}
{{- else -}}
{{- include "logstore.ingress.singleBinaryServicePaths" . }}
{{- end -}}
{{- end -}}

{{/*
Ingress service paths for scalable deployment
*/}}
{{- define "logstore.ingress.scalableServicePaths" -}}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "svcName" "read" "paths" .Values.ingress.paths.read )}}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "svcName" "write" "paths" .Values.ingress.paths.write )}}
{{- end -}}

{{/*
Ingress service paths for single binary deployment
*/}}
{{- define "logstore.ingress.singleBinaryServicePaths" -}}
{{- include "logstore.ingress.servicePath" (dict "ctx" . "svcName" "singleBinary" "paths" .Values.ingress.paths.singleBinary )}}
{{- end -}}

{{/*
Ingress service path helper function
Params:
  ctx = . context
  svcName = service name without the "logstore.fullname" part (ie. read, write)
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
    {{- $serviceName := include "logstore.ingress.serviceName" (dict "ctx" $.ctx "svcName" $.svcName) }}
    service:
      name: {{ $serviceName }}
      port:
        number: 3100
    {{- else }}
    serviceName: {{ $serviceName }}
    servicePort: 3100
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Ingress service name helper function
Params:
  ctx = . context
  svcName = service name without the "logstore.fullname" part (ie. read, write)
*/}}
{{- define "logstore.ingress.serviceName" -}}
{{- if (eq .svcName "singleBinary") }}
{{- printf "%s" (include "logstore.fullname" .ctx) }}
{{- else }}
{{- printf "%s-%s" (include "logstore.fullname" .ctx) .svcName }}
{{- end -}}
{{- end -}}

{{/*
Create the service endpoint including port for MinIO.
*/}}
{{- define "logstore.minio" -}}
{{- if .Values.minio.enabled -}}
{{- printf "%s-%s.%s.svc:%s" .Release.Name "minio" .Release.Namespace (.Values.minio.service.port | toString) -}}
{{- end -}}
{{- end -}}

{{/* Return the appropriate apiVersion for PodDisruptionBudget. */}}
{{- define "logstore.podDisruptionBudget.apiVersion" -}}
  {{- if and (.Capabilities.APIVersions.Has "policy/v1") (semverCompare ">= 1.21-0" .Capabilities.KubeVersion.Version) -}}
    {{- print "policy/v1" -}}
  {{- else -}}
    {{- print "policy/v1beta1" -}}
  {{- end -}}
{{- end -}}

{{/* Determine if deployment is using object storage */}}
{{- define "logstore.isUsingObjectStorage" -}}
{{- or (eq .Values.logstore.storage.type "gcs") (eq .Values.logstore.storage.type "s3") (eq .Values.logstore.storage.type "azure") -}}
{{- end -}}

{{/* Configure the correct name for the memberlist service */}}
{{- define "logstore.memberlist" -}}
{{ include "logstore.name" . }}-memberlist
{{- end -}}

{{/* Determine the public host for the Logstore cluster */}}
{{- define "logstore.host" -}}
{{- $isSingleBinary := eq (include "logstore.deployment.isSingleBinary" .) "true" -}}
{{- $url := printf "%s.%s.svc.%s." (include "logstore.gatewayFullname" .) .Release.Namespace .Values.global.clusterDomain }}
{{- if and $isSingleBinary (not .Values.gateway.enabled)  }}
  {{- $url = printf "%s.%s.svc.%s.:3100" (include "logstore.singleBinaryFullname" .) .Release.Namespace .Values.global.clusterDomain }}
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

  client_max_body_size 4M;

  proxy_read_timeout    600; ## 6 minutes
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
  resolver {{ .Values.global.dnsService }}.{{ .Values.global.dnsNamespace }}.svc.{{ .Values.global.clusterDomain }}.;

  {{- with .Values.gateway.nginxConfig.httpSnippet }}
  {{- tpl . $ | nindent 2 }}
  {{- end }}

  server {
    listen             8080;

    {{- if .Values.gateway.basicAuth.enabled }}
    auth_basic           "Logstore";
    auth_basic_user_file /etc/nginx/secrets/.htpasswd;
    {{- end }}

    location = / {
      return 200 'OK';
      auth_basic off;
    }

    {{- $backendHost := include "logstore.backendFullname" .}}
    {{- $readHost := include "logstore.readFullname" .}}
    {{- $writeHost := include "logstore.writeFullname" .}}

    {{- if .Values.read.legacyReadTarget }}
    {{- $backendHost = include "logstore.readFullname" . }}
    {{- end }}

    {{- if gt (int .Values.singleBinary.replicas) 0 }}
    {{- $backendHost = include "logstore.singleBinaryFullname" . }}
    {{- $readHost = include "logstore.singleBinaryFullname" .}}
    {{- $writeHost = include "logstore.singleBinaryFullname" .}}
    {{- end }}

    {{- $writeUrl    := printf "http://%s.%s.svc.%s:3100" $writeHost   .Release.Namespace .Values.global.clusterDomain }}
    {{- $readUrl     := printf "http://%s.%s.svc.%s:3100" $readHost    .Release.Namespace .Values.global.clusterDomain }}
    {{- $backendUrl  := printf "http://%s.%s.svc.%s:3100" $backendHost .Release.Namespace .Values.global.clusterDomain }}

    {{- if .Values.gateway.nginxConfig.customWriteUrl }}
    {{- $writeUrl  = .Values.gateway.nginxConfig.customWriteUrl }}
    {{- end }}
    {{- if .Values.gateway.nginxConfig.customReadUrl }}
    {{- $readUrl = .Values.gateway.nginxConfig.customReadUrl }}
    {{- end }}
    {{- if .Values.gateway.nginxConfig.customBackendUrl }}
    {{- $backendUrl = .Values.gateway.nginxConfig.customBackendUrl }}
    {{- end }}

    location = /api/prom/push {
      proxy_pass       {{ $writeUrl }}$request_uri;
    }

    location = /api/prom/tail {
      proxy_pass       {{ $readUrl }}$request_uri;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection "upgrade";
    }

    location ~ /api/prom/.* {
      proxy_pass       {{ $readUrl }}$request_uri;
    }

    location ~ /prometheus/api/v1/alerts.* {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }
    location ~ /prometheus/api/v1/rules.* {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }
    location ~ /ruler/.* {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }

    location = /logstore/api/v1/push {
      proxy_pass       {{ $writeUrl }}$request_uri;
    }

    location = /logstore/api/v1/tail {
      proxy_pass       {{ $readUrl }}$request_uri;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection "upgrade";
    }

    location ~ /compactor/.* {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }

    location ~ /distributor/.* {
      proxy_pass       {{ $writeUrl }}$request_uri;
    }

    location ~ /ring {
      proxy_pass       {{ $writeUrl }}$request_uri;
    }

    location ~ /ingester/.* {
      proxy_pass       {{ $writeUrl }}$request_uri;
    }

    location ~ /store-gateway/.* {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }

    location ~ /query-scheduler/.* {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }
    location ~ /scheduler/.* {
      proxy_pass       {{ $backendUrl }}$request_uri;
    }

    location ~ /logstore/api/.* {
      proxy_pass       {{ $readUrl }}$request_uri;
    }

    location ~ /admin/api/.* {
      proxy_pass       {{ $writeUrl }}$request_uri;
    }

    {{- with .Values.gateway.nginxConfig.serverSnippet }}
    {{ . | nindent 4 }}
    {{- end }}
  }
}
{{- end }}

{{/* Configure enableServiceLinks in pod */}}
{{- define "logstore.enableServiceLinks" -}}
{{- if semverCompare ">=1.13-0" .Capabilities.KubeVersion.Version -}}
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
{{- $compactorAddress := include "logstore.backendFullname" . -}}
{{- if and $isSimpleScalable .Values.read.legacyReadTarget -}}
{{/* 2 target configuration */}}
{{- $compactorAddress = include "logstore.readFullname" . -}}
{{- else if (not $isSimpleScalable) -}}
{{/* single binary */}}
{{- $compactorAddress = include "logstore.singleBinaryFullname" . -}}
{{- end -}}
{{- printf "%s" $compactorAddress }}
{{- end }}

