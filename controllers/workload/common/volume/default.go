package volume

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// DefaultDataVolumeClaim is the default data volume claim for Elasticsearch pods.
	// We default to a 1GB persistent volume, using the default storage class.
	DefaultDataVolumeClaim = corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: DataVolumeName,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse("1Gi"),
				},
			},
		},
	}
	DefaultDataVolumeMount = corev1.VolumeMount{
		Name:      DataVolumeName,
		MountPath: DataMountPath,
	}

	// DefaultVolumeClaimTemplates is the default volume claim templates for Elasticsearch pods
	DefaultVolumeClaimTemplates = []corev1.PersistentVolumeClaim{DefaultDataVolumeClaim}

	// DefaultLogsVolume is the default EmptyDir logs volume for Elasticsearch pods.
	DefaultLogsVolume = corev1.Volume{
		Name: LogsVolumeName,
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	}
	// DefaultLogsVolumeMount is the default logs volume mount for the Elasticsearch container.
	DefaultLogsVolumeMount = corev1.VolumeMount{
		Name:      LogsVolumeName,
		MountPath: LogsMountPath,
	}
)
