// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

/*
Package resources is a generated protocol buffer package.

It is generated from these files:
	resources.proto

It has these top-level messages:
	Data
	None
*/
package resources

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

type Data struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type None struct {
}

func (m *None) Reset()                    { *m = None{} }
func (m *None) String() string            { return proto.CompactTextString(m) }
func (*None) ProtoMessage()               {}
func (*None) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Data)(nil), "resources.Data")
	proto.RegisterType((*None)(nil), "resources.None")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Resources service

type ResourcesClient interface {
	GetData(ctx context.Context, in *None, opts ...grpc.CallOption) (*Data, error)
	StreamListData(ctx context.Context, in *None, opts ...grpc.CallOption) (Resources_StreamListDataClient, error)
	StreamRandData(ctx context.Context, in *None, opts ...grpc.CallOption) (Resources_StreamRandDataClient, error)
}

type resourcesClient struct {
	cc *grpc.ClientConn
}

func NewResourcesClient(cc *grpc.ClientConn) ResourcesClient {
	return &resourcesClient{cc}
}

func (c *resourcesClient) GetData(ctx context.Context, in *None, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := grpc.Invoke(ctx, "/resources.Resources/GetData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) StreamListData(ctx context.Context, in *None, opts ...grpc.CallOption) (Resources_StreamListDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Resources_serviceDesc.Streams[0], c.cc, "/resources.Resources/StreamListData", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourcesStreamListDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Resources_StreamListDataClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type resourcesStreamListDataClient struct {
	grpc.ClientStream
}

func (x *resourcesStreamListDataClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resourcesClient) StreamRandData(ctx context.Context, in *None, opts ...grpc.CallOption) (Resources_StreamRandDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Resources_serviceDesc.Streams[1], c.cc, "/resources.Resources/StreamRandData", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourcesStreamRandDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Resources_StreamRandDataClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type resourcesStreamRandDataClient struct {
	grpc.ClientStream
}

func (x *resourcesStreamRandDataClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Resources service

type ResourcesServer interface {
	GetData(context.Context, *None) (*Data, error)
	StreamListData(*None, Resources_StreamListDataServer) error
	StreamRandData(*None, Resources_StreamRandDataServer) error
}

func RegisterResourcesServer(s *grpc.Server, srv ResourcesServer) {
	s.RegisterService(&_Resources_serviceDesc, srv)
}

func _Resources_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).GetData(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_StreamListData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(None)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourcesServer).StreamListData(m, &resourcesStreamListDataServer{stream})
}

type Resources_StreamListDataServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type resourcesStreamListDataServer struct {
	grpc.ServerStream
}

func (x *resourcesStreamListDataServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _Resources_StreamRandData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(None)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourcesServer).StreamRandData(m, &resourcesStreamRandDataServer{stream})
}

type Resources_StreamRandDataServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type resourcesStreamRandDataServer struct {
	grpc.ServerStream
}

func (x *resourcesStreamRandDataServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

var _Resources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resources.Resources",
	HandlerType: (*ResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _Resources_GetData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamListData",
			Handler:       _Resources_StreamListData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamRandData",
			Handler:       _Resources_StreamRandData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "resources.proto",
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x49, 0x71, 0xb1, 0xb8, 0x24, 0x96, 0x24, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x6c, 0x5c, 0x2c, 0x7e, 0xf9, 0x79, 0xa9,
	0x46, 0xab, 0x18, 0xb9, 0x38, 0x83, 0x60, 0x3a, 0x84, 0x74, 0xb9, 0xd8, 0xdd, 0x53, 0x4b, 0xc0,
	0x9a, 0xf8, 0xf5, 0x10, 0x26, 0x83, 0x54, 0x4a, 0x21, 0x0b, 0x80, 0x54, 0x28, 0x31, 0x08, 0x99,
	0x71, 0xf1, 0x05, 0x97, 0x14, 0xa5, 0x26, 0xe6, 0xfa, 0x64, 0x16, 0x13, 0xad, 0xcb, 0x80, 0x11,
	0xa1, 0x2f, 0x28, 0x31, 0x2f, 0x85, 0x78, 0x7d, 0x49, 0x6c, 0x60, 0x2f, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x82, 0xf0, 0x3e, 0x45, 0xf5, 0x00, 0x00, 0x00,
}
