// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rules.proto

package chremoas_roles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Rule struct {
	Name    string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Role    string `protobuf:"bytes,2,opt,name=Role" json:"Role,omitempty"`
	FilterA string `protobuf:"bytes,3,opt,name=filterA" json:"filterA,omitempty"`
	FilterB string `protobuf:"bytes,4,opt,name=filterB" json:"filterB,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Rule) GetFilterA() string {
	if m != nil {
		return m.FilterA
	}
	return ""
}

func (m *Rule) GetFilterB() string {
	if m != nil {
		return m.FilterB
	}
	return ""
}

type RemoveRuleRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *RemoveRuleRequest) Reset()                    { *m = RemoveRuleRequest{} }
func (m *RemoveRuleRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRuleRequest) ProtoMessage()               {}
func (*RemoveRuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *RemoveRuleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RulesList struct {
	Rules []*Rule `protobuf:"bytes,1,rep,name=Rules" json:"Rules,omitempty"`
}

func (m *RulesList) Reset()                    { *m = RulesList{} }
func (m *RulesList) String() string            { return proto.CompactTextString(m) }
func (*RulesList) ProtoMessage()               {}
func (*RulesList) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *RulesList) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "chremoas.roles.Rule")
	proto.RegisterType((*RemoveRuleRequest)(nil), "chremoas.roles.RemoveRuleRequest")
	proto.RegisterType((*RulesList)(nil), "chremoas.roles.RulesList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Rules service

type RulesClient interface {
	AddRule(ctx context.Context, in *Rule, opts ...client.CallOption) (*NilMessage, error)
	UpdateRule(ctx context.Context, in *Rule, opts ...client.CallOption) (*NilMessage, error)
	RemoveRule(ctx context.Context, in *RemoveRuleRequest, opts ...client.CallOption) (*NilMessage, error)
	GetRules(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*RulesList, error)
}

type rulesClient struct {
	c           client.Client
	serviceName string
}

func NewRulesClient(serviceName string, c client.Client) RulesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "chremoas.roles"
	}
	return &rulesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *rulesClient) AddRule(ctx context.Context, in *Rule, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Rules.AddRule", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) UpdateRule(ctx context.Context, in *Rule, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Rules.UpdateRule", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) RemoveRule(ctx context.Context, in *RemoveRuleRequest, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Rules.RemoveRule", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) GetRules(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*RulesList, error) {
	req := c.c.NewRequest(c.serviceName, "Rules.GetRules", in)
	out := new(RulesList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rules service

type RulesHandler interface {
	AddRule(context.Context, *Rule, *NilMessage) error
	UpdateRule(context.Context, *Rule, *NilMessage) error
	RemoveRule(context.Context, *RemoveRuleRequest, *NilMessage) error
	GetRules(context.Context, *NilMessage, *RulesList) error
}

func RegisterRulesHandler(s server.Server, hdlr RulesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Rules{hdlr}, opts...))
}

type Rules struct {
	RulesHandler
}

func (h *Rules) AddRule(ctx context.Context, in *Rule, out *NilMessage) error {
	return h.RulesHandler.AddRule(ctx, in, out)
}

func (h *Rules) UpdateRule(ctx context.Context, in *Rule, out *NilMessage) error {
	return h.RulesHandler.UpdateRule(ctx, in, out)
}

func (h *Rules) RemoveRule(ctx context.Context, in *RemoveRuleRequest, out *NilMessage) error {
	return h.RulesHandler.RemoveRule(ctx, in, out)
}

func (h *Rules) GetRules(ctx context.Context, in *NilMessage, out *RulesList) error {
	return h.RulesHandler.GetRules(ctx, in, out)
}

func init() { proto.RegisterFile("rules.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x6d, 0xda, 0x40, 0xe9, 0x15, 0x21, 0x61, 0x31, 0x98, 0x4c, 0xc5, 0x0b, 0x15, 0x43, 0x86,
	0x32, 0x30, 0x21, 0xd1, 0x32, 0xb0, 0xd0, 0x0e, 0x96, 0xf8, 0x80, 0xb4, 0x39, 0x20, 0x92, 0xcd,
	0x15, 0x9f, 0xc3, 0xd7, 0xf0, 0xb1, 0x28, 0x36, 0xa8, 0x50, 0x22, 0x90, 0xba, 0xdd, 0xbd, 0xf7,
	0xfc, 0xce, 0xef, 0x0e, 0x86, 0xae, 0x36, 0xc8, 0xf9, 0xda, 0x91, 0x27, 0x71, 0xb4, 0x7a, 0x76,
	0x68, 0xa9, 0xe0, 0xdc, 0x91, 0x41, 0xce, 0x0e, 0x57, 0x64, 0x2d, 0xbd, 0x44, 0x56, 0x2d, 0x21,
	0xd5, 0xb5, 0x41, 0x21, 0x20, 0x5d, 0x14, 0x16, 0x65, 0x32, 0x4a, 0xc6, 0x03, 0x1d, 0xea, 0x06,
	0xd3, 0x64, 0x50, 0x76, 0x23, 0xd6, 0xd4, 0x42, 0x42, 0xff, 0xb1, 0x32, 0x1e, 0xdd, 0x54, 0xf6,
	0x02, 0xfc, 0xd5, 0x6e, 0x98, 0x99, 0x4c, 0xbf, 0x33, 0x33, 0x75, 0x0e, 0xc7, 0x1a, 0x2d, 0xbd,
	0x61, 0x33, 0x49, 0xe3, 0x6b, 0x8d, 0xec, 0xdb, 0x06, 0xaa, 0x2b, 0x18, 0x34, 0x12, 0xbe, 0xaf,
	0xd8, 0x8b, 0x0b, 0xd8, 0x0b, 0x8d, 0x4c, 0x46, 0xbd, 0xf1, 0x70, 0x72, 0x92, 0xff, 0xcc, 0x91,
	0x07, 0xb3, 0x28, 0x99, 0xbc, 0x77, 0x3f, 0xc5, 0xe2, 0x1a, 0xfa, 0xd3, 0xb2, 0x0c, 0x91, 0x5a,
	0x5f, 0x64, 0xd9, 0x36, 0xba, 0xa8, 0xcc, 0x1c, 0x99, 0x8b, 0x27, 0x54, 0x1d, 0x71, 0x03, 0xf0,
	0xb0, 0x2e, 0x0b, 0x8f, 0x3b, 0x3b, 0xcc, 0x01, 0x36, 0x61, 0xc5, 0xd9, 0x2f, 0x87, 0xed, 0x45,
	0xfc, 0x63, 0x77, 0x0b, 0x07, 0x77, 0xe8, 0x63, 0xb6, 0x3f, 0x94, 0xd9, 0x69, 0xdb, 0x57, 0xc3,
	0x22, 0x55, 0x67, 0xb9, 0x1f, 0x6e, 0x7d, 0xf9, 0x11, 0x00, 0x00, 0xff, 0xff, 0xed, 0x17, 0x64,
	0xef, 0x18, 0x02, 0x00, 0x00,
}
