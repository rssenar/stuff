// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	User
*/
package internal

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

type User struct {
	Name    string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=Address" json:"Address,omitempty"`
	Age     int64  `protobuf:"varint,3,opt,name=Age" json:"Age,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "internal.User")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xdc,
	0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0x76, 0xc7, 0x94, 0x94, 0xa2, 0xd4,
	0xe2, 0x62, 0x09, 0x26, 0xb0, 0x30, 0x8c, 0x2b, 0x24, 0xc0, 0xc5, 0xec, 0x98, 0x9e, 0x2a, 0xc1,
	0xac, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0x62, 0x26, 0xb1, 0x81, 0x0d, 0x36, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xa3, 0x5a, 0x5e, 0xbb, 0x6a, 0x00, 0x00, 0x00,
}