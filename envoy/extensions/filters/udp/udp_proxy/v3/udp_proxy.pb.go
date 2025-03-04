// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: envoy/extensions/filters/udp/udp_proxy/v3/udp_proxy.proto

package udp_proxyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v32 "github.com/cncf/xds/go/xds/type/matcher/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Configuration for the UDP proxy filter.
// [#next-free-field: 10]
type UdpProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stat prefix used when emitting UDP proxy filter stats.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are assignable to RouteSpecifier:
	//	*UdpProxyConfig_Cluster
	//	*UdpProxyConfig_Matcher
	RouteSpecifier isUdpProxyConfig_RouteSpecifier `protobuf_oneof:"route_specifier"`
	// The idle timeout for sessions. Idle is defined as no datagrams between received or sent by
	// the session. The default if not specified is 1 minute.
	IdleTimeout *duration.Duration `protobuf:"bytes,3,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// Use the remote downstream IP address as the sender IP address when sending packets to upstream hosts.
	// This option requires Envoy to be run with the ``CAP_NET_ADMIN`` capability on Linux.
	// And the IPv6 stack must be enabled on Linux kernel.
	// This option does not preserve the remote downstream port.
	// If this option is enabled, the IP address of sent datagrams will be changed to the remote downstream IP address.
	// This means that Envoy will not receive packets that are sent by upstream hosts because the upstream hosts
	// will send the packets with the remote downstream IP address as the destination. All packets will be routed
	// to the remote downstream directly if there are route rules on the upstream host side.
	// There are two options to return the packets back to the remote downstream.
	// The first one is to use DSR (Direct Server Return).
	// The other one is to configure routing rules on the upstream hosts to forward
	// all packets back to Envoy and configure iptables rules on the host running Envoy to
	// forward all packets from upstream hosts to the Envoy process so that Envoy can forward the packets to the downstream.
	// If the platform does not support this option, Envoy will raise a configuration error.
	UseOriginalSrcIp bool `protobuf:"varint,4,opt,name=use_original_src_ip,json=useOriginalSrcIp,proto3" json:"use_original_src_ip,omitempty"`
	// Optional configuration for UDP proxy hash policies. If hash_policies is not set, the hash-based
	// load balancing algorithms will select a host randomly. Currently the number of hash policies is
	// limited to 1.
	HashPolicies []*UdpProxyConfig_HashPolicy `protobuf:"bytes,5,rep,name=hash_policies,json=hashPolicies,proto3" json:"hash_policies,omitempty"`
	// UDP socket configuration for upstream sockets. The default for
	// :ref:`prefer_gro <envoy_v3_api_field_config.core.v3.UdpSocketConfig.prefer_gro>` is true for upstream
	// sockets as the assumption is datagrams will be received from a single source.
	UpstreamSocketConfig *v3.UdpSocketConfig `protobuf:"bytes,6,opt,name=upstream_socket_config,json=upstreamSocketConfig,proto3" json:"upstream_socket_config,omitempty"`
	// Perform per packet load balancing (upstream host selection) on each received data chunk.
	// The default if not specified is false, that means each data chunk is forwarded
	// to upstream host selected on first chunk receival for that "session" (identified by source IP/port and local IP/port).
	UsePerPacketLoadBalancing bool `protobuf:"varint,7,opt,name=use_per_packet_load_balancing,json=usePerPacketLoadBalancing,proto3" json:"use_per_packet_load_balancing,omitempty"`
	// Configuration for access logs emitted by the UDP proxy. Note that certain UDP specific data is emitted as :ref:`Dynamic Metadata <config_access_log_format_dynamic_metadata>`.
	AccessLog []*v31.AccessLog `protobuf:"bytes,8,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
}

func (x *UdpProxyConfig) Reset() {
	*x = UdpProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpProxyConfig) ProtoMessage() {}

func (x *UdpProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpProxyConfig.ProtoReflect.Descriptor instead.
func (*UdpProxyConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *UdpProxyConfig) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (m *UdpProxyConfig) GetRouteSpecifier() isUdpProxyConfig_RouteSpecifier {
	if m != nil {
		return m.RouteSpecifier
	}
	return nil
}

// Deprecated: Do not use.
func (x *UdpProxyConfig) GetCluster() string {
	if x, ok := x.GetRouteSpecifier().(*UdpProxyConfig_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *UdpProxyConfig) GetMatcher() *v32.Matcher {
	if x, ok := x.GetRouteSpecifier().(*UdpProxyConfig_Matcher); ok {
		return x.Matcher
	}
	return nil
}

func (x *UdpProxyConfig) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *UdpProxyConfig) GetUseOriginalSrcIp() bool {
	if x != nil {
		return x.UseOriginalSrcIp
	}
	return false
}

func (x *UdpProxyConfig) GetHashPolicies() []*UdpProxyConfig_HashPolicy {
	if x != nil {
		return x.HashPolicies
	}
	return nil
}

func (x *UdpProxyConfig) GetUpstreamSocketConfig() *v3.UdpSocketConfig {
	if x != nil {
		return x.UpstreamSocketConfig
	}
	return nil
}

func (x *UdpProxyConfig) GetUsePerPacketLoadBalancing() bool {
	if x != nil {
		return x.UsePerPacketLoadBalancing
	}
	return false
}

func (x *UdpProxyConfig) GetAccessLog() []*v31.AccessLog {
	if x != nil {
		return x.AccessLog
	}
	return nil
}

type isUdpProxyConfig_RouteSpecifier interface {
	isUdpProxyConfig_RouteSpecifier()
}

type UdpProxyConfig_Cluster struct {
	// The upstream cluster to connect to.
	// This field is deprecated in favor of
	// :ref:`matcher <envoy_v3_api_field_extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.matcher>`.
	//
	// Deprecated: Do not use.
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type UdpProxyConfig_Matcher struct {
	// The match tree to use when resolving route actions for incoming requests.
	// See :ref:`Routing <config_udp_listener_filters_udp_proxy_routing>` for more information.
	Matcher *v32.Matcher `protobuf:"bytes,9,opt,name=matcher,proto3,oneof"`
}

func (*UdpProxyConfig_Cluster) isUdpProxyConfig_RouteSpecifier() {}

func (*UdpProxyConfig_Matcher) isUdpProxyConfig_RouteSpecifier() {}

// Specifies the UDP hash policy.
// The packets can be routed by hash policy.
type UdpProxyConfig_HashPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PolicySpecifier:
	//	*UdpProxyConfig_HashPolicy_SourceIp
	//	*UdpProxyConfig_HashPolicy_Key
	PolicySpecifier isUdpProxyConfig_HashPolicy_PolicySpecifier `protobuf_oneof:"policy_specifier"`
}

func (x *UdpProxyConfig_HashPolicy) Reset() {
	*x = UdpProxyConfig_HashPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpProxyConfig_HashPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpProxyConfig_HashPolicy) ProtoMessage() {}

func (x *UdpProxyConfig_HashPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpProxyConfig_HashPolicy.ProtoReflect.Descriptor instead.
func (*UdpProxyConfig_HashPolicy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescGZIP(), []int{0, 0}
}

func (m *UdpProxyConfig_HashPolicy) GetPolicySpecifier() isUdpProxyConfig_HashPolicy_PolicySpecifier {
	if m != nil {
		return m.PolicySpecifier
	}
	return nil
}

func (x *UdpProxyConfig_HashPolicy) GetSourceIp() bool {
	if x, ok := x.GetPolicySpecifier().(*UdpProxyConfig_HashPolicy_SourceIp); ok {
		return x.SourceIp
	}
	return false
}

func (x *UdpProxyConfig_HashPolicy) GetKey() string {
	if x, ok := x.GetPolicySpecifier().(*UdpProxyConfig_HashPolicy_Key); ok {
		return x.Key
	}
	return ""
}

type isUdpProxyConfig_HashPolicy_PolicySpecifier interface {
	isUdpProxyConfig_HashPolicy_PolicySpecifier()
}

type UdpProxyConfig_HashPolicy_SourceIp struct {
	// The source IP will be used to compute the hash used by hash-based load balancing algorithms.
	SourceIp bool `protobuf:"varint,1,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

type UdpProxyConfig_HashPolicy_Key struct {
	// A given key will be used to compute the hash used by hash-based load balancing algorithms.
	// In certain cases there is a need to direct different UDP streams jointly towards the selected set of endpoints.
	// A possible use-case is VoIP telephony, where media (RTP) and its corresponding control (RTCP) belong to the same logical session,
	// although they travel in separate streams. To ensure that these pair of streams are load-balanced on session level
	// (instead of individual stream level), dynamically created listeners can use the same hash key for each stream in the session.
	Key string `protobuf:"bytes,2,opt,name=key,proto3,oneof"`
}

func (*UdpProxyConfig_HashPolicy_SourceIp) isUdpProxyConfig_HashPolicy_PolicySpecifier() {}

func (*UdpProxyConfig_HashPolicy_Key) isUdpProxyConfig_HashPolicy_PolicySpecifier() {}

var File_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x75,
	0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x06, 0x0a, 0x0e, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x12, 0x18, 0x01, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x92, 0xc7, 0x86, 0xd8,
	0x04, 0x03, 0x33, 0x2e, 0x30, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x42, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42,
	0x08, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x75, 0x73, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x72, 0x63, 0x49,
	0x70, 0x12, 0x73, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x16, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x64,
	0x70, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x40, 0x0a, 0x1d, 0x75, 0x73, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x75, 0x73, 0x65, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6c, 0x6f, 0x67, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x1a, 0x6a, 0x0a, 0x0a, 0x48, 0x61,
	0x73, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x26, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x6a, 0x02, 0x08, 0x01, 0x48, 0x00, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70,
	0x12, 0x1b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x17, 0x0a,
	0x10, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x3a, 0x3f, 0x9a, 0xc5, 0x88, 0x1e, 0x3a, 0x0a, 0x38, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x16, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42,
	0xb0, 0x01, 0x0a, 0x37, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x75,
	0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x55, 0x64, 0x70,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64,
	0x70, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x75,
	0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData = file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_goTypes = []interface{}{
	(*UdpProxyConfig)(nil),            // 0: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig
	(*UdpProxyConfig_HashPolicy)(nil), // 1: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.HashPolicy
	(*v32.Matcher)(nil),               // 2: xds.type.matcher.v3.Matcher
	(*duration.Duration)(nil),         // 3: google.protobuf.Duration
	(*v3.UdpSocketConfig)(nil),        // 4: envoy.config.core.v3.UdpSocketConfig
	(*v31.AccessLog)(nil),             // 5: envoy.config.accesslog.v3.AccessLog
}
var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.matcher:type_name -> xds.type.matcher.v3.Matcher
	3, // 1: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.idle_timeout:type_name -> google.protobuf.Duration
	1, // 2: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.hash_policies:type_name -> envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.HashPolicy
	4, // 3: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.upstream_socket_config:type_name -> envoy.config.core.v3.UdpSocketConfig
	5, // 4: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.access_log:type_name -> envoy.config.accesslog.v3.AccessLog
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_init() }
func file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_init() {
	if File_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpProxyConfig); i {
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
		file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpProxyConfig_HashPolicy); i {
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
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UdpProxyConfig_Cluster)(nil),
		(*UdpProxyConfig_Matcher)(nil),
	}
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UdpProxyConfig_HashPolicy_SourceIp)(nil),
		(*UdpProxyConfig_HashPolicy_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto = out.File
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_depIdxs = nil
}
