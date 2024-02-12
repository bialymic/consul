// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/destinations_configuration.proto

package meshv2beta1

import (
	v2beta1 "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DestinationsConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selection of workloads these destinations should apply to.
	// These can be prefixes or specific workload names.
	Workloads *v2beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=workloads,proto3" json:"workloads,omitempty"`
	// DefaultConfig applies to all destinations for the workloads selected by this resource.
	DefaultConfig *DestinationConfiguration `protobuf:"bytes,2,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	// ConfigOverrides provides per-destination or per-destination-port config overrides.
	ConfigOverrides []*DestinationConfigOverrides `protobuf:"bytes,3,rep,name=config_overrides,json=configOverrides,proto3" json:"config_overrides,omitempty"`
}

func (x *DestinationsConfiguration) Reset() {
	*x = DestinationsConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationsConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationsConfiguration) ProtoMessage() {}

func (x *DestinationsConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationsConfiguration.ProtoReflect.Descriptor instead.
func (*DestinationsConfiguration) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_destinations_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *DestinationsConfiguration) GetWorkloads() *v2beta1.WorkloadSelector {
	if x != nil {
		return x.Workloads
	}
	return nil
}

func (x *DestinationsConfiguration) GetDefaultConfig() *DestinationConfiguration {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

func (x *DestinationsConfiguration) GetConfigOverrides() []*DestinationConfigOverrides {
	if x != nil {
		return x.ConfigOverrides
	}
	return nil
}

// DestinationConfigOverrides allow to override destination configuration per destination_ref/port/datacenter.
// In that sense, those three fields (destination_ref, destination_port and datacenter) are treated
// sort of like map keys and config is a like a map value for that key.
type DestinationConfigOverrides struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DestinationRef is the reference to an destination service that this configuration applies to.
	// This has to be pbcatalog.Service type.
	DestinationRef *pbresource.Reference `protobuf:"bytes,1,opt,name=destination_ref,json=destinationRef,proto3" json:"destination_ref,omitempty"`
	// DestinationPort is the port of the destination service.
	//
	// For more details on potential values of this field, see documentation for Service.ServicePort.
	//
	// If not provided, this configuration will apply to all ports of an destination.
	DestinationPort string `protobuf:"bytes,2,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// Datacenter is the datacenter for where this destination service lives.
	Datacenter string `protobuf:"bytes,3,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	// Config is the configuration that should apply to this destination.
	Config *DestinationConfiguration `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *DestinationConfigOverrides) Reset() {
	*x = DestinationConfigOverrides{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationConfigOverrides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationConfigOverrides) ProtoMessage() {}

func (x *DestinationConfigOverrides) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationConfigOverrides.ProtoReflect.Descriptor instead.
func (*DestinationConfigOverrides) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_destinations_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *DestinationConfigOverrides) GetDestinationRef() *pbresource.Reference {
	if x != nil {
		return x.DestinationRef
	}
	return nil
}

func (x *DestinationConfigOverrides) GetDestinationPort() string {
	if x != nil {
		return x.DestinationPort
	}
	return ""
}

func (x *DestinationConfigOverrides) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *DestinationConfigOverrides) GetConfig() *DestinationConfiguration {
	if x != nil {
		return x.Config
	}
	return nil
}

type DestinationConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Protocol overrides destination's port protocol. If no port for an destination is specified
	// or if used in the default configuration, this protocol will be used for all ports
	// or for all ports of all destinations respectively.
	Protocol v2beta1.Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=hashicorp.consul.catalog.v2beta1.Protocol" json:"protocol,omitempty"`
	// ConnectTimeout is the timeout used when making a new
	// connection to this destination. Defaults to 5 seconds if not set.
	ConnectTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
	// Limits are the set of limits that are applied to the proxy for a specific destination.
	Limits *DestinationLimits `protobuf:"bytes,3,opt,name=limits,proto3" json:"limits,omitempty"`
	// PassiveHealthCheck configuration determines how destination proxy instances will
	// be monitored for removal from the load balancing pool.
	PassiveHealthCheck *PassiveHealthCheck `protobuf:"bytes,4,opt,name=passive_health_check,json=passiveHealthCheck,proto3" json:"passive_health_check,omitempty"`
	// BalanceOutboundConnections indicates how the proxy should attempt to distribute
	// connections across worker threads.
	BalanceOutboundConnections BalanceConnections `protobuf:"varint,5,opt,name=balance_outbound_connections,json=balanceOutboundConnections,proto3,enum=hashicorp.consul.mesh.v2beta1.BalanceConnections" json:"balance_outbound_connections,omitempty"`
	// MeshGatewayMode is the Mesh Gateway routing mode.
	MeshGatewayMode MeshGatewayMode `protobuf:"varint,6,opt,name=mesh_gateway_mode,json=meshGatewayMode,proto3,enum=hashicorp.consul.mesh.v2beta1.MeshGatewayMode" json:"mesh_gateway_mode,omitempty"`
}

func (x *DestinationConfiguration) Reset() {
	*x = DestinationConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationConfiguration) ProtoMessage() {}

func (x *DestinationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationConfiguration.ProtoReflect.Descriptor instead.
func (*DestinationConfiguration) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_destinations_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *DestinationConfiguration) GetProtocol() v2beta1.Protocol {
	if x != nil {
		return x.Protocol
	}
	return v2beta1.Protocol(0)
}

func (x *DestinationConfiguration) GetConnectTimeout() *durationpb.Duration {
	if x != nil {
		return x.ConnectTimeout
	}
	return nil
}

func (x *DestinationConfiguration) GetLimits() *DestinationLimits {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *DestinationConfiguration) GetPassiveHealthCheck() *PassiveHealthCheck {
	if x != nil {
		return x.PassiveHealthCheck
	}
	return nil
}

func (x *DestinationConfiguration) GetBalanceOutboundConnections() BalanceConnections {
	if x != nil {
		return x.BalanceOutboundConnections
	}
	return BalanceConnections_BALANCE_CONNECTIONS_DEFAULT
}

func (x *DestinationConfiguration) GetMeshGatewayMode() MeshGatewayMode {
	if x != nil {
		return x.MeshGatewayMode
	}
	return MeshGatewayMode_MESH_GATEWAY_MODE_UNSPECIFIED
}

// UpstreamLimits describes the limits that are associated with a specific
// destination of a service instance.
type DestinationLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MaxConnections is the maximum number of connections the local proxy can
	// make to the destination service.
	MaxConnections *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// MaxPendingRequests is the maximum number of requests that will be queued
	// waiting for an available connection. This is mostly applicable to HTTP/1.1
	// clusters since all HTTP/2 requests are streamed over a single
	// connection.
	MaxPendingRequests *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// MaxConcurrentRequests is the maximum number of in-flight requests that will be allowed
	// to the destination cluster at a point in time. This is mostly applicable to HTTP/2
	// clusters since all HTTP/1.1 requests are limited by MaxConnections.
	MaxConcurrentRequests *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_concurrent_requests,json=maxConcurrentRequests,proto3" json:"max_concurrent_requests,omitempty"`
}

func (x *DestinationLimits) Reset() {
	*x = DestinationLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationLimits) ProtoMessage() {}

func (x *DestinationLimits) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationLimits.ProtoReflect.Descriptor instead.
func (*DestinationLimits) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_destinations_configuration_proto_rawDescGZIP(), []int{3}
}

func (x *DestinationLimits) GetMaxConnections() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxConnections
	}
	return nil
}

func (x *DestinationLimits) GetMaxPendingRequests() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxPendingRequests
	}
	return nil
}

func (x *DestinationLimits) GetMaxConcurrentRequests() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxConcurrentRequests
	}
	return nil
}

type PassiveHealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interval between health check analysis sweeps. Each sweep may remove
	// hosts or return hosts to the pool.
	Interval *durationpb.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// MaxFailures is the count of consecutive failures that results in a host
	// being removed from the pool.
	MaxFailures uint32 `protobuf:"varint,2,opt,name=max_failures,json=maxFailures,proto3" json:"max_failures,omitempty"`
	// EnforcingConsecutive5XX is the % chance that a host will be actually ejected
	// when an outlier status is detected through consecutive 5xx.
	// This setting can be used to disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx uint32 `protobuf:"varint,3,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
}

func (x *PassiveHealthCheck) Reset() {
	*x = PassiveHealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassiveHealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassiveHealthCheck) ProtoMessage() {}

func (x *PassiveHealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassiveHealthCheck.ProtoReflect.Descriptor instead.
func (*PassiveHealthCheck) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_destinations_configuration_proto_rawDescGZIP(), []int{4}
}

func (x *PassiveHealthCheck) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *PassiveHealthCheck) GetMaxFailures() uint32 {
	if x != nil {
		return x.MaxFailures
	}
	return 0
}

func (x *PassiveHealthCheck) GetEnforcingConsecutive_5Xx() uint32 {
	if x != nil {
		return x.EnforcingConsecutive_5Xx
	}
	return 0
}

var File_pbmesh_v2beta1_destinations_configuration_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_destinations_configuration_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a,
	0x19, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x5e, 0x0a, 0x0e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x64, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x73, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x73, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x22, 0x87, 0x02, 0x0a, 0x1a, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0xa6, 0x04, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x46, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x42, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x48, 0x0a,
	0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52,
	0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x63, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x69,
	0x76, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x69, 0x76,
	0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x73, 0x0a, 0x1c,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x1a, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x75,
	0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x5a, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x68, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x6d, 0x65,
	0x73, 0x68, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x80, 0x02,
	0x0a, 0x11, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x14, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x54, 0x0a, 0x17, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x43, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x22, 0xaa, 0x01, 0x0a, 0x12, 0x50, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x73, 0x12, 0x3a, 0x0a, 0x19, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x35, 0x78, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x35, 0x78, 0x78, 0x42, 0x9f, 0x02,
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x1e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d,
	0x65, 0x73, 0x68, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d,
	0xaa, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x29, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_destinations_configuration_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_destinations_configuration_proto_rawDescData = file_pbmesh_v2beta1_destinations_configuration_proto_rawDesc
)

func file_pbmesh_v2beta1_destinations_configuration_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_destinations_configuration_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_destinations_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_destinations_configuration_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_destinations_configuration_proto_rawDescData
}

var file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pbmesh_v2beta1_destinations_configuration_proto_goTypes = []interface{}{
	(*DestinationsConfiguration)(nil),  // 0: hashicorp.consul.mesh.v2beta1.DestinationsConfiguration
	(*DestinationConfigOverrides)(nil), // 1: hashicorp.consul.mesh.v2beta1.DestinationConfigOverrides
	(*DestinationConfiguration)(nil),   // 2: hashicorp.consul.mesh.v2beta1.DestinationConfiguration
	(*DestinationLimits)(nil),          // 3: hashicorp.consul.mesh.v2beta1.DestinationLimits
	(*PassiveHealthCheck)(nil),         // 4: hashicorp.consul.mesh.v2beta1.PassiveHealthCheck
	(*v2beta1.WorkloadSelector)(nil),   // 5: hashicorp.consul.catalog.v2beta1.WorkloadSelector
	(*pbresource.Reference)(nil),       // 6: hashicorp.consul.resource.Reference
	(v2beta1.Protocol)(0),              // 7: hashicorp.consul.catalog.v2beta1.Protocol
	(*durationpb.Duration)(nil),        // 8: google.protobuf.Duration
	(BalanceConnections)(0),            // 9: hashicorp.consul.mesh.v2beta1.BalanceConnections
	(MeshGatewayMode)(0),               // 10: hashicorp.consul.mesh.v2beta1.MeshGatewayMode
	(*wrapperspb.UInt32Value)(nil),     // 11: google.protobuf.UInt32Value
}
var file_pbmesh_v2beta1_destinations_configuration_proto_depIdxs = []int32{
	5,  // 0: hashicorp.consul.mesh.v2beta1.DestinationsConfiguration.workloads:type_name -> hashicorp.consul.catalog.v2beta1.WorkloadSelector
	2,  // 1: hashicorp.consul.mesh.v2beta1.DestinationsConfiguration.default_config:type_name -> hashicorp.consul.mesh.v2beta1.DestinationConfiguration
	1,  // 2: hashicorp.consul.mesh.v2beta1.DestinationsConfiguration.config_overrides:type_name -> hashicorp.consul.mesh.v2beta1.DestinationConfigOverrides
	6,  // 3: hashicorp.consul.mesh.v2beta1.DestinationConfigOverrides.destination_ref:type_name -> hashicorp.consul.resource.Reference
	2,  // 4: hashicorp.consul.mesh.v2beta1.DestinationConfigOverrides.config:type_name -> hashicorp.consul.mesh.v2beta1.DestinationConfiguration
	7,  // 5: hashicorp.consul.mesh.v2beta1.DestinationConfiguration.protocol:type_name -> hashicorp.consul.catalog.v2beta1.Protocol
	8,  // 6: hashicorp.consul.mesh.v2beta1.DestinationConfiguration.connect_timeout:type_name -> google.protobuf.Duration
	3,  // 7: hashicorp.consul.mesh.v2beta1.DestinationConfiguration.limits:type_name -> hashicorp.consul.mesh.v2beta1.DestinationLimits
	4,  // 8: hashicorp.consul.mesh.v2beta1.DestinationConfiguration.passive_health_check:type_name -> hashicorp.consul.mesh.v2beta1.PassiveHealthCheck
	9,  // 9: hashicorp.consul.mesh.v2beta1.DestinationConfiguration.balance_outbound_connections:type_name -> hashicorp.consul.mesh.v2beta1.BalanceConnections
	10, // 10: hashicorp.consul.mesh.v2beta1.DestinationConfiguration.mesh_gateway_mode:type_name -> hashicorp.consul.mesh.v2beta1.MeshGatewayMode
	11, // 11: hashicorp.consul.mesh.v2beta1.DestinationLimits.max_connections:type_name -> google.protobuf.UInt32Value
	11, // 12: hashicorp.consul.mesh.v2beta1.DestinationLimits.max_pending_requests:type_name -> google.protobuf.UInt32Value
	11, // 13: hashicorp.consul.mesh.v2beta1.DestinationLimits.max_concurrent_requests:type_name -> google.protobuf.UInt32Value
	8,  // 14: hashicorp.consul.mesh.v2beta1.PassiveHealthCheck.interval:type_name -> google.protobuf.Duration
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_destinations_configuration_proto_init() }
func file_pbmesh_v2beta1_destinations_configuration_proto_init() {
	if File_pbmesh_v2beta1_destinations_configuration_proto != nil {
		return
	}
	file_pbmesh_v2beta1_connection_proto_init()
	file_pbmesh_v2beta1_routing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationsConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationConfigOverrides); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationLimits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassiveHealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmesh_v2beta1_destinations_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_destinations_configuration_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_destinations_configuration_proto_depIdxs,
		MessageInfos:      file_pbmesh_v2beta1_destinations_configuration_proto_msgTypes,
	}.Build()
	File_pbmesh_v2beta1_destinations_configuration_proto = out.File
	file_pbmesh_v2beta1_destinations_configuration_proto_rawDesc = nil
	file_pbmesh_v2beta1_destinations_configuration_proto_goTypes = nil
	file_pbmesh_v2beta1_destinations_configuration_proto_depIdxs = nil
}
