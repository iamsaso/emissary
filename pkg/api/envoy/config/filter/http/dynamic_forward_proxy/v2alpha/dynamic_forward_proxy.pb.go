// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v2alpha "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/config/common/dynamic_forward_proxy/v2alpha"
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

// Configuration for the dynamic forward proxy HTTP filter. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#extension: envoy.filters.http.dynamic_forward_proxy]
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DNS cache configuration that the filter will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy cluster configuration
	// <envoy_api_field_config.cluster.dynamic_forward_proxy.v2alpha.ClusterConfig.dns_cache_config>`.
	DnsCacheConfig *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if x != nil {
		return x.DnsCacheConfig
	}
	return nil
}

// Per route Configuration for the dynamic forward proxy HTTP filter.
type PerRouteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to HostRewriteSpecifier:
	//	*PerRouteConfig_HostRewrite
	//	*PerRouteConfig_AutoHostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
}

func (x *PerRouteConfig) Reset() {
	*x = PerRouteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerRouteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerRouteConfig) ProtoMessage() {}

func (x *PerRouteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerRouteConfig.ProtoReflect.Descriptor instead.
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescGZIP(), []int{1}
}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (x *PerRouteConfig) GetHostRewrite() string {
	if x, ok := x.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewrite); ok {
		return x.HostRewrite
	}
	return ""
}

func (x *PerRouteConfig) GetAutoHostRewriteHeader() string {
	if x, ok := x.GetHostRewriteSpecifier().(*PerRouteConfig_AutoHostRewriteHeader); ok {
		return x.AutoHostRewriteHeader
	}
	return ""
}

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
}

type PerRouteConfig_HostRewrite struct {
	// Indicates that before DNS lookup, the host header will be swapped with
	// this value. If not set or empty, the original host header value
	// will be used and no rewrite will happen.
	//
	// Note: this rewrite affects both DNS lookup and host header forwarding. However, this
	// option shouldn't be used with
	// :ref:`HCM host rewrite <envoy_api_field_route.RouteAction.host_rewrite>` given that the
	// value set here would be used for DNS lookups whereas the value set in the HCM would be used
	// for host header forwarding which is not the desired outcome.
	HostRewrite string `protobuf:"bytes,1,opt,name=host_rewrite,json=hostRewrite,proto3,oneof"`
}

type PerRouteConfig_AutoHostRewriteHeader struct {
	// Indicates that before DNS lookup, the host header will be swapped with
	// the value of this header. If not set or empty, the original host header
	// value will be used and no rewrite will happen.
	//
	// Note: this rewrite affects both DNS lookup and host header forwarding. However, this
	// option shouldn't be used with
	// :ref:`HCM host rewrite header <envoy_api_field_route.RouteAction.auto_host_rewrite_header>`
	// given that the value set here would be used for DNS lookups whereas the value set in the HCM
	// would be used for host header forwarding which is not the desired outcome.
	//
	// .. note::
	//
	//   If the header appears multiple times only the first value is used.
	AutoHostRewriteHeader string `protobuf:"bytes,2,opt,name=auto_host_rewrite_header,json=autoHostRewriteHeader,proto3,oneof"`
}

func (*PerRouteConfig_HostRewrite) isPerRouteConfig_HostRewriteSpecifier() {}

func (*PerRouteConfig_AutoHostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier() {}

var File_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDesc = []byte{
	0x0a, 0x52, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x41, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x75, 0x0a, 0x10, 0x64, 0x6e, 0x73, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0e, 0x64, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0xc5, 0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x41, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x16,
	0x0a, 0x14, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x56, 0x0a, 0x18, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x15, 0x0a,
	0x13, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x15, 0x61, 0x75, 0x74, 0x6f, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x18, 0x0a,
	0x16, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x87, 0x02, 0x0a, 0x44, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x18, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xf2, 0x98, 0xfe, 0x8f, 0x05,
	0x38, 0x12, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescData = file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDesc
)

func file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescData)
	})
	return file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDescData
}

var file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_goTypes = []interface{}{
	(*FilterConfig)(nil),           // 0: envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig
	(*PerRouteConfig)(nil),         // 1: envoy.config.filter.http.dynamic_forward_proxy.v2alpha.PerRouteConfig
	(*v2alpha.DnsCacheConfig)(nil), // 2: envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig
}
var file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_depIdxs = []int32{
	2, // 0: envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig.dns_cache_config:type_name -> envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_init()
}
func file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_init() {
	if File_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
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
		file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerRouteConfig); i {
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
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PerRouteConfig_HostRewrite)(nil),
		(*PerRouteConfig_AutoHostRewriteHeader)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto = out.File
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_rawDesc = nil
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_goTypes = nil
	file_envoy_config_filter_http_dynamic_forward_proxy_v2alpha_dynamic_forward_proxy_proto_depIdxs = nil
}
