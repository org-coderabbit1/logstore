package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagementStateType defines the type for CR management states.
//
// +kubebuilder:validation:Enum=Managed;Unmanaged
type ManagementStateType string

const (
	// ManagementStateManaged when the LogstoreStack custom resource should be
	// reconciled by the operator.
	ManagementStateManaged ManagementStateType = "Managed"

	// ManagementStateUnmanaged when the LogstoreStack custom resource should not be
	// reconciled by the operator.
	ManagementStateUnmanaged ManagementStateType = "Unmanaged"
)

// LogstoreStackSizeType declares the type for logstore cluster scale outs.
//
// +kubebuilder:validation:Enum="1x.extra-small";"1x.small";"1x.medium"
type LogstoreStackSizeType string

const (
	// SizeOneXExtraSmall defines the size of a single Logstore deployment
	// with minimal resource requirements and without HA support.
	//
	// This is ONLY for development, testing, or demos on limited single-node clusters.
	// There are NO performance guarantees.
	// LogstoreStack will use whatever resources are available,
	// and WILL NOT FUNCTION CORRECTLY if there is not enough memory or CPU.
	//
	// DO NOT USE THIS IN PRODUCTION!
	SizeOneXExtraSmall LogstoreStackSizeType = "1x.extra-small"

	// SizeOneXSmall defines the size of a single Logstore deployment
	// with small resources/limits requirements and HA support for all
	// Logstore components. This size is dedicated for setup **without** the
	// requirement for single replication factor and auto-compaction.
	//
	// FIXME: Add clear description of ingestion/query performance expectations.
	SizeOneXSmall LogstoreStackSizeType = "1x.small"

	// SizeOneXMedium defines the size of a single Logstore deployment
	// with small resources/limits requirements and HA support for all
	// Logstore components. This size is dedicated for setup **with** the
	// requirement for single replication factor and auto-compaction.
	//
	// FIXME: Add clear description of ingestion/query performance expectations.
	SizeOneXMedium LogstoreStackSizeType = "1x.medium"
)

// SubjectKind is a kind of LogstoreStack Gateway RBAC subject.
//
// +kubebuilder:validation:Enum=user;group
type SubjectKind string

const (
	// User represents a subject that is a user.
	User SubjectKind = "user"
	// Group represents a subject that is a group.
	Group SubjectKind = "group"
)

// Subject represents a subject that has been bound to a role.
type Subject struct {
	Name string      `json:"name"`
	Kind SubjectKind `json:"kind"`
}

// RoleBindingsSpec binds a set of roles to a set of subjects.
type RoleBindingsSpec struct {
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
	Roles    []string  `json:"roles"`
}

// PermissionType is a LogstoreStack Gateway RBAC permission.
//
// +kubebuilder:validation:Enum=read;write
type PermissionType string

const (
	// Write gives access to write data to a tenant.
	Write PermissionType = "write"
	// Read gives access to read data from a tenant.
	Read PermissionType = "read"
)

// RoleSpec describes a set of permissions to interact with a tenant.
type RoleSpec struct {
	Name        string           `json:"name"`
	Resources   []string         `json:"resources"`
	Tenants     []string         `json:"tenants"`
	Permissions []PermissionType `json:"permissions"`
}

// OPASpec defines the opa configuration spec for logstoreStack Gateway component.
type OPASpec struct {
	// URL defines the third-party endpoint for authorization.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OpenPolicyAgent URL"
	URL string `json:"url"`
}

// AuthorizationSpec defines the opa, role bindings and roles
// configuration per tenant for logstoreStack Gateway component.
type AuthorizationSpec struct {
	// OPA defines the spec for the third-party endpoint for tenant's authorization.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OPA Configuration"
	OPA *OPASpec `json:"opa"`
	// Roles defines a set of permissions to interact with a tenant.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Static Roles"
	Roles []RoleSpec `json:"roles"`
	// RoleBindings defines configuration to bind a set of roles to a set of subjects.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Static Role Bindings"
	RoleBindings []RoleBindingsSpec `json:"roleBindings"`
}

// TenantSecretSpec is a secret reference containing name only
// for a secret living in the same namespace as the LogstoreStack custom resource.
type TenantSecretSpec struct {
	// Name of a secret in the namespace configured for tenant secrets.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:io.kubernetes:Secret",displayName="Tenant Secret Name"
	Name string `json:"name"`
}

// OIDCSpec defines the oidc configuration spec for logstoreStack Gateway component.
type OIDCSpec struct {
	// Secret defines the spec for the clientID, clientSecret and issuerCAPath for tenant's authentication.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Tenant Secret"
	Secret *TenantSecretSpec `json:"secret"`
	// IssuerURL defines the URL for issuer.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Issuer URL"
	IssuerURL string `json:"issuerURL"`
	// RedirectURL defines the URL for redirect.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Redirect URL"
	RedirectURL string `json:"redirectURL,omitempty"`
	// Group claim field from ID Token
	//
	// +optional
	// +kubebuilder:validation:Optional
	GroupClaim string `json:"groupClaim,omitempty"`
	// User claim field from ID Token
	//
	// +optional
	// +kubebuilder:validation:Optional
	UsernameClaim string `json:"usernameClaim,omitempty"`
}

// AuthenticationSpec defines the oidc configuration per tenant for logstoreStack Gateway component.
type AuthenticationSpec struct {
	// TenantName defines the name of the tenant.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Tenant Name"
	TenantName string `json:"tenantName"`
	// TenantID defines the id of the tenant.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Tenant ID"
	TenantID string `json:"tenantId"`
	// OIDC defines the spec for the OIDC tenant's authentication.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OIDC Configuration"
	OIDC *OIDCSpec `json:"oidc"`
}

// ModeType is the authentication/authorization mode in which LogstoreStack Gateway will be configured.
//
// +kubebuilder:validation:Enum=static;dynamic;openshift-logging;openshift-network
type ModeType string

const (
	// Static mode asserts the Authorization Spec's Roles and RoleBindings
	// using an in-process OpenPolicyAgent Rego authorizer.
	Static ModeType = "static"
	// Dynamic mode delegates the authorization to a third-party OPA-compatible endpoint.
	Dynamic ModeType = "dynamic"
	// OpenshiftLogging mode provides fully automatic OpenShift in-cluster authentication and authorization support for application, infrastructure and audit logs.
	OpenshiftLogging ModeType = "openshift-logging"
	// OpenshiftNetwork mode provides fully automatic OpenShift in-cluster authentication and authorization support for network logs only.
	OpenshiftNetwork ModeType = "openshift-network"
)

// TenantsSpec defines the mode, authentication and authorization
// configuration of the logstoreStack gateway component.
type TenantsSpec struct {
	// Mode defines the mode in which logstorestack-gateway component will be configured.
	//
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=openshift-logging
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:static","urn:alm:descriptor:com.tectonic.ui:select:dynamic","urn:alm:descriptor:com.tectonic.ui:select:openshift-logging","urn:alm:descriptor:com.tectonic.ui:select:openshift-network"},displayName="Mode"
	Mode ModeType `json:"mode"`
	// Authentication defines the logstorestack-gateway component authentication configuration spec per tenant.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Authentication"
	Authentication []AuthenticationSpec `json:"authentication,omitempty"`
	// Authorization defines the logstorestack-gateway component authorization configuration spec per tenant.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Authorization"
	Authorization *AuthorizationSpec `json:"authorization,omitempty"`
}

// LogstoreComponentSpec defines the requirements to configure scheduling
// of each logstore component individually.
type LogstoreComponentSpec struct {
	// Replicas defines the number of replica pods of the component.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:hidden"
	Replicas int32 `json:"replicas,omitempty"`

	// NodeSelector defines the labels required by a node to schedule
	// the component onto it.
	//
	// +optional
	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Tolerations defines the tolerations required by a node to schedule
	// the component onto it.
	//
	// +optional
	// +kubebuilder:validation:Optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
}

// LogstoreTemplateSpec defines the template of all requirements to configure
// scheduling of all Logstore components to be deployed.
type LogstoreTemplateSpec struct {
	// Compactor defines the compaction component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Compactor pods"
	Compactor *LogstoreComponentSpec `json:"compactor,omitempty"`

	// Distributor defines the distributor component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Distributor pods"
	Distributor *LogstoreComponentSpec `json:"distributor,omitempty"`

	// Ingester defines the ingester component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Ingester pods"
	Ingester *LogstoreComponentSpec `json:"ingester,omitempty"`

	// Querier defines the querier component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Querier pods"
	Querier *LogstoreComponentSpec `json:"querier,omitempty"`

	// QueryFrontend defines the query frontend component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Query Frontend pods"
	QueryFrontend *LogstoreComponentSpec `json:"queryFrontend,omitempty"`

	// Gateway defines the logstorestack gateway component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Gateway pods"
	Gateway *LogstoreComponentSpec `json:"gateway,omitempty"`

	// IndexGateway defines the index gateway component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Index Gateway pods"
	IndexGateway *LogstoreComponentSpec `json:"indexGateway,omitempty"`

	// Ruler defines the ruler component spec.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Ruler pods"
	Ruler *LogstoreComponentSpec `json:"ruler,omitempty"`
}

// ClusterProxy is the Proxy configuration when the cluster is behind a Proxy.
type ClusterProxy struct {
	// HTTPProxy configures the HTTP_PROXY/http_proxy env variable.
	//
	// +optional
	// +kubebuilder:validation:optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="HTTPProxy"
	HTTPProxy string `json:"httpProxy,omitempty"`
	// HTTPSProxy configures the HTTPS_PROXY/https_proxy env variable.
	//
	// +optional
	// +kubebuilder:validation:optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="HTTPSProxy"
	HTTPSProxy string `json:"httpsProxy,omitempty"`
	// NoProxy configures the NO_PROXY/no_proxy env variable.
	//
	// +optional
	// +kubebuilder:validation:optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="NoProxy"
	NoProxy string `json:"noProxy,omitempty"`
}

// HashRingType defines the type of hash ring which can be used with the Logstore cluster.
//
// +kubebuilder:validation:Enum=memberlist
type HashRingType string

const (
	// HashRingMemberList when using memberlist for the distributed hash ring.
	HashRingMemberList HashRingType = "memberlist"
)

// InstanceAddrType defines the type of pod network to use for advertising IPs to the ring.
//
// +kubebuilder:validation:Enum=default;podIP
type InstanceAddrType string

const (
	// InstanceAddrDefault when using the first from any private network interfaces (RFC 1918 and RFC 6598).
	InstanceAddrDefault InstanceAddrType = "default"
	// InstanceAddrPodIP when using the public pod IP from the cluster's pod network.
	InstanceAddrPodIP InstanceAddrType = "podIP"
)

// MemberListSpec defines the configuration for the memberlist based hash ring.
type MemberListSpec struct {
	// InstanceAddrType defines the type of address to use to advertise to the ring.
	// Defaults to the first address from any private network interfaces of the current pod.
	// Alternatively the public pod IP can be used in case private networks (RFC 1918 and RFC 6598)
	// are not available.
	//
	// +optional
	// +kubebuilder:validation:optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:default","urn:alm:descriptor:com.tectonic.ui:select:podIP"},displayName="Instance Address"
	InstanceAddrType InstanceAddrType `json:"instanceAddrType,omitempty"`
}

// HashRingSpec defines the hash ring configuration
type HashRingSpec struct {
	// Type of hash ring implementation that should be used
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:memberlist"},displayName="Type"
	// +kubebuilder:default:=memberlist
	Type HashRingType `json:"type"`

	// MemberList configuration spec
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Memberlist Config"
	MemberList *MemberListSpec `json:"memberlist,omitempty"`
}

// ObjectStorageTLSSpec is the TLS configuration for reaching the object storage endpoint.
type ObjectStorageTLSSpec struct {
	// Key is the data key of a ConfigMap containing a CA certificate.
	// It needs to be in the same namespace as the LogstoreStack custom resource.
	// If empty, it defaults to "service-ca.crt".
	//
	// +optional
	// +kubebuilder:validation:optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="CA ConfigMap Key"
	CAKey string `json:"caKey,omitempty"`
	// CA is the name of a ConfigMap containing a CA certificate.
	// It needs to be in the same namespace as the LogstoreStack custom resource.
	//
	// +required
	// +kubebuilder:validation:required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:io.kubernetes:ConfigMap",displayName="CA ConfigMap Name"
	CA string `json:"caName"`
}

// ObjectStorageSecretType defines the type of storage which can be used with the Logstore cluster.
//
// +kubebuilder:validation:Enum=azure;gcs;s3;swift
type ObjectStorageSecretType string

const (
	// ObjectStorageSecretAzure when using Azure for Logstore storage
	ObjectStorageSecretAzure ObjectStorageSecretType = "azure"

	// ObjectStorageSecretGCS when using GCS for Logstore storage
	ObjectStorageSecretGCS ObjectStorageSecretType = "gcs"

	// ObjectStorageSecretS3 when using S3 for Logstore storage
	ObjectStorageSecretS3 ObjectStorageSecretType = "s3"

	// ObjectStorageSecretSwift when using Swift for Logstore storage
	ObjectStorageSecretSwift ObjectStorageSecretType = "swift"
)

// ObjectStorageSecretSpec is a secret reference containing name only, no namespace.
type ObjectStorageSecretSpec struct {
	// Type of object storage that should be used
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:azure","urn:alm:descriptor:com.tectonic.ui:select:gcs","urn:alm:descriptor:com.tectonic.ui:select:s3","urn:alm:descriptor:com.tectonic.ui:select:swift"},displayName="Object Storage Secret Type"
	Type ObjectStorageSecretType `json:"type"`

	// Name of a secret in the namespace configured for object storage secrets.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:io.kubernetes:Secret",displayName="Object Storage Secret Name"
	Name string `json:"name"`
}

// ObjectStorageSchemaVersion defines the storage schema version which will be
// used with the Logstore cluster.
//
// +kubebuilder:validation:Enum=v11;v12
type ObjectStorageSchemaVersion string

const (
	// ObjectStorageSchemaV11 when using v11 for the storage schema
	ObjectStorageSchemaV11 ObjectStorageSchemaVersion = "v11"

	// ObjectStorageSchemaV12 when using v12 for the storage schema
	ObjectStorageSchemaV12 ObjectStorageSchemaVersion = "v12"
)

// ObjectStorageSchema defines the requirements needed to configure a new
// storage schema.
type ObjectStorageSchema struct {
	// Version for writing and reading logs.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:v11","urn:alm:descriptor:com.tectonic.ui:select:v12"},displayName="Version"
	Version ObjectStorageSchemaVersion `json:"version"`

	// EffectiveDate is the date in UTC that the schema will be applied on.
	// To ensure readibility of logs, this date should be before the current
	// date in UTC.
	//
	// +required
	// +kubebuilder:validation:Required
	EffectiveDate StorageSchemaEffectiveDate `json:"effectiveDate"`
}

// ObjectStorageSpec defines the requirements to access the object
// storage bucket to persist logs by the ingester component.
type ObjectStorageSpec struct {
	// Schemas for reading and writing logs.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:default:={{version:v11,effectiveDate:"2020-10-11"}}
	Schemas []ObjectStorageSchema `json:"schemas"`

	// Secret for object storage authentication.
	// Name of a secret in the same namespace as the LogstoreStack custom resource.
	//
	// +required
	// +kubebuilder:validation:Required
	Secret ObjectStorageSecretSpec `json:"secret"`

	// TLS configuration for reaching the object storage endpoint.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="TLS Config"
	TLS *ObjectStorageTLSSpec `json:"tls,omitempty"`
}

// QueryLimitSpec defines the limits applies at the query path.
type QueryLimitSpec struct {
	// MaxEntriesLimitsPerQuery defines the maximum number of log entries
	// that will be returned for a query.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Entries Limit per Query"
	MaxEntriesLimitPerQuery int32 `json:"maxEntriesLimitPerQuery,omitempty"`

	// MaxChunksPerQuery defines the maximum number of chunks
	// that can be fetched by a single query.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Chunk per Query"
	MaxChunksPerQuery int32 `json:"maxChunksPerQuery,omitempty"`

	// MaxQuerySeries defines the the maximum of unique series
	// that is returned by a metric query.
	//
	// + optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Query Series"
	MaxQuerySeries int32 `json:"maxQuerySeries,omitempty"`

	// Timeout when querying ingesters or storage during the execution of a query request.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="1m"
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Query Timeout"
	QueryTimeout string `json:"queryTimeout,omitempty"`
}

// IngestionLimitSpec defines the limits applied at the ingestion path.
type IngestionLimitSpec struct {
	// IngestionRate defines the sample size per second. Units MB.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Ingestion Rate (in MB)"
	IngestionRate int32 `json:"ingestionRate,omitempty"`

	// IngestionBurstSize defines the local rate-limited sample size per
	// distributor replica. It should be set to the set at least to the
	// maximum logs size expected in a single push request.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Ingestion Burst Size (in MB)"
	IngestionBurstSize int32 `json:"ingestionBurstSize,omitempty"`

	// MaxLabelNameLength defines the maximum number of characters allowed
	// for label keys in log streams.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Label Name Length"
	MaxLabelNameLength int32 `json:"maxLabelNameLength,omitempty"`

	// MaxLabelValueLength defines the maximum number of characters allowed
	// for label values in log streams.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Label Value Length"
	MaxLabelValueLength int32 `json:"maxLabelValueLength,omitempty"`

	// MaxLabelNamesPerSeries defines the maximum number of label names per series
	// in each log stream.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Labels Names per Series"
	MaxLabelNamesPerSeries int32 `json:"maxLabelNamesPerSeries,omitempty"`

	// MaxGlobalStreamsPerTenant defines the maximum number of active streams
	// per tenant, across the cluster.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Global Streams per  Tenant"
	MaxGlobalStreamsPerTenant int32 `json:"maxGlobalStreamsPerTenant,omitempty"`

	// MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Max Line Size"
	MaxLineSize int32 `json:"maxLineSize,omitempty"`
}

// RetentionStreamSpec defines a log stream with separate retention time.
type RetentionStreamSpec struct {
	// Days contains the number of days logs are kept.
	//
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum:=1
	Days uint `json:"days"`

	// Priority defines the priority of this selector compared to other retention rules.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=1
	Priority uint32 `json:"priority,omitempty"`

	// Selector contains the LogQL query used to define the log stream.
	//
	// +required
	// +kubebuilder:validation:Required
	Selector string `json:"selector"`
}

// RetentionLimitSpec controls how long logs will be kept in storage.
type RetentionLimitSpec struct {
	// Days contains the number of days logs are kept.
	//
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum:=1
	Days uint `json:"days"`

	// Stream defines the log stream.
	//
	// +optional
	// +kubebuilder:validation:Optional
	Streams []*RetentionStreamSpec `json:"streams,omitempty"`
}

// LimitsTemplateSpec defines the limits  applied at ingestion or query path.
type LimitsTemplateSpec struct {
	// IngestionLimits defines the limits applied on ingested log streams.
	//
	// +optional
	// +kubebuilder:validation:Optional
	IngestionLimits *IngestionLimitSpec `json:"ingestion,omitempty"`

	// QueryLimits defines the limit applied on querying log streams.
	//
	// +optional
	// +kubebuilder:validation:Optional
	QueryLimits *QueryLimitSpec `json:"queries,omitempty"`

	// Retention defines how long logs are kept in storage.
	//
	// +optional
	// +kubebuilder:validation:Optional
	Retention *RetentionLimitSpec `json:"retention,omitempty"`
}

// LimitsSpec defines the spec for limits applied at ingestion or query
// path across the cluster or per tenant.
type LimitsSpec struct {
	// Global defines the limits applied globally across the cluster.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Global Limits"
	Global *LimitsTemplateSpec `json:"global,omitempty"`

	// Tenants defines the limits applied per tenant.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Limits per Tenant"
	Tenants map[string]LimitsTemplateSpec `json:"tenants,omitempty"`
}

// RulesSpec defines the spec for the ruler component.
type RulesSpec struct {
	// Enabled defines a flag to enable/disable the ruler component
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch",displayName="Enable"
	Enabled bool `json:"enabled"`

	// A selector to select which LogstoreRules to mount for loading alerting/recording
	// rules from.
	//
	// +optional
	// +kubebuilder:validation:optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Selector"
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// Namespaces to be selected for PrometheusRules discovery. If unspecified, only
	// the same namespace as the LogstoreStack object is in is used.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace Selector"
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
}

// LogstoreStackSpec defines the desired state of LogstoreStack
type LogstoreStackSpec struct {
	// ManagementState defines if the CR should be managed by the operator or not.
	// Default is managed.
	//
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=Managed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Managed","urn:alm:descriptor:com.tectonic.ui:select:Unmanaged"},displayName="Management State"
	ManagementState ManagementStateType `json:"managementState,omitempty"`

	// Size defines one of the support Logstore deployment scale out sizes.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:1x.extra-small","urn:alm:descriptor:com.tectonic.ui:select:1x.small","urn:alm:descriptor:com.tectonic.ui:select:1x.medium"},displayName="LogstoreStack Size"
	Size LogstoreStackSizeType `json:"size"`

	// HashRing defines the spec for the distributed hash ring configuration.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:advanced",displayName="Hash Ring"
	HashRing *HashRingSpec `json:"hashRing,omitempty"`

	// Storage defines the spec for the object storage endpoint to store logs.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Object Storage"
	Storage ObjectStorageSpec `json:"storage"`

	// Storage class name defines the storage class for ingester/querier PVCs.
	//
	// +required
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:io.kubernetes:StorageClass",displayName="Storage Class Name"
	StorageClassName string `json:"storageClassName"`

	// Proxy defines the spec for the object proxy to configure cluster proxy information.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Cluster Proxy"
	Proxy *ClusterProxy `json:"proxy,omitempty"`

	// ReplicationFactor defines the policy for log stream replication.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum:=1
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number",displayName="Replication Factor"
	ReplicationFactor int32 `json:"replicationFactor,omitempty"`

	// Rules defines the spec for the ruler component
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:advanced",displayName="Rules"
	Rules *RulesSpec `json:"rules,omitempty"`

	// Limits defines the limits to be applied to log stream processing.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:advanced",displayName="Rate Limiting"
	Limits *LimitsSpec `json:"limits,omitempty"`

	// Template defines the resource/limits/tolerations/nodeselectors per component
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:advanced",displayName="Node Placement"
	Template *LogstoreTemplateSpec `json:"template,omitempty"`

	// Tenants defines the per-tenant authentication and authorization spec for the logstorestack-gateway component.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Tenants Configuration"
	Tenants *TenantsSpec `json:"tenants,omitempty"`
}

// LogstoreStackConditionType deifnes the type of condition types of a Logstore deployment.
type LogstoreStackConditionType string

const (
	// ConditionReady defines the condition that all components in the Logstore deployment are ready.
	ConditionReady LogstoreStackConditionType = "Ready"

	// ConditionPending defines the condition that some or all components are in pending state.
	ConditionPending LogstoreStackConditionType = "Pending"

	// ConditionFailed defines the condition that components in the Logstore deployment failed to roll out.
	ConditionFailed LogstoreStackConditionType = "Failed"

	// ConditionDegraded defines the condition that some or all components in the Logstore deployment
	// are degraded or the cluster cannot connect to object storage.
	ConditionDegraded LogstoreStackConditionType = "Degraded"
)

// LogstoreStackConditionReason defines the type for valid reasons of a Logstore deployment conditions.
type LogstoreStackConditionReason string

const (
	// ReasonFailedComponents when all/some LogstoreStack components fail to roll out.
	ReasonFailedComponents LogstoreStackConditionReason = "FailedComponents"
	// ReasonPendingComponents when all/some LogstoreStack components pending dependencies
	ReasonPendingComponents LogstoreStackConditionReason = "PendingComponents"
	// ReasonReadyComponents when all LogstoreStack components are ready to serve traffic.
	ReasonReadyComponents LogstoreStackConditionReason = "ReadyComponents"
	// ReasonMissingObjectStorageSecret when the required secret to store logs to object
	// storage is missing.
	ReasonMissingObjectStorageSecret LogstoreStackConditionReason = "MissingObjectStorageSecret"
	// ReasonInvalidObjectStorageSecret when the format of the secret is invalid.
	ReasonInvalidObjectStorageSecret LogstoreStackConditionReason = "InvalidObjectStorageSecret"
	// ReasonInvalidObjectStorageSchema when the spec contains an invalid schema(s).
	ReasonInvalidObjectStorageSchema LogstoreStackConditionReason = "InvalidObjectStorageSchema"
	// ReasonMissingObjectStorageCAConfigMap when the required configmap to verify object storage
	// certificates is missing.
	ReasonMissingObjectStorageCAConfigMap LogstoreStackConditionReason = "MissingObjectStorageCAConfigMap"
	// ReasonInvalidObjectStorageCAConfigMap when the format of the CA configmap is invalid.
	ReasonInvalidObjectStorageCAConfigMap LogstoreStackConditionReason = "InvalidObjectStorageCAConfigMap"
	// ReasonMissingRulerSecret when the required secret to authorization remote write connections
	// for the ruler is missing.
	ReasonMissingRulerSecret LogstoreStackConditionReason = "MissingRulerSecret"
	// ReasonInvalidRulerSecret when the format of the ruler remote write authorization secret is invalid.
	ReasonInvalidRulerSecret LogstoreStackConditionReason = "InvalidRulerSecret"
	// ReasonInvalidReplicationConfiguration when the configurated replication factor is not valid
	// with the select cluster size.
	ReasonInvalidReplicationConfiguration LogstoreStackConditionReason = "InvalidReplicationConfiguration"
	// ReasonMissingGatewayTenantSecret when the required tenant secret
	// for authentication is missing.
	ReasonMissingGatewayTenantSecret LogstoreStackConditionReason = "MissingGatewayTenantSecret"
	// ReasonInvalidGatewayTenantSecret when the format of the secret is invalid.
	ReasonInvalidGatewayTenantSecret LogstoreStackConditionReason = "InvalidGatewayTenantSecret"
	// ReasonInvalidTenantsConfiguration when the tenant configuration provided is invalid.
	ReasonInvalidTenantsConfiguration LogstoreStackConditionReason = "InvalidTenantsConfiguration"
	// ReasonMissingGatewayOpenShiftBaseDomain when the reconciler cannot lookup the OpenShift DNS base domain.
	ReasonMissingGatewayOpenShiftBaseDomain LogstoreStackConditionReason = "MissingGatewayOpenShiftBaseDomain"
	// ReasonFailedCertificateRotation when the reconciler cannot rotate any of the required TLS certificates.
	ReasonFailedCertificateRotation LogstoreStackConditionReason = "FailedCertificateRotation"
)

// PodStatusMap defines the type for mapping pod status to pod name.
type PodStatusMap map[corev1.PodPhase][]string

// LogstoreStackComponentStatus defines the map of per pod status per LogstoreStack component.
// Each component is represented by a separate map of v1.Phase to a list of pods.
type LogstoreStackComponentStatus struct {
	// Compactor is a map to the pod status of the compactor pod.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="Compactor",order=5
	Compactor PodStatusMap `json:"compactor,omitempty"`

	// Distributor is a map to the per pod status of the distributor deployment
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="Distributor",order=1
	Distributor PodStatusMap `json:"distributor,omitempty"`

	// IndexGateway is a map to the per pod status of the index gateway statefulset
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="IndexGateway",order=6
	IndexGateway PodStatusMap `json:"indexGateway,omitempty"`

	// Ingester is a map to the per pod status of the ingester statefulset
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="Ingester",order=2
	Ingester PodStatusMap `json:"ingester,omitempty"`

	// Querier is a map to the per pod status of the querier deployment
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="Querier",order=3
	Querier PodStatusMap `json:"querier,omitempty"`

	// QueryFrontend is a map to the per pod status of the query frontend deployment
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="Query Frontend",order=4
	QueryFrontend PodStatusMap `json:"queryFrontend,omitempty"`

	// Gateway is a map to the per pod status of the logstorestack gateway deployment.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="Gateway",order=5
	Gateway PodStatusMap `json:"gateway,omitempty"`

	// Ruler is a map to the per pod status of the logstorestack ruler statefulset.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:podStatuses",displayName="Ruler",order=6
	Ruler PodStatusMap `json:"ruler,omitempty"`
}

// LogstoreStackStorageStatus defines the observed state of
// the Logstore storage configuration.
type LogstoreStackStorageStatus struct {
	// Schemas is a list of schemas which have been applied
	// to the LogstoreStack.
	//
	// +optional
	// +kubebuilder:validation:Optional
	Schemas []ObjectStorageSchema `json:"schemas,omitempty"`
}

// LogstoreStackStatus defines the observed state of LogstoreStack
type LogstoreStackStatus struct {
	// Components provides summary of all Logstore pod status grouped
	// per component.
	//
	// +optional
	// +kubebuilder:validation:Optional
	Components LogstoreStackComponentStatus `json:"components,omitempty"`

	// Storage provides summary of all changes that have occurred
	// to the storage configuration.
	//
	// +optional
	// +kubebuilder:validation:Optional
	Storage LogstoreStackStorageStatus `json:"storage,omitempty"`

	// Conditions of the Logstore deployment health.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:io.kubernetes.conditions"
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:categories=logging
// +kubebuilder:webhook:path=/validate-logstore-acme-com-v1-logstorestack,mutating=false,failurePolicy=fail,sideEffects=None,groups=logstore.acme.com,resources=logstorestacks,verbs=create;update,versions=v1,name=vlogstorestack.logstore.acme.com,admissionReviewVersions=v1

// LogstoreStack is the Schema for the logstorestacks API
//
// +operator-sdk:csv:customresourcedefinitions:displayName="LogstoreStack",resources={{Deployment,v1},{StatefulSet,v1},{ConfigMap,v1},{Ingress,v1},{Service,v1},{ServiceAccount,v1},{PersistentVolumeClaims,v1},{Route,v1},{ServiceMonitor,v1}}
type LogstoreStack struct {
	// LogstoreStack CR spec field.
	Spec LogstoreStackSpec `json:"spec,omitempty"`
	// LogstoreStack CR spec Status.
	Status            LogstoreStackStatus `json:"status,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	metav1.TypeMeta   `json:",inline"`
}

// +kubebuilder:object:root=true

// LogstoreStackList contains a list of LogstoreStack
type LogstoreStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogstoreStack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogstoreStack{}, &LogstoreStackList{})
}

// Hub declares the v1.LogstoreStack as the hub CRD version.
func (*LogstoreStack) Hub() {}
