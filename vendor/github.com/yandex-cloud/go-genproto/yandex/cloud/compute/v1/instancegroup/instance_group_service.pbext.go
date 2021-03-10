// Code generated by protoc-gen-goext. DO NOT EDIT.

package instancegroup

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetInstanceGroupRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *GetInstanceGroupRequest) SetView(v InstanceGroupView) {
	m.View = v
}

func (m *CreateInstanceGroupRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateInstanceGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateInstanceGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateInstanceGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateInstanceGroupRequest) SetInstanceTemplate(v *InstanceTemplate) {
	m.InstanceTemplate = v
}

func (m *CreateInstanceGroupRequest) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

func (m *CreateInstanceGroupRequest) SetDeployPolicy(v *DeployPolicy) {
	m.DeployPolicy = v
}

func (m *CreateInstanceGroupRequest) SetAllocationPolicy(v *AllocationPolicy) {
	m.AllocationPolicy = v
}

func (m *CreateInstanceGroupRequest) SetLoadBalancerSpec(v *LoadBalancerSpec) {
	m.LoadBalancerSpec = v
}

func (m *CreateInstanceGroupRequest) SetHealthChecksSpec(v *HealthChecksSpec) {
	m.HealthChecksSpec = v
}

func (m *CreateInstanceGroupRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateInstanceGroupRequest) SetVariables(v []*Variable) {
	m.Variables = v
}

func (m *CreateInstanceGroupRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateInstanceGroupFromYamlRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateInstanceGroupFromYamlRequest) SetInstanceGroupYaml(v string) {
	m.InstanceGroupYaml = v
}

func (m *CreateInstanceGroupMetadata) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *UpdateInstanceGroupRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *UpdateInstanceGroupRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateInstanceGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateInstanceGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateInstanceGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateInstanceGroupRequest) SetInstanceTemplate(v *InstanceTemplate) {
	m.InstanceTemplate = v
}

func (m *UpdateInstanceGroupRequest) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

func (m *UpdateInstanceGroupRequest) SetDeployPolicy(v *DeployPolicy) {
	m.DeployPolicy = v
}

func (m *UpdateInstanceGroupRequest) SetAllocationPolicy(v *AllocationPolicy) {
	m.AllocationPolicy = v
}

func (m *UpdateInstanceGroupRequest) SetHealthChecksSpec(v *HealthChecksSpec) {
	m.HealthChecksSpec = v
}

func (m *UpdateInstanceGroupRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateInstanceGroupRequest) SetLoadBalancerSpec(v *LoadBalancerSpec) {
	m.LoadBalancerSpec = v
}

func (m *UpdateInstanceGroupRequest) SetVariables(v []*Variable) {
	m.Variables = v
}

func (m *UpdateInstanceGroupRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateInstanceGroupFromYamlRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *UpdateInstanceGroupFromYamlRequest) SetInstanceGroupYaml(v string) {
	m.InstanceGroupYaml = v
}

func (m *UpdateInstanceGroupMetadata) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *StartInstanceGroupRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *StartInstanceGroupMetadata) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *StopInstanceGroupRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *StopInstanceGroupMetadata) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *DeleteInstanceGroupRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *DeleteInstanceGroupMetadata) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *DeleteInstancesMetadata) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *StopInstancesMetadata) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *ListInstanceGroupsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListInstanceGroupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListInstanceGroupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListInstanceGroupsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListInstanceGroupsRequest) SetView(v InstanceGroupView) {
	m.View = v
}

func (m *ListInstanceGroupsResponse) SetInstanceGroups(v []*InstanceGroup) {
	m.InstanceGroups = v
}

func (m *ListInstanceGroupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListInstanceGroupInstancesRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *ListInstanceGroupInstancesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListInstanceGroupInstancesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListInstanceGroupInstancesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListInstanceGroupInstancesResponse) SetInstances(v []*ManagedInstance) {
	m.Instances = v
}

func (m *ListInstanceGroupInstancesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *DeleteInstancesRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *DeleteInstancesRequest) SetManagedInstanceIds(v []string) {
	m.ManagedInstanceIds = v
}

func (m *DeleteInstancesRequest) SetCreateAnother(v bool) {
	m.CreateAnother = v
}

func (m *StopInstancesRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *StopInstancesRequest) SetManagedInstanceIds(v []string) {
	m.ManagedInstanceIds = v
}

func (m *ListInstanceGroupOperationsRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *ListInstanceGroupOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListInstanceGroupOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListInstanceGroupOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListInstanceGroupOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListInstanceGroupOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListInstanceGroupLogRecordsRequest) SetInstanceGroupId(v string) {
	m.InstanceGroupId = v
}

func (m *ListInstanceGroupLogRecordsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListInstanceGroupLogRecordsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListInstanceGroupLogRecordsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListInstanceGroupLogRecordsResponse) SetLogRecords(v []*LogRecord) {
	m.LogRecords = v
}

func (m *ListInstanceGroupLogRecordsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
