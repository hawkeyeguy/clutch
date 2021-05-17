// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: api/v1/annotations.proto

package apiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_api_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Reference)(nil),
		Field:         58901,
		Name:          "clutch.api.v1.reference",
		Tag:           "bytes,58901,opt,name=reference",
		Filename:      "api/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Identifier)(nil),
		Field:         58902,
		Name:          "clutch.api.v1.id",
		Tag:           "bytes,58902,opt,name=id",
		Filename:      "api/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         58903,
		Name:          "clutch.api.v1.redacted",
		Tag:           "varint,58903,opt,name=redacted",
		Filename:      "api/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Action)(nil),
		Field:         58901,
		Name:          "clutch.api.v1.action",
		Tag:           "bytes,58901,opt,name=action",
		Filename:      "api/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         58902,
		Name:          "clutch.api.v1.disable_audit",
		Tag:           "varint,58902,opt,name=disable_audit",
		Filename:      "api/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Use a random high number that won't conflict with annotations from other
	// libraries.
	//
	// optional clutch.api.v1.Reference reference = 58901;
	E_Reference = &file_api_v1_annotations_proto_extTypes[0]
	// optional clutch.api.v1.Identifier id = 58902;
	E_Id = &file_api_v1_annotations_proto_extTypes[1]
	// optional bool redacted = 58903;
	E_Redacted = &file_api_v1_annotations_proto_extTypes[2]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// Use a random high number that won't conflict with annotations from other
	// libraries.
	//
	// optional clutch.api.v1.Action action = 58901;
	E_Action = &file_api_v1_annotations_proto_extTypes[3]
	// optional bool disable_audit = 58902;
	E_DisableAudit = &file_api_v1_annotations_proto_extTypes[4]
)

var File_api_v1_annotations_proto protoreflect.FileDescriptor

var file_api_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3a, 0x59, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x95,
	0xcc, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x3a, 0x4c, 0x0a, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x96, 0xcc, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x64,
	0x61, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x97, 0xcc, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x65, 0x64, 0x3a, 0x4f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x95, 0xcc, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x45, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x96, 0xcc, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*Reference)(nil),                   // 2: clutch.api.v1.Reference
	(*Identifier)(nil),                  // 3: clutch.api.v1.Identifier
	(*Action)(nil),                      // 4: clutch.api.v1.Action
}
var file_api_v1_annotations_proto_depIdxs = []int32{
	0, // 0: clutch.api.v1.reference:extendee -> google.protobuf.MessageOptions
	0, // 1: clutch.api.v1.id:extendee -> google.protobuf.MessageOptions
	0, // 2: clutch.api.v1.redacted:extendee -> google.protobuf.MessageOptions
	1, // 3: clutch.api.v1.action:extendee -> google.protobuf.MethodOptions
	1, // 4: clutch.api.v1.disable_audit:extendee -> google.protobuf.MethodOptions
	2, // 5: clutch.api.v1.reference:type_name -> clutch.api.v1.Reference
	3, // 6: clutch.api.v1.id:type_name -> clutch.api.v1.Identifier
	4, // 7: clutch.api.v1.action:type_name -> clutch.api.v1.Action
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	5, // [5:8] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_annotations_proto_init() }
func file_api_v1_annotations_proto_init() {
	if File_api_v1_annotations_proto != nil {
		return
	}
	file_api_v1_schema_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_annotations_proto_goTypes,
		DependencyIndexes: file_api_v1_annotations_proto_depIdxs,
		ExtensionInfos:    file_api_v1_annotations_proto_extTypes,
	}.Build()
	File_api_v1_annotations_proto = out.File
	file_api_v1_annotations_proto_rawDesc = nil
	file_api_v1_annotations_proto_goTypes = nil
	file_api_v1_annotations_proto_depIdxs = nil
}
