// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/config/core/v4alpha/substitution_format_string.proto

package envoy_config_core_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type SubstitutionFormatString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Format:
	//	*SubstitutionFormatString_TextFormat
	//	*SubstitutionFormatString_JsonFormat
	//	*SubstitutionFormatString_TextFormatSource
	Format      isSubstitutionFormatString_Format `protobuf_oneof:"format"`
	ContentType string                            `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *SubstitutionFormatString) Reset() {
	*x = SubstitutionFormatString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v4alpha_substitution_format_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubstitutionFormatString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubstitutionFormatString) ProtoMessage() {}

func (x *SubstitutionFormatString) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v4alpha_substitution_format_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubstitutionFormatString.ProtoReflect.Descriptor instead.
func (*SubstitutionFormatString) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescGZIP(), []int{0}
}

func (m *SubstitutionFormatString) GetFormat() isSubstitutionFormatString_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (x *SubstitutionFormatString) GetTextFormat() string {
	if x, ok := x.GetFormat().(*SubstitutionFormatString_TextFormat); ok {
		return x.TextFormat
	}
	return ""
}

func (x *SubstitutionFormatString) GetJsonFormat() *_struct.Struct {
	if x, ok := x.GetFormat().(*SubstitutionFormatString_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

func (x *SubstitutionFormatString) GetTextFormatSource() *DataSource {
	if x, ok := x.GetFormat().(*SubstitutionFormatString_TextFormatSource); ok {
		return x.TextFormatSource
	}
	return nil
}

func (x *SubstitutionFormatString) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type isSubstitutionFormatString_Format interface {
	isSubstitutionFormatString_Format()
}

type SubstitutionFormatString_TextFormat struct {
	TextFormat string `protobuf:"bytes,1,opt,name=text_format,json=textFormat,proto3,oneof"`
}

type SubstitutionFormatString_JsonFormat struct {
	JsonFormat *_struct.Struct `protobuf:"bytes,2,opt,name=json_format,json=jsonFormat,proto3,oneof"`
}

type SubstitutionFormatString_TextFormatSource struct {
	TextFormatSource *DataSource `protobuf:"bytes,5,opt,name=text_format_source,json=textFormatSource,proto3,oneof"`
}

func (*SubstitutionFormatString_TextFormat) isSubstitutionFormatString_Format() {}

func (*SubstitutionFormatString_JsonFormat) isSubstitutionFormatString_Format() {}

func (*SubstitutionFormatString_TextFormatSource) isSubstitutionFormatString_Format() {}

var File_envoy_config_core_v4alpha_substitution_format_string_proto protoreflect.FileDescriptor

var file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x65, 0x78, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x44, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00,
	0x52, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x55, 0x0a, 0x12,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48,
	0x00, 0x52, 0x10, 0x74, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x34, 0x9a, 0xc5, 0x88, 0x1e, 0x2f, 0x0a, 0x2d, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0d, 0x0a, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x52, 0x0a, 0x27, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescOnce sync.Once
	file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescData = file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDesc
)

func file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescData)
	})
	return file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDescData
}

var file_envoy_config_core_v4alpha_substitution_format_string_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_core_v4alpha_substitution_format_string_proto_goTypes = []interface{}{
	(*SubstitutionFormatString)(nil), // 0: envoy.config.core.v4alpha.SubstitutionFormatString
	(*_struct.Struct)(nil),           // 1: google.protobuf.Struct
	(*DataSource)(nil),               // 2: envoy.config.core.v4alpha.DataSource
}
var file_envoy_config_core_v4alpha_substitution_format_string_proto_depIdxs = []int32{
	1, // 0: envoy.config.core.v4alpha.SubstitutionFormatString.json_format:type_name -> google.protobuf.Struct
	2, // 1: envoy.config.core.v4alpha.SubstitutionFormatString.text_format_source:type_name -> envoy.config.core.v4alpha.DataSource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v4alpha_substitution_format_string_proto_init() }
func file_envoy_config_core_v4alpha_substitution_format_string_proto_init() {
	if File_envoy_config_core_v4alpha_substitution_format_string_proto != nil {
		return
	}
	file_envoy_config_core_v4alpha_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v4alpha_substitution_format_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubstitutionFormatString); i {
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
	file_envoy_config_core_v4alpha_substitution_format_string_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SubstitutionFormatString_TextFormat)(nil),
		(*SubstitutionFormatString_JsonFormat)(nil),
		(*SubstitutionFormatString_TextFormatSource)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v4alpha_substitution_format_string_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v4alpha_substitution_format_string_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v4alpha_substitution_format_string_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v4alpha_substitution_format_string_proto = out.File
	file_envoy_config_core_v4alpha_substitution_format_string_proto_rawDesc = nil
	file_envoy_config_core_v4alpha_substitution_format_string_proto_goTypes = nil
	file_envoy_config_core_v4alpha_substitution_format_string_proto_depIdxs = nil
}
