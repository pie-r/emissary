// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/type/matcher/v4alpha/metadata.proto

package envoy_type_matcher_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type MetadataMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter string                         `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Path   []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Value  *ValueMatcher                  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MetadataMatcher) Reset() {
	*x = MetadataMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataMatcher) ProtoMessage() {}

func (x *MetadataMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataMatcher.ProtoReflect.Descriptor instead.
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v4alpha_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataMatcher) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *MetadataMatcher) GetPath() []*MetadataMatcher_PathSegment {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *MetadataMatcher) GetValue() *ValueMatcher {
	if x != nil {
		return x.Value
	}
	return nil
}

type MetadataMatcher_PathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Segment:
	//	*MetadataMatcher_PathSegment_Key
	Segment isMetadataMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
}

func (x *MetadataMatcher_PathSegment) Reset() {
	*x = MetadataMatcher_PathSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataMatcher_PathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataMatcher_PathSegment) ProtoMessage() {}

func (x *MetadataMatcher_PathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataMatcher_PathSegment.ProtoReflect.Descriptor instead.
func (*MetadataMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v4alpha_metadata_proto_rawDescGZIP(), []int{0, 0}
}

func (m *MetadataMatcher_PathSegment) GetSegment() isMetadataMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (x *MetadataMatcher_PathSegment) GetKey() string {
	if x, ok := x.GetSegment().(*MetadataMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*MetadataMatcher_PathSegment_Key) isMetadataMatcher_PathSegment_Segment() {}

var File_envoy_type_matcher_v4alpha_metadata_proto protoreflect.FileDescriptor

var file_envoy_type_matcher_v4alpha_metadata_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x0f, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x55, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x48, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x74, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x3a, 0x38, 0x9a, 0xc5,
	0x88, 0x1e, 0x33, 0x0a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x42, 0x43, 0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_type_matcher_v4alpha_metadata_proto_rawDescOnce sync.Once
	file_envoy_type_matcher_v4alpha_metadata_proto_rawDescData = file_envoy_type_matcher_v4alpha_metadata_proto_rawDesc
)

func file_envoy_type_matcher_v4alpha_metadata_proto_rawDescGZIP() []byte {
	file_envoy_type_matcher_v4alpha_metadata_proto_rawDescOnce.Do(func() {
		file_envoy_type_matcher_v4alpha_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_matcher_v4alpha_metadata_proto_rawDescData)
	})
	return file_envoy_type_matcher_v4alpha_metadata_proto_rawDescData
}

var file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_type_matcher_v4alpha_metadata_proto_goTypes = []interface{}{
	(*MetadataMatcher)(nil),             // 0: envoy.type.matcher.v4alpha.MetadataMatcher
	(*MetadataMatcher_PathSegment)(nil), // 1: envoy.type.matcher.v4alpha.MetadataMatcher.PathSegment
	(*ValueMatcher)(nil),                // 2: envoy.type.matcher.v4alpha.ValueMatcher
}
var file_envoy_type_matcher_v4alpha_metadata_proto_depIdxs = []int32{
	1, // 0: envoy.type.matcher.v4alpha.MetadataMatcher.path:type_name -> envoy.type.matcher.v4alpha.MetadataMatcher.PathSegment
	2, // 1: envoy.type.matcher.v4alpha.MetadataMatcher.value:type_name -> envoy.type.matcher.v4alpha.ValueMatcher
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_type_matcher_v4alpha_metadata_proto_init() }
func file_envoy_type_matcher_v4alpha_metadata_proto_init() {
	if File_envoy_type_matcher_v4alpha_metadata_proto != nil {
		return
	}
	file_envoy_type_matcher_v4alpha_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataMatcher); i {
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
		file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataMatcher_PathSegment); i {
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
	file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_matcher_v4alpha_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_matcher_v4alpha_metadata_proto_goTypes,
		DependencyIndexes: file_envoy_type_matcher_v4alpha_metadata_proto_depIdxs,
		MessageInfos:      file_envoy_type_matcher_v4alpha_metadata_proto_msgTypes,
	}.Build()
	File_envoy_type_matcher_v4alpha_metadata_proto = out.File
	file_envoy_type_matcher_v4alpha_metadata_proto_rawDesc = nil
	file_envoy_type_matcher_v4alpha_metadata_proto_goTypes = nil
	file_envoy_type_matcher_v4alpha_metadata_proto_depIdxs = nil
}
