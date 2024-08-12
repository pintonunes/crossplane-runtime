//
//Copyright 2024 The Crossplane Authors.
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: changelogs/proto/v1alpha1/changelog.proto

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// OperationType represents the type of operation that was performed on a
// resource.
type OperationType int32

const (
	OperationType_OPERATION_TYPE_UNSPECIFIED OperationType = 0
	OperationType_OPERATION_TYPE_CREATE      OperationType = 1
	OperationType_OPERATION_TYPE_UPDATE      OperationType = 2
	OperationType_OPERATION_TYPE_DELETE      OperationType = 3
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_TYPE_UNSPECIFIED",
		1: "OPERATION_TYPE_CREATE",
		2: "OPERATION_TYPE_UPDATE",
		3: "OPERATION_TYPE_DELETE",
	}
	OperationType_value = map[string]int32{
		"OPERATION_TYPE_UNSPECIFIED": 0,
		"OPERATION_TYPE_CREATE":      1,
		"OPERATION_TYPE_UPDATE":      2,
		"OPERATION_TYPE_DELETE":      3,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_changelogs_proto_v1alpha1_changelog_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_changelogs_proto_v1alpha1_changelog_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_changelogs_proto_v1alpha1_changelog_proto_rawDescGZIP(), []int{0}
}

// SendChangeLogRequest represents a request to send a single change log entry.
type SendChangeLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The change log entry to send as part of this request.
	Entry *ChangeLogEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *SendChangeLogRequest) Reset() {
	*x = SendChangeLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChangeLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChangeLogRequest) ProtoMessage() {}

func (x *SendChangeLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChangeLogRequest.ProtoReflect.Descriptor instead.
func (*SendChangeLogRequest) Descriptor() ([]byte, []int) {
	return file_changelogs_proto_v1alpha1_changelog_proto_rawDescGZIP(), []int{0}
}

func (x *SendChangeLogRequest) GetEntry() *ChangeLogEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

// ChangeLogEntry represents a single change log entry, with detailed information
// about the resource that was changed.
type ChangeLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp at which the change occurred.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The name and version of the provider that is making the change to the
	// resource.
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	// The API version of the resource that was changed, e.g. Group/Version.
	ApiVersion string `protobuf:"bytes,3,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The kind of the resource that was changed.
	Kind string `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	// The name of the resource that was changed.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// The external name of the resource that was changed.
	ExternalName string `protobuf:"bytes,6,opt,name=external_name,json=externalName,proto3" json:"external_name,omitempty"`
	// The type of operation that was performed on the resource, e.g. Create,
	// Update, or Delete.
	Operation OperationType `protobuf:"varint,7,opt,name=operation,proto3,enum=changelogs.proto.v1alpha1.OperationType" json:"operation,omitempty"`
	// A full snapshot of the resource's state, as observed directly before the
	// resource was changed.
	Snapshot *structpb.Struct `protobuf:"bytes,8,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	// An optional error message that describes any error encountered while
	// performing the operation on the resource.
	ErrorMessage *string `protobuf:"bytes,9,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
	// An optional additional details that can be provided for further context
	// about the change.
	AdditionalDetails map[string]string `protobuf:"bytes,10,rep,name=additional_details,json=additionalDetails,proto3" json:"additional_details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChangeLogEntry) Reset() {
	*x = ChangeLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLogEntry) ProtoMessage() {}

func (x *ChangeLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLogEntry.ProtoReflect.Descriptor instead.
func (*ChangeLogEntry) Descriptor() ([]byte, []int) {
	return file_changelogs_proto_v1alpha1_changelog_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeLogEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ChangeLogEntry) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ChangeLogEntry) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ChangeLogEntry) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ChangeLogEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChangeLogEntry) GetExternalName() string {
	if x != nil {
		return x.ExternalName
	}
	return ""
}

func (x *ChangeLogEntry) GetOperation() OperationType {
	if x != nil {
		return x.Operation
	}
	return OperationType_OPERATION_TYPE_UNSPECIFIED
}

func (x *ChangeLogEntry) GetSnapshot() *structpb.Struct {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

func (x *ChangeLogEntry) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *ChangeLogEntry) GetAdditionalDetails() map[string]string {
	if x != nil {
		return x.AdditionalDetails
	}
	return nil
}

// SendChangeLogResponse is the response returned by the ChangeLogService after
// a change log entry is sent. Currently, this is an empty message as the only
// useful information expected to sent back at this time will be through errors.
type SendChangeLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendChangeLogResponse) Reset() {
	*x = SendChangeLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChangeLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChangeLogResponse) ProtoMessage() {}

func (x *SendChangeLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChangeLogResponse.ProtoReflect.Descriptor instead.
func (*SendChangeLogResponse) Descriptor() ([]byte, []int) {
	return file_changelogs_proto_v1alpha1_changelog_proto_rawDescGZIP(), []int{2}
}

var File_changelogs_proto_v1alpha1_changelog_proto protoreflect.FileDescriptor

var file_changelogs_proto_v1alpha1_changelog_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xc4,
	0x04, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x6f, 0x0a, 0x12,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x44, 0x0a,
	0x16, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x80,
	0x01, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x03, 0x32, 0x88, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x2f, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x49, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_changelogs_proto_v1alpha1_changelog_proto_rawDescOnce sync.Once
	file_changelogs_proto_v1alpha1_changelog_proto_rawDescData = file_changelogs_proto_v1alpha1_changelog_proto_rawDesc
)

func file_changelogs_proto_v1alpha1_changelog_proto_rawDescGZIP() []byte {
	file_changelogs_proto_v1alpha1_changelog_proto_rawDescOnce.Do(func() {
		file_changelogs_proto_v1alpha1_changelog_proto_rawDescData = protoimpl.X.CompressGZIP(file_changelogs_proto_v1alpha1_changelog_proto_rawDescData)
	})
	return file_changelogs_proto_v1alpha1_changelog_proto_rawDescData
}

var file_changelogs_proto_v1alpha1_changelog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_changelogs_proto_v1alpha1_changelog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_changelogs_proto_v1alpha1_changelog_proto_goTypes = []any{
	(OperationType)(0),            // 0: changelogs.proto.v1alpha1.OperationType
	(*SendChangeLogRequest)(nil),  // 1: changelogs.proto.v1alpha1.SendChangeLogRequest
	(*ChangeLogEntry)(nil),        // 2: changelogs.proto.v1alpha1.ChangeLogEntry
	(*SendChangeLogResponse)(nil), // 3: changelogs.proto.v1alpha1.SendChangeLogResponse
	nil,                           // 4: changelogs.proto.v1alpha1.ChangeLogEntry.AdditionalDetailsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 6: google.protobuf.Struct
}
var file_changelogs_proto_v1alpha1_changelog_proto_depIdxs = []int32{
	2, // 0: changelogs.proto.v1alpha1.SendChangeLogRequest.entry:type_name -> changelogs.proto.v1alpha1.ChangeLogEntry
	5, // 1: changelogs.proto.v1alpha1.ChangeLogEntry.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: changelogs.proto.v1alpha1.ChangeLogEntry.operation:type_name -> changelogs.proto.v1alpha1.OperationType
	6, // 3: changelogs.proto.v1alpha1.ChangeLogEntry.snapshot:type_name -> google.protobuf.Struct
	4, // 4: changelogs.proto.v1alpha1.ChangeLogEntry.additional_details:type_name -> changelogs.proto.v1alpha1.ChangeLogEntry.AdditionalDetailsEntry
	1, // 5: changelogs.proto.v1alpha1.ChangeLogService.SendChangeLog:input_type -> changelogs.proto.v1alpha1.SendChangeLogRequest
	3, // 6: changelogs.proto.v1alpha1.ChangeLogService.SendChangeLog:output_type -> changelogs.proto.v1alpha1.SendChangeLogResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_changelogs_proto_v1alpha1_changelog_proto_init() }
func file_changelogs_proto_v1alpha1_changelog_proto_init() {
	if File_changelogs_proto_v1alpha1_changelog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendChangeLogRequest); i {
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
		file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeLogEntry); i {
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
		file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SendChangeLogResponse); i {
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
	file_changelogs_proto_v1alpha1_changelog_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_changelogs_proto_v1alpha1_changelog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_changelogs_proto_v1alpha1_changelog_proto_goTypes,
		DependencyIndexes: file_changelogs_proto_v1alpha1_changelog_proto_depIdxs,
		EnumInfos:         file_changelogs_proto_v1alpha1_changelog_proto_enumTypes,
		MessageInfos:      file_changelogs_proto_v1alpha1_changelog_proto_msgTypes,
	}.Build()
	File_changelogs_proto_v1alpha1_changelog_proto = out.File
	file_changelogs_proto_v1alpha1_changelog_proto_rawDesc = nil
	file_changelogs_proto_v1alpha1_changelog_proto_goTypes = nil
	file_changelogs_proto_v1alpha1_changelog_proto_depIdxs = nil
}
