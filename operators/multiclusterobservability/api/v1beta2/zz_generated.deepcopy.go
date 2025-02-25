// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/open-cluster-management/multicluster-observability-operator/operators/multiclusterobservability/api/shared"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedConfig) DeepCopyInto(out *AdvancedConfig) {
	*out = *in
	if in.RetentionConfig != nil {
		in, out := &in.RetentionConfig, &out.RetentionConfig
		*out = new(RetentionConfig)
		**out = **in
	}
	if in.RBACQueryProxy != nil {
		in, out := &in.RBACQueryProxy, &out.RBACQueryProxy
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Alertmanager != nil {
		in, out := &in.Alertmanager, &out.Alertmanager
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StoreMemcached != nil {
		in, out := &in.StoreMemcached, &out.StoreMemcached
		*out = new(CacheConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryFrontendMemcached != nil {
		in, out := &in.QueryFrontendMemcached, &out.QueryFrontendMemcached
		*out = new(CacheConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ObservatoriumAPI != nil {
		in, out := &in.ObservatoriumAPI, &out.ObservatoriumAPI
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryFrontend != nil {
		in, out := &in.QueryFrontend, &out.QueryFrontend
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Compact != nil {
		in, out := &in.Compact, &out.Compact
		*out = new(CompactSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Receive != nil {
		in, out := &in.Receive, &out.Receive
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Store != nil {
		in, out := &in.Store, &out.Store
		*out = new(CommonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedConfig.
func (in *AdvancedConfig) DeepCopy() *AdvancedConfig {
	if in == nil {
		return nil
	}
	out := new(AdvancedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheConfig) DeepCopyInto(out *CacheConfig) {
	*out = *in
	if in.MemoryLimitMB != nil {
		in, out := &in.MemoryLimitMB, &out.MemoryLimitMB
		*out = new(int32)
		**out = **in
	}
	if in.ConnectionLimit != nil {
		in, out := &in.ConnectionLimit, &out.ConnectionLimit
		*out = new(int32)
		**out = **in
	}
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheConfig.
func (in *CacheConfig) DeepCopy() *CacheConfig {
	if in == nil {
		return nil
	}
	out := new(CacheConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSpec) DeepCopyInto(out *CommonSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSpec.
func (in *CommonSpec) DeepCopy() *CommonSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompactSpec) DeepCopyInto(out *CompactSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactSpec.
func (in *CompactSpec) DeepCopy() *CompactSpec {
	if in == nil {
		return nil
	}
	out := new(CompactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservability) DeepCopyInto(out *MultiClusterObservability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservability.
func (in *MultiClusterObservability) DeepCopy() *MultiClusterObservability {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterObservability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilityList) DeepCopyInto(out *MultiClusterObservabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterObservability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilityList.
func (in *MultiClusterObservabilityList) DeepCopy() *MultiClusterObservabilityList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterObservabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilitySpec) DeepCopyInto(out *MultiClusterObservabilitySpec) {
	*out = *in
	if in.AdvancedConfig != nil {
		in, out := &in.AdvancedConfig, &out.AdvancedConfig
		*out = new(AdvancedConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageConfig != nil {
		in, out := &in.StorageConfig, &out.StorageConfig
		*out = new(StorageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ObservabilityAddonSpec != nil {
		in, out := &in.ObservabilityAddonSpec, &out.ObservabilityAddonSpec
		*out = new(shared.ObservabilityAddonSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilitySpec.
func (in *MultiClusterObservabilitySpec) DeepCopy() *MultiClusterObservabilitySpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilityStatus) DeepCopyInto(out *MultiClusterObservabilityStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]shared.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilityStatus.
func (in *MultiClusterObservabilityStatus) DeepCopy() *MultiClusterObservabilityStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionConfig) DeepCopyInto(out *RetentionConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionConfig.
func (in *RetentionConfig) DeepCopy() *RetentionConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.MetricObjectStorage != nil {
		in, out := &in.MetricObjectStorage, &out.MetricObjectStorage
		*out = new(shared.PreConfiguredStorage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}
