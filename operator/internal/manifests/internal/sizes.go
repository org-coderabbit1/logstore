package internal

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

// ComponentResources is a map of component->requests/limits
type ComponentResources struct {
	IndexGateway ResourceRequirements
	Ingester     ResourceRequirements
	Compactor    ResourceRequirements
	Ruler        ResourceRequirements
	WALStorage   ResourceRequirements
	// these two don't need a PVCSize
	Querier       corev1.ResourceRequirements
	Distributor   corev1.ResourceRequirements
	QueryFrontend corev1.ResourceRequirements
	Gateway       corev1.ResourceRequirements
}

func (c ComponentResources) DeepCopy() ComponentResources {
	return ComponentResources{
		IndexGateway:  *c.IndexGateway.DeepCopy(),
		Ingester:      *c.Ingester.DeepCopy(),
		Compactor:     *c.Compactor.DeepCopy(),
		Ruler:         *c.Ruler.DeepCopy(),
		WALStorage:    *c.WALStorage.DeepCopy(),
		Querier:       *c.Querier.DeepCopy(),
		Distributor:   *c.Distributor.DeepCopy(),
		QueryFrontend: *c.QueryFrontend.DeepCopy(),
		Gateway:       *c.Gateway.DeepCopy(),
	}
}

// ResourceRequirements sets CPU, Memory, and PVC requirements for a component
type ResourceRequirements struct {
	Limits          corev1.ResourceList
	Requests        corev1.ResourceList
	PVCSize         resource.Quantity
	PDBMinAvailable int
}

func (r *ResourceRequirements) DeepCopy() *ResourceRequirements {
	return &ResourceRequirements{
		Limits:          r.Limits.DeepCopy(),
		Requests:        r.Requests.DeepCopy(),
		PVCSize:         r.PVCSize.DeepCopy(),
		PDBMinAvailable: r.PDBMinAvailable,
	}
}

// resourceRequirementsTable defines the default resource requests and limits for each size
var resourceRequirementsTable = map[logstorev1.LogstoreStackSizeType]ComponentResources{
	logstorev1.SizeOneXDemo: {
		Ruler: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
		IndexGateway: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
	},
	logstorev1.SizeOneXPico: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("750m"),
				corev1.ResourceMemory: resource.MustParse("1.5Gi"),
			},
		},
		Ruler: ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
			PVCSize: resource.MustParse("10Gi"),
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("3Gi"),
			},
			PDBMinAvailable: 2,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		IndexGateway: ResourceRequirements{
			PVCSize: resource.MustParse("50Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("150m"),
				corev1.ResourceMemory: resource.MustParse("250Mi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
	logstorev1.SizeOneXExtraSmall: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1.5"),
				corev1.ResourceMemory: resource.MustParse("3Gi"),
			},
		},
		Ruler: ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
			PVCSize: resource.MustParse("10Gi"),
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("8Gi"),
			},
			PDBMinAvailable: 1,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		IndexGateway: ResourceRequirements{
			PVCSize: resource.MustParse("50Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
	logstorev1.SizeOneXSmall: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
		Ruler: ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("8Gi"),
			},
			PVCSize: resource.MustParse("10Gi"),
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("20Gi"),
			},
			PDBMinAvailable: 1,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("2.5Gi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		IndexGateway: ResourceRequirements{
			PVCSize: resource.MustParse("50Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
	logstorev1.SizeOneXMedium: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("6"),
				corev1.ResourceMemory: resource.MustParse("10Gi"),
			},
		},
		Ruler: ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("8"),
				corev1.ResourceMemory: resource.MustParse("16Gi"),
			},
			PVCSize: resource.MustParse("10Gi"),
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("6"),
				corev1.ResourceMemory: resource.MustParse("30Gi"),
			},
			PDBMinAvailable: 2,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("2.5Gi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		IndexGateway: ResourceRequirements{
			PVCSize: resource.MustParse("50Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
}

// ResourceRequirementsForSize returns the resource configuration for a specific LogstoreStack size.
func ResourceRequirementsForSize(size logstorev1.LogstoreStackSizeType, useRequestsAsLimits bool) ComponentResources {
	resources := resourceRequirementsTable[size].DeepCopy()
	if useRequestsAsLimits {
		resources.IndexGateway.Limits = resources.IndexGateway.Requests.DeepCopy()
		resources.Ingester.Limits = resources.Ingester.Requests.DeepCopy()
		resources.Compactor.Limits = resources.Compactor.Requests.DeepCopy()
		resources.Ruler.Limits = resources.Ruler.Requests.DeepCopy()
		resources.WALStorage.Limits = resources.WALStorage.Requests.DeepCopy()
		resources.Querier.Limits = resources.Querier.Requests.DeepCopy()
		resources.Distributor.Limits = resources.Distributor.Requests.DeepCopy()
		resources.QueryFrontend.Limits = resources.QueryFrontend.Requests.DeepCopy()
		resources.Gateway.Limits = resources.Gateway.Requests.DeepCopy()
	}
	return resources
}

// StackSizeTable defines the default configurations for each size
var StackSizeTable = map[logstorev1.LogstoreStackSizeType]logstorev1.LogstoreStackSpec{
	logstorev1.SizeOneXDemo: {
		Size: logstorev1.SizeOneXDemo,
		Replication: &logstorev1.ReplicationSpec{
			Factor: 1,
		},
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Defaults from Logstore docs
					IngestionRate:           4,
					IngestionBurstSize:      6,
					MaxLabelNameLength:      1024,
					MaxLabelValueLength:     2048,
					MaxLabelNamesPerSeries:  30,
					MaxLineSize:             256000,
					PerStreamDesiredRate:    3,
					PerStreamRateLimit:      5,
					PerStreamRateLimitBurst: 15,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &logstorev1.LogstoreTemplateSpec{
			Compactor: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Distributor: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Ingester: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Querier: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			QueryFrontend: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Gateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Ruler: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
		},
	},

	logstorev1.SizeOneXPico: {
		Size: logstorev1.SizeOneXPico,
		Replication: &logstorev1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Defaults from Logstore docs
					IngestionRate:             4,
					IngestionBurstSize:        6,
					MaxGlobalStreamsPerTenant: 10000,
					MaxLabelNameLength:        1024,
					MaxLabelValueLength:       2048,
					MaxLabelNamesPerSeries:    30,
					MaxLineSize:               256000,
					PerStreamDesiredRate:      3,
					PerStreamRateLimit:        5,
					PerStreamRateLimitBurst:   15,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &logstorev1.LogstoreTemplateSpec{
			Compactor: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Distributor: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ingester: &logstorev1.LogstoreComponentSpec{
				Replicas: 3,
			},
			Querier: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			QueryFrontend: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Gateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ruler: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
		},
	},

	logstorev1.SizeOneXExtraSmall: {
		Size: logstorev1.SizeOneXExtraSmall,
		Replication: &logstorev1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Defaults from Logstore docs
					IngestionRate:             4,
					IngestionBurstSize:        6,
					MaxGlobalStreamsPerTenant: 10000,
					MaxLabelNameLength:        1024,
					MaxLabelValueLength:       2048,
					MaxLabelNamesPerSeries:    30,
					MaxLineSize:               256000,
					PerStreamDesiredRate:      3,
					PerStreamRateLimit:        5,
					PerStreamRateLimitBurst:   15,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &logstorev1.LogstoreTemplateSpec{
			Compactor: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Distributor: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ingester: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Querier: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			QueryFrontend: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Gateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ruler: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
		},
	},

	logstorev1.SizeOneXSmall: {
		Size: logstorev1.SizeOneXSmall,
		Replication: &logstorev1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Custom for 1x.small
					IngestionRate:             15,
					IngestionBurstSize:        20,
					MaxGlobalStreamsPerTenant: 10000,
					// Defaults from Logstore docs
					MaxLabelNameLength:      1024,
					MaxLabelValueLength:     2048,
					MaxLabelNamesPerSeries:  30,
					MaxLineSize:             256000,
					PerStreamDesiredRate:    3,
					PerStreamRateLimit:      5,
					PerStreamRateLimitBurst: 15,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &logstorev1.LogstoreTemplateSpec{
			Compactor: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Distributor: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ingester: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Querier: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			QueryFrontend: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Gateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ruler: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
		},
	},

	logstorev1.SizeOneXMedium: {
		Size: logstorev1.SizeOneXMedium,
		Replication: &logstorev1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Custom for 1x.medium
					IngestionRate:             50,
					IngestionBurstSize:        20,
					MaxGlobalStreamsPerTenant: 25000,
					// Defaults from Logstore docs
					MaxLabelNameLength:      1024,
					MaxLabelValueLength:     2048,
					MaxLabelNamesPerSeries:  30,
					MaxLineSize:             256000,
					PerStreamDesiredRate:    3,
					PerStreamRateLimit:      5,
					PerStreamRateLimitBurst: 15,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &logstorev1.LogstoreTemplateSpec{
			Compactor: &logstorev1.LogstoreComponentSpec{
				Replicas: 1,
			},
			Distributor: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ingester: &logstorev1.LogstoreComponentSpec{
				Replicas: 3,
			},
			Querier: &logstorev1.LogstoreComponentSpec{
				Replicas: 3,
			},
			QueryFrontend: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Gateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
			Ruler: &logstorev1.LogstoreComponentSpec{
				Replicas: 2,
			},
		},
	},
}
