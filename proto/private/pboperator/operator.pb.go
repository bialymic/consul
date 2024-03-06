// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: private/pboperator/operator.proto

package pboperator

import (
	_ "github.com/hashicorp/consul/proto-public/annotations/ratelimit"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransferLeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TransferLeaderRequest) Reset() {
	*x = TransferLeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pboperator_operator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferLeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferLeaderRequest) ProtoMessage() {}

func (x *TransferLeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_pboperator_operator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferLeaderRequest.ProtoReflect.Descriptor instead.
func (*TransferLeaderRequest) Descriptor() ([]byte, []int) {
	return file_private_pboperator_operator_proto_rawDescGZIP(), []int{0}
}

func (x *TransferLeaderRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// mog annotation:
//
// target=github.com/hashicorp/consul/api.TransferLeaderResponse
// output=operator.gen.go
// name=API
type TransferLeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// true if the transfer is a success
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TransferLeaderResponse) Reset() {
	*x = TransferLeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pboperator_operator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferLeaderResponse) ProtoMessage() {}

func (x *TransferLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_pboperator_operator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferLeaderResponse.ProtoReflect.Descriptor instead.
func (*TransferLeaderResponse) Descriptor() ([]byte, []int) {
	return file_private_pboperator_operator_proto_rawDescGZIP(), []int{1}
}

func (x *TransferLeaderResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_private_pboperator_operator_proto protoreflect.FileDescriptor

var file_private_pboperator_operator_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x25, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27,
	0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x0f,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x91, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x39, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xe2, 0x86, 0x04, 0x04, 0x08,
	0x01, 0x10, 0x0a, 0x42, 0x99, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0d,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0xa2, 0x02, 0x04, 0x48, 0x43, 0x49, 0x4f, 0xaa, 0x02, 0x22, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0xca, 0x02, 0x22, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0xe2, 0x02, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x3a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_private_pboperator_operator_proto_rawDescOnce sync.Once
	file_private_pboperator_operator_proto_rawDescData = file_private_pboperator_operator_proto_rawDesc
)

func file_private_pboperator_operator_proto_rawDescGZIP() []byte {
	file_private_pboperator_operator_proto_rawDescOnce.Do(func() {
		file_private_pboperator_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_pboperator_operator_proto_rawDescData)
	})
	return file_private_pboperator_operator_proto_rawDescData
}

var file_private_pboperator_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_private_pboperator_operator_proto_goTypes = []interface{}{
	(*TransferLeaderRequest)(nil),  // 0: hashicorp.consul.internal.operator.TransferLeaderRequest
	(*TransferLeaderResponse)(nil), // 1: hashicorp.consul.internal.operator.TransferLeaderResponse
}
var file_private_pboperator_operator_proto_depIdxs = []int32{
	0, // 0: hashicorp.consul.internal.operator.OperatorService.TransferLeader:input_type -> hashicorp.consul.internal.operator.TransferLeaderRequest
	1, // 1: hashicorp.consul.internal.operator.OperatorService.TransferLeader:output_type -> hashicorp.consul.internal.operator.TransferLeaderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_private_pboperator_operator_proto_init() }
func file_private_pboperator_operator_proto_init() {
	if File_private_pboperator_operator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_private_pboperator_operator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferLeaderRequest); i {
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
		file_private_pboperator_operator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferLeaderResponse); i {
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
			RawDescriptor: file_private_pboperator_operator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_private_pboperator_operator_proto_goTypes,
		DependencyIndexes: file_private_pboperator_operator_proto_depIdxs,
		MessageInfos:      file_private_pboperator_operator_proto_msgTypes,
	}.Build()
	File_private_pboperator_operator_proto = out.File
	file_private_pboperator_operator_proto_rawDesc = nil
	file_private_pboperator_operator_proto_goTypes = nil
	file_private_pboperator_operator_proto_depIdxs = nil
}
