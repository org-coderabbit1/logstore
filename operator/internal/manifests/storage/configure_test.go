package storage_test

import (
	"testing"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests/storage"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestConfigureDeploymentForStorageType(t *testing.T) {
	type tt struct {
		desc string
		opts storage.Options
		dpl  *appsv1.Deployment
		want *appsv1.Deployment
	}

	tc := []tt{
		{
			desc: "object storage other than GCS",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretS3,
			},
			dpl: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
			want: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "object storage GCS",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretGCS,
			},
			dpl: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
			want: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      "test",
											ReadOnly:  false,
											MountPath: "/etc/storage/secrets",
										},
									},
									Env: []corev1.EnvVar{
										{
											Name:  storage.EnvGoogleApplicationCredentials,
											Value: "/etc/storage/secrets/key.json",
										},
									},
								},
							},
							Volumes: []corev1.Volume{
								{
									Name: "test",
									VolumeSource: corev1.VolumeSource{
										Secret: &corev1.SecretVolumeSource{
											SecretName: "test",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tc {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			err := storage.ConfigureDeployment(tc.dpl, tc.opts)
			require.NoError(t, err)
			require.Equal(t, tc.want, tc.dpl)
		})
	}
}

func TestConfigureStatefulSetForStorageType(t *testing.T) {
	type tt struct {
		desc string
		opts storage.Options
		sts  *appsv1.StatefulSet
		want *appsv1.StatefulSet
	}

	tc := []tt{
		{
			desc: "object storage other than GCS",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretS3,
			},
			sts: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
			want: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "object storage GCS",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretGCS,
			},
			sts: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
			want: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      "test",
											ReadOnly:  false,
											MountPath: "/etc/storage/secrets",
										},
									},
									Env: []corev1.EnvVar{
										{
											Name:  storage.EnvGoogleApplicationCredentials,
											Value: "/etc/storage/secrets/key.json",
										},
									},
								},
							},
							Volumes: []corev1.Volume{
								{
									Name: "test",
									VolumeSource: corev1.VolumeSource{
										Secret: &corev1.SecretVolumeSource{
											SecretName: "test",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tc {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			err := storage.ConfigureStatefulSet(tc.sts, tc.opts)
			require.NoError(t, err)
			require.Equal(t, tc.want, tc.sts)
		})
	}
}

func TestConfigureDeploymentForStorageCA(t *testing.T) {
	type tt struct {
		desc string
		opts storage.Options
		dpl  *appsv1.Deployment
		want *appsv1.Deployment
	}

	tc := []tt{
		{
			desc: "object storage other than S3",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretAzure,
			},
			dpl: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-querier",
								},
							},
						},
					},
				},
			},
			want: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-querier",
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "object storage S3",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretS3,
				TLS: &storage.TLSConfig{
					CA:  "test",
					Key: "service-ca.crt",
				},
			},
			dpl: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-querier",
								},
							},
						},
					},
				},
			},
			want: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-querier",
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      "test",
											ReadOnly:  false,
											MountPath: "/etc/storage/ca",
										},
									},
									Args: []string{
										"-s3.http.ca-file=/etc/storage/ca/service-ca.crt",
									},
								},
							},
							Volumes: []corev1.Volume{
								{
									Name: "test",
									VolumeSource: corev1.VolumeSource{
										ConfigMap: &corev1.ConfigMapVolumeSource{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "test",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tc {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			err := storage.ConfigureDeployment(tc.dpl, tc.opts)
			require.NoError(t, err)
			require.Equal(t, tc.want, tc.dpl)
		})
	}
}

func TestConfigureStatefulSetForStorageCA(t *testing.T) {
	type tt struct {
		desc string
		opts storage.Options
		sts  *appsv1.StatefulSet
		want *appsv1.StatefulSet
	}

	tc := []tt{
		{
			desc: "object storage other than S3",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretAzure,
				TLS: &storage.TLSConfig{
					CA: "test",
				},
			},
			sts: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
			want: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "object storage S3",
			opts: storage.Options{
				SecretName:  "test",
				SharedStore: logstorev1.ObjectStorageSecretS3,
				TLS: &storage.TLSConfig{
					CA:  "test",
					Key: "service-ca.crt",
				},
			},
			sts: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
								},
							},
						},
					},
				},
			},
			want: &appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "logstore-ingester",
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      "test",
											ReadOnly:  false,
											MountPath: "/etc/storage/ca",
										},
									},
									Args: []string{
										"-s3.http.ca-file=/etc/storage/ca/service-ca.crt",
									},
								},
							},
							Volumes: []corev1.Volume{
								{
									Name: "test",
									VolumeSource: corev1.VolumeSource{
										ConfigMap: &corev1.ConfigMapVolumeSource{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "test",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tc {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			err := storage.ConfigureStatefulSet(tc.sts, tc.opts)
			require.NoError(t, err)
			require.Equal(t, tc.want, tc.sts)
		})
	}
}
