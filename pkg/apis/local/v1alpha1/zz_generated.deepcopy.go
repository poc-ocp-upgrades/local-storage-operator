package v1alpha1

import (
	operatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *DiskMakerImageVersion) DeepCopyInto(out *DiskMakerImageVersion) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *DiskMakerImageVersion) DeepCopy() *DiskMakerImageVersion {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DiskMakerImageVersion)
	in.DeepCopyInto(out)
	return out
}
func (in *LocalProvisionerImageVersion) DeepCopyInto(out *LocalProvisionerImageVersion) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *LocalProvisionerImageVersion) DeepCopy() *LocalProvisionerImageVersion {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LocalProvisionerImageVersion)
	in.DeepCopyInto(out)
	return out
}
func (in *LocalVolume) DeepCopyInto(out *LocalVolume) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *LocalVolume) DeepCopy() *LocalVolume {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LocalVolume)
	in.DeepCopyInto(out)
	return out
}
func (in *LocalVolume) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *LocalVolumeList) DeepCopyInto(out *LocalVolumeList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *LocalVolumeList) DeepCopy() *LocalVolumeList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LocalVolumeList)
	in.DeepCopyInto(out)
	return out
}
func (in *LocalVolumeList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *LocalVolumeSpec) DeepCopyInto(out *LocalVolumeSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClassDevices != nil {
		in, out := &in.StorageClassDevices, &out.StorageClassDevices
		*out = make([]StorageClassDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LocalProvisionerImageVersion = in.LocalProvisionerImageVersion
	out.DiskMakerImageVersion = in.DiskMakerImageVersion
	return
}
func (in *LocalVolumeSpec) DeepCopy() *LocalVolumeSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LocalVolumeSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *LocalVolumeStatus) DeepCopyInto(out *LocalVolumeStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]operatorv1alpha1.GenerationHistory, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]operatorv1alpha1.OperatorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *LocalVolumeStatus) DeepCopy() *LocalVolumeStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LocalVolumeStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *StorageClassDevice) DeepCopyInto(out *StorageClassDevice) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.DeviceNames != nil {
		in, out := &in.DeviceNames, &out.DeviceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeviceIDs != nil {
		in, out := &in.DeviceIDs, &out.DeviceIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *StorageClassDevice) DeepCopy() *StorageClassDevice {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(StorageClassDevice)
	in.DeepCopyInto(out)
	return out
}
