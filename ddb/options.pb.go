// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ddb/options.proto

package ddb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OneofOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OneofOptions) Reset() {
	*x = OneofOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofOptions) ProtoMessage() {}

func (x *OneofOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofOptions.ProtoReflect.Descriptor instead.
func (*OneofOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{0}
}

type EnumOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnumOptions) Reset() {
	*x = EnumOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumOptions) ProtoMessage() {}

func (x *EnumOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumOptions.ProtoReflect.Descriptor instead.
func (*EnumOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{1}
}

type EnumValueOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnumValueOptions) Reset() {
	*x = EnumValueOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumValueOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumValueOptions) ProtoMessage() {}

func (x *EnumValueOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumValueOptions.ProtoReflect.Descriptor instead.
func (*EnumValueOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{2}
}

type ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceOptions) Reset() {
	*x = ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOptions) ProtoMessage() {}

func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{3}
}

type MethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethodOptions) Reset() {
	*x = MethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptions) ProtoMessage() {}

func (x *MethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptions.ProtoReflect.Descriptor instead.
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{4}
}

type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{5}
}

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set this to true to indicate that this message represents a DynamoDB document
	// and protoc-gen-dynamodb should generate code for you.
	Document bool `protobuf:"varint,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{6}
}

func (x *MessageOptions) GetDocument() bool {
	if x != nil {
		return x.Document
	}
	return false
}

type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddb_options_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ddb_options_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_ddb_options_proto_rawDescGZIP(), []int{7}
}

var file_ddb_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         50300,
		Name:          "ddb.file",
		Tag:           "bytes,50300,opt,name=file",
		Filename:      "ddb/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         50301,
		Name:          "ddb.message",
		Tag:           "bytes,50301,opt,name=message",
		Filename:      "ddb/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         50302,
		Name:          "ddb.field",
		Tag:           "bytes,50302,opt,name=field",
		Filename:      "ddb/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*OneofOptions)(nil),
		Field:         50303,
		Name:          "ddb.oneof",
		Tag:           "bytes,50303,opt,name=oneof",
		Filename:      "ddb/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*EnumOptions)(nil),
		Field:         50304,
		Name:          "ddb.enum",
		Tag:           "bytes,50304,opt,name=enum",
		Filename:      "ddb/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*EnumValueOptions)(nil),
		Field:         50305,
		Name:          "ddb.enum_value",
		Tag:           "bytes,50305,opt,name=enum_value",
		Filename:      "ddb/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceOptions)(nil),
		Field:         50306,
		Name:          "ddb.service",
		Tag:           "bytes,50306,opt,name=service",
		Filename:      "ddb/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptions)(nil),
		Field:         50307,
		Name:          "ddb.method",
		Tag:           "bytes,50307,opt,name=method",
		Filename:      "ddb/options.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional ddb.FileOptions file = 50300;
	E_File = &file_ddb_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional ddb.MessageOptions message = 50301;
	E_Message = &file_ddb_options_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional ddb.FieldOptions field = 50302;
	E_Field = &file_ddb_options_proto_extTypes[2]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional ddb.OneofOptions oneof = 50303;
	E_Oneof = &file_ddb_options_proto_extTypes[3]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional ddb.EnumOptions enum = 50304;
	E_Enum = &file_ddb_options_proto_extTypes[4]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional ddb.EnumValueOptions enum_value = 50305;
	E_EnumValue = &file_ddb_options_proto_extTypes[5]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional ddb.ServiceOptions service = 50306;
	E_Service = &file_ddb_options_proto_extTypes[6]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional ddb.MethodOptions method = 50307;
	E_Method = &file_ddb_options_proto_extTypes[7]
)

var File_ddb_options_proto protoreflect.FileDescriptor

var file_ddb_options_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x64, 0x62, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x64, 0x62, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x4f, 0x6e,
	0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6e,
	0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x45, 0x6e, 0x75,
	0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x10, 0x0a,
	0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x0f, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x2c, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x0e, 0x0a,
	0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x44, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xfc, 0x88, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x64,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x3a, 0x50, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xfd, 0x88, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x64, 0x62, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x48, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0x88,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x64, 0x62, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a,
	0x48, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xff, 0x88, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x64, 0x64, 0x62, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x3a, 0x44, 0x0a, 0x04, 0x65, 0x6e, 0x75,
	0x6d, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x80, 0x89, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x64, 0x62, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x3a,
	0x59, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x81, 0x89, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x64, 0x62, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x09, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x50, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x82, 0x89, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x64, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x4c, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x83, 0x89, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x64, 0x64, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x36, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x2f, 0x64, 0x64, 0x62, 0x3b, 0x64, 0x64, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ddb_options_proto_rawDescOnce sync.Once
	file_ddb_options_proto_rawDescData = file_ddb_options_proto_rawDesc
)

func file_ddb_options_proto_rawDescGZIP() []byte {
	file_ddb_options_proto_rawDescOnce.Do(func() {
		file_ddb_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_ddb_options_proto_rawDescData)
	})
	return file_ddb_options_proto_rawDescData
}

var file_ddb_options_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ddb_options_proto_goTypes = []interface{}{
	(*OneofOptions)(nil),                  // 0: ddb.OneofOptions
	(*EnumOptions)(nil),                   // 1: ddb.EnumOptions
	(*EnumValueOptions)(nil),              // 2: ddb.EnumValueOptions
	(*ServiceOptions)(nil),                // 3: ddb.ServiceOptions
	(*MethodOptions)(nil),                 // 4: ddb.MethodOptions
	(*FileOptions)(nil),                   // 5: ddb.FileOptions
	(*MessageOptions)(nil),                // 6: ddb.MessageOptions
	(*FieldOptions)(nil),                  // 7: ddb.FieldOptions
	(*descriptorpb.FileOptions)(nil),      // 8: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil),   // 9: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),     // 10: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil),     // 11: google.protobuf.OneofOptions
	(*descriptorpb.EnumOptions)(nil),      // 12: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 13: google.protobuf.EnumValueOptions
	(*descriptorpb.ServiceOptions)(nil),   // 14: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),    // 15: google.protobuf.MethodOptions
}
var file_ddb_options_proto_depIdxs = []int32{
	8,  // 0: ddb.file:extendee -> google.protobuf.FileOptions
	9,  // 1: ddb.message:extendee -> google.protobuf.MessageOptions
	10, // 2: ddb.field:extendee -> google.protobuf.FieldOptions
	11, // 3: ddb.oneof:extendee -> google.protobuf.OneofOptions
	12, // 4: ddb.enum:extendee -> google.protobuf.EnumOptions
	13, // 5: ddb.enum_value:extendee -> google.protobuf.EnumValueOptions
	14, // 6: ddb.service:extendee -> google.protobuf.ServiceOptions
	15, // 7: ddb.method:extendee -> google.protobuf.MethodOptions
	5,  // 8: ddb.file:type_name -> ddb.FileOptions
	6,  // 9: ddb.message:type_name -> ddb.MessageOptions
	7,  // 10: ddb.field:type_name -> ddb.FieldOptions
	0,  // 11: ddb.oneof:type_name -> ddb.OneofOptions
	1,  // 12: ddb.enum:type_name -> ddb.EnumOptions
	2,  // 13: ddb.enum_value:type_name -> ddb.EnumValueOptions
	3,  // 14: ddb.service:type_name -> ddb.ServiceOptions
	4,  // 15: ddb.method:type_name -> ddb.MethodOptions
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	8,  // [8:16] is the sub-list for extension type_name
	0,  // [0:8] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ddb_options_proto_init() }
func file_ddb_options_proto_init() {
	if File_ddb_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ddb_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofOptions); i {
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
		file_ddb_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumOptions); i {
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
		file_ddb_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumValueOptions); i {
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
		file_ddb_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOptions); i {
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
		file_ddb_options_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptions); i {
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
		file_ddb_options_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOptions); i {
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
		file_ddb_options_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_ddb_options_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
			RawDescriptor: file_ddb_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 8,
			NumServices:   0,
		},
		GoTypes:           file_ddb_options_proto_goTypes,
		DependencyIndexes: file_ddb_options_proto_depIdxs,
		MessageInfos:      file_ddb_options_proto_msgTypes,
		ExtensionInfos:    file_ddb_options_proto_extTypes,
	}.Build()
	File_ddb_options_proto = out.File
	file_ddb_options_proto_rawDesc = nil
	file_ddb_options_proto_goTypes = nil
	file_ddb_options_proto_depIdxs = nil
}
