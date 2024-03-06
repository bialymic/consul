// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: private/pbconnect/connect.proto

package pbconnect

import (
	pbcommon "github.com/hashicorp/consul/proto/private/pbcommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CARoots is the list of all currently trusted CA Roots.
//
// mog annotation:
//
// target=github.com/hashicorp/consul/agent/structs.IndexedCARoots
// output=connect.gen.go
// name=StructsIndexedCARoots
type CARoots struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ActiveRootID is the ID of a root in Roots that is the active CA root.
	// Other roots are still valid if they're in the Roots list but are in
	// the process of being rotated out.
	ActiveRootID string `protobuf:"bytes,1,opt,name=ActiveRootID,proto3" json:"ActiveRootID,omitempty"`
	// TrustDomain is the identification root for this Consul cluster. All
	// certificates signed by the cluster's CA must have their identifying URI in
	// this domain.
	//
	// This does not include the protocol (currently spiffe://) since we may
	// implement other protocols in future with equivalent semantics. It should be
	// compared against the "authority" section of a URI (i.e. host:port).
	//
	// We need to support migrating a cluster between trust domains to support
	// Multi-DC migration in Enterprise. In this case the current trust domain is
	// here but entries in Roots may also have ExternalTrustDomain set to a
	// non-empty value implying they were previous roots that are still trusted
	// but under a different trust domain.
	//
	// Note that we DON'T validate trust domain during AuthZ since it causes
	// issues of loss of connectivity during migration between trust domains. The
	// only time the additional validation adds value is where the cluster shares
	// an external root (e.g. organization-wide root) with another distinct Consul
	// cluster or PKI system. In this case, x509 Name Constraints can be added to
	// enforce that Consul's CA can only validly sign or trust certs within the
	// same trust-domain. Name constraints as enforced by TLS handshake also allow
	// seamless rotation between trust domains thanks to cross-signing.
	TrustDomain string `protobuf:"bytes,2,opt,name=TrustDomain,proto3" json:"TrustDomain,omitempty"`
	// Roots is a list of root CA certs to trust.
	Roots []*CARoot `protobuf:"bytes,3,rep,name=Roots,proto3" json:"Roots,omitempty"`
	// QueryMeta here is mainly used to contain the latest Raft Index that could
	// be used to perform a blocking query.
	// mog: func-to=QueryMetaTo func-from=QueryMetaFrom
	QueryMeta *pbcommon.QueryMeta `protobuf:"bytes,4,opt,name=QueryMeta,proto3" json:"QueryMeta,omitempty"`
}

func (x *CARoots) Reset() {
	*x = CARoots{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconnect_connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CARoots) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CARoots) ProtoMessage() {}

func (x *CARoots) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconnect_connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CARoots.ProtoReflect.Descriptor instead.
func (*CARoots) Descriptor() ([]byte, []int) {
	return file_private_pbconnect_connect_proto_rawDescGZIP(), []int{0}
}

func (x *CARoots) GetActiveRootID() string {
	if x != nil {
		return x.ActiveRootID
	}
	return ""
}

func (x *CARoots) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

func (x *CARoots) GetRoots() []*CARoot {
	if x != nil {
		return x.Roots
	}
	return nil
}

func (x *CARoots) GetQueryMeta() *pbcommon.QueryMeta {
	if x != nil {
		return x.QueryMeta
	}
	return nil
}

// CARoot is the trusted CA Root.
//
// mog annotation:
//
// target=github.com/hashicorp/consul/agent/structs.CARoot
// output=connect.gen.go
// name=StructsCARoot
type CARoot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is a globally unique ID (UUID) representing this CA root.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name is a human-friendly name for this CA root. This value is
	// opaque to Consul and is not used for anything internally.
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// SerialNumber is the x509 serial number of the certificate.
	SerialNumber uint64 `protobuf:"varint,3,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	// SigningKeyID is the ID of the public key that corresponds to the private
	// key used to sign leaf certificates. Is is the HexString format of the
	// raw AuthorityKeyID bytes.
	SigningKeyID string `protobuf:"bytes,4,opt,name=SigningKeyID,proto3" json:"SigningKeyID,omitempty"`
	// ExternalTrustDomain is the trust domain this root was generated under. It
	// is usually empty implying "the current cluster trust-domain". It is set
	// only in the case that a cluster changes trust domain and then all old roots
	// that are still trusted have the old trust domain set here.
	//
	// We currently DON'T validate these trust domains explicitly anywhere, see
	// IndexedRoots.TrustDomain doc. We retain this information for debugging and
	// future flexibility.
	ExternalTrustDomain string `protobuf:"bytes,5,opt,name=ExternalTrustDomain,proto3" json:"ExternalTrustDomain,omitempty"`
	// Time validity bounds.
	// mog: func-to=structs.TimeFromProto func-from=structs.TimeToProto
	NotBefore *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=NotBefore,proto3" json:"NotBefore,omitempty"`
	// mog: func-to=structs.TimeFromProto func-from=structs.TimeToProto
	NotAfter *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=NotAfter,proto3" json:"NotAfter,omitempty"`
	// RootCert is the PEM-encoded public certificate.
	RootCert string `protobuf:"bytes,8,opt,name=RootCert,proto3" json:"RootCert,omitempty"`
	// IntermediateCerts is a list of PEM-encoded intermediate certs to
	// attach to any leaf certs signed by this CA.
	IntermediateCerts []string `protobuf:"bytes,9,rep,name=IntermediateCerts,proto3" json:"IntermediateCerts,omitempty"`
	// SigningCert is the PEM-encoded signing certificate and SigningKey
	// is the PEM-encoded private key for the signing certificate. These
	// may actually be empty if the CA plugin in use manages these for us.
	SigningCert string `protobuf:"bytes,10,opt,name=SigningCert,proto3" json:"SigningCert,omitempty"`
	SigningKey  string `protobuf:"bytes,11,opt,name=SigningKey,proto3" json:"SigningKey,omitempty"`
	// Active is true if this is the current active CA. This must only
	// be true for exactly one CA. For any method that modifies roots in the
	// state store, tests should be written to verify that multiple roots
	// cannot be active.
	Active bool `protobuf:"varint,12,opt,name=Active,proto3" json:"Active,omitempty"`
	// RotatedOutAt is the time at which this CA was removed from the state.
	// This will only be set on roots that have been rotated out from being the
	// active root.
	// mog: func-to=structs.TimeFromProto func-from=structs.TimeToProto
	RotatedOutAt *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=RotatedOutAt,proto3" json:"RotatedOutAt,omitempty"`
	// PrivateKeyType is the type of the private key used to sign certificates. It
	// may be "rsa" or "ec". This is provided as a convenience to avoid parsing
	// the public key to from the certificate to infer the type.
	PrivateKeyType string `protobuf:"bytes,14,opt,name=PrivateKeyType,proto3" json:"PrivateKeyType,omitempty"`
	// PrivateKeyBits is the length of the private key used to sign certificates.
	// This is provided as a convenience to avoid parsing the public key from the
	// certificate to infer the type.
	// mog: func-to=int func-from=int32
	PrivateKeyBits int32 `protobuf:"varint,15,opt,name=PrivateKeyBits,proto3" json:"PrivateKeyBits,omitempty"`
	// mog: func-to=RaftIndexTo func-from=RaftIndexFrom
	RaftIndex *pbcommon.RaftIndex `protobuf:"bytes,16,opt,name=RaftIndex,proto3" json:"RaftIndex,omitempty"`
}

func (x *CARoot) Reset() {
	*x = CARoot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconnect_connect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CARoot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CARoot) ProtoMessage() {}

func (x *CARoot) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconnect_connect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CARoot.ProtoReflect.Descriptor instead.
func (*CARoot) Descriptor() ([]byte, []int) {
	return file_private_pbconnect_connect_proto_rawDescGZIP(), []int{1}
}

func (x *CARoot) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CARoot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CARoot) GetSerialNumber() uint64 {
	if x != nil {
		return x.SerialNumber
	}
	return 0
}

func (x *CARoot) GetSigningKeyID() string {
	if x != nil {
		return x.SigningKeyID
	}
	return ""
}

func (x *CARoot) GetExternalTrustDomain() string {
	if x != nil {
		return x.ExternalTrustDomain
	}
	return ""
}

func (x *CARoot) GetNotBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

func (x *CARoot) GetNotAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.NotAfter
	}
	return nil
}

func (x *CARoot) GetRootCert() string {
	if x != nil {
		return x.RootCert
	}
	return ""
}

func (x *CARoot) GetIntermediateCerts() []string {
	if x != nil {
		return x.IntermediateCerts
	}
	return nil
}

func (x *CARoot) GetSigningCert() string {
	if x != nil {
		return x.SigningCert
	}
	return ""
}

func (x *CARoot) GetSigningKey() string {
	if x != nil {
		return x.SigningKey
	}
	return ""
}

func (x *CARoot) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *CARoot) GetRotatedOutAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RotatedOutAt
	}
	return nil
}

func (x *CARoot) GetPrivateKeyType() string {
	if x != nil {
		return x.PrivateKeyType
	}
	return ""
}

func (x *CARoot) GetPrivateKeyBits() int32 {
	if x != nil {
		return x.PrivateKeyBits
	}
	return 0
}

func (x *CARoot) GetRaftIndex() *pbcommon.RaftIndex {
	if x != nil {
		return x.RaftIndex
	}
	return nil
}

// RaftIndex is used to track the index used while creating
// or modifying a given struct type.
//
// mog annotation:
//
// target=github.com/hashicorp/consul/agent/structs.IssuedCert
// output=connect.gen.go
// name=StructsIssuedCert
type IssuedCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SerialNumber is the unique serial number for this certificate.
	// This is encoded in standard hex separated by :.
	SerialNumber string `protobuf:"bytes,1,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	// CertPEM and PrivateKeyPEM are the PEM-encoded certificate and private
	// key for that cert, respectively. This should not be stored in the
	// state store, but is present in the sign API response.
	CertPEM       string `protobuf:"bytes,2,opt,name=CertPEM,proto3" json:"CertPEM,omitempty"`
	PrivateKeyPEM string `protobuf:"bytes,3,opt,name=PrivateKeyPEM,proto3" json:"PrivateKeyPEM,omitempty"`
	// Service is the name of the service for which the cert was issued.
	Service string `protobuf:"bytes,4,opt,name=Service,proto3" json:"Service,omitempty"`
	// ServiceURI is the cert URI value.
	ServiceURI string `protobuf:"bytes,5,opt,name=ServiceURI,proto3" json:"ServiceURI,omitempty"`
	// Agent is the name of the node for which the cert was issued.
	Agent string `protobuf:"bytes,6,opt,name=Agent,proto3" json:"Agent,omitempty"`
	// AgentURI is the cert URI value.
	AgentURI string `protobuf:"bytes,7,opt,name=AgentURI,proto3" json:"AgentURI,omitempty"`
	// Kind is the kind of service for which the cert was issued.
	// mog: func-to=structs.ServiceKind func-from=string
	Kind string `protobuf:"bytes,12,opt,name=Kind,proto3" json:"Kind,omitempty"`
	// KindURI is the cert URI value.
	KindURI string `protobuf:"bytes,13,opt,name=KindURI,proto3" json:"KindURI,omitempty"`
	// ServerURI is the URI value of a cert issued for a server agent.
	// The same URI is shared by all servers in a Consul datacenter.
	ServerURI string `protobuf:"bytes,14,opt,name=ServerURI,proto3" json:"ServerURI,omitempty"`
	// ValidAfter and ValidBefore are the validity periods for the
	// certificate.
	// mog: func-to=structs.TimeFromProto func-from=structs.TimeToProto
	ValidAfter *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=ValidAfter,proto3" json:"ValidAfter,omitempty"`
	// mog: func-to=structs.TimeFromProto func-from=structs.TimeToProto
	ValidBefore *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=ValidBefore,proto3" json:"ValidBefore,omitempty"`
	// EnterpriseMeta is the Consul Enterprise specific metadata
	// mog: func-to=EnterpriseMetaTo func-from=EnterpriseMetaFrom
	EnterpriseMeta *pbcommon.EnterpriseMeta `protobuf:"bytes,10,opt,name=EnterpriseMeta,proto3" json:"EnterpriseMeta,omitempty"`
	// mog: func-to=RaftIndexTo func-from=RaftIndexFrom
	RaftIndex *pbcommon.RaftIndex `protobuf:"bytes,11,opt,name=RaftIndex,proto3" json:"RaftIndex,omitempty"`
}

func (x *IssuedCert) Reset() {
	*x = IssuedCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconnect_connect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuedCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuedCert) ProtoMessage() {}

func (x *IssuedCert) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconnect_connect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuedCert.ProtoReflect.Descriptor instead.
func (*IssuedCert) Descriptor() ([]byte, []int) {
	return file_private_pbconnect_connect_proto_rawDescGZIP(), []int{2}
}

func (x *IssuedCert) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *IssuedCert) GetCertPEM() string {
	if x != nil {
		return x.CertPEM
	}
	return ""
}

func (x *IssuedCert) GetPrivateKeyPEM() string {
	if x != nil {
		return x.PrivateKeyPEM
	}
	return ""
}

func (x *IssuedCert) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *IssuedCert) GetServiceURI() string {
	if x != nil {
		return x.ServiceURI
	}
	return ""
}

func (x *IssuedCert) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *IssuedCert) GetAgentURI() string {
	if x != nil {
		return x.AgentURI
	}
	return ""
}

func (x *IssuedCert) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *IssuedCert) GetKindURI() string {
	if x != nil {
		return x.KindURI
	}
	return ""
}

func (x *IssuedCert) GetServerURI() string {
	if x != nil {
		return x.ServerURI
	}
	return ""
}

func (x *IssuedCert) GetValidAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidAfter
	}
	return nil
}

func (x *IssuedCert) GetValidBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidBefore
	}
	return nil
}

func (x *IssuedCert) GetEnterpriseMeta() *pbcommon.EnterpriseMeta {
	if x != nil {
		return x.EnterpriseMeta
	}
	return nil
}

func (x *IssuedCert) GetRaftIndex() *pbcommon.RaftIndex {
	if x != nil {
		return x.RaftIndex
	}
	return nil
}

var File_private_pbconnect_connect_proto protoreflect.FileDescriptor

var file_private_pbconnect_connect_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x07, 0x43, 0x41, 0x52, 0x6f, 0x6f, 0x74, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3f, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x41, 0x52, 0x6f, 0x6f, 0x74,
	0x52, 0x05, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x49, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x61, 0x22, 0x97, 0x05, 0x0a, 0x06, 0x43, 0x41, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x4b, 0x65, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x4e,
	0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x4e, 0x6f, 0x74, 0x42,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x4e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x43, 0x65, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x41,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x41,
	0x74, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x69, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x69, 0x74,
	0x73, 0x12, 0x49, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x09, 0x52, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xc7, 0x04, 0x0a,
	0x0a, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x65, 0x72, 0x74, 0x50, 0x45, 0x4d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x65, 0x72, 0x74, 0x50, 0x45, 0x4d, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x45, 0x4d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x45, 0x4d, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x52, 0x49, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x52, 0x49, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x49, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x49, 0x12, 0x12, 0x0a, 0x04, 0x4b,
	0x69, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x4b, 0x69, 0x6e, 0x64, 0x55, 0x52, 0x49, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4b, 0x69, 0x6e, 0x64, 0x55, 0x52, 0x49, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x55, 0x52, 0x49, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x55, 0x52, 0x49, 0x12, 0x3a, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x12, 0x58, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0e, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x09, 0x52,
	0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x09, 0x52, 0x61, 0x66,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x92, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x42, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0xa2, 0x02, 0x04, 0x48, 0x43, 0x49, 0x43, 0xaa, 0x02, 0x21, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0xca, 0x02, 0x21, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0xe2, 0x02, 0x2d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_private_pbconnect_connect_proto_rawDescOnce sync.Once
	file_private_pbconnect_connect_proto_rawDescData = file_private_pbconnect_connect_proto_rawDesc
)

func file_private_pbconnect_connect_proto_rawDescGZIP() []byte {
	file_private_pbconnect_connect_proto_rawDescOnce.Do(func() {
		file_private_pbconnect_connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_pbconnect_connect_proto_rawDescData)
	})
	return file_private_pbconnect_connect_proto_rawDescData
}

var file_private_pbconnect_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_private_pbconnect_connect_proto_goTypes = []interface{}{
	(*CARoots)(nil),                 // 0: hashicorp.consul.internal.connect.CARoots
	(*CARoot)(nil),                  // 1: hashicorp.consul.internal.connect.CARoot
	(*IssuedCert)(nil),              // 2: hashicorp.consul.internal.connect.IssuedCert
	(*pbcommon.QueryMeta)(nil),      // 3: hashicorp.consul.internal.common.QueryMeta
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
	(*pbcommon.RaftIndex)(nil),      // 5: hashicorp.consul.internal.common.RaftIndex
	(*pbcommon.EnterpriseMeta)(nil), // 6: hashicorp.consul.internal.common.EnterpriseMeta
}
var file_private_pbconnect_connect_proto_depIdxs = []int32{
	1,  // 0: hashicorp.consul.internal.connect.CARoots.Roots:type_name -> hashicorp.consul.internal.connect.CARoot
	3,  // 1: hashicorp.consul.internal.connect.CARoots.QueryMeta:type_name -> hashicorp.consul.internal.common.QueryMeta
	4,  // 2: hashicorp.consul.internal.connect.CARoot.NotBefore:type_name -> google.protobuf.Timestamp
	4,  // 3: hashicorp.consul.internal.connect.CARoot.NotAfter:type_name -> google.protobuf.Timestamp
	4,  // 4: hashicorp.consul.internal.connect.CARoot.RotatedOutAt:type_name -> google.protobuf.Timestamp
	5,  // 5: hashicorp.consul.internal.connect.CARoot.RaftIndex:type_name -> hashicorp.consul.internal.common.RaftIndex
	4,  // 6: hashicorp.consul.internal.connect.IssuedCert.ValidAfter:type_name -> google.protobuf.Timestamp
	4,  // 7: hashicorp.consul.internal.connect.IssuedCert.ValidBefore:type_name -> google.protobuf.Timestamp
	6,  // 8: hashicorp.consul.internal.connect.IssuedCert.EnterpriseMeta:type_name -> hashicorp.consul.internal.common.EnterpriseMeta
	5,  // 9: hashicorp.consul.internal.connect.IssuedCert.RaftIndex:type_name -> hashicorp.consul.internal.common.RaftIndex
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_private_pbconnect_connect_proto_init() }
func file_private_pbconnect_connect_proto_init() {
	if File_private_pbconnect_connect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_private_pbconnect_connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CARoots); i {
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
		file_private_pbconnect_connect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CARoot); i {
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
		file_private_pbconnect_connect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuedCert); i {
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
			RawDescriptor: file_private_pbconnect_connect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_private_pbconnect_connect_proto_goTypes,
		DependencyIndexes: file_private_pbconnect_connect_proto_depIdxs,
		MessageInfos:      file_private_pbconnect_connect_proto_msgTypes,
	}.Build()
	File_private_pbconnect_connect_proto = out.File
	file_private_pbconnect_connect_proto_rawDesc = nil
	file_private_pbconnect_connect_proto_goTypes = nil
	file_private_pbconnect_connect_proto_depIdxs = nil
}
