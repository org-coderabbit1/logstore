package internal

import (
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
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

// ResourceRequirements sets CPU, Memory, and PVC requirements for a component
type ResourceRequirements struct {
	Limits   corev1.ResourceList
	Requests corev1.ResourceList
	PVCSize  resource.Quantity
}

// ResourceRequirementsTable defines the default resource requests and limits for each size
var ResourceRequirementsTable = map[logstorev1.LogstoreStackSizeType]ComponentResources{
	logstorev1.SizeOneXExtraSmall: {
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

// StackSizeTable defines the default configurations for each size
var StackSizeTable = map[logstorev1.LogstoreStackSizeType]logstorev1.LogstoreStackSpec{
	logstorev1.SizeOneXExtraSmall: {
		Size:              logstorev1.SizeOneXExtraSmall,
		ReplicationFactor: 1,
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Defaults from Logstore docs
					IngestionRate:          4,
					IngestionBurstSize:     6,
					MaxLabelNameLength:     1024,
					MaxLabelValueLength:    2048,
					MaxLabelNamesPerSeries: 30,
					MaxLineSize:            256000,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
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

	logstorev1.SizeOneXSmall: {
		Size:              logstorev1.SizeOneXSmall,
		ReplicationFactor: 2,
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Custom for 1x.small
					IngestionRate:             15,
					IngestionBurstSize:        20,
					MaxGlobalStreamsPerTenant: 10000,
					// Defaults from Logstore docs
					MaxLabelNameLength:     1024,
					MaxLabelValueLength:    2048,
					MaxLabelNamesPerSeries: 30,
					MaxLineSize:            256000,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
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
		Size:              logstorev1.SizeOneXMedium,
		ReplicationFactor: 3,
		Limits: &logstorev1.LimitsSpec{
			Global: &logstorev1.LimitsTemplateSpec{
				IngestionLimits: &logstorev1.IngestionLimitSpec{
					// Custom for 1x.medium
					IngestionRate:             50,
					IngestionBurstSize:        20,
					MaxGlobalStreamsPerTenant: 25000,
					// Defaults from Logstore docs
					MaxLabelNameLength:     1024,
					MaxLabelValueLength:    2048,
					MaxLabelNamesPerSeries: 30,
					MaxLineSize:            256000,
				},
				QueryLimits: &logstorev1.QueryLimitSpec{
					// Defaults from Logstore docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
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
