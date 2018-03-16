// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package chremoas_roles is a generated protocol buffer package.

It is generated from these files:
	common.proto
	filters.proto
	roles.proto
	rules.proto

It has these top-level messages:
	NilMessage
	FilterList
	Filter
	Members
	MemberList
	StringList
	Role
	UpdateInfo
	GetRolesResponse
	SyncRoles
	SyncRolesResponse
	Rule
	RuleInfo
	RulesList
*/
package chremoas_roles

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

type NilMessage struct {
}

func (m *NilMessage) Reset()                    { *m = NilMessage{} }
func (m *NilMessage) String() string            { return proto.CompactTextString(m) }
func (*NilMessage) ProtoMessage()               {}
func (*NilMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*NilMessage)(nil), "chremoas.roles.NilMessage")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0xce, 0x28, 0x4a, 0xcd, 0xcd,
	0x4f, 0x2c, 0xd6, 0x2b, 0xca, 0xcf, 0x49, 0x2d, 0x56, 0xe2, 0xe1, 0xe2, 0xf2, 0xcb, 0xcc, 0xf1,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x4d, 0x62, 0x03, 0x2b, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x82, 0x14, 0x41, 0x5d, 0x34, 0x00, 0x00, 0x00,
}
