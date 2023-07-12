package manifests

import (
	"fmt"

	"github.com/imdario/mergo"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
)

// BuildLogstoreGossipRingService creates a k8s service for the gossip/memberlist members of the cluster
func BuildLogstoreGossipRingService(stackName string) *corev1.Service {
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   fmt.Sprintf("%s-gossip-ring", stackName),
			Labels: commonLabels(stackName),
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Ports: []corev1.ServicePort{
				{
					Name:       logstoreGossipPortName,
					Port:       gossipPort,
					Protocol:   protocolTCP,
					TargetPort: intstr.IntOrString{IntVal: gossipPort},
				},
			},
			Selector: commonLabels(stackName),
		},
	}
}

func configureHashRingEnv(p *corev1.PodSpec, opts Options) error {
	resetProxyVar(p, gossipInstanceAddrEnvVarName)
	hashRing := opts.Stack.HashRing

	if hashRing == nil {
		return nil
	}

	if hashRing.Type != logstorev1.HashRingMemberList {
		return nil
	}

	if hashRing.MemberList == nil {
		return nil
	}

	switch hashRing.MemberList.InstanceAddrType {
	case "", logstorev1.InstanceAddrDefault:
		return nil
	default:
		// Proceed with appending env var
	}

	src := corev1.Container{
		Env: []corev1.EnvVar{
			{
				Name: gossipInstanceAddrEnvVarName,
				ValueFrom: &corev1.EnvVarSource{
					FieldRef: &corev1.ObjectFieldSelector{
						APIVersion: "v1",
						FieldPath:  "status.podIP",
					},
				},
			},
		},
	}

	for i, dst := range p.Containers {
		if err := mergo.Merge(&dst, src, mergo.WithAppendSlice); err != nil {
			return err
		}
		p.Containers[i] = dst
	}

	return nil
}
