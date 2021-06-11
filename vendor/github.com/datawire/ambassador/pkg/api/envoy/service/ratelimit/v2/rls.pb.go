// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/service/ratelimit/v2/rls.proto

package envoy_service_ratelimit_v2

import (
	context "context"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	ratelimit "github.com/datawire/ambassador/pkg/api/envoy/api/v2/ratelimit"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RateLimitResponse_Code int32

const (
	RateLimitResponse_UNKNOWN    RateLimitResponse_Code = 0
	RateLimitResponse_OK         RateLimitResponse_Code = 1
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

// Enum value maps for RateLimitResponse_Code.
var (
	RateLimitResponse_Code_name = map[int32]string{
		0: "UNKNOWN",
		1: "OK",
		2: "OVER_LIMIT",
	}
	RateLimitResponse_Code_value = map[string]int32{
		"UNKNOWN":    0,
		"OK":         1,
		"OVER_LIMIT": 2,
	}
)

func (x RateLimitResponse_Code) Enum() *RateLimitResponse_Code {
	p := new(RateLimitResponse_Code)
	*p = x
	return p
}

func (x RateLimitResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_service_ratelimit_v2_rls_proto_enumTypes[0].Descriptor()
}

func (RateLimitResponse_Code) Type() protoreflect.EnumType {
	return &file_envoy_service_ratelimit_v2_rls_proto_enumTypes[0]
}

func (x RateLimitResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitResponse_Code.Descriptor instead.
func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	RateLimitResponse_RateLimit_SECOND  RateLimitResponse_RateLimit_Unit = 1
	RateLimitResponse_RateLimit_MINUTE  RateLimitResponse_RateLimit_Unit = 2
	RateLimitResponse_RateLimit_HOUR    RateLimitResponse_RateLimit_Unit = 3
	RateLimitResponse_RateLimit_DAY     RateLimitResponse_RateLimit_Unit = 4
)

// Enum value maps for RateLimitResponse_RateLimit_Unit.
var (
	RateLimitResponse_RateLimit_Unit_name = map[int32]string{
		0: "UNKNOWN",
		1: "SECOND",
		2: "MINUTE",
		3: "HOUR",
		4: "DAY",
	}
	RateLimitResponse_RateLimit_Unit_value = map[string]int32{
		"UNKNOWN": 0,
		"SECOND":  1,
		"MINUTE":  2,
		"HOUR":    3,
		"DAY":     4,
	}
)

func (x RateLimitResponse_RateLimit_Unit) Enum() *RateLimitResponse_RateLimit_Unit {
	p := new(RateLimitResponse_RateLimit_Unit)
	*p = x
	return p
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitResponse_RateLimit_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_service_ratelimit_v2_rls_proto_enumTypes[1].Descriptor()
}

func (RateLimitResponse_RateLimit_Unit) Type() protoreflect.EnumType {
	return &file_envoy_service_ratelimit_v2_rls_proto_enumTypes[1]
}

func (x RateLimitResponse_RateLimit_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitResponse_RateLimit_Unit.Descriptor instead.
func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 0, 0}
}

type RateLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain      string                           `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	HitsAddend  uint32                           `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
}

func (x *RateLimitRequest) Reset() {
	*x = RateLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitRequest) ProtoMessage() {}

func (x *RateLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitRequest.ProtoReflect.Descriptor instead.
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RateLimitRequest) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

func (x *RateLimitRequest) GetHitsAddend() uint32 {
	if x != nil {
		return x.HitsAddend
	}
	return 0
}

type RateLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverallCode         RateLimitResponse_Code                `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"overall_code,omitempty"`
	Statuses            []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	Headers             []*core.HeaderValue                   `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	RequestHeadersToAdd []*core.HeaderValue                   `protobuf:"bytes,4,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RawBody             []byte                                `protobuf:"bytes,5,opt,name=raw_body,json=rawBody,proto3" json:"raw_body,omitempty"`
	DynamicMetadata     *_struct.Struct                       `protobuf:"bytes,6,opt,name=dynamic_metadata,json=dynamicMetadata,proto3" json:"dynamic_metadata,omitempty"`
}

func (x *RateLimitResponse) Reset() {
	*x = RateLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitResponse) ProtoMessage() {}

func (x *RateLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitResponse.ProtoReflect.Descriptor instead.
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if x != nil {
		return x.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (x *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

func (x *RateLimitResponse) GetHeaders() []*core.HeaderValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *RateLimitResponse) GetRequestHeadersToAdd() []*core.HeaderValue {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

func (x *RateLimitResponse) GetRawBody() []byte {
	if x != nil {
		return x.RawBody
	}
	return nil
}

func (x *RateLimitResponse) GetDynamicMetadata() *_struct.Struct {
	if x != nil {
		return x.DynamicMetadata
	}
	return nil
}

type RateLimitResponse_RateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	RequestsPerUnit uint32                           `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	Unit            RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
}

func (x *RateLimitResponse_RateLimit) Reset() {
	*x = RateLimitResponse_RateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitResponse_RateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitResponse_RateLimit) ProtoMessage() {}

func (x *RateLimitResponse_RateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitResponse_RateLimit.ProtoReflect.Descriptor instead.
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RateLimitResponse_RateLimit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if x != nil {
		return x.RequestsPerUnit
	}
	return 0
}

func (x *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if x != nil {
		return x.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           RateLimitResponse_Code       `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"code,omitempty"`
	CurrentLimit   *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	LimitRemaining uint32                       `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
}

func (x *RateLimitResponse_DescriptorStatus) Reset() {
	*x = RateLimitResponse_DescriptorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitResponse_DescriptorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitResponse_DescriptorStatus) ProtoMessage() {}

func (x *RateLimitResponse_DescriptorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitResponse_DescriptorStatus.ProtoReflect.Descriptor instead.
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 1}
}

func (x *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if x != nil {
		return x.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (x *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if x != nil {
		return x.CurrentLimit
	}
	return nil
}

func (x *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if x != nil {
		return x.LimitRemaining
	}
	return 0
}

var File_envoy_service_ratelimit_v2_rls_proto protoreflect.FileDescriptor

var file_envoy_service_ratelimit_v2_rls_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a,
	0x01, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x4d, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x69,
	0x74, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x68, 0x69, 0x74, 0x73, 0x41, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x22, 0xc6, 0x07, 0x0a, 0x11,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x55, 0x0a, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x5a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1f, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x19, 0x0a, 0x17, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x53, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54,
	0x6f, 0x41, 0x64, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x42, 0x0a, 0x10, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0xdd, 0x01, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69,
	0x74, 0x12, 0x50, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x3c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x22, 0x3e, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41,
	0x59, 0x10, 0x04, 0x1a, 0xe1, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x46, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x5c, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x2b, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4d,
	0x49, 0x54, 0x10, 0x02, 0x32, 0x84, 0x01, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x0f, 0x53, 0x68, 0x6f,
	0x75, 0x6c, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x0a, 0x28, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x42, 0x08, 0x52, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x88, 0x01, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_ratelimit_v2_rls_proto_rawDescOnce sync.Once
	file_envoy_service_ratelimit_v2_rls_proto_rawDescData = file_envoy_service_ratelimit_v2_rls_proto_rawDesc
)

func file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP() []byte {
	file_envoy_service_ratelimit_v2_rls_proto_rawDescOnce.Do(func() {
		file_envoy_service_ratelimit_v2_rls_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_ratelimit_v2_rls_proto_rawDescData)
	})
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescData
}

var file_envoy_service_ratelimit_v2_rls_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_service_ratelimit_v2_rls_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_service_ratelimit_v2_rls_proto_goTypes = []interface{}{
	(RateLimitResponse_Code)(0),                // 0: envoy.service.ratelimit.v2.RateLimitResponse.Code
	(RateLimitResponse_RateLimit_Unit)(0),      // 1: envoy.service.ratelimit.v2.RateLimitResponse.RateLimit.Unit
	(*RateLimitRequest)(nil),                   // 2: envoy.service.ratelimit.v2.RateLimitRequest
	(*RateLimitResponse)(nil),                  // 3: envoy.service.ratelimit.v2.RateLimitResponse
	(*RateLimitResponse_RateLimit)(nil),        // 4: envoy.service.ratelimit.v2.RateLimitResponse.RateLimit
	(*RateLimitResponse_DescriptorStatus)(nil), // 5: envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus
	(*ratelimit.RateLimitDescriptor)(nil),      // 6: envoy.api.v2.ratelimit.RateLimitDescriptor
	(*core.HeaderValue)(nil),                   // 7: envoy.api.v2.core.HeaderValue
	(*_struct.Struct)(nil),                     // 8: google.protobuf.Struct
}
var file_envoy_service_ratelimit_v2_rls_proto_depIdxs = []int32{
	6,  // 0: envoy.service.ratelimit.v2.RateLimitRequest.descriptors:type_name -> envoy.api.v2.ratelimit.RateLimitDescriptor
	0,  // 1: envoy.service.ratelimit.v2.RateLimitResponse.overall_code:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.Code
	5,  // 2: envoy.service.ratelimit.v2.RateLimitResponse.statuses:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus
	7,  // 3: envoy.service.ratelimit.v2.RateLimitResponse.headers:type_name -> envoy.api.v2.core.HeaderValue
	7,  // 4: envoy.service.ratelimit.v2.RateLimitResponse.request_headers_to_add:type_name -> envoy.api.v2.core.HeaderValue
	8,  // 5: envoy.service.ratelimit.v2.RateLimitResponse.dynamic_metadata:type_name -> google.protobuf.Struct
	1,  // 6: envoy.service.ratelimit.v2.RateLimitResponse.RateLimit.unit:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.RateLimit.Unit
	0,  // 7: envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus.code:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.Code
	4,  // 8: envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus.current_limit:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.RateLimit
	2,  // 9: envoy.service.ratelimit.v2.RateLimitService.ShouldRateLimit:input_type -> envoy.service.ratelimit.v2.RateLimitRequest
	3,  // 10: envoy.service.ratelimit.v2.RateLimitService.ShouldRateLimit:output_type -> envoy.service.ratelimit.v2.RateLimitResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_envoy_service_ratelimit_v2_rls_proto_init() }
func file_envoy_service_ratelimit_v2_rls_proto_init() {
	if File_envoy_service_ratelimit_v2_rls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitRequest); i {
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
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitResponse); i {
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
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitResponse_RateLimit); i {
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
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitResponse_DescriptorStatus); i {
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
			RawDescriptor: file_envoy_service_ratelimit_v2_rls_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_ratelimit_v2_rls_proto_goTypes,
		DependencyIndexes: file_envoy_service_ratelimit_v2_rls_proto_depIdxs,
		EnumInfos:         file_envoy_service_ratelimit_v2_rls_proto_enumTypes,
		MessageInfos:      file_envoy_service_ratelimit_v2_rls_proto_msgTypes,
	}.Build()
	File_envoy_service_ratelimit_v2_rls_proto = out.File
	file_envoy_service_ratelimit_v2_rls_proto_rawDesc = nil
	file_envoy_service_ratelimit_v2_rls_proto_goTypes = nil
	file_envoy_service_ratelimit_v2_rls_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateLimitServiceClient(cc grpc.ClientConnInterface) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

// UnimplementedRateLimitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitServiceServer struct {
}

func (*UnimplementedRateLimitServiceServer) ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShouldRateLimit not implemented")
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v2.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v2/rls.proto",
}
