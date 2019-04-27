package v1alpha1

import (
	operatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultDiskMakerImageVersion	= "registry.svc.ci.openshift.org/openshift/origin-v4.0:local-storage-diskmaker"
	defaultProvisionImage		= "quay.io/external_storage/local-volume-provisioner:v2.3.0"
)

type LocalVolumeList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`
	Items		[]LocalVolume	`json:"items"`
}
type LocalVolume struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata"`
	Spec			LocalVolumeSpec		`json:"spec"`
	Status			LocalVolumeStatus	`json:"status,omitempty"`
}
type LocalVolumeSpec struct {
	NodeSelector		*corev1.NodeSelector	`json:"nodeSelector,omitempty"`
	StorageClassDevices	[]StorageClassDevice	`json:"storageClassDevices,omitempty"`
	LocalProvisionerImageVersion
	DiskMakerImageVersion
}
type PersistentVolumeMode string

const (
	PersistentVolumeBlock		PersistentVolumeMode	= "Block"
	PersistentVolumeFilesystem	PersistentVolumeMode	= "Filesystem"
)

type StorageClassDevice struct {
	StorageClassName	string			`json:"storageClassName"`
	VolumeMode		PersistentVolumeMode	`json:"volumeMode"`
	FSType			string			`json:"fsType"`
	DeviceNames		[]string		`json:"deviceNames,omitempty"`
	DeviceIDs		[]string		`json:"deviceIDs,omitempty"`
}
type LocalProvisionerImageVersion struct {
	ProvisionerImage string `json:"provisionerImage,omitempty"`
}
type DiskMakerImageVersion struct {
	DiskMakerImage string `json:"diskMakerImage,omitempty"`
}
type LocalVolumeStatus struct {
	ObservedGeneration	*int64					`json:"observedGeneration,omitempty"`
	Children		[]operatorv1alpha1.GenerationHistory	`json:"children,omitempty"`
	State			operatorv1alpha1.ManagementState	`json:"state,omitempty"`
	Conditions		[]operatorv1alpha1.OperatorCondition
}

func (local *LocalVolume) SetDefaults() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(local.Spec.DiskMakerImageVersion.DiskMakerImage) == 0 {
		local.Spec.DiskMakerImageVersion = DiskMakerImageVersion{defaultDiskMakerImageVersion}
	}
	if len(local.Spec.LocalProvisionerImageVersion.ProvisionerImage) == 0 {
		local.Spec.LocalProvisionerImageVersion = LocalProvisionerImageVersion{defaultProvisionImage}
	}
}
