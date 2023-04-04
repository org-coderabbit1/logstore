package manifests

import (
	"testing"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
)

func TestConfigureHashRingEnv_UseDefaults_NoHashRingSpec(t *testing.T) {
	opt := Options{
		Name:      "test",
		Namespace: "test",
		Image:     "test",
		Stack: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
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
					Replicas: 1,
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	}

	wantEnvVar := v1.EnvVar{
		ValueFrom: &v1.EnvVarSource{
			FieldRef: &v1.ObjectFieldSelector{
				APIVersion: "v1",
				FieldPath:  "status.podIP",
			},
		},
	}

	for _, cs := range logstoreContainers(t, opt) {
		for _, c := range cs {
			require.NotContains(t, c.Env, wantEnvVar, "contains envVar %s for: %s", gossipInstanceAddrEnvVarName, c.Name)
		}
	}
}

func TestConfigureHashRingEnv_UseDefaults_WithCustomHashRingSpec(t *testing.T) {
	opt := Options{
		Name:      "test",
		Namespace: "test",
		Image:     "test",
		Stack: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			HashRing: &logstorev1.HashRingSpec{
				Type: logstorev1.HashRingMemberList,
				MemberList: &logstorev1.MemberListSpec{
					InstanceAddrType: logstorev1.InstanceAddrDefault,
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
					Replicas: 1,
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	}

	wantEnvVar := v1.EnvVar{
		ValueFrom: &v1.EnvVarSource{
			FieldRef: &v1.ObjectFieldSelector{
				APIVersion: "v1",
				FieldPath:  "status.podIP",
			},
		},
	}

	for _, cs := range logstoreContainers(t, opt) {
		for _, c := range cs {
			require.NotContains(t, c.Env, wantEnvVar, "contains envVar %s for: %s", gossipInstanceAddrEnvVarName, c.Name)
		}
	}
}

func TestConfigureHashRingEnv_UseInstanceAddrPodIP(t *testing.T) {
	opt := Options{
		Name:      "test",
		Namespace: "test",
		Image:     "test",
		Stack: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			HashRing: &logstorev1.HashRingSpec{
				Type: logstorev1.HashRingMemberList,
				MemberList: &logstorev1.MemberListSpec{
					InstanceAddrType: logstorev1.InstanceAddrPodIP,
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
					Replicas: 1,
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	}

	wantEnvVar := v1.EnvVar{
		Name: gossipInstanceAddrEnvVarName,
		ValueFrom: &v1.EnvVarSource{
			FieldRef: &v1.ObjectFieldSelector{
				APIVersion: "v1",
				FieldPath:  "status.podIP",
			},
		},
	}

	for _, cs := range logstoreContainers(t, opt) {
		for _, c := range cs {
			require.Contains(t, c.Env, wantEnvVar, "missing envVar %s for: %s", gossipInstanceAddrEnvVarName, c.Name)
		}
	}
}
