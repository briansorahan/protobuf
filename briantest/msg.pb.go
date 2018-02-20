// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

/*
Package briantest is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	MsgTags
	Msg
*/
package briantest

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MsgTags struct {
	AField string `protobuf:"bytes,1,opt,name=a_field,json=aField,proto3" json:"a"`
	BField string `protobuf:"bytes,2,opt,name=b_field,json=bField,proto3" json:"b"`
}

func (m *MsgTags) Reset()                    { *m = MsgTags{} }
func (m *MsgTags) String() string            { return proto.CompactTextString(m) }
func (*MsgTags) ProtoMessage()               {}
func (*MsgTags) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{0} }

func (m *MsgTags) GetAField() string {
	if m != nil {
		return m.AField
	}
	return ""
}

func (m *MsgTags) GetBField() string {
	if m != nil {
		return m.BField
	}
	return ""
}

type Msg struct {
	AField string `protobuf:"bytes,1,opt,name=a_field,json=aField,proto3" json:"a_field,omitempty"`
	BField string `protobuf:"bytes,2,opt,name=b_field,json=bField,proto3" json:"b_field,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{1} }

func (m *Msg) GetAField() string {
	if m != nil {
		return m.AField
	}
	return ""
}

func (m *Msg) GetBField() string {
	if m != nil {
		return m.BField
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgTags)(nil), "briantest.MsgTags")
	proto.RegisterType((*Msg)(nil), "briantest.Msg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptorMsg) }

var fileDescriptorMsg = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2a, 0xca, 0x4c, 0xcc, 0x2b, 0x49, 0x2d, 0x2e,
	0x91, 0xe2, 0x4a, 0xcf, 0x4f, 0xcf, 0x87, 0x08, 0x2b, 0x79, 0x72, 0xb1, 0xfb, 0x16, 0xa7, 0x87,
	0x24, 0xa6, 0x17, 0x0b, 0xc9, 0x71, 0xb1, 0x27, 0xc6, 0xa7, 0x65, 0xa6, 0xe6, 0xa4, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xb1, 0xbe, 0xba, 0x27, 0xcf, 0x98, 0x18, 0xc4, 0x96, 0xe8, 0x06,
	0x12, 0x04, 0xc9, 0x27, 0x41, 0xe5, 0x99, 0x10, 0xf2, 0x49, 0x41, 0x6c, 0x49, 0x60, 0x79, 0x25,
	0x73, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0x21, 0x71, 0x34, 0x63, 0xe0, 0xfa, 0xc5, 0xd1, 0xf4, 0xc3,
	0x34, 0x26, 0xb1, 0x81, 0x9d, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xed, 0xb6, 0x24, 0x45,
	0xae, 0x00, 0x00, 0x00,
}
