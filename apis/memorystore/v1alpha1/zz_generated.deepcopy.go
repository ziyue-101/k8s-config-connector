//go:build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryEndpoint) DeepCopyInto(out *DiscoveryEndpoint) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryEndpoint.
func (in *DiscoveryEndpoint) DeepCopy() *DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := new(DiscoveryEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateInfo != nil {
		in, out := &in.StateInfo, &out.StateInfo
		*out = new(Instance_StateInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.AuthorizationMode != nil {
		in, out := &in.AuthorizationMode, &out.AuthorizationMode
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int32)
		**out = **in
	}
	if in.DiscoveryEndpoints != nil {
		in, out := &in.DiscoveryEndpoints, &out.DiscoveryEndpoints
		*out = make([]DiscoveryEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = new(PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.EngineConfigs != nil {
		in, out := &in.EngineConfigs, &out.EngineConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = new(NodeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneDistributionConfig != nil {
		in, out := &in.ZoneDistributionConfig, &out.ZoneDistributionConfig
		*out = new(ZoneDistributionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PscAutoConnections != nil {
		in, out := &in.PscAutoConnections, &out.PscAutoConnections
		*out = make([]PscAutoConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance_StateInfo) DeepCopyInto(out *Instance_StateInfo) {
	*out = *in
	if in.UpdateInfo != nil {
		in, out := &in.UpdateInfo, &out.UpdateInfo
		*out = new(Instance_StateInfo_UpdateInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance_StateInfo.
func (in *Instance_StateInfo) DeepCopy() *Instance_StateInfo {
	if in == nil {
		return nil
	}
	out := new(Instance_StateInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance_StateInfo_UpdateInfo) DeepCopyInto(out *Instance_StateInfo_UpdateInfo) {
	*out = *in
	if in.TargetShardCount != nil {
		in, out := &in.TargetShardCount, &out.TargetShardCount
		*out = new(int32)
		**out = **in
	}
	if in.TargetReplicaCount != nil {
		in, out := &in.TargetReplicaCount, &out.TargetReplicaCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance_StateInfo_UpdateInfo.
func (in *Instance_StateInfo_UpdateInfo) DeepCopy() *Instance_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := new(Instance_StateInfo_UpdateInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstance) DeepCopyInto(out *MemorystoreInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorystoreInstance.
func (in *MemorystoreInstance) DeepCopy() *MemorystoreInstance {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemorystoreInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstanceList) DeepCopyInto(out *MemorystoreInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MemorystoreInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorystoreInstanceList.
func (in *MemorystoreInstanceList) DeepCopy() *MemorystoreInstanceList {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemorystoreInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstanceObservedState) DeepCopyInto(out *MemorystoreInstanceObservedState) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateInfo != nil {
		in, out := &in.StateInfo, &out.StateInfo
		*out = new(Instance_StateInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.AuthorizationMode != nil {
		in, out := &in.AuthorizationMode, &out.AuthorizationMode
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryEndpoints != nil {
		in, out := &in.DiscoveryEndpoints, &out.DiscoveryEndpoints
		*out = make([]DiscoveryEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = new(NodeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneDistributionConfig != nil {
		in, out := &in.ZoneDistributionConfig, &out.ZoneDistributionConfig
		*out = new(ZoneDistributionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PscAutoConnections != nil {
		in, out := &in.PscAutoConnections, &out.PscAutoConnections
		*out = make([]PscAutoConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorystoreInstanceObservedState.
func (in *MemorystoreInstanceObservedState) DeepCopy() *MemorystoreInstanceObservedState {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstanceObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstanceParent) DeepCopyInto(out *MemorystoreInstanceParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorystoreInstanceParent.
func (in *MemorystoreInstanceParent) DeepCopy() *MemorystoreInstanceParent {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstanceParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstanceRef) DeepCopyInto(out *MemorystoreInstanceRef) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(MemorystoreInstanceParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorystoreInstanceRef.
func (in *MemorystoreInstanceRef) DeepCopy() *MemorystoreInstanceRef {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstanceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstanceSpec) DeepCopyInto(out *MemorystoreInstanceSpec) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	in.Parent.DeepCopyInto(&out.Parent)
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.AuthorizationMode != nil {
		in, out := &in.AuthorizationMode, &out.AuthorizationMode
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int32)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = new(PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.EngineConfigs != nil {
		in, out := &in.EngineConfigs, &out.EngineConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ZoneDistributionConfig != nil {
		in, out := &in.ZoneDistributionConfig, &out.ZoneDistributionConfig
		*out = new(ZoneDistributionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PscAutoConnectionsSpec != nil {
		in, out := &in.PscAutoConnectionsSpec, &out.PscAutoConnectionsSpec
		*out = make([]PscAutoConnectionSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorystoreInstanceSpec.
func (in *MemorystoreInstanceSpec) DeepCopy() *MemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstanceStatus) DeepCopyInto(out *MemorystoreInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(MemorystoreInstanceObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorystoreInstanceStatus.
func (in *MemorystoreInstanceStatus) DeepCopy() *MemorystoreInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	if in.SizeGb != nil {
		in, out := &in.SizeGb, &out.SizeGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfig) DeepCopyInto(out *PersistenceConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.RdbConfig != nil {
		in, out := &in.RdbConfig, &out.RdbConfig
		*out = new(PersistenceConfig_RDBConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AofConfig != nil {
		in, out := &in.AofConfig, &out.AofConfig
		*out = new(PersistenceConfig_AOFConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfig.
func (in *PersistenceConfig) DeepCopy() *PersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfig_AOFConfig) DeepCopyInto(out *PersistenceConfig_AOFConfig) {
	*out = *in
	if in.AppendFsync != nil {
		in, out := &in.AppendFsync, &out.AppendFsync
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfig_AOFConfig.
func (in *PersistenceConfig_AOFConfig) DeepCopy() *PersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfig_AOFConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfig_RDBConfig) DeepCopyInto(out *PersistenceConfig_RDBConfig) {
	*out = *in
	if in.RdbSnapshotPeriod != nil {
		in, out := &in.RdbSnapshotPeriod, &out.RdbSnapshotPeriod
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotStartTime != nil {
		in, out := &in.RdbSnapshotStartTime, &out.RdbSnapshotStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfig_RDBConfig.
func (in *PersistenceConfig_RDBConfig) DeepCopy() *PersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfig_RDBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PscAutoConnection) DeepCopyInto(out *PscAutoConnection) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.PscConnectionID != nil {
		in, out := &in.PscConnectionID, &out.PscConnectionID
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.ForwardingRule != nil {
		in, out := &in.ForwardingRule, &out.ForwardingRule
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.ServiceAttachment != nil {
		in, out := &in.ServiceAttachment, &out.ServiceAttachment
		*out = new(string)
		**out = **in
	}
	if in.PscConnectionStatus != nil {
		in, out := &in.PscConnectionStatus, &out.PscConnectionStatus
		*out = new(string)
		**out = **in
	}
	if in.ConnectionType != nil {
		in, out := &in.ConnectionType, &out.ConnectionType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PscAutoConnection.
func (in *PscAutoConnection) DeepCopy() *PscAutoConnection {
	if in == nil {
		return nil
	}
	out := new(PscAutoConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PscAutoConnectionSpec) DeepCopyInto(out *PscAutoConnectionSpec) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(v1beta1.ComputeNetworkRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PscAutoConnectionSpec.
func (in *PscAutoConnectionSpec) DeepCopy() *PscAutoConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(PscAutoConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneDistributionConfig) DeepCopyInto(out *ZoneDistributionConfig) {
	*out = *in
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneDistributionConfig.
func (in *ZoneDistributionConfig) DeepCopy() *ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := new(ZoneDistributionConfig)
	in.DeepCopyInto(out)
	return out
}
