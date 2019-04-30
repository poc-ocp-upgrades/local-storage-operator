package v1alpha1

import (
	sdkK8sutil "github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	version		= "v1alpha1"
	groupName	= "local.storage.openshift.io"
	LocalVolumeKind	= "LocalVolume"
)

var (
	SchemeBuilder		= runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme		= SchemeBuilder.AddToScheme
	SchemeGroupVersion	= schema.GroupVersion{Group: groupName, Version: version}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sdkK8sutil.AddToSDKScheme(AddToScheme)
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(SchemeGroupVersion, &LocalVolume{}, &LocalVolumeList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
