// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: envoytriage/v1/envoytriage_api.proto

package envoytriagev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*ReadOperation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetOperations() []*ReadOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type ReadOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Include *ReadOperation_Include `protobuf:"bytes,2,opt,name=include,proto3" json:"include,omitempty"`
}

func (x *ReadOperation) Reset() {
	*x = ReadOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOperation) ProtoMessage() {}

func (x *ReadOperation) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOperation.ProtoReflect.Descriptor instead.
func (*ReadOperation) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{1}
}

func (x *ReadOperation) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ReadOperation) GetInclude() *ReadOperation_Include {
	if x != nil {
		return x.Include
	}
	return nil
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{2}
}

func (x *ReadResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Address) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      *Address       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	NodeMetadata *NodeMetadata  `protobuf:"bytes,2,opt,name=node_metadata,json=nodeMetadata,proto3" json:"node_metadata,omitempty"`
	Output       *Result_Output `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Result) GetNodeMetadata() *NodeMetadata {
	if x != nil {
		return x.NodeMetadata
	}
	return nil
}

func (x *Result) GetOutput() *Result_Output {
	if x != nil {
		return x.Output
	}
	return nil
}

type NodeMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceNode    string `protobuf:"bytes,1,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	ServiceCluster string `protobuf:"bytes,2,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	ServiceZone    string `protobuf:"bytes,3,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	Version        string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *NodeMetadata) Reset() {
	*x = NodeMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetadata) ProtoMessage() {}

func (x *NodeMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetadata.ProtoReflect.Descriptor instead.
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{5}
}

func (x *NodeMetadata) GetServiceNode() string {
	if x != nil {
		return x.ServiceNode
	}
	return ""
}

func (x *NodeMetadata) GetServiceCluster() string {
	if x != nil {
		return x.ServiceCluster
	}
	return ""
}

func (x *NodeMetadata) GetServiceZone() string {
	if x != nil {
		return x.ServiceZone
	}
	return ""
}

func (x *NodeMetadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ReadOperation_Include struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters   bool `protobuf:"varint,1,opt,name=clusters,proto3" json:"clusters,omitempty"`
	ConfigDump bool `protobuf:"varint,2,opt,name=config_dump,json=configDump,proto3" json:"config_dump,omitempty"`
	Listeners  bool `protobuf:"varint,3,opt,name=listeners,proto3" json:"listeners,omitempty"`
	Runtime    bool `protobuf:"varint,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Stats      bool `protobuf:"varint,5,opt,name=stats,proto3" json:"stats,omitempty"`
	ServerInfo bool `protobuf:"varint,6,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
}

func (x *ReadOperation_Include) Reset() {
	*x = ReadOperation_Include{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOperation_Include) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOperation_Include) ProtoMessage() {}

func (x *ReadOperation_Include) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOperation_Include.ProtoReflect.Descriptor instead.
func (*ReadOperation_Include) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ReadOperation_Include) GetClusters() bool {
	if x != nil {
		return x.Clusters
	}
	return false
}

func (x *ReadOperation_Include) GetConfigDump() bool {
	if x != nil {
		return x.ConfigDump
	}
	return false
}

func (x *ReadOperation_Include) GetListeners() bool {
	if x != nil {
		return x.Listeners
	}
	return false
}

func (x *ReadOperation_Include) GetRuntime() bool {
	if x != nil {
		return x.Runtime
	}
	return false
}

func (x *ReadOperation_Include) GetStats() bool {
	if x != nil {
		return x.Stats
	}
	return false
}

func (x *ReadOperation_Include) GetServerInfo() bool {
	if x != nil {
		return x.ServerInfo
	}
	return false
}

type Result_Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters   *Clusters   `protobuf:"bytes,1,opt,name=clusters,proto3" json:"clusters,omitempty"`
	ConfigDump *ConfigDump `protobuf:"bytes,2,opt,name=config_dump,json=configDump,proto3" json:"config_dump,omitempty"`
	Listeners  *Listeners  `protobuf:"bytes,3,opt,name=listeners,proto3" json:"listeners,omitempty"`
	Runtime    *Runtime    `protobuf:"bytes,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Stats      *Stats      `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	ServerInfo *ServerInfo `protobuf:"bytes,6,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
}

func (x *Result_Output) Reset() {
	*x = Result_Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result_Output) ProtoMessage() {}

func (x *Result_Output) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_envoytriage_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result_Output.ProtoReflect.Descriptor instead.
func (*Result_Output) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Result_Output) GetClusters() *Clusters {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *Result_Output) GetConfigDump() *ConfigDump {
	if x != nil {
		return x.ConfigDump
	}
	return nil
}

func (x *Result_Output) GetListeners() *Listeners {
	if x != nil {
		return x.Listeners
	}
	return nil
}

func (x *Result_Output) GetRuntime() *Runtime {
	if x != nil {
		return x.Runtime
	}
	return nil
}

func (x *Result_Output) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Result_Output) GetServerInfo() *ServerInfo {
	if x != nil {
		return x.ServerInfo
	}
	return nil
}

var File_envoytriage_v1_envoytriage_api_proto protoreflect.FileDescriptor

var file_envoytriage_v1_envoytriage_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72,
	0x69, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0b, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xd3, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x1a, 0xb5,
	0x01, 0x0a, 0x07, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x64, 0x75, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22,
	0x46, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xa8,
	0x01, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0xff, 0xff,
	0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xc8, 0x04, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x48, 0x0a, 0x0d,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x1a, 0xfb, 0x02, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x3b, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x44, 0x75, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70,
	0x12, 0x3e, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x38, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x42,
	0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x88, 0x01, 0x0a,
	0x0e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x54, 0x72, 0x69, 0x61, 0x67, 0x65, 0x41, 0x50, 0x49, 0x12,
	0x76, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x3a, 0x01,
	0x2a, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoytriage_v1_envoytriage_api_proto_rawDescOnce sync.Once
	file_envoytriage_v1_envoytriage_api_proto_rawDescData = file_envoytriage_v1_envoytriage_api_proto_rawDesc
)

func file_envoytriage_v1_envoytriage_api_proto_rawDescGZIP() []byte {
	file_envoytriage_v1_envoytriage_api_proto_rawDescOnce.Do(func() {
		file_envoytriage_v1_envoytriage_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoytriage_v1_envoytriage_api_proto_rawDescData)
	})
	return file_envoytriage_v1_envoytriage_api_proto_rawDescData
}

var file_envoytriage_v1_envoytriage_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_envoytriage_v1_envoytriage_api_proto_goTypes = []interface{}{
	(*ReadRequest)(nil),           // 0: clutch.envoytriage.v1.ReadRequest
	(*ReadOperation)(nil),         // 1: clutch.envoytriage.v1.ReadOperation
	(*ReadResponse)(nil),          // 2: clutch.envoytriage.v1.ReadResponse
	(*Address)(nil),               // 3: clutch.envoytriage.v1.Address
	(*Result)(nil),                // 4: clutch.envoytriage.v1.Result
	(*NodeMetadata)(nil),          // 5: clutch.envoytriage.v1.NodeMetadata
	(*ReadOperation_Include)(nil), // 6: clutch.envoytriage.v1.ReadOperation.Include
	(*Result_Output)(nil),         // 7: clutch.envoytriage.v1.Result.Output
	(*Clusters)(nil),              // 8: clutch.envoytriage.v1.Clusters
	(*ConfigDump)(nil),            // 9: clutch.envoytriage.v1.ConfigDump
	(*Listeners)(nil),             // 10: clutch.envoytriage.v1.Listeners
	(*Runtime)(nil),               // 11: clutch.envoytriage.v1.Runtime
	(*Stats)(nil),                 // 12: clutch.envoytriage.v1.Stats
	(*ServerInfo)(nil),            // 13: clutch.envoytriage.v1.ServerInfo
}
var file_envoytriage_v1_envoytriage_api_proto_depIdxs = []int32{
	1,  // 0: clutch.envoytriage.v1.ReadRequest.operations:type_name -> clutch.envoytriage.v1.ReadOperation
	3,  // 1: clutch.envoytriage.v1.ReadOperation.address:type_name -> clutch.envoytriage.v1.Address
	6,  // 2: clutch.envoytriage.v1.ReadOperation.include:type_name -> clutch.envoytriage.v1.ReadOperation.Include
	4,  // 3: clutch.envoytriage.v1.ReadResponse.results:type_name -> clutch.envoytriage.v1.Result
	3,  // 4: clutch.envoytriage.v1.Result.address:type_name -> clutch.envoytriage.v1.Address
	5,  // 5: clutch.envoytriage.v1.Result.node_metadata:type_name -> clutch.envoytriage.v1.NodeMetadata
	7,  // 6: clutch.envoytriage.v1.Result.output:type_name -> clutch.envoytriage.v1.Result.Output
	8,  // 7: clutch.envoytriage.v1.Result.Output.clusters:type_name -> clutch.envoytriage.v1.Clusters
	9,  // 8: clutch.envoytriage.v1.Result.Output.config_dump:type_name -> clutch.envoytriage.v1.ConfigDump
	10, // 9: clutch.envoytriage.v1.Result.Output.listeners:type_name -> clutch.envoytriage.v1.Listeners
	11, // 10: clutch.envoytriage.v1.Result.Output.runtime:type_name -> clutch.envoytriage.v1.Runtime
	12, // 11: clutch.envoytriage.v1.Result.Output.stats:type_name -> clutch.envoytriage.v1.Stats
	13, // 12: clutch.envoytriage.v1.Result.Output.server_info:type_name -> clutch.envoytriage.v1.ServerInfo
	0,  // 13: clutch.envoytriage.v1.EnvoyTriageAPI.Read:input_type -> clutch.envoytriage.v1.ReadRequest
	2,  // 14: clutch.envoytriage.v1.EnvoyTriageAPI.Read:output_type -> clutch.envoytriage.v1.ReadResponse
	14, // [14:15] is the sub-list for method output_type
	13, // [13:14] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_envoytriage_v1_envoytriage_api_proto_init() }
func file_envoytriage_v1_envoytriage_api_proto_init() {
	if File_envoytriage_v1_envoytriage_api_proto != nil {
		return
	}
	file_envoytriage_v1_output_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadOperation); i {
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
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetadata); i {
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
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadOperation_Include); i {
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
		file_envoytriage_v1_envoytriage_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result_Output); i {
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
			RawDescriptor: file_envoytriage_v1_envoytriage_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoytriage_v1_envoytriage_api_proto_goTypes,
		DependencyIndexes: file_envoytriage_v1_envoytriage_api_proto_depIdxs,
		MessageInfos:      file_envoytriage_v1_envoytriage_api_proto_msgTypes,
	}.Build()
	File_envoytriage_v1_envoytriage_api_proto = out.File
	file_envoytriage_v1_envoytriage_api_proto_rawDesc = nil
	file_envoytriage_v1_envoytriage_api_proto_goTypes = nil
	file_envoytriage_v1_envoytriage_api_proto_depIdxs = nil
}
