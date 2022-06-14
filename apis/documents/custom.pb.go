// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: documents/custom.proto

package documents

import (
	idl_common "github.com/krafton-hq/red-fox/apis/idl_common"
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

type CustomDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string                 `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string                 `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *idl_common.ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	RawSpec    string                 `protobuf:"bytes,4,opt,name=rawSpec,proto3" json:"rawSpec,omitempty"`
}

func (x *CustomDocument) Reset() {
	*x = CustomDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDocument) ProtoMessage() {}

func (x *CustomDocument) ProtoReflect() protoreflect.Message {
	mi := &file_documents_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDocument.ProtoReflect.Descriptor instead.
func (*CustomDocument) Descriptor() ([]byte, []int) {
	return file_documents_custom_proto_rawDescGZIP(), []int{0}
}

func (x *CustomDocument) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *CustomDocument) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CustomDocument) GetMetadata() *idl_common.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CustomDocument) GetRawSpec() string {
	if x != nil {
		return x.RawSpec
	}
	return ""
}

type GetCustomDocumentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonRes *idl_common.CommonRes `protobuf:"bytes,1,opt,name=commonRes,proto3" json:"commonRes,omitempty"`
	Endpoint  *CustomDocument       `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *GetCustomDocumentRes) Reset() {
	*x = GetCustomDocumentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomDocumentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomDocumentRes) ProtoMessage() {}

func (x *GetCustomDocumentRes) ProtoReflect() protoreflect.Message {
	mi := &file_documents_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomDocumentRes.ProtoReflect.Descriptor instead.
func (*GetCustomDocumentRes) Descriptor() ([]byte, []int) {
	return file_documents_custom_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomDocumentRes) GetCommonRes() *idl_common.CommonRes {
	if x != nil {
		return x.CommonRes
	}
	return nil
}

func (x *GetCustomDocumentRes) GetEndpoint() *CustomDocument {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

type ListCustomDocumentsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonRes *idl_common.CommonRes `protobuf:"bytes,1,opt,name=commonRes,proto3" json:"commonRes,omitempty"`
	Endpoints []*CustomDocument     `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *ListCustomDocumentsRes) Reset() {
	*x = ListCustomDocumentsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCustomDocumentsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomDocumentsRes) ProtoMessage() {}

func (x *ListCustomDocumentsRes) ProtoReflect() protoreflect.Message {
	mi := &file_documents_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomDocumentsRes.ProtoReflect.Descriptor instead.
func (*ListCustomDocumentsRes) Descriptor() ([]byte, []int) {
	return file_documents_custom_proto_rawDescGZIP(), []int{2}
}

func (x *ListCustomDocumentsRes) GetCommonRes() *idl_common.CommonRes {
	if x != nil {
		return x.CommonRes
	}
	return nil
}

func (x *ListCustomDocumentsRes) GetEndpoints() []*CustomDocument {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type DesiredCustomDocumentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonRes *idl_common.CommonReq `protobuf:"bytes,1,opt,name=commonRes,proto3" json:"commonRes,omitempty"`
	Endpoint  *CustomDocument       `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *DesiredCustomDocumentReq) Reset() {
	*x = DesiredCustomDocumentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesiredCustomDocumentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredCustomDocumentReq) ProtoMessage() {}

func (x *DesiredCustomDocumentReq) ProtoReflect() protoreflect.Message {
	mi := &file_documents_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredCustomDocumentReq.ProtoReflect.Descriptor instead.
func (*DesiredCustomDocumentReq) Descriptor() ([]byte, []int) {
	return file_documents_custom_proto_rawDescGZIP(), []int{3}
}

func (x *DesiredCustomDocumentReq) GetCommonRes() *idl_common.CommonReq {
	if x != nil {
		return x.CommonRes
	}
	return nil
}

func (x *DesiredCustomDocumentReq) GetEndpoint() *CustomDocument {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

var File_documents_custom_proto protoreflect.FileDescriptor

var file_documents_custom_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x16,
	0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x3d, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x6c,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x61, 0x77, 0x53, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x61, 0x77, 0x53, 0x70, 0x65, 0x63, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x3e, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12,
	0x40, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x6c,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x22, 0x9c, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3e, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x40, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x32,
	0xaa, 0x04, 0x0a, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e,
	0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x72, 0x65, 0x64,
	0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x2c, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x6a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x26, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x66, 0x74,
	0x6f, 0x6e, 0x2d, 0x68, 0x71, 0x2f, 0x72, 0x65, 0x64, 0x2d, 0x66, 0x6f, 0x78, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_documents_custom_proto_rawDescOnce sync.Once
	file_documents_custom_proto_rawDescData = file_documents_custom_proto_rawDesc
)

func file_documents_custom_proto_rawDescGZIP() []byte {
	file_documents_custom_proto_rawDescOnce.Do(func() {
		file_documents_custom_proto_rawDescData = protoimpl.X.CompressGZIP(file_documents_custom_proto_rawDescData)
	})
	return file_documents_custom_proto_rawDescData
}

var file_documents_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_documents_custom_proto_goTypes = []interface{}{
	(*CustomDocument)(nil),             // 0: redfox.api.documents.CustomDocument
	(*GetCustomDocumentRes)(nil),       // 1: redfox.api.documents.GetCustomDocumentRes
	(*ListCustomDocumentsRes)(nil),     // 2: redfox.api.documents.ListCustomDocumentsRes
	(*DesiredCustomDocumentReq)(nil),   // 3: redfox.api.documents.DesiredCustomDocumentReq
	(*idl_common.ObjectMeta)(nil),      // 4: redfox.api.idl_common.ObjectMeta
	(*idl_common.CommonRes)(nil),       // 5: redfox.api.idl_common.CommonRes
	(*idl_common.CommonReq)(nil),       // 6: redfox.api.idl_common.CommonReq
	(*idl_common.SingleObjectReq)(nil), // 7: redfox.api.idl_common.SingleObjectReq
	(*idl_common.ListObjectReq)(nil),   // 8: redfox.api.idl_common.ListObjectReq
}
var file_documents_custom_proto_depIdxs = []int32{
	4,  // 0: redfox.api.documents.CustomDocument.metadata:type_name -> redfox.api.idl_common.ObjectMeta
	5,  // 1: redfox.api.documents.GetCustomDocumentRes.commonRes:type_name -> redfox.api.idl_common.CommonRes
	0,  // 2: redfox.api.documents.GetCustomDocumentRes.endpoint:type_name -> redfox.api.documents.CustomDocument
	5,  // 3: redfox.api.documents.ListCustomDocumentsRes.commonRes:type_name -> redfox.api.idl_common.CommonRes
	0,  // 4: redfox.api.documents.ListCustomDocumentsRes.endpoints:type_name -> redfox.api.documents.CustomDocument
	6,  // 5: redfox.api.documents.DesiredCustomDocumentReq.commonRes:type_name -> redfox.api.idl_common.CommonReq
	0,  // 6: redfox.api.documents.DesiredCustomDocumentReq.endpoint:type_name -> redfox.api.documents.CustomDocument
	7,  // 7: redfox.api.documents.CustomDocumentServer.GetCustomDocument:input_type -> redfox.api.idl_common.SingleObjectReq
	8,  // 8: redfox.api.documents.CustomDocumentServer.ListCustomDocuments:input_type -> redfox.api.idl_common.ListObjectReq
	3,  // 9: redfox.api.documents.CustomDocumentServer.CreateCustomDocument:input_type -> redfox.api.documents.DesiredCustomDocumentReq
	3,  // 10: redfox.api.documents.CustomDocumentServer.UpdateCustomDocument:input_type -> redfox.api.documents.DesiredCustomDocumentReq
	7,  // 11: redfox.api.documents.CustomDocumentServer.DeleteCustomDocument:input_type -> redfox.api.idl_common.SingleObjectReq
	1,  // 12: redfox.api.documents.CustomDocumentServer.GetCustomDocument:output_type -> redfox.api.documents.GetCustomDocumentRes
	2,  // 13: redfox.api.documents.CustomDocumentServer.ListCustomDocuments:output_type -> redfox.api.documents.ListCustomDocumentsRes
	5,  // 14: redfox.api.documents.CustomDocumentServer.CreateCustomDocument:output_type -> redfox.api.idl_common.CommonRes
	5,  // 15: redfox.api.documents.CustomDocumentServer.UpdateCustomDocument:output_type -> redfox.api.idl_common.CommonRes
	5,  // 16: redfox.api.documents.CustomDocumentServer.DeleteCustomDocument:output_type -> redfox.api.idl_common.CommonRes
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_documents_custom_proto_init() }
func file_documents_custom_proto_init() {
	if File_documents_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_documents_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomDocument); i {
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
		file_documents_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomDocumentRes); i {
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
		file_documents_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCustomDocumentsRes); i {
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
		file_documents_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesiredCustomDocumentReq); i {
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
			RawDescriptor: file_documents_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_documents_custom_proto_goTypes,
		DependencyIndexes: file_documents_custom_proto_depIdxs,
		MessageInfos:      file_documents_custom_proto_msgTypes,
	}.Build()
	File_documents_custom_proto = out.File
	file_documents_custom_proto_rawDesc = nil
	file_documents_custom_proto_goTypes = nil
	file_documents_custom_proto_depIdxs = nil
}