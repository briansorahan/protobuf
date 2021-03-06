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
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

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
	AField string                       `protobuf:"bytes,1,opt,name=a_field,json=aField,proto3" json:"a"`
	BField string                       `protobuf:"bytes,2,opt,name=b_field,json=bField,proto3" json:"b"`
	CField *google_protobuf1.Int64Value `protobuf:"bytes,3,opt,name=c_field,json=cField" json:"c"`
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

func (m *MsgTags) GetCField() *google_protobuf1.Int64Value {
	if m != nil {
		return m.CField
	}
	return nil
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
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2a, 0xca, 0x4c, 0xcc, 0x2b, 0x49, 0x2d, 0x2e,
	0x91, 0xe2, 0x4a, 0xcf, 0x4f, 0xcf, 0x87, 0x08, 0x4b, 0x59, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x83, 0x25, 0x92, 0x4a,
	0xd3, 0xf4, 0x0b, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0xcb, 0x8b, 0x12, 0x0b, 0x0a, 0x52, 0x8b,
	0x10, 0x0c, 0x88, 0x56, 0xa5, 0x36, 0x46, 0x2e, 0x76, 0xdf, 0xe2, 0xf4, 0x90, 0xc4, 0xf4, 0x62,
	0x21, 0x39, 0x2e, 0xf6, 0xc4, 0xf8, 0xb4, 0xcc, 0xd4, 0x9c, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x4e, 0x27, 0xd6, 0x57, 0xf7, 0xe4, 0x19, 0x13, 0x83, 0xd8, 0x12, 0xdd, 0x40, 0x82, 0x20, 0xf9,
	0x24, 0xa8, 0x3c, 0x13, 0x42, 0x3e, 0x29, 0x88, 0x2d, 0x09, 0x22, 0x6f, 0xcd, 0xc5, 0x9e, 0x0c,
	0x95, 0x67, 0x56, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd6, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x83, 0xb9, 0x46, 0xcf, 0x33, 0xaf, 0xc4, 0xcc, 0x24, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0xa2, 0x39,
	0x39, 0x88, 0x2d, 0x19, 0xac, 0x59, 0xc9, 0x9c, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0x48, 0x1c, 0xcd,
	0x0d, 0x70, 0xcb, 0xc5, 0xd1, 0x2c, 0x87, 0xd9, 0x9a, 0xc4, 0x06, 0x36, 0xdc, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0xb1, 0x86, 0x4d, 0x27, 0x01, 0x00, 0x00,
}
