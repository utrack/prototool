// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uber/foo/v1/hello_api.proto

package foov1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FooRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooRequest) Reset()         { *m = FooRequest{} }
func (m *FooRequest) String() string { return proto.CompactTextString(m) }
func (*FooRequest) ProtoMessage()    {}
func (*FooRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a081c6ba99dbfe3, []int{0}
}

func (m *FooRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooRequest.Unmarshal(m, b)
}
func (m *FooRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooRequest.Marshal(b, m, deterministic)
}
func (m *FooRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooRequest.Merge(m, src)
}
func (m *FooRequest) XXX_Size() int {
	return xxx_messageInfo_FooRequest.Size(m)
}
func (m *FooRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FooRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FooRequest proto.InternalMessageInfo

type FooResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooResponse) Reset()         { *m = FooResponse{} }
func (m *FooResponse) String() string { return proto.CompactTextString(m) }
func (*FooResponse) ProtoMessage()    {}
func (*FooResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a081c6ba99dbfe3, []int{1}
}

func (m *FooResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooResponse.Unmarshal(m, b)
}
func (m *FooResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooResponse.Marshal(b, m, deterministic)
}
func (m *FooResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooResponse.Merge(m, src)
}
func (m *FooResponse) XXX_Size() int {
	return xxx_messageInfo_FooResponse.Size(m)
}
func (m *FooResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FooResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FooResponse proto.InternalMessageInfo

type BarRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarRequest) Reset()         { *m = BarRequest{} }
func (m *BarRequest) String() string { return proto.CompactTextString(m) }
func (*BarRequest) ProtoMessage()    {}
func (*BarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a081c6ba99dbfe3, []int{2}
}

func (m *BarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarRequest.Unmarshal(m, b)
}
func (m *BarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarRequest.Marshal(b, m, deterministic)
}
func (m *BarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRequest.Merge(m, src)
}
func (m *BarRequest) XXX_Size() int {
	return xxx_messageInfo_BarRequest.Size(m)
}
func (m *BarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BarRequest proto.InternalMessageInfo

func (m *BarRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BarResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarResponse) Reset()         { *m = BarResponse{} }
func (m *BarResponse) String() string { return proto.CompactTextString(m) }
func (*BarResponse) ProtoMessage()    {}
func (*BarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a081c6ba99dbfe3, []int{3}
}

func (m *BarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarResponse.Unmarshal(m, b)
}
func (m *BarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarResponse.Marshal(b, m, deterministic)
}
func (m *BarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarResponse.Merge(m, src)
}
func (m *BarResponse) XXX_Size() int {
	return xxx_messageInfo_BarResponse.Size(m)
}
func (m *BarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BarResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FooRequest)(nil), "uber.foo.v1.FooRequest")
	proto.RegisterType((*FooResponse)(nil), "uber.foo.v1.FooResponse")
	proto.RegisterType((*BarRequest)(nil), "uber.foo.v1.BarRequest")
	proto.RegisterType((*BarResponse)(nil), "uber.foo.v1.BarResponse")
}

func init() { proto.RegisterFile("uber/foo/v1/hello_api.proto", fileDescriptor_3a081c6ba99dbfe3) }

var fileDescriptor_3a081c6ba99dbfe3 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4d, 0x4a, 0x2d,
	0xd2, 0x4f, 0xcb, 0xcf, 0xd7, 0x2f, 0x33, 0xd4, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x8f, 0x4f, 0x2c,
	0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x49, 0xea, 0xa5, 0xe5, 0xe7, 0xeb,
	0x95, 0x19, 0x2a, 0xf1, 0x70, 0x71, 0xb9, 0xe5, 0xe7, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x28, 0xf1, 0x72, 0x71, 0x83, 0x79, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x32, 0x5c, 0x5c,
	0x4e, 0x89, 0x45, 0x50, 0x49, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xe6, 0x20, 0xa6, 0xcc, 0x14, 0x90, 0x62, 0xb0, 0x2c, 0x44, 0xb1, 0x51, 0x1d, 0x17, 0x87, 0x07,
	0xc8, 0x26, 0xc7, 0x00, 0x4f, 0x21, 0x0b, 0x2e, 0x66, 0xb7, 0xfc, 0x7c, 0x21, 0x71, 0x3d, 0x24,
	0xab, 0xf4, 0x10, 0xf6, 0x48, 0x49, 0x60, 0x4a, 0x40, 0x4c, 0x01, 0xe9, 0x74, 0x4a, 0x2c, 0x42,
	0xd3, 0x89, 0x70, 0x04, 0x9a, 0x4e, 0x24, 0xfb, 0x9d, 0xbc, 0xb9, 0xf8, 0x93, 0xf3, 0x73, 0x91,
	0xa5, 0x9d, 0x78, 0x21, 0x0e, 0x2a, 0xc8, 0x0c, 0x00, 0x79, 0x3c, 0x80, 0x31, 0x8a, 0x35, 0x2d,
	0x3f, 0xbf, 0xcc, 0x70, 0x11, 0x13, 0x73, 0xa8, 0x5b, 0xc4, 0x2a, 0x26, 0xee, 0x50, 0x90, 0x5a,
	0xb7, 0xfc, 0x7c, 0xbd, 0x30, 0xc3, 0x53, 0x10, 0x5e, 0x8c, 0x5b, 0x7e, 0x7e, 0x4c, 0x98, 0x61,
	0x12, 0x1b, 0x38, 0xa8, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x4c, 0x48, 0x35, 0x49,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloAPIClient is the client API for HelloAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloAPIClient interface {
	// Foo does a foo.
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
	// Bar does a bar.
	Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error)
}

type helloAPIClient struct {
	cc *grpc.ClientConn
}

func NewHelloAPIClient(cc *grpc.ClientConn) HelloAPIClient {
	return &helloAPIClient{cc}
}

func (c *helloAPIClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, "/uber.foo.v1.HelloAPI/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloAPIClient) Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error) {
	out := new(BarResponse)
	err := c.cc.Invoke(ctx, "/uber.foo.v1.HelloAPI/Bar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloAPIServer is the server API for HelloAPI service.
type HelloAPIServer interface {
	// Foo does a foo.
	Foo(context.Context, *FooRequest) (*FooResponse, error)
	// Bar does a bar.
	Bar(context.Context, *BarRequest) (*BarResponse, error)
}

func RegisterHelloAPIServer(s *grpc.Server, srv HelloAPIServer) {
	s.RegisterService(&_HelloAPI_serviceDesc, srv)
}

func _HelloAPI_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloAPIServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.foo.v1.HelloAPI/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloAPIServer).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloAPI_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloAPIServer).Bar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.foo.v1.HelloAPI/Bar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloAPIServer).Bar(ctx, req.(*BarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uber.foo.v1.HelloAPI",
	HandlerType: (*HelloAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _HelloAPI_Foo_Handler,
		},
		{
			MethodName: "Bar",
			Handler:    _HelloAPI_Bar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uber/foo/v1/hello_api.proto",
}
