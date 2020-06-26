// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package proto

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

type EchoResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResp) Reset()         { *m = EchoResp{} }
func (m *EchoResp) String() string { return proto.CompactTextString(m) }
func (*EchoResp) ProtoMessage()    {}
func (*EchoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *EchoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResp.Unmarshal(m, b)
}
func (m *EchoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResp.Marshal(b, m, deterministic)
}
func (m *EchoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResp.Merge(m, src)
}
func (m *EchoResp) XXX_Size() int {
	return xxx_messageInfo_EchoResp.Size(m)
}
func (m *EchoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResp.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResp proto.InternalMessageInfo

func (m *EchoResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EchoReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoReq) Reset()         { *m = EchoReq{} }
func (m *EchoReq) String() string { return proto.CompactTextString(m) }
func (*EchoReq) ProtoMessage()    {}
func (*EchoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{1}
}

func (m *EchoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoReq.Unmarshal(m, b)
}
func (m *EchoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoReq.Marshal(b, m, deterministic)
}
func (m *EchoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoReq.Merge(m, src)
}
func (m *EchoReq) XXX_Size() int {
	return xxx_messageInfo_EchoReq.Size(m)
}
func (m *EchoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoReq.DiscardUnknown(m)
}

var xxx_messageInfo_EchoReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EchoResp)(nil), "proto.EchoResp")
	proto.RegisterType((*EchoReq)(nil), "proto.EchoReq")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x32, 0x5c, 0x1c, 0xae, 0xc9,
	0x19, 0xf9, 0x41, 0xa9, 0xc5, 0x05, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x12, 0x27, 0x17, 0x3b, 0x44, 0xb6, 0xd0, 0xc8, 0x04, 0xc2,
	0x0c, 0x2e, 0x4b, 0x16, 0xd2, 0xe4, 0x62, 0x01, 0x31, 0x85, 0xf8, 0x20, 0x46, 0xe9, 0x41, 0x95,
	0x48, 0xf1, 0xa3, 0xf0, 0x8b, 0x0b, 0x94, 0x18, 0x9c, 0x0c, 0xb8, 0xa4, 0x33, 0xf3, 0xf5, 0xd2,
	0x8b, 0x0a, 0x92, 0xf5, 0x52, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0x8b, 0xf5, 0x32, 0x52, 0x73,
	0x72, 0xf2, 0xcb, 0xf3, 0x8b, 0x72, 0x52, 0x9c, 0xf8, 0x3d, 0x40, 0xec, 0x70, 0x10, 0x3b, 0x00,
	0xa4, 0x35, 0x80, 0x31, 0x89, 0x0d, 0x6c, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x52, 0x45,
	0x7c, 0xc0, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoSvcClient is the client API for EchoSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoSvcClient interface {
	Echo(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoResp, error)
}

type echoSvcClient struct {
	cc *grpc.ClientConn
}

func NewEchoSvcClient(cc *grpc.ClientConn) EchoSvcClient {
	return &echoSvcClient{cc}
}

func (c *echoSvcClient) Echo(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoResp, error) {
	out := new(EchoResp)
	err := c.cc.Invoke(ctx, "/proto.EchoSvc/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoSvcServer is the server API for EchoSvc service.
type EchoSvcServer interface {
	Echo(context.Context, *EchoReq) (*EchoResp, error)
}

func RegisterEchoSvcServer(s *grpc.Server, srv EchoSvcServer) {
	s.RegisterService(&_EchoSvc_serviceDesc, srv)
}

func _EchoSvc_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoSvcServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EchoSvc/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoSvcServer).Echo(ctx, req.(*EchoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EchoSvc",
	HandlerType: (*EchoSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoSvc_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}
