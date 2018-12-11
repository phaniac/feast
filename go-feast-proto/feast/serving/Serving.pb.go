// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/serving/Serving.proto

package serving // import "github.com/gojektech/feast/go-feast-proto/feast/serving"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gojektech/feast/go-feast-proto/feast/types"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestType int32

const (
	// LAST : return only the latest value based on timestamp. (default)
	RequestType_LAST RequestType = 0
	// LIST : return list of historical data sorted by timestamp.
	RequestType_LIST RequestType = 1
)

var RequestType_name = map[int32]string{
	0: "LAST",
	1: "LIST",
}
var RequestType_value = map[string]int32{
	"LAST": 0,
	"LIST": 1,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{0}
}

type QueryFeatures struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryFeatures) Reset()         { *m = QueryFeatures{} }
func (m *QueryFeatures) String() string { return proto.CompactTextString(m) }
func (*QueryFeatures) ProtoMessage()    {}
func (*QueryFeatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{0}
}
func (m *QueryFeatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFeatures.Unmarshal(m, b)
}
func (m *QueryFeatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFeatures.Marshal(b, m, deterministic)
}
func (dst *QueryFeatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeatures.Merge(dst, src)
}
func (m *QueryFeatures) XXX_Size() int {
	return xxx_messageInfo_QueryFeatures.Size(m)
}
func (m *QueryFeatures) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeatures.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeatures proto.InternalMessageInfo

type QueryFeatures_Request struct {
	// e.g. "driver", "customer", "city".
	EntityName string `protobuf:"bytes,1,opt,name=entityName,proto3" json:"entityName,omitempty"`
	// List of entity ID.
	EntityId []string `protobuf:"bytes,2,rep,name=entityId,proto3" json:"entityId,omitempty"`
	// List of request details, contains: featureId, type of query, and limit.
	RequestDetails []*RequestDetail `protobuf:"bytes,3,rep,name=requestDetails,proto3" json:"requestDetails,omitempty"`
	// Filter specifying only to retrieve features having timestamp within this range.
	TimestampRange       *TimestampRange `protobuf:"bytes,4,opt,name=timestampRange,proto3" json:"timestampRange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *QueryFeatures_Request) Reset()         { *m = QueryFeatures_Request{} }
func (m *QueryFeatures_Request) String() string { return proto.CompactTextString(m) }
func (*QueryFeatures_Request) ProtoMessage()    {}
func (*QueryFeatures_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{0, 0}
}
func (m *QueryFeatures_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFeatures_Request.Unmarshal(m, b)
}
func (m *QueryFeatures_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFeatures_Request.Marshal(b, m, deterministic)
}
func (dst *QueryFeatures_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeatures_Request.Merge(dst, src)
}
func (m *QueryFeatures_Request) XXX_Size() int {
	return xxx_messageInfo_QueryFeatures_Request.Size(m)
}
func (m *QueryFeatures_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeatures_Request.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeatures_Request proto.InternalMessageInfo

func (m *QueryFeatures_Request) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *QueryFeatures_Request) GetEntityId() []string {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *QueryFeatures_Request) GetRequestDetails() []*RequestDetail {
	if m != nil {
		return m.RequestDetails
	}
	return nil
}

func (m *QueryFeatures_Request) GetTimestampRange() *TimestampRange {
	if m != nil {
		return m.TimestampRange
	}
	return nil
}

type QueryFeatures_Response struct {
	// e.g. "driver", "customer", "city".
	EntityName string `protobuf:"bytes,1,opt,name=entityName,proto3" json:"entityName,omitempty"`
	// map of entity ID and its entity's properties.
	Entities             map[string]*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QueryFeatures_Response) Reset()         { *m = QueryFeatures_Response{} }
func (m *QueryFeatures_Response) String() string { return proto.CompactTextString(m) }
func (*QueryFeatures_Response) ProtoMessage()    {}
func (*QueryFeatures_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{0, 1}
}
func (m *QueryFeatures_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFeatures_Response.Unmarshal(m, b)
}
func (m *QueryFeatures_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFeatures_Response.Marshal(b, m, deterministic)
}
func (dst *QueryFeatures_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeatures_Response.Merge(dst, src)
}
func (m *QueryFeatures_Response) XXX_Size() int {
	return xxx_messageInfo_QueryFeatures_Response.Size(m)
}
func (m *QueryFeatures_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeatures_Response.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeatures_Response proto.InternalMessageInfo

func (m *QueryFeatures_Response) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *QueryFeatures_Response) GetEntities() map[string]*Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

type RequestDetail struct {
	// feature ID to be included in the query.
	// feature ID is in the form of [entity_name].[granularity].[feature_name]
	// e.g: "driver.day.total_accepted_booking"
	// all requested feature ID shall have same entity name.
	FeatureId string `protobuf:"bytes,1,opt,name=featureId,proto3" json:"featureId,omitempty"`
	// request type either LAST or LIST.
	// LAST : return only the latest value based on timestamp.
	// LIST : return list of historical data sorted by timestamp.
	Type RequestType `protobuf:"varint,2,opt,name=type,proto3,enum=feast.serving.RequestType" json:"type,omitempty"`
	// only applicable to LIST.
	// length of the returned list <= limit.
	// default = 0
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestDetail) Reset()         { *m = RequestDetail{} }
func (m *RequestDetail) String() string { return proto.CompactTextString(m) }
func (*RequestDetail) ProtoMessage()    {}
func (*RequestDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{1}
}
func (m *RequestDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestDetail.Unmarshal(m, b)
}
func (m *RequestDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestDetail.Marshal(b, m, deterministic)
}
func (dst *RequestDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestDetail.Merge(dst, src)
}
func (m *RequestDetail) XXX_Size() int {
	return xxx_messageInfo_RequestDetail.Size(m)
}
func (m *RequestDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RequestDetail proto.InternalMessageInfo

func (m *RequestDetail) GetFeatureId() string {
	if m != nil {
		return m.FeatureId
	}
	return ""
}

func (m *RequestDetail) GetType() RequestType {
	if m != nil {
		return m.Type
	}
	return RequestType_LAST
}

func (m *RequestDetail) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// range of timestamp for querying
// valid timestamp range is having start <= end
type TimestampRange struct {
	// start time of the range query.
	Start *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// end time of the range query.
	End                  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimestampRange) Reset()         { *m = TimestampRange{} }
func (m *TimestampRange) String() string { return proto.CompactTextString(m) }
func (*TimestampRange) ProtoMessage()    {}
func (*TimestampRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{2}
}
func (m *TimestampRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimestampRange.Unmarshal(m, b)
}
func (m *TimestampRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimestampRange.Marshal(b, m, deterministic)
}
func (dst *TimestampRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimestampRange.Merge(dst, src)
}
func (m *TimestampRange) XXX_Size() int {
	return xxx_messageInfo_TimestampRange.Size(m)
}
func (m *TimestampRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimestampRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimestampRange proto.InternalMessageInfo

func (m *TimestampRange) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TimestampRange) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type Entity struct {
	// map of feature ID and its feature value.
	Features             map[string]*FeatureValueList `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{3}
}
func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (dst *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(dst, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetFeatures() map[string]*FeatureValueList {
	if m != nil {
		return m.Features
	}
	return nil
}

type FeatureValueList struct {
	// list of feature value
	// if "type" in "requestDetail" is "LAST", then the length will always be 1.
	// if "type" in "requestDetail" is "LIST", then the length is <= "limit".
	ValueList *types.ValueList `protobuf:"bytes,1,opt,name=valueList,proto3" json:"valueList,omitempty"`
	// list of timestamp of the value.
	// the i-th timestamps correspond to the i-th value.
	TimestampList        *types.TimestampList `protobuf:"bytes,2,opt,name=timestampList,proto3" json:"timestampList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeatureValueList) Reset()         { *m = FeatureValueList{} }
func (m *FeatureValueList) String() string { return proto.CompactTextString(m) }
func (*FeatureValueList) ProtoMessage()    {}
func (*FeatureValueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Serving_fd09065a28d57eb8, []int{4}
}
func (m *FeatureValueList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureValueList.Unmarshal(m, b)
}
func (m *FeatureValueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureValueList.Marshal(b, m, deterministic)
}
func (dst *FeatureValueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureValueList.Merge(dst, src)
}
func (m *FeatureValueList) XXX_Size() int {
	return xxx_messageInfo_FeatureValueList.Size(m)
}
func (m *FeatureValueList) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureValueList.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureValueList proto.InternalMessageInfo

func (m *FeatureValueList) GetValueList() *types.ValueList {
	if m != nil {
		return m.ValueList
	}
	return nil
}

func (m *FeatureValueList) GetTimestampList() *types.TimestampList {
	if m != nil {
		return m.TimestampList
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryFeatures)(nil), "feast.serving.QueryFeatures")
	proto.RegisterType((*QueryFeatures_Request)(nil), "feast.serving.QueryFeatures.Request")
	proto.RegisterType((*QueryFeatures_Response)(nil), "feast.serving.QueryFeatures.Response")
	proto.RegisterMapType((map[string]*Entity)(nil), "feast.serving.QueryFeatures.Response.EntitiesEntry")
	proto.RegisterType((*RequestDetail)(nil), "feast.serving.RequestDetail")
	proto.RegisterType((*TimestampRange)(nil), "feast.serving.TimestampRange")
	proto.RegisterType((*Entity)(nil), "feast.serving.Entity")
	proto.RegisterMapType((map[string]*FeatureValueList)(nil), "feast.serving.Entity.FeaturesEntry")
	proto.RegisterType((*FeatureValueList)(nil), "feast.serving.FeatureValueList")
	proto.RegisterEnum("feast.serving.RequestType", RequestType_name, RequestType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServingAPIClient is the client API for ServingAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServingAPIClient interface {
	// Query features from Feast
	QueryFeatures(ctx context.Context, in *QueryFeatures_Request, opts ...grpc.CallOption) (*QueryFeatures_Response, error)
}

type servingAPIClient struct {
	cc *grpc.ClientConn
}

func NewServingAPIClient(cc *grpc.ClientConn) ServingAPIClient {
	return &servingAPIClient{cc}
}

func (c *servingAPIClient) QueryFeatures(ctx context.Context, in *QueryFeatures_Request, opts ...grpc.CallOption) (*QueryFeatures_Response, error) {
	out := new(QueryFeatures_Response)
	err := c.cc.Invoke(ctx, "/feast.serving.ServingAPI/QueryFeatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServingAPIServer is the server API for ServingAPI service.
type ServingAPIServer interface {
	// Query features from Feast
	QueryFeatures(context.Context, *QueryFeatures_Request) (*QueryFeatures_Response, error)
}

func RegisterServingAPIServer(s *grpc.Server, srv ServingAPIServer) {
	s.RegisterService(&_ServingAPI_serviceDesc, srv)
}

func _ServingAPI_QueryFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeatures_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingAPIServer).QueryFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.serving.ServingAPI/QueryFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingAPIServer).QueryFeatures(ctx, req.(*QueryFeatures_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServingAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feast.serving.ServingAPI",
	HandlerType: (*ServingAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryFeatures",
			Handler:    _ServingAPI_QueryFeatures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feast/serving/Serving.proto",
}

func init() {
	proto.RegisterFile("feast/serving/Serving.proto", fileDescriptor_Serving_fd09065a28d57eb8)
}

var fileDescriptor_Serving_fd09065a28d57eb8 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xef, 0x6e, 0xd3, 0x3e,
	0x14, 0x5d, 0x96, 0x75, 0xbf, 0xf6, 0x56, 0xe9, 0xaf, 0xb2, 0xf8, 0x13, 0x85, 0xc1, 0x4a, 0x01,
	0xa9, 0x02, 0xe6, 0xa0, 0x0c, 0x04, 0xe2, 0x0b, 0x6c, 0xda, 0x90, 0x2a, 0x4d, 0x30, 0xbc, 0x0a,
	0x09, 0x84, 0x90, 0xb2, 0xed, 0x36, 0x0b, 0x6b, 0x93, 0x10, 0x3b, 0x95, 0xf2, 0x0a, 0xbc, 0x08,
	0x6f, 0xc0, 0x6b, 0x20, 0xf1, 0x44, 0xa8, 0xb6, 0xd3, 0x26, 0x51, 0xc5, 0xf8, 0x14, 0xdb, 0xf7,
	0x9c, 0x73, 0x73, 0xef, 0xb1, 0x2f, 0xdc, 0x1a, 0xa3, 0xcf, 0x85, 0xcb, 0x31, 0x9d, 0x85, 0x51,
	0xe0, 0x9e, 0xa8, 0x2f, 0x4d, 0xd2, 0x58, 0xc4, 0xc4, 0x92, 0x41, 0xaa, 0x83, 0xce, 0x76, 0x10,
	0xc7, 0xc1, 0x04, 0x5d, 0x19, 0x3c, 0xcd, 0xc6, 0xae, 0x08, 0xa7, 0xc8, 0x85, 0x3f, 0x4d, 0x14,
	0xde, 0xb9, 0xa9, 0xc4, 0x44, 0x9e, 0x20, 0x77, 0x3f, 0xf8, 0x93, 0x0c, 0x55, 0xa0, 0xff, 0xd3,
	0x04, 0xeb, 0x7d, 0x86, 0x69, 0xfe, 0x06, 0x7d, 0x91, 0xa5, 0xc8, 0x9d, 0xdf, 0x06, 0xfc, 0xc7,
	0xf0, 0x5b, 0x86, 0x5c, 0x90, 0x3b, 0x00, 0x18, 0x89, 0x50, 0xe4, 0x6f, 0xfd, 0x29, 0xda, 0x46,
	0xcf, 0x18, 0xb4, 0x58, 0xe9, 0x84, 0x38, 0xd0, 0x54, 0xbb, 0xe1, 0xb9, 0xbd, 0xde, 0x33, 0x07,
	0x2d, 0xb6, 0xd8, 0x93, 0x03, 0xe8, 0xa4, 0x4a, 0xe6, 0x00, 0x85, 0x1f, 0x4e, 0xb8, 0x6d, 0xf6,
	0xcc, 0x41, 0xdb, 0xdb, 0xa2, 0x95, 0x7f, 0xa7, 0xac, 0x0c, 0x62, 0x35, 0x0e, 0x39, 0x84, 0xce,
	0xa2, 0x16, 0xe6, 0x47, 0x01, 0xda, 0x1b, 0x3d, 0x63, 0xd0, 0xf6, 0x6e, 0xd7, 0x54, 0x46, 0x15,
	0x10, 0xab, 0x91, 0x9c, 0x5f, 0x06, 0x34, 0x19, 0xf2, 0x24, 0x8e, 0x38, 0x5e, 0x59, 0xd5, 0x3b,
	0x5d, 0x55, 0x88, 0x5c, 0x56, 0xd5, 0xf6, 0x76, 0x6b, 0xd9, 0x2a, 0x1d, 0xa3, 0x85, 0x30, 0x3d,
	0xd4, 0xac, 0xc3, 0x48, 0xa4, 0x39, 0x5b, 0x88, 0x38, 0x0c, 0xac, 0x4a, 0x88, 0x74, 0xc1, 0xbc,
	0xc4, 0x5c, 0xa7, 0x9e, 0x2f, 0xc9, 0x23, 0x68, 0xcc, 0xe6, 0xb6, 0xd8, 0xeb, 0xb2, 0xbc, 0xeb,
	0xb5, 0x84, 0x92, 0x9e, 0x33, 0x85, 0x79, 0xb9, 0xfe, 0xc2, 0xe8, 0x73, 0xb0, 0x2a, 0x9d, 0x23,
	0x5b, 0xd0, 0x1a, 0xab, 0x3f, 0x1a, 0x9e, 0x6b, 0xe5, 0xe5, 0x01, 0xa1, 0xb0, 0x31, 0x37, 0x5f,
	0xca, 0x77, 0x3c, 0x67, 0xb5, 0x07, 0xa3, 0x3c, 0x41, 0x26, 0x71, 0xe4, 0x1a, 0x34, 0x26, 0xe1,
	0x34, 0x14, 0xb6, 0xd9, 0x33, 0x06, 0x0d, 0xa6, 0x36, 0xfd, 0x04, 0x3a, 0xd5, 0x46, 0x93, 0x27,
	0xd0, 0xe0, 0xc2, 0x4f, 0x85, 0xcc, 0xd8, 0xf6, 0x1c, 0xaa, 0x6e, 0x22, 0x2d, 0x6e, 0x62, 0xc9,
	0x18, 0x05, 0x24, 0x8f, 0xc1, 0xc4, 0xe8, 0x5c, 0xd7, 0xf9, 0x37, 0xfc, 0x1c, 0xd6, 0xff, 0x61,
	0xc0, 0xa6, 0x2a, 0x9e, 0xbc, 0x82, 0xa6, 0xae, 0x87, 0xdb, 0x86, 0xb4, 0xe5, 0xde, 0xca, 0x2e,
	0xd1, 0xc2, 0x18, 0x6d, 0x43, 0x41, 0x72, 0x3e, 0x83, 0x55, 0x09, 0xad, 0xb0, 0xe1, 0x59, 0xd5,
	0x86, 0xed, 0x5a, 0x02, 0x4d, 0x97, 0x0f, 0xe8, 0x28, 0xe4, 0xa2, 0x6c, 0xc8, 0x77, 0x03, 0xba,
	0xf5, 0x38, 0x79, 0x0a, 0xad, 0x59, 0xb1, 0xd1, 0x2d, 0xba, 0xa1, 0x35, 0xe5, 0x5b, 0xa4, 0x4b,
	0xa9, 0x25, 0x90, 0xbc, 0x06, 0x6b, 0x71, 0x7f, 0x25, 0xb3, 0x68, 0x56, 0x99, 0x39, 0x2a, 0x23,
	0x58, 0x95, 0xf0, 0xf0, 0x2e, 0xb4, 0x4b, 0x9e, 0x92, 0x26, 0x6c, 0x1c, 0xed, 0x9d, 0x8c, 0xba,
	0x6b, 0x72, 0x35, 0x3c, 0x19, 0x75, 0x0d, 0x6f, 0x02, 0xa0, 0x67, 0xca, 0xde, 0xf1, 0x90, 0x7c,
	0xa9, 0x8d, 0x01, 0x72, 0xff, 0x8a, 0x2b, 0x2f, 0xc5, 0x9d, 0x07, 0xff, 0xf4, 0x30, 0xfa, 0x6b,
	0xfb, 0x1f, 0xa1, 0x3a, 0xb2, 0xf6, 0xff, 0x5f, 0x26, 0x3f, 0x9e, 0x7b, 0xff, 0xe9, 0x79, 0x10,
	0x8a, 0x8b, 0xec, 0x94, 0x9e, 0xc5, 0x53, 0x37, 0x88, 0xbf, 0xe2, 0xa5, 0xc0, 0xb3, 0x0b, 0x57,
	0xcd, 0xad, 0x20, 0xde, 0x91, 0x8b, 0x1d, 0x79, 0x4d, 0xdc, 0xca, 0x64, 0x3c, 0xdd, 0x94, 0x87,
	0xbb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x47, 0x65, 0xae, 0x98, 0x31, 0x05, 0x00, 0x00,
}
