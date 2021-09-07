// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: app/observatory/multiObservatory/config.proto

package multiObservatory

import (
	_ "github.com/v2fly/v2ray-core/v4/common/protoext"
	taggedfeatures "github.com/v2fly/v2ray-core/v4/common/taggedfeatures"
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

	Holders *taggedfeatures.Config `protobuf:"bytes,1,opt,name=holders,proto3" json:"holders,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_observatory_multiObservatory_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_observatory_multiObservatory_config_proto_msgTypes[0]
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
	return file_app_observatory_multiObservatory_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetHolders() *taggedfeatures.Config {
	if x != nil {
		return x.Holders
	}
	return nil
}

var File_app_observatory_multiObservatory_config_proto protoreflect.FileDescriptor

var file_app_observatory_multiObservatory_config_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2b, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x24, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x2f, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42,
	0x0a, 0x07, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x3a, 0x23, 0x82, 0xb5, 0x18, 0x09, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x82, 0xb5, 0x18, 0x12, 0x12, 0x10, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x42, 0xa2, 0x01, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x01, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0xaa, 0x02,
	0x2b, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_observatory_multiObservatory_config_proto_rawDescOnce sync.Once
	file_app_observatory_multiObservatory_config_proto_rawDescData = file_app_observatory_multiObservatory_config_proto_rawDesc
)

func file_app_observatory_multiObservatory_config_proto_rawDescGZIP() []byte {
	file_app_observatory_multiObservatory_config_proto_rawDescOnce.Do(func() {
		file_app_observatory_multiObservatory_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_observatory_multiObservatory_config_proto_rawDescData)
	})
	return file_app_observatory_multiObservatory_config_proto_rawDescData
}

var file_app_observatory_multiObservatory_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_observatory_multiObservatory_config_proto_goTypes = []interface{}{
	(*Config)(nil),                // 0: v2ray.core.app.observatory.multiObservatory.Config
	(*taggedfeatures.Config)(nil), // 1: v2ray.core.common.taggedfeatures.Config
}
var file_app_observatory_multiObservatory_config_proto_depIdxs = []int32{
	1, // 0: v2ray.core.app.observatory.multiObservatory.Config.holders:type_name -> v2ray.core.common.taggedfeatures.Config
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_observatory_multiObservatory_config_proto_init() }
func file_app_observatory_multiObservatory_config_proto_init() {
	if File_app_observatory_multiObservatory_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_observatory_multiObservatory_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_app_observatory_multiObservatory_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_observatory_multiObservatory_config_proto_goTypes,
		DependencyIndexes: file_app_observatory_multiObservatory_config_proto_depIdxs,
		MessageInfos:      file_app_observatory_multiObservatory_config_proto_msgTypes,
	}.Build()
	File_app_observatory_multiObservatory_config_proto = out.File
	file_app_observatory_multiObservatory_config_proto_rawDesc = nil
	file_app_observatory_multiObservatory_config_proto_goTypes = nil
	file_app_observatory_multiObservatory_config_proto_depIdxs = nil
}
