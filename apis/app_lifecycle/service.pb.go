// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: app_lifecycle/service.proto

package app_lifecycle

import (
	idl_common "github.com/krafton-hq/red-fox/apis/idl_common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_app_lifecycle_service_proto protoreflect.FileDescriptor

var file_app_lifecycle_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x72,
	0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x5f, 0x6c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x1a, 0x16, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x86, 0x02, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x05, 0x4c, 0x69, 0x76,
	0x65, 0x7a, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x7a, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x66, 0x6f, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x66, 0x74, 0x6f, 0x6e, 0x2d, 0x68,
	0x71, 0x2f, 0x72, 0x65, 0x64, 0x2d, 0x66, 0x6f, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x70, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_app_lifecycle_service_proto_goTypes = []interface{}{
	(*idl_common.CommonReq)(nil), // 0: redfox.api.idl_common.CommonReq
	(*idl_common.CommonRes)(nil), // 1: redfox.api.idl_common.CommonRes
}
var file_app_lifecycle_service_proto_depIdxs = []int32{
	0, // 0: redfox.api.app_lifecycle.ApplicationLifecycle.Version:input_type -> redfox.api.idl_common.CommonReq
	0, // 1: redfox.api.app_lifecycle.ApplicationLifecycle.Livez:input_type -> redfox.api.idl_common.CommonReq
	0, // 2: redfox.api.app_lifecycle.ApplicationLifecycle.Readyz:input_type -> redfox.api.idl_common.CommonReq
	1, // 3: redfox.api.app_lifecycle.ApplicationLifecycle.Version:output_type -> redfox.api.idl_common.CommonRes
	1, // 4: redfox.api.app_lifecycle.ApplicationLifecycle.Livez:output_type -> redfox.api.idl_common.CommonRes
	1, // 5: redfox.api.app_lifecycle.ApplicationLifecycle.Readyz:output_type -> redfox.api.idl_common.CommonRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_lifecycle_service_proto_init() }
func file_app_lifecycle_service_proto_init() {
	if File_app_lifecycle_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_lifecycle_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_lifecycle_service_proto_goTypes,
		DependencyIndexes: file_app_lifecycle_service_proto_depIdxs,
	}.Build()
	File_app_lifecycle_service_proto = out.File
	file_app_lifecycle_service_proto_rawDesc = nil
	file_app_lifecycle_service_proto_goTypes = nil
	file_app_lifecycle_service_proto_depIdxs = nil
}
