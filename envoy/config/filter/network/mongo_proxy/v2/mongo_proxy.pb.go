// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto

package mongo_proxyv2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type MongoProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mongo_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated :ref:`runtime <config_network_filters_mongo_proxy_runtime>`.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Inject a fixed delay before proxying a Mongo operation. Delays are
	// applied to the following MongoDB operations: Query, Insert, GetMore,
	// and KillCursors. Once an active delay is in progress, all incoming
	// data up until the timer event fires will be a part of the delay.
	Delay *v2.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
	// Flag to specify whether :ref:`dynamic metadata
	// <config_network_filters_mongo_proxy_dynamic_metadata>` should be emitted. Defaults to false.
	EmitDynamicMetadata bool `protobuf:"varint,4,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
}

func (x *MongoProxy) Reset() {
	*x = MongoProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoProxy) ProtoMessage() {}

func (x *MongoProxy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoProxy.ProtoReflect.Descriptor instead.
func (*MongoProxy) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *MongoProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *MongoProxy) GetAccessLog() string {
	if x != nil {
		return x.AccessLog
	}
	return ""
}

func (x *MongoProxy) GetDelay() *v2.FaultDelay {
	if x != nil {
		return x.Delay
	}
	return nil
}

func (x *MongoProxy) GetEmitDynamicMetadata() bool {
	if x != nil {
		return x.EmitDynamicMetadata
	}
	return false
}

var File_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto protoreflect.FileDescriptor

var file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a,
	0x0a, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x28, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x12, 0x3e, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x6d, 0x69, 0x74, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xed, 0x01, 0x0a, 0x38, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x32, 0x3b, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x32, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x31, 0x12,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescOnce sync.Once
	file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescData = file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDesc
)

func file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescData)
	})
	return file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDescData
}

var file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_goTypes = []interface{}{
	(*MongoProxy)(nil),    // 0: envoy.config.filter.network.mongo_proxy.v2.MongoProxy
	(*v2.FaultDelay)(nil), // 1: envoy.config.filter.fault.v2.FaultDelay
}
var file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_depIdxs = []int32{
	1, // 0: envoy.config.filter.network.mongo_proxy.v2.MongoProxy.delay:type_name -> envoy.config.filter.fault.v2.FaultDelay
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_init() }
func file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_init() {
	if File_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoProxy); i {
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
			RawDescriptor: file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto = out.File
	file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_rawDesc = nil
	file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_goTypes = nil
	file_envoy_config_filter_network_mongo_proxy_v2_mongo_proxy_proto_depIdxs = nil
}
