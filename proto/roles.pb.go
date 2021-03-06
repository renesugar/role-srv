// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roles.proto

/*
Package chremoas_roles is a generated protocol buffer package.

It is generated from these files:
	roles.proto

It has these top-level messages:
	NilMessage
	RoleMembershipRequest
	RoleMembershipResponse
	ListUserRolesRequest
	ListUserRolesResponse
	GetDiscordUserRequest
	GetDiscordUserListResponse
	GetDiscordUserResponse
	SyncRequest
	StringList
	Role
	UpdateInfo
	GetRolesResponse
	FilterList
	Filter
	Members
	MemberList
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

type RoleMembershipRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *RoleMembershipRequest) Reset()                    { *m = RoleMembershipRequest{} }
func (m *RoleMembershipRequest) String() string            { return proto.CompactTextString(m) }
func (*RoleMembershipRequest) ProtoMessage()               {}
func (*RoleMembershipRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RoleMembershipRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RoleMembershipResponse struct {
	Members []string `protobuf:"bytes,1,rep,name=Members" json:"Members,omitempty"`
}

func (m *RoleMembershipResponse) Reset()                    { *m = RoleMembershipResponse{} }
func (m *RoleMembershipResponse) String() string            { return proto.CompactTextString(m) }
func (*RoleMembershipResponse) ProtoMessage()               {}
func (*RoleMembershipResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RoleMembershipResponse) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ListUserRolesRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *ListUserRolesRequest) Reset()                    { *m = ListUserRolesRequest{} }
func (m *ListUserRolesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListUserRolesRequest) ProtoMessage()               {}
func (*ListUserRolesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListUserRolesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListUserRolesResponse struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=Roles" json:"Roles,omitempty"`
}

func (m *ListUserRolesResponse) Reset()                    { *m = ListUserRolesResponse{} }
func (m *ListUserRolesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListUserRolesResponse) ProtoMessage()               {}
func (*ListUserRolesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListUserRolesResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type GetDiscordUserRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *GetDiscordUserRequest) Reset()                    { *m = GetDiscordUserRequest{} }
func (m *GetDiscordUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDiscordUserRequest) ProtoMessage()               {}
func (*GetDiscordUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetDiscordUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetDiscordUserListResponse struct {
	Users []*GetDiscordUserResponse `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *GetDiscordUserListResponse) Reset()                    { *m = GetDiscordUserListResponse{} }
func (m *GetDiscordUserListResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDiscordUserListResponse) ProtoMessage()               {}
func (*GetDiscordUserListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetDiscordUserListResponse) GetUsers() []*GetDiscordUserResponse {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetDiscordUserResponse struct {
	Id            string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	Discriminator string `protobuf:"bytes,3,opt,name=Discriminator" json:"Discriminator,omitempty"`
	Avatar        string `protobuf:"bytes,4,opt,name=Avatar" json:"Avatar,omitempty"`
	Bot           bool   `protobuf:"varint,5,opt,name=Bot" json:"Bot,omitempty"`
	MfaEnabled    bool   `protobuf:"varint,6,opt,name=MfaEnabled" json:"MfaEnabled,omitempty"`
	Verified      bool   `protobuf:"varint,7,opt,name=Verified" json:"Verified,omitempty"`
	Email         string `protobuf:"bytes,8,opt,name=Email" json:"Email,omitempty"`
	Nick          string `protobuf:"bytes,9,opt,name=Nick" json:"Nick,omitempty"`
}

func (m *GetDiscordUserResponse) Reset()                    { *m = GetDiscordUserResponse{} }
func (m *GetDiscordUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDiscordUserResponse) ProtoMessage()               {}
func (*GetDiscordUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetDiscordUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDiscordUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetDiscordUserResponse) GetDiscriminator() string {
	if m != nil {
		return m.Discriminator
	}
	return ""
}

func (m *GetDiscordUserResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *GetDiscordUserResponse) GetBot() bool {
	if m != nil {
		return m.Bot
	}
	return false
}

func (m *GetDiscordUserResponse) GetMfaEnabled() bool {
	if m != nil {
		return m.MfaEnabled
	}
	return false
}

func (m *GetDiscordUserResponse) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *GetDiscordUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetDiscordUserResponse) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

type SyncRequest struct {
	ChannelId   string `protobuf:"bytes,1,opt,name=ChannelId" json:"ChannelId,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=UserId" json:"UserId,omitempty"`
	SendMessage bool   `protobuf:"varint,3,opt,name=SendMessage" json:"SendMessage,omitempty"`
}

func (m *SyncRequest) Reset()                    { *m = SyncRequest{} }
func (m *SyncRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()               {}
func (*SyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SyncRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SyncRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SyncRequest) GetSendMessage() bool {
	if m != nil {
		return m.SendMessage
	}
	return false
}

type StringList struct {
	Value []string `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty"`
}

func (m *StringList) Reset()                    { *m = StringList{} }
func (m *StringList) String() string            { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()               {}
func (*StringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type Role struct {
	Type      string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
	ShortName string `protobuf:"bytes,2,opt,name=ShortName" json:"ShortName,omitempty"`
	FilterA   string `protobuf:"bytes,3,opt,name=FilterA" json:"FilterA,omitempty"`
	FilterB   string `protobuf:"bytes,4,opt,name=FilterB" json:"FilterB,omitempty"`
	Sig       bool   `protobuf:"varint,5,opt,name=Sig" json:"Sig,omitempty"`
	Joinable  bool   `protobuf:"varint,6,opt,name=Joinable" json:"Joinable,omitempty"`
	// Discord
	Name        string `protobuf:"bytes,20,opt,name=Name" json:"Name,omitempty"`
	Color       int32  `protobuf:"varint,21,opt,name=Color" json:"Color,omitempty"`
	Hoist       bool   `protobuf:"varint,22,opt,name=Hoist" json:"Hoist,omitempty"`
	Position    int32  `protobuf:"varint,23,opt,name=Position" json:"Position,omitempty"`
	Permissions int32  `protobuf:"varint,24,opt,name=Permissions" json:"Permissions,omitempty"`
	Managed     bool   `protobuf:"varint,25,opt,name=Managed" json:"Managed,omitempty"`
	Mentionable bool   `protobuf:"varint,26,opt,name=Mentionable" json:"Mentionable,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Role) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Role) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *Role) GetFilterA() string {
	if m != nil {
		return m.FilterA
	}
	return ""
}

func (m *Role) GetFilterB() string {
	if m != nil {
		return m.FilterB
	}
	return ""
}

func (m *Role) GetSig() bool {
	if m != nil {
		return m.Sig
	}
	return false
}

func (m *Role) GetJoinable() bool {
	if m != nil {
		return m.Joinable
	}
	return false
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *Role) GetHoist() bool {
	if m != nil {
		return m.Hoist
	}
	return false
}

func (m *Role) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Role) GetPermissions() int32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

func (m *Role) GetManaged() bool {
	if m != nil {
		return m.Managed
	}
	return false
}

func (m *Role) GetMentionable() bool {
	if m != nil {
		return m.Mentionable
	}
	return false
}

type UpdateInfo struct {
	Name  string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=Value" json:"Value,omitempty"`
}

func (m *UpdateInfo) Reset()                    { *m = UpdateInfo{} }
func (m *UpdateInfo) String() string            { return proto.CompactTextString(m) }
func (*UpdateInfo) ProtoMessage()               {}
func (*UpdateInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpdateInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetRolesResponse struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=Roles" json:"Roles,omitempty"`
}

func (m *GetRolesResponse) Reset()                    { *m = GetRolesResponse{} }
func (m *GetRolesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRolesResponse) ProtoMessage()               {}
func (*GetRolesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetRolesResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type FilterList struct {
	FilterList []*Filter `protobuf:"bytes,1,rep,name=FilterList" json:"FilterList,omitempty"`
}

func (m *FilterList) Reset()                    { *m = FilterList{} }
func (m *FilterList) String() string            { return proto.CompactTextString(m) }
func (*FilterList) ProtoMessage()               {}
func (*FilterList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *FilterList) GetFilterList() []*Filter {
	if m != nil {
		return m.FilterList
	}
	return nil
}

type Filter struct {
	Name        string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Members struct {
	Name   []string `protobuf:"bytes,1,rep,name=Name" json:"Name,omitempty"`
	Filter string   `protobuf:"bytes,2,opt,name=Filter" json:"Filter,omitempty"`
}

func (m *Members) Reset()                    { *m = Members{} }
func (m *Members) String() string            { return proto.CompactTextString(m) }
func (*Members) ProtoMessage()               {}
func (*Members) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Members) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Members) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type MemberList struct {
	Members []string `protobuf:"bytes,1,rep,name=Members" json:"Members,omitempty"`
}

func (m *MemberList) Reset()                    { *m = MemberList{} }
func (m *MemberList) String() string            { return proto.CompactTextString(m) }
func (*MemberList) ProtoMessage()               {}
func (*MemberList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *MemberList) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*NilMessage)(nil), "chremoas.roles.NilMessage")
	proto.RegisterType((*RoleMembershipRequest)(nil), "chremoas.roles.RoleMembershipRequest")
	proto.RegisterType((*RoleMembershipResponse)(nil), "chremoas.roles.RoleMembershipResponse")
	proto.RegisterType((*ListUserRolesRequest)(nil), "chremoas.roles.ListUserRolesRequest")
	proto.RegisterType((*ListUserRolesResponse)(nil), "chremoas.roles.ListUserRolesResponse")
	proto.RegisterType((*GetDiscordUserRequest)(nil), "chremoas.roles.GetDiscordUserRequest")
	proto.RegisterType((*GetDiscordUserListResponse)(nil), "chremoas.roles.GetDiscordUserListResponse")
	proto.RegisterType((*GetDiscordUserResponse)(nil), "chremoas.roles.GetDiscordUserResponse")
	proto.RegisterType((*SyncRequest)(nil), "chremoas.roles.SyncRequest")
	proto.RegisterType((*StringList)(nil), "chremoas.roles.StringList")
	proto.RegisterType((*Role)(nil), "chremoas.roles.Role")
	proto.RegisterType((*UpdateInfo)(nil), "chremoas.roles.UpdateInfo")
	proto.RegisterType((*GetRolesResponse)(nil), "chremoas.roles.GetRolesResponse")
	proto.RegisterType((*FilterList)(nil), "chremoas.roles.FilterList")
	proto.RegisterType((*Filter)(nil), "chremoas.roles.Filter")
	proto.RegisterType((*Members)(nil), "chremoas.roles.Members")
	proto.RegisterType((*MemberList)(nil), "chremoas.roles.MemberList")
}

func init() { proto.RegisterFile("roles.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0x96, 0x64, 0xcb, 0x96, 0x46, 0xb6, 0x91, 0x10, 0xb6, 0xb2, 0xdd, 0x16, 0x85, 0x40, 0x24,
	0x86, 0x91, 0x02, 0x2a, 0xe0, 0xa2, 0xed, 0xa5, 0x0d, 0x2a, 0xcb, 0xf1, 0x4f, 0x52, 0x1b, 0xc1,
	0x2a, 0xc9, 0xa1, 0x87, 0x02, 0x6b, 0xed, 0x58, 0x22, 0xb2, 0x5a, 0xaa, 0x24, 0x63, 0x40, 0xef,
	0xd9, 0x37, 0xe8, 0x5b, 0xf4, 0x54, 0xf0, 0x67, 0x7f, 0x24, 0xad, 0xa5, 0xd4, 0xbe, 0x71, 0x86,
	0x33, 0xdf, 0xcc, 0x7c, 0x9c, 0x21, 0x09, 0x2d, 0xc1, 0x63, 0x94, 0xdd, 0xa9, 0xe0, 0x8a, 0x93,
	0xbd, 0xe1, 0x58, 0xe0, 0x84, 0x87, 0xb2, 0x6b, 0xb4, 0x74, 0x07, 0xe0, 0x9a, 0xc5, 0x57, 0x28,
	0x65, 0x38, 0x42, 0xfa, 0x1d, 0x1c, 0x04, 0x3c, 0xc6, 0x2b, 0x9c, 0xdc, 0xa0, 0x90, 0x63, 0x36,
	0x0d, 0xf0, 0xaf, 0xcf, 0x28, 0x15, 0x21, 0xb0, 0x79, 0x1d, 0x4e, 0xd0, 0xab, 0x76, 0xaa, 0x47,
	0xcd, 0xc0, 0xac, 0xe9, 0x31, 0xb4, 0x17, 0x8d, 0xe5, 0x94, 0x27, 0x12, 0x89, 0x07, 0xdb, 0x4e,
	0xeb, 0x55, 0x3b, 0x1b, 0x47, 0xcd, 0x20, 0x15, 0x69, 0x17, 0xf6, 0x7f, 0x67, 0x52, 0x7d, 0x90,
	0x28, 0xb4, 0xaf, 0x4c, 0xf1, 0xdb, 0xb0, 0xa5, 0x75, 0x97, 0x91, 0x8b, 0xe0, 0x24, 0xda, 0x87,
	0x83, 0x05, 0x7b, 0x17, 0xe2, 0x25, 0xd4, 0x8d, 0xc2, 0x04, 0x68, 0x1d, 0xef, 0x77, 0xe7, 0xeb,
	0xea, 0xea, 0xcd, 0xc0, 0x9a, 0xd0, 0xef, 0xe1, 0xe0, 0x1c, 0xd5, 0x29, 0x93, 0x43, 0x2e, 0x22,
	0x03, 0xb5, 0x26, 0xea, 0x1f, 0xe0, 0xcf, 0x3b, 0xe8, 0x1c, 0xb2, 0xd0, 0xbf, 0x40, 0x5d, 0xeb,
	0xd2, 0xd0, 0x87, 0x8b, 0xa1, 0x17, 0x63, 0x59, 0xb7, 0xc0, 0x3a, 0xd1, 0x7f, 0xab, 0xd0, 0x2e,
	0xb7, 0x20, 0x7b, 0x50, 0xcb, 0x52, 0xa9, 0x5d, 0x46, 0xc4, 0x87, 0x86, 0xde, 0x4f, 0x34, 0xf1,
	0x35, 0xa3, 0xcd, 0x64, 0xf2, 0x1c, 0x76, 0x35, 0x84, 0x60, 0x13, 0x96, 0x84, 0x8a, 0x0b, 0x6f,
	0xc3, 0x18, 0xcc, 0x2b, 0x75, 0x81, 0xbd, 0xbb, 0x50, 0x85, 0xc2, 0xdb, 0xb4, 0x05, 0x5a, 0x89,
	0x3c, 0x81, 0x8d, 0x13, 0xae, 0xbc, 0x7a, 0xa7, 0x7a, 0xd4, 0x08, 0xf4, 0x92, 0x7c, 0x0b, 0x70,
	0x75, 0x1b, 0xbe, 0x4e, 0xc2, 0x9b, 0x18, 0x23, 0x6f, 0xcb, 0x6c, 0x14, 0x34, 0x3a, 0x97, 0x8f,
	0x28, 0xd8, 0x2d, 0xc3, 0xc8, 0xdb, 0x36, 0xbb, 0x99, 0x4c, 0xf6, 0xa1, 0xfe, 0x7a, 0x12, 0xb2,
	0xd8, 0x6b, 0x98, 0x20, 0x56, 0x30, 0x2d, 0xc3, 0x86, 0x9f, 0xbc, 0xa6, 0x6b, 0x19, 0x36, 0xfc,
	0x44, 0x11, 0x5a, 0x83, 0x59, 0x32, 0x4c, 0xf9, 0xff, 0x06, 0x9a, 0xfd, 0x71, 0x98, 0x24, 0x18,
	0x67, 0x75, 0xe7, 0x8a, 0xc2, 0xe9, 0xd4, 0x8a, 0xa7, 0x43, 0x3a, 0xd0, 0x1a, 0x60, 0x12, 0xb9,
	0x9e, 0x35, 0x85, 0x37, 0x82, 0xa2, 0x8a, 0x52, 0x80, 0x81, 0x12, 0x2c, 0x19, 0xe9, 0x73, 0xd3,
	0xe9, 0x7d, 0x0c, 0xe3, 0xcf, 0xe8, 0x7a, 0xd1, 0x0a, 0xf4, 0xef, 0x1a, 0x6c, 0xea, 0xf6, 0xd0,
	0x79, 0xbe, 0x9f, 0x4d, 0xb3, 0xd6, 0xd6, 0x6b, 0x9d, 0xd8, 0x60, 0xcc, 0x85, 0xba, 0xce, 0xa9,
	0xcf, 0x15, 0xba, 0xbd, 0xcf, 0x58, 0xac, 0x50, 0xf4, 0x1c, 0xeb, 0xa9, 0x98, 0xef, 0x9c, 0x38,
	0xc2, 0x53, 0x51, 0x33, 0x3e, 0x60, 0xa3, 0x94, 0xf1, 0x01, 0x1b, 0x69, 0x46, 0xdf, 0x70, 0x66,
	0xe8, 0x75, 0x7c, 0x67, 0x72, 0x36, 0x6e, 0xfb, 0xf9, 0xb8, 0xe9, 0x32, 0xfa, 0x3c, 0xe6, 0xc2,
	0x3b, 0xe8, 0x54, 0x8f, 0xea, 0x81, 0x15, 0xb4, 0xf6, 0x82, 0x33, 0xa9, 0xbc, 0xb6, 0x81, 0xb0,
	0x82, 0xc6, 0x7e, 0xc7, 0x25, 0x53, 0x8c, 0x27, 0xde, 0x33, 0x63, 0x9e, 0xc9, 0x9a, 0xbe, 0x77,
	0x28, 0x26, 0x4c, 0x4a, 0xc6, 0x13, 0xe9, 0x79, 0x66, 0xbb, 0xa8, 0x32, 0xe3, 0x1b, 0x26, 0xe1,
	0x08, 0x23, 0xef, 0x2b, 0x83, 0x9a, 0x8a, 0xda, 0xf7, 0x0a, 0x13, 0x0d, 0x63, 0xd2, 0xf6, 0x2d,
	0xf5, 0x05, 0x15, 0xbd, 0x00, 0xf8, 0x30, 0x8d, 0x42, 0x85, 0x97, 0xc9, 0x2d, 0x2f, 0xbb, 0x36,
	0x34, 0x13, 0x6f, 0x71, 0xe6, 0x58, 0xd5, 0xcb, 0xfc, 0x80, 0x2c, 0x9b, 0xee, 0x80, 0x5e, 0xc1,
	0x93, 0x73, 0x54, 0x0f, 0x9f, 0xfa, 0x53, 0x00, 0x4b, 0xbe, 0x69, 0x82, 0x9f, 0x8a, 0x92, 0x73,
	0x6f, 0x2f, 0xba, 0x5b, 0x8b, 0xa0, 0x60, 0x49, 0x5f, 0xc1, 0x96, 0x95, 0x4a, 0x6b, 0xe9, 0x40,
	0xeb, 0x14, 0xf5, 0xc0, 0x4d, 0x0d, 0xd5, 0xb6, 0xa6, 0xa2, 0x8a, 0xfe, 0x98, 0x5d, 0x85, 0x05,
	0x80, 0x8d, 0x0c, 0xa0, 0x9d, 0xc2, 0xa7, 0x3d, 0x6e, 0x25, 0x7a, 0x08, 0x60, 0xdd, 0x4c, 0xf2,
	0xf7, 0xde, 0xa7, 0xc7, 0xff, 0x80, 0x63, 0x84, 0xfc, 0x0a, 0xdb, 0xbd, 0x28, 0x32, 0x1d, 0x5d,
	0x4a, 0x8b, 0xef, 0x2f, 0x6a, 0x0b, 0xf7, 0x7e, 0x85, 0x9c, 0xa5, 0xe7, 0x66, 0x10, 0x96, 0x6c,
	0xf3, 0x33, 0x5d, 0x83, 0xf3, 0x1b, 0x40, 0x80, 0x13, 0x7e, 0x87, 0x0f, 0xce, 0xe4, 0x0d, 0x34,
	0xd2, 0x73, 0x27, 0x2b, 0x2c, 0xfd, 0x4e, 0xc9, 0xbd, 0x3b, 0xd7, 0x2d, 0xb4, 0x42, 0x7e, 0x86,
	0x6d, 0xa7, 0xbd, 0x27, 0x95, 0x52, 0x2d, 0xad, 0x90, 0x73, 0x68, 0x39, 0xc7, 0xb7, 0x38, 0x5b,
	0x9d, 0xc7, 0xd2, 0x5e, 0x7e, 0xf5, 0xd0, 0x0a, 0xb9, 0x80, 0x1d, 0x07, 0xa4, 0x2f, 0x96, 0xc7,
	0x20, 0x45, 0xf0, 0xd4, 0x21, 0xe5, 0x2f, 0x2e, 0x79, 0x51, 0x96, 0xff, 0xd2, 0xf3, 0xed, 0x1f,
	0xae, 0x33, 0xcb, 0x18, 0xfb, 0x13, 0x76, 0xe7, 0x1e, 0x5c, 0xf2, 0x7c, 0xd1, 0xb5, 0xec, 0xfd,
	0xf6, 0x5f, 0xac, 0xb1, 0xca, 0xf0, 0xcf, 0x00, 0xce, 0x51, 0xd9, 0x2e, 0xff, 0x9f, 0x6c, 0x14,
	0xa6, 0xb2, 0x42, 0x7a, 0xd0, 0xec, 0x45, 0x91, 0x1b, 0xcd, 0x7b, 0x06, 0x79, 0x4d, 0xa3, 0x9d,
	0xc2, 0x8e, 0x6d, 0xd5, 0x47, 0xa1, 0x9c, 0x98, 0x82, 0xd2, 0x19, 0xff, 0x62, 0x8c, 0x7c, 0xba,
	0x69, 0x85, 0xf4, 0x01, 0x7a, 0x51, 0x94, 0x62, 0x3c, 0x2b, 0xb7, 0x95, 0x6b, 0x27, 0x78, 0xd7,
	0x96, 0xf3, 0x48, 0x9c, 0x6b, 0x78, 0xaa, 0xdf, 0xe8, 0xf7, 0xbc, 0x3f, 0x0e, 0xd5, 0x00, 0xc5,
	0x1d, 0x1b, 0x22, 0xf9, 0x7a, 0xa9, 0x35, 0xf3, 0x67, 0x7c, 0x0d, 0x5e, 0x08, 0x7b, 0xf3, 0xff,
	0x9d, 0xe5, 0xa6, 0x2d, 0xfd, 0x9d, 0xf9, 0x5f, 0xf8, 0xb1, 0x32, 0x4d, 0x4b, 0x96, 0xff, 0x6b,
	0x2b, 0x9b, 0xeb, 0xe5, 0x6a, 0xec, 0xe2, 0x7f, 0x8f, 0x56, 0x6e, 0xb6, 0xcc, 0xdf, 0xf9, 0x87,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x6f, 0x96, 0x3d, 0x4a, 0x0b, 0x00, 0x00,
}
