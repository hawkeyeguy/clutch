// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: config/module/sourcecontrol/v1/sourcecontrol.proto

package sourcecontrolv1

import (
	v1 "github.com/lyft/clutch/backend/api/sourcecontrol/v1"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the owner for new repositories.
	Owners []string `protobuf:"bytes,1,rep,name=owners,proto3" json:"owners,omitempty"`
	// Visibility options for new repositories.
	VisibilityOptions []v1.Visibility `protobuf:"varint,2,rep,packed,name=visibility_options,json=visibilityOptions,proto3,enum=clutch.sourcecontrol.v1.Visibility" json:"visibility_options,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_module_sourcecontrol_v1_sourcecontrol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_module_sourcecontrol_v1_sourcecontrol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetOwners() []string {
	if x != nil {
		return x.Owners
	}
	return nil
}

func (x *Config) GetVisibilityOptions() []v1.Visibility {
	if x != nil {
		return x.VisibilityOptions
	}
	return nil
}

var File_config_module_sourcecontrol_v1_sourcecontrol_proto protoreflect.FileDescriptor

var file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x74, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x12, 0x52, 0x0a, 0x12, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x11, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescOnce sync.Once
	file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescData = file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDesc
)

func file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescGZIP() []byte {
	file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescOnce.Do(func() {
		file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescData)
	})
	return file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDescData
}

var file_config_module_sourcecontrol_v1_sourcecontrol_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_module_sourcecontrol_v1_sourcecontrol_proto_goTypes = []interface{}{
	(*Config)(nil),     // 0: clutch.config.module.sourcecontrol.v1.Config
	(v1.Visibility)(0), // 1: clutch.sourcecontrol.v1.Visibility
}
var file_config_module_sourcecontrol_v1_sourcecontrol_proto_depIdxs = []int32{
	1, // 0: clutch.config.module.sourcecontrol.v1.Config.visibility_options:type_name -> clutch.sourcecontrol.v1.Visibility
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_module_sourcecontrol_v1_sourcecontrol_proto_init() }
func file_config_module_sourcecontrol_v1_sourcecontrol_proto_init() {
	if File_config_module_sourcecontrol_v1_sourcecontrol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_module_sourcecontrol_v1_sourcecontrol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_module_sourcecontrol_v1_sourcecontrol_proto_goTypes,
		DependencyIndexes: file_config_module_sourcecontrol_v1_sourcecontrol_proto_depIdxs,
		MessageInfos:      file_config_module_sourcecontrol_v1_sourcecontrol_proto_msgTypes,
	}.Build()
	File_config_module_sourcecontrol_v1_sourcecontrol_proto = out.File
	file_config_module_sourcecontrol_v1_sourcecontrol_proto_rawDesc = nil
	file_config_module_sourcecontrol_v1_sourcecontrol_proto_goTypes = nil
	file_config_module_sourcecontrol_v1_sourcecontrol_proto_depIdxs = nil
}
