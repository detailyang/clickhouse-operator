// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiCluster) DeepCopyInto(out *ChiCluster) {
	*out = *in
	in.Layout.DeepCopyInto(&out.Layout)
	out.Templates = in.Templates
	out.Address = in.Address
	if in.Chi != nil {
		in, out := &in.Chi, &out.Chi
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiCluster.
func (in *ChiCluster) DeepCopy() *ChiCluster {
	if in == nil {
		return nil
	}
	out := new(ChiCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterAddress) DeepCopyInto(out *ChiClusterAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterAddress.
func (in *ChiClusterAddress) DeepCopy() *ChiClusterAddress {
	if in == nil {
		return nil
	}
	out := new(ChiClusterAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiConfiguration) DeepCopyInto(out *ChiConfiguration) {
	*out = *in
	in.Zookeeper.DeepCopyInto(&out.Zookeeper)
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Quotas != nil {
		in, out := &in.Quotas, &out.Quotas
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ChiCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiConfiguration.
func (in *ChiConfiguration) DeepCopy() *ChiConfiguration {
	if in == nil {
		return nil
	}
	out := new(ChiConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDefaults) DeepCopyInto(out *ChiDefaults) {
	*out = *in
	out.DistributedDDL = in.DistributedDDL
	out.Templates = in.Templates
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDefaults.
func (in *ChiDefaults) DeepCopy() *ChiDefaults {
	if in == nil {
		return nil
	}
	out := new(ChiDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDistributedDDL) DeepCopyInto(out *ChiDistributedDDL) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDistributedDDL.
func (in *ChiDistributedDDL) DeepCopy() *ChiDistributedDDL {
	if in == nil {
		return nil
	}
	out := new(ChiDistributedDDL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHost) DeepCopyInto(out *ChiHost) {
	*out = *in
	out.Templates = in.Templates
	out.Address = in.Address
	out.Config = in.Config
	if in.Chi != nil {
		in, out := &in.Chi, &out.Chi
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHost.
func (in *ChiHost) DeepCopy() *ChiHost {
	if in == nil {
		return nil
	}
	out := new(ChiHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHostAddress) DeepCopyInto(out *ChiHostAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHostAddress.
func (in *ChiHostAddress) DeepCopy() *ChiHostAddress {
	if in == nil {
		return nil
	}
	out := new(ChiHostAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHostConfig) DeepCopyInto(out *ChiHostConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHostConfig.
func (in *ChiHostConfig) DeepCopy() *ChiHostConfig {
	if in == nil {
		return nil
	}
	out := new(ChiHostConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiLayout) DeepCopyInto(out *ChiLayout) {
	*out = *in
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make([]ChiShard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiLayout.
func (in *ChiLayout) DeepCopy() *ChiLayout {
	if in == nil {
		return nil
	}
	out := new(ChiLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPodTemplate) DeepCopyInto(out *ChiPodTemplate) {
	*out = *in
	in.Zone.DeepCopyInto(&out.Zone)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPodTemplate.
func (in *ChiPodTemplate) DeepCopy() *ChiPodTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiPodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPodTemplateZone) DeepCopyInto(out *ChiPodTemplateZone) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPodTemplateZone.
func (in *ChiPodTemplateZone) DeepCopy() *ChiPodTemplateZone {
	if in == nil {
		return nil
	}
	out := new(ChiPodTemplateZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiServiceTemplate) DeepCopyInto(out *ChiServiceTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiServiceTemplate.
func (in *ChiServiceTemplate) DeepCopy() *ChiServiceTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiServiceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiShard) DeepCopyInto(out *ChiShard) {
	*out = *in
	out.Templates = in.Templates
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]ChiHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Address = in.Address
	if in.Chi != nil {
		in, out := &in.Chi, &out.Chi
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiShard.
func (in *ChiShard) DeepCopy() *ChiShard {
	if in == nil {
		return nil
	}
	out := new(ChiShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiShardAddress) DeepCopyInto(out *ChiShardAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiShardAddress.
func (in *ChiShardAddress) DeepCopy() *ChiShardAddress {
	if in == nil {
		return nil
	}
	out := new(ChiShardAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiSpec) DeepCopyInto(out *ChiSpec) {
	*out = *in
	out.Defaults = in.Defaults
	in.Configuration.DeepCopyInto(&out.Configuration)
	in.Templates.DeepCopyInto(&out.Templates)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiSpec.
func (in *ChiSpec) DeepCopy() *ChiSpec {
	if in == nil {
		return nil
	}
	out := new(ChiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiStatus) DeepCopyInto(out *ChiStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiStatus.
func (in *ChiStatus) DeepCopy() *ChiStatus {
	if in == nil {
		return nil
	}
	out := new(ChiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplateNames) DeepCopyInto(out *ChiTemplateNames) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplateNames.
func (in *ChiTemplateNames) DeepCopy() *ChiTemplateNames {
	if in == nil {
		return nil
	}
	out := new(ChiTemplateNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplates) DeepCopyInto(out *ChiTemplates) {
	*out = *in
	if in.PodTemplates != nil {
		in, out := &in.PodTemplates, &out.PodTemplates
		*out = make([]ChiPodTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]ChiVolumeClaimTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceTemplates != nil {
		in, out := &in.ServiceTemplates, &out.ServiceTemplates
		*out = make([]ChiServiceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodTemplatesIndex != nil {
		in, out := &in.PodTemplatesIndex, &out.PodTemplatesIndex
		*out = make(map[string]*ChiPodTemplate, len(*in))
		for key, val := range *in {
			var outVal *ChiPodTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ChiPodTemplate)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.VolumeClaimTemplatesIndex != nil {
		in, out := &in.VolumeClaimTemplatesIndex, &out.VolumeClaimTemplatesIndex
		*out = make(map[string]*ChiVolumeClaimTemplate, len(*in))
		for key, val := range *in {
			var outVal *ChiVolumeClaimTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ChiVolumeClaimTemplate)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ServiceTemplatesIndex != nil {
		in, out := &in.ServiceTemplatesIndex, &out.ServiceTemplatesIndex
		*out = make(map[string]*ChiServiceTemplate, len(*in))
		for key, val := range *in {
			var outVal *ChiServiceTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ChiServiceTemplate)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplates.
func (in *ChiTemplates) DeepCopy() *ChiTemplates {
	if in == nil {
		return nil
	}
	out := new(ChiTemplates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiVolumeClaimTemplate) DeepCopyInto(out *ChiVolumeClaimTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiVolumeClaimTemplate.
func (in *ChiVolumeClaimTemplate) DeepCopy() *ChiVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiZookeeperConfig) DeepCopyInto(out *ChiZookeeperConfig) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]ChiZookeeperNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiZookeeperConfig.
func (in *ChiZookeeperConfig) DeepCopy() *ChiZookeeperConfig {
	if in == nil {
		return nil
	}
	out := new(ChiZookeeperConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiZookeeperNode) DeepCopyInto(out *ChiZookeeperNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiZookeeperNode.
func (in *ChiZookeeperNode) DeepCopy() *ChiZookeeperNode {
	if in == nil {
		return nil
	}
	out := new(ChiZookeeperNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallation) DeepCopyInto(out *ClickHouseInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallation.
func (in *ClickHouseInstallation) DeepCopy() *ClickHouseInstallation {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationList) DeepCopyInto(out *ClickHouseInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationList.
func (in *ClickHouseInstallationList) DeepCopy() *ClickHouseInstallationList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationTemplate) DeepCopyInto(out *ClickHouseInstallationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationTemplate.
func (in *ClickHouseInstallationTemplate) DeepCopy() *ClickHouseInstallationTemplate {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationTemplateList) DeepCopyInto(out *ClickHouseInstallationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseInstallationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationTemplateList.
func (in *ClickHouseInstallationTemplateList) DeepCopy() *ClickHouseInstallationTemplateList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
