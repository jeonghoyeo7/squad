// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/liquidstaking/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a37bd8b89a8d11ee, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a37bd8b89a8d11ee, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryLiquidValidatorsRequest is the request type for the Query/LiquidValidators RPC method.
type QueryLiquidValidatorsRequest struct {
	// TODO: status, etc query strings
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryLiquidValidatorsRequest) Reset()         { *m = QueryLiquidValidatorsRequest{} }
func (m *QueryLiquidValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidValidatorsRequest) ProtoMessage()    {}
func (*QueryLiquidValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a37bd8b89a8d11ee, []int{2}
}
func (m *QueryLiquidValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidValidatorsRequest.Merge(m, src)
}
func (m *QueryLiquidValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidValidatorsRequest proto.InternalMessageInfo

func (m *QueryLiquidValidatorsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryLiquidValidatorsResponse is the response type for the Query/LiquidValidators RPC method.
type QueryLiquidValidatorsResponse struct {
	LiquidValidators []LiquidValidator `protobuf:"bytes,1,rep,name=liquid_validators,json=liquidValidators,proto3" json:"liquid_validators"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryLiquidValidatorsResponse) Reset()         { *m = QueryLiquidValidatorsResponse{} }
func (m *QueryLiquidValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidValidatorsResponse) ProtoMessage()    {}
func (*QueryLiquidValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a37bd8b89a8d11ee, []int{3}
}
func (m *QueryLiquidValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidValidatorsResponse.Merge(m, src)
}
func (m *QueryLiquidValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidValidatorsResponse proto.InternalMessageInfo

func (m *QueryLiquidValidatorsResponse) GetLiquidValidators() []LiquidValidator {
	if m != nil {
		return m.LiquidValidators
	}
	return nil
}

func (m *QueryLiquidValidatorsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "crescent.liquidstaking.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "crescent.liquidstaking.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryLiquidValidatorsRequest)(nil), "crescent.liquidstaking.v1beta1.QueryLiquidValidatorsRequest")
	proto.RegisterType((*QueryLiquidValidatorsResponse)(nil), "crescent.liquidstaking.v1beta1.QueryLiquidValidatorsResponse")
}

func init() {
	proto.RegisterFile("crescent/liquidstaking/v1beta1/query.proto", fileDescriptor_a37bd8b89a8d11ee)
}

var fileDescriptor_a37bd8b89a8d11ee = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x4e, 0xd4, 0x4e,
	0x1c, 0xdf, 0xee, 0xef, 0x07, 0x87, 0xe1, 0x82, 0x23, 0x07, 0xb2, 0xc1, 0x3a, 0xe9, 0x01, 0x91,
	0xd8, 0x8e, 0x14, 0x13, 0x13, 0xd4, 0x03, 0xc4, 0xe0, 0xc5, 0x44, 0xdd, 0x18, 0x0f, 0x7a, 0x20,
	0xd3, 0xf6, 0x4b, 0xb7, 0xa1, 0x9d, 0xe9, 0x76, 0xa6, 0x2b, 0x78, 0xf4, 0x09, 0x0c, 0x3e, 0x85,
	0x17, 0x9f, 0x83, 0x83, 0x89, 0x44, 0x2f, 0x26, 0x26, 0xc6, 0x00, 0xcf, 0xe0, 0xd9, 0xec, 0x4c,
	0xa1, 0xbb, 0x4b, 0xb0, 0xfe, 0x39, 0xed, 0x64, 0xf6, 0xfb, 0xf9, 0x3b, 0x33, 0x45, 0xcb, 0x61,
	0x01, 0x32, 0x04, 0xae, 0x68, 0x9a, 0xf4, 0xcb, 0x24, 0x92, 0x8a, 0xed, 0x24, 0x3c, 0xa6, 0x83,
	0x95, 0x00, 0x14, 0x5b, 0xa1, 0xfd, 0x12, 0x8a, 0x3d, 0x2f, 0x2f, 0x84, 0x12, 0xd8, 0x3e, 0x9d,
	0xf5, 0xc6, 0x66, 0xbd, 0x6a, 0xb6, 0xb3, 0x10, 0x0b, 0x11, 0xa7, 0x40, 0x59, 0x9e, 0x50, 0xc6,
	0xb9, 0x50, 0x4c, 0x25, 0x82, 0x4b, 0x83, 0xee, 0x2c, 0x87, 0x42, 0x66, 0x42, 0xd2, 0x80, 0x49,
	0x30, 0xb4, 0x67, 0x22, 0x39, 0x8b, 0x13, 0xae, 0x87, 0xab, 0x59, 0xbf, 0xc1, 0xd5, 0xb8, 0xbe,
	0xc1, 0xcc, 0xc5, 0x22, 0x16, 0x7a, 0x49, 0x87, 0xab, 0x6a, 0xd7, 0xfc, 0x84, 0x6e, 0x0c, 0xdc,
	0x15, 0x39, 0x70, 0x96, 0x27, 0x03, 0x9f, 0x8a, 0x5c, 0x3b, 0x3b, 0xef, 0xd2, 0x99, 0x43, 0xf8,
	0xc9, 0xd0, 0xdb, 0x63, 0x56, 0xb0, 0x4c, 0x76, 0xa1, 0x5f, 0x82, 0x54, 0xce, 0x0b, 0x74, 0x79,
	0x6c, 0x57, 0xe6, 0x82, 0x4b, 0xc0, 0xf7, 0xd1, 0x74, 0xae, 0x77, 0xe6, 0x2d, 0x62, 0x2d, 0xcd,
	0xf8, 0x8b, 0xde, 0xaf, 0x1b, 0xf2, 0x0c, 0x7e, 0xe3, 0xff, 0x83, 0x6f, 0x57, 0x5b, 0xdd, 0x0a,
	0xeb, 0x6c, 0xa3, 0x05, 0x4d, 0xfe, 0x50, 0x43, 0x9e, 0xb1, 0x34, 0x89, 0x98, 0x12, 0xc5, 0xa9,
	0x38, 0xde, 0x44, 0xa8, 0x2e, 0xa8, 0x56, 0xd2, 0x6d, 0x7a, 0xc3, 0x36, 0x3d, 0x73, 0x48, 0xb5,
	0x48, 0x0c, 0x15, 0xb6, 0x3b, 0x82, 0x74, 0x3e, 0x58, 0xe8, 0xca, 0x05, 0x42, 0x55, 0x9e, 0x00,
	0x5d, 0x32, 0xbe, 0xb7, 0x06, 0x67, 0x7f, 0xce, 0x5b, 0xe4, 0xbf, 0xa5, 0x19, 0x9f, 0x36, 0x45,
	0x9b, 0x20, 0xad, 0x32, 0xce, 0xa6, 0x13, 0x5a, 0xf8, 0xc1, 0x58, 0x9a, 0xb6, 0x4e, 0x73, 0xad,
	0x31, 0x8d, 0x31, 0x38, 0x1a, 0xc7, 0xff, 0x34, 0x85, 0xa6, 0x74, 0x1c, 0xfc, 0xb1, 0x8d, 0xa6,
	0x4d, 0xb3, 0xd8, 0x6f, 0xb2, 0x79, 0xfe, 0x70, 0x3b, 0xab, 0x7f, 0x84, 0x31, 0x4e, 0x9c, 0xaf,
	0xd6, 0xfe, 0xfa, 0x3b, 0xab, 0x73, 0xab, 0x0b, 0xaa, 0x2c, 0xb8, 0x24, 0x2c, 0x4d, 0x89, 0x3e,
	0x4f, 0x50, 0x50, 0x48, 0x22, 0xb6, 0x89, 0xea, 0x01, 0x31, 0x7c, 0xa4, 0x22, 0x24, 0x99, 0x88,
	0xca, 0x14, 0x3c, 0xa7, 0x8f, 0xec, 0xcd, 0x84, 0x47, 0x44, 0x94, 0x8a, 0x64, 0xa2, 0x00, 0xc2,
	0x82, 0xe1, 0x72, 0x88, 0x30, 0x77, 0x02, 0x3f, 0xea, 0x29, 0x95, 0xcb, 0x35, 0x4a, 0xe3, 0x44,
	0xf5, 0xca, 0xc0, 0x0b, 0x45, 0x46, 0x03, 0xb7, 0xc7, 0x8a, 0x01, 0x48, 0x45, 0xe3, 0x82, 0x0d,
	0x12, 0xb5, 0xe7, 0x46, 0xb0, 0xeb, 0xbe, 0x12, 0x1c, 0xe8, 0xee, 0xc4, 0x3b, 0x91, 0x39, 0x84,
	0xf4, 0xe6, 0xed, 0x2d, 0xc3, 0xe6, 0x65, 0xd1, 0xeb, 0xcf, 0x27, 0x6f, 0xdb, 0x4b, 0x78, 0x91,
	0x36, 0x3c, 0xac, 0x4a, 0xfe, 0x47, 0x1b, 0xcd, 0x4e, 0xde, 0x12, 0x7c, 0xf7, 0xb7, 0x7a, 0xba,
	0xe0, 0x16, 0x77, 0xee, 0xfd, 0x25, 0xba, 0xea, 0xfb, 0xc4, 0xda, 0x5f, 0x7f, 0x6f, 0x75, 0xee,
	0x8c, 0xf6, 0x5d, 0xb5, 0x5b, 0xdf, 0xd5, 0x86, 0xda, 0x05, 0xba, 0x7e, 0x51, 0xed, 0xe7, 0xa8,
	0xf0, 0xc6, 0xbf, 0x9f, 0x80, 0x2e, 0xfd, 0x06, 0x5e, 0x6e, 0x2a, 0xbd, 0x56, 0xdd, 0x78, 0x7a,
	0x70, 0x64, 0x5b, 0x87, 0x47, 0xb6, 0xf5, 0xfd, 0xc8, 0xb6, 0xde, 0x1c, 0xdb, 0xad, 0xc3, 0x63,
	0xbb, 0xf5, 0xe5, 0xd8, 0x6e, 0x3d, 0x5f, 0x1b, 0xf1, 0x71, 0xca, 0xe7, 0x72, 0x50, 0x2f, 0x45,
	0xb1, 0x53, 0x0b, 0x4c, 0xda, 0x50, 0x7b, 0x39, 0xc8, 0x60, 0x5a, 0x7f, 0xdb, 0x56, 0x7f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x91, 0x51, 0xd1, 0xed, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns parameters of the liquidstaking module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// LiquidValidators returns liquid validators of the liquidstaking module.
	LiquidValidators(ctx context.Context, in *QueryLiquidValidatorsRequest, opts ...grpc.CallOption) (*QueryLiquidValidatorsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/crescent.liquidstaking.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidValidators(ctx context.Context, in *QueryLiquidValidatorsRequest, opts ...grpc.CallOption) (*QueryLiquidValidatorsResponse, error) {
	out := new(QueryLiquidValidatorsResponse)
	err := c.cc.Invoke(ctx, "/crescent.liquidstaking.v1beta1.Query/LiquidValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns parameters of the liquidstaking module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// LiquidValidators returns liquid validators of the liquidstaking module.
	LiquidValidators(context.Context, *QueryLiquidValidatorsRequest) (*QueryLiquidValidatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) LiquidValidators(ctx context.Context, req *QueryLiquidValidatorsRequest) (*QueryLiquidValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidValidators not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crescent.liquidstaking.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crescent.liquidstaking.v1beta1.Query/LiquidValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidValidators(ctx, req.(*QueryLiquidValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crescent.liquidstaking.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "LiquidValidators",
			Handler:    _Query_LiquidValidators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crescent/liquidstaking/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryLiquidValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLiquidValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.LiquidValidators) > 0 {
		for iNdEx := len(m.LiquidValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryLiquidValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLiquidValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LiquidValidators) > 0 {
		for _, e := range m.LiquidValidators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLiquidValidatorsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLiquidValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidValidators = append(m.LiquidValidators, LiquidValidator{})
			if err := m.LiquidValidators[len(m.LiquidValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
