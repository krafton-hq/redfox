//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]ClusterRole, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	out.Apiserver = in.Apiserver
	if in.AwsIamIdps != nil {
		in, out := &in.AwsIamIdps, &out.AwsIamIdps
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatusApiserver) DeepCopyInto(out *ClusterStatusApiserver) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatusApiserver.
func (in *ClusterStatusApiserver) DeepCopy() *ClusterStatusApiserver {
	if in == nil {
		return nil
	}
	out := new(ClusterStatusApiserver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressAddress) DeepCopyInto(out *IngressAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressAddress.
func (in *IngressAddress) DeepCopy() *IngressAddress {
	if in == nil {
		return nil
	}
	out := new(IngressAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressAddressList) DeepCopyInto(out *IngressAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressAddressList.
func (in *IngressAddressList) DeepCopy() *IngressAddressList {
	if in == nil {
		return nil
	}
	out := new(IngressAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressAddressSpec) DeepCopyInto(out *IngressAddressSpec) {
	*out = *in
	out.Default = in.Default
	in.Specs.DeepCopyInto(&out.Specs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressAddressSpec.
func (in *IngressAddressSpec) DeepCopy() *IngressAddressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressAddressSpecDefault) DeepCopyInto(out *IngressAddressSpecDefault) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressAddressSpecDefault.
func (in *IngressAddressSpecDefault) DeepCopy() *IngressAddressSpecDefault {
	if in == nil {
		return nil
	}
	out := new(IngressAddressSpecDefault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressAddressSpecs) DeepCopyInto(out *IngressAddressSpecs) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressAddressSpecs.
func (in *IngressAddressSpecs) DeepCopy() *IngressAddressSpecs {
	if in == nil {
		return nil
	}
	out := new(IngressAddressSpecs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestCommit) DeepCopyInto(out *LatestCommit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestCommit.
func (in *LatestCommit) DeepCopy() *LatestCommit {
	if in == nil {
		return nil
	}
	out := new(LatestCommit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LatestCommit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestCommitGitRef) DeepCopyInto(out *LatestCommitGitRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestCommitGitRef.
func (in *LatestCommitGitRef) DeepCopy() *LatestCommitGitRef {
	if in == nil {
		return nil
	}
	out := new(LatestCommitGitRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestCommitList) DeepCopyInto(out *LatestCommitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LatestCommit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestCommitList.
func (in *LatestCommitList) DeepCopy() *LatestCommitList {
	if in == nil {
		return nil
	}
	out := new(LatestCommitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LatestCommitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestCommitSpec) DeepCopyInto(out *LatestCommitSpec) {
	*out = *in
	out.GitRef = in.GitRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestCommitSpec.
func (in *LatestCommitSpec) DeepCopy() *LatestCommitSpec {
	if in == nil {
		return nil
	}
	out := new(LatestCommitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestCommitStatus) DeepCopyInto(out *LatestCommitStatus) {
	*out = *in
	in.CommitTimestamp.DeepCopyInto(&out.CommitTimestamp)
	in.SyncTimestamp.DeepCopyInto(&out.SyncTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestCommitStatus.
func (in *LatestCommitStatus) DeepCopy() *LatestCommitStatus {
	if in == nil {
		return nil
	}
	out := new(LatestCommitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestVersion) DeepCopyInto(out *LatestVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestVersion.
func (in *LatestVersion) DeepCopy() *LatestVersion {
	if in == nil {
		return nil
	}
	out := new(LatestVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LatestVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestVersionGitRef) DeepCopyInto(out *LatestVersionGitRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestVersionGitRef.
func (in *LatestVersionGitRef) DeepCopy() *LatestVersionGitRef {
	if in == nil {
		return nil
	}
	out := new(LatestVersionGitRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestVersionList) DeepCopyInto(out *LatestVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LatestVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestVersionList.
func (in *LatestVersionList) DeepCopy() *LatestVersionList {
	if in == nil {
		return nil
	}
	out := new(LatestVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LatestVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestVersionSpec) DeepCopyInto(out *LatestVersionSpec) {
	*out = *in
	out.GitRef = in.GitRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestVersionSpec.
func (in *LatestVersionSpec) DeepCopy() *LatestVersionSpec {
	if in == nil {
		return nil
	}
	out := new(LatestVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestVersionStatus) DeepCopyInto(out *LatestVersionStatus) {
	*out = *in
	out.VersionRef = in.VersionRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestVersionStatus.
func (in *LatestVersionStatus) DeepCopy() *LatestVersionStatus {
	if in == nil {
		return nil
	}
	out := new(LatestVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatestVersionVersionRef) DeepCopyInto(out *LatestVersionVersionRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatestVersionVersionRef.
func (in *LatestVersionVersionRef) DeepCopy() *LatestVersionVersionRef {
	if in == nil {
		return nil
	}
	out := new(LatestVersionVersionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatIp) DeepCopyInto(out *NatIp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatIp.
func (in *NatIp) DeepCopy() *NatIp {
	if in == nil {
		return nil
	}
	out := new(NatIp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NatIp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatIpList) DeepCopyInto(out *NatIpList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NatIp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatIpList.
func (in *NatIpList) DeepCopy() *NatIpList {
	if in == nil {
		return nil
	}
	out := new(NatIpList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NatIpList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatIpSpec) DeepCopyInto(out *NatIpSpec) {
	*out = *in
	if in.Cidrs != nil {
		in, out := &in.Cidrs, &out.Cidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatIpSpec.
func (in *NatIpSpec) DeepCopy() *NatIpSpec {
	if in == nil {
		return nil
	}
	out := new(NatIpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionMetadata) DeepCopyInto(out *RegionMetadata) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionMetadata.
func (in *RegionMetadata) DeepCopy() *RegionMetadata {
	if in == nil {
		return nil
	}
	out := new(RegionMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegionMetadata) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionMetadataArgocd) DeepCopyInto(out *RegionMetadataArgocd) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionMetadataArgocd.
func (in *RegionMetadataArgocd) DeepCopy() *RegionMetadataArgocd {
	if in == nil {
		return nil
	}
	out := new(RegionMetadataArgocd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionMetadataIdentity) DeepCopyInto(out *RegionMetadataIdentity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionMetadataIdentity.
func (in *RegionMetadataIdentity) DeepCopy() *RegionMetadataIdentity {
	if in == nil {
		return nil
	}
	out := new(RegionMetadataIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionMetadataList) DeepCopyInto(out *RegionMetadataList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RegionMetadata, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionMetadataList.
func (in *RegionMetadataList) DeepCopy() *RegionMetadataList {
	if in == nil {
		return nil
	}
	out := new(RegionMetadataList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegionMetadataList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionMetadataSpec) DeepCopyInto(out *RegionMetadataSpec) {
	*out = *in
	out.Argocd = in.Argocd
	if in.ContainerRegistries != nil {
		in, out := &in.ContainerRegistries, &out.ContainerRegistries
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Identity = in.Identity
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionMetadataSpec.
func (in *RegionMetadataSpec) DeepCopy() *RegionMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(RegionMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Version) DeepCopyInto(out *Version) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Version.
func (in *Version) DeepCopy() *Version {
	if in == nil {
		return nil
	}
	out := new(Version)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Version) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionCount) DeepCopyInto(out *VersionCount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionCount.
func (in *VersionCount) DeepCopy() *VersionCount {
	if in == nil {
		return nil
	}
	out := new(VersionCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VersionCount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionCountList) DeepCopyInto(out *VersionCountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VersionCount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionCountList.
func (in *VersionCountList) DeepCopy() *VersionCountList {
	if in == nil {
		return nil
	}
	out := new(VersionCountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VersionCountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionCountSpec) DeepCopyInto(out *VersionCountSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionCountSpec.
func (in *VersionCountSpec) DeepCopy() *VersionCountSpec {
	if in == nil {
		return nil
	}
	out := new(VersionCountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionDetail) DeepCopyInto(out *VersionDetail) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionDetail.
func (in *VersionDetail) DeepCopy() *VersionDetail {
	if in == nil {
		return nil
	}
	out := new(VersionDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionGitRef) DeepCopyInto(out *VersionGitRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionGitRef.
func (in *VersionGitRef) DeepCopy() *VersionGitRef {
	if in == nil {
		return nil
	}
	out := new(VersionGitRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionList) DeepCopyInto(out *VersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Version, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionList.
func (in *VersionList) DeepCopy() *VersionList {
	if in == nil {
		return nil
	}
	out := new(VersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionSpec) DeepCopyInto(out *VersionSpec) {
	*out = *in
	out.VersionDetail = in.VersionDetail
	out.GitRef = in.GitRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionSpec.
func (in *VersionSpec) DeepCopy() *VersionSpec {
	if in == nil {
		return nil
	}
	out := new(VersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionStatus) DeepCopyInto(out *VersionStatus) {
	*out = *in
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make([]VersionStatusArtifact, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionStatus.
func (in *VersionStatus) DeepCopy() *VersionStatus {
	if in == nil {
		return nil
	}
	out := new(VersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionStatusArtifact) DeepCopyInto(out *VersionStatusArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionStatusArtifact.
func (in *VersionStatusArtifact) DeepCopy() *VersionStatusArtifact {
	if in == nil {
		return nil
	}
	out := new(VersionStatusArtifact)
	in.DeepCopyInto(out)
	return out
}
