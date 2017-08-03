// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

/*
Package resources is a generated protocol buffer package.

It is generated from these files:
	resources.proto

It has these top-level messages:
	None
	User
*/
package resources

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type None struct {
}

func (m *None) Reset()                    { *m = None{} }
func (m *None) String() string            { return proto.CompactTextString(m) }
func (*None) ProtoMessage()               {}
func (*None) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age       string `protobuf:"bytes,2,opt,name=age" json:"age,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname" json:"firstname,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() string {
	if m != nil {
		return m.Age
	}
	return ""
}

func (m *User) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func init() {
	proto.RegisterType((*None)(nil), "resources.None")
	proto.RegisterType((*User)(nil), "resources.User")
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xb1, 0x71, 0xb1, 0xf8, 0xe5, 0xe7, 0xa5, 0x2a, 0x79, 0x71, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x42, 0x02, 0x5c, 0xcc, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x21, 0x10, 0x53, 0x48, 0x86, 0x8b,
	0x33, 0x2d, 0xb3, 0xa8, 0xb8, 0x04, 0xac, 0x94, 0x19, 0x2c, 0x8e, 0x10, 0x30, 0xb2, 0xe4, 0xe2,
	0x08, 0x82, 0x5a, 0x20, 0xa4, 0xcb, 0xc5, 0xee, 0x9e, 0x5a, 0x02, 0x36, 0x9a, 0x5f, 0x0f, 0xe1,
	0x0e, 0x90, 0x9d, 0x52, 0xc8, 0x02, 0x20, 0x15, 0x4a, 0x0c, 0x4e, 0x3a, 0x5c, 0x12, 0x99, 0xf9,
	0x7a, 0xe9, 0x45, 0x05, 0xc9, 0x7a, 0xe5, 0x19, 0x95, 0x79, 0xf9, 0x25, 0x08, 0x25, 0x4e, 0x7c,
	0x30, 0x43, 0x8b, 0x03, 0x40, 0xbe, 0x08, 0x60, 0x4c, 0x62, 0x03, 0x7b, 0xc7, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x00, 0x43, 0xc6, 0x25, 0xe1, 0x00, 0x00, 0x00,
}
