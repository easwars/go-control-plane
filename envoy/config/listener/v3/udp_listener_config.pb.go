// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/listener/v3/udp_listener_config.proto

package envoy_config_listener_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#next-free-field: 7]
type UdpListenerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to create a specific UDP listener factory. If not specified the default UDP listener is
	// used.
	// [#comment:#extension-category: envoy.udp_listeners]
	// [#not-implemented-hide:]
	// [#comment:TODO(#12829): Remove this as an extension point.]
	ListenerConfig *v3.TypedExtensionConfig `protobuf:"bytes,4,opt,name=listener_config,json=listenerConfig,proto3" json:"listener_config,omitempty"`
	// The maximum size of received downstream UDP datagrams. Using a larger size will cause Envoy to allocate
	// more memory per listener. Received datagrams above this size will be dropped. If not set
	// defaults to 1500 bytes.
	MaxDownstreamRxDatagramSize *wrappers.UInt64Value `protobuf:"bytes,5,opt,name=max_downstream_rx_datagram_size,json=maxDownstreamRxDatagramSize,proto3" json:"max_downstream_rx_datagram_size,omitempty"`
	// If the protocol in the listener socket address in :ref:`protocol
	// <envoy_api_field_config.core.v3.SocketAddress.protocol>` is :ref:`UDP
	// <envoy_api_enum_value_config.core.v3.SocketAddress.Protocol.UDP>`, this field specifies the
	// actual UDP writer to create. If not specified the default UDP writer is used.
	// [#comment:#extension-category: envoy.udp_packet_writers]
	// [#not-implemented-hide:]
	// [#comment:TODO(#12829): Remove this as an extension point.]
	WriterConfig *v3.TypedExtensionConfig `protobuf:"bytes,6,opt,name=writer_config,json=writerConfig,proto3" json:"writer_config,omitempty"`
	// Types that are assignable to ConfigType:
	//	*UdpListenerConfig_HiddenEnvoyDeprecatedConfig
	ConfigType isUdpListenerConfig_ConfigType `protobuf_oneof:"config_type"`
}

func (x *UdpListenerConfig) Reset() {
	*x = UdpListenerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpListenerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpListenerConfig) ProtoMessage() {}

func (x *UdpListenerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpListenerConfig.ProtoReflect.Descriptor instead.
func (*UdpListenerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP(), []int{0}
}

func (x *UdpListenerConfig) GetListenerConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.ListenerConfig
	}
	return nil
}

func (x *UdpListenerConfig) GetMaxDownstreamRxDatagramSize() *wrappers.UInt64Value {
	if x != nil {
		return x.MaxDownstreamRxDatagramSize
	}
	return nil
}

func (x *UdpListenerConfig) GetWriterConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.WriterConfig
	}
	return nil
}

func (m *UdpListenerConfig) GetConfigType() isUdpListenerConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (x *UdpListenerConfig) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := x.GetConfigType().(*UdpListenerConfig_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

type isUdpListenerConfig_ConfigType interface {
	isUdpListenerConfig_ConfigType()
}

type UdpListenerConfig_HiddenEnvoyDeprecatedConfig struct {
	// Deprecated: Do not use.
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

func (*UdpListenerConfig_HiddenEnvoyDeprecatedConfig) isUdpListenerConfig_ConfigType() {}

type ActiveRawUdpListenerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActiveRawUdpListenerConfig) Reset() {
	*x = ActiveRawUdpListenerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRawUdpListenerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRawUdpListenerConfig) ProtoMessage() {}

func (x *ActiveRawUdpListenerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRawUdpListenerConfig.ProtoReflect.Descriptor instead.
func (*ActiveRawUdpListenerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP(), []int{1}
}

var File_envoy_config_listener_v3_udp_listener_config_proto protoreflect.FileDescriptor

var file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x24,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a, 0x11, 0x55, 0x64, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x53,
	0x0a, 0x0f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x6f, 0x0a, 0x1f, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x32,
	0x06, 0x10, 0x80, 0x80, 0x04, 0x20, 0x00, 0x52, 0x1b, 0x6d, 0x61, 0x78, 0x44, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x78, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x62, 0x0a, 0x1e, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x1b, 0x68, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x2e, 0x9a, 0xc5, 0x88, 0x1e, 0x29,
	0x0a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04,
	0x08, 0x03, 0x10, 0x04, 0x22, 0x55, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61,
	0x77, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x77, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x4a, 0x0a, 0x26, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescOnce sync.Once
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData = file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc
)

func file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP() []byte {
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData)
	})
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData
}

var file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_listener_v3_udp_listener_config_proto_goTypes = []interface{}{
	(*UdpListenerConfig)(nil),          // 0: envoy.config.listener.v3.UdpListenerConfig
	(*ActiveRawUdpListenerConfig)(nil), // 1: envoy.config.listener.v3.ActiveRawUdpListenerConfig
	(*v3.TypedExtensionConfig)(nil),    // 2: envoy.config.core.v3.TypedExtensionConfig
	(*wrappers.UInt64Value)(nil),       // 3: google.protobuf.UInt64Value
	(*_struct.Struct)(nil),             // 4: google.protobuf.Struct
}
var file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs = []int32{
	2, // 0: envoy.config.listener.v3.UdpListenerConfig.listener_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	3, // 1: envoy.config.listener.v3.UdpListenerConfig.max_downstream_rx_datagram_size:type_name -> google.protobuf.UInt64Value
	2, // 2: envoy.config.listener.v3.UdpListenerConfig.writer_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	4, // 3: envoy.config.listener.v3.UdpListenerConfig.hidden_envoy_deprecated_config:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_config_listener_v3_udp_listener_config_proto_init() }
func file_envoy_config_listener_v3_udp_listener_config_proto_init() {
	if File_envoy_config_listener_v3_udp_listener_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpListenerConfig); i {
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
		file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRawUdpListenerConfig); i {
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
	file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UdpListenerConfig_HiddenEnvoyDeprecatedConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_listener_v3_udp_listener_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes,
	}.Build()
	File_envoy_config_listener_v3_udp_listener_config_proto = out.File
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc = nil
	file_envoy_config_listener_v3_udp_listener_config_proto_goTypes = nil
	file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs = nil
}
