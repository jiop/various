// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

/*
Package request is a generated protocol buffer package.

It is generated from these files:
	request.proto

It has these top-level messages:
	CityMessage
	PostalCodeMessage
*/
package request

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type CityMessage struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CityMessage) Reset()                    { *m = CityMessage{} }
func (m *CityMessage) String() string            { return proto.CompactTextString(m) }
func (*CityMessage) ProtoMessage()               {}
func (*CityMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CityMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PostalCodeMessage struct {
	Postalcode string `protobuf:"bytes,1,opt,name=postalcode" json:"postalcode,omitempty"`
}

func (m *PostalCodeMessage) Reset()                    { *m = PostalCodeMessage{} }
func (m *PostalCodeMessage) String() string            { return proto.CompactTextString(m) }
func (*PostalCodeMessage) ProtoMessage()               {}
func (*PostalCodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PostalCodeMessage) GetPostalcode() string {
	if m != nil {
		return m.Postalcode
	}
	return ""
}

func init() {
	proto.RegisterType((*CityMessage)(nil), "request.CityMessage")
	proto.RegisterType((*PostalCodeMessage)(nil), "request.PostalCodeMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Request service

type RequestClient interface {
	GetCP(ctx context.Context, in *CityMessage, opts ...grpc.CallOption) (*PostalCodeMessage, error)
}

type requestClient struct {
	cc *grpc.ClientConn
}

func NewRequestClient(cc *grpc.ClientConn) RequestClient {
	return &requestClient{cc}
}

func (c *requestClient) GetCP(ctx context.Context, in *CityMessage, opts ...grpc.CallOption) (*PostalCodeMessage, error) {
	out := new(PostalCodeMessage)
	err := grpc.Invoke(ctx, "/request.Request/GetCP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Request service

type RequestServer interface {
	GetCP(context.Context, *CityMessage) (*PostalCodeMessage, error)
}

func RegisterRequestServer(s *grpc.Server, srv RequestServer) {
	s.RegisterService(&_Request_serviceDesc, srv)
}

func _Request_GetCP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServer).GetCP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.Request/GetCP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServer).GetCP(ctx, req.(*CityMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Request_serviceDesc = grpc.ServiceDesc{
	ServiceName: "request.Request",
	HandlerType: (*RequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCP",
			Handler:    _Request_GetCP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "request.proto",
}

func init() { proto.RegisterFile("request.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x14, 0xb9,
	0xb8, 0x9d, 0x33, 0x4b, 0x2a, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x84, 0xb8, 0x58,
	0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x63, 0x2e,
	0xc1, 0x80, 0xfc, 0xe2, 0x92, 0xc4, 0x1c, 0xe7, 0xfc, 0x94, 0x54, 0x98, 0x42, 0x39, 0x2e, 0xae,
	0x02, 0xb0, 0x60, 0x72, 0x7e, 0x0a, 0x4c, 0x39, 0x92, 0x88, 0x91, 0x1b, 0x17, 0x7b, 0x10, 0xc4,
	0x0a, 0x21, 0x6b, 0x2e, 0x56, 0xf7, 0xd4, 0x12, 0xe7, 0x00, 0x21, 0x11, 0x3d, 0x98, 0x23, 0x90,
	0xac, 0x94, 0x92, 0x82, 0x8b, 0x62, 0xd8, 0xa2, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xaf, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x58, 0x06, 0xbb, 0xef, 0xc0, 0x00, 0x00, 0x00,
}
