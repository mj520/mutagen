// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/prompting/prompting.proto

package prompting

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PromptRequest struct {
	Prompter             string   `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	Prompt               string   `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromptRequest) Reset()         { *m = PromptRequest{} }
func (m *PromptRequest) String() string { return proto.CompactTextString(m) }
func (*PromptRequest) ProtoMessage()    {}
func (*PromptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed1be0639533e2a, []int{0}
}

func (m *PromptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromptRequest.Unmarshal(m, b)
}
func (m *PromptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromptRequest.Marshal(b, m, deterministic)
}
func (m *PromptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromptRequest.Merge(m, src)
}
func (m *PromptRequest) XXX_Size() int {
	return xxx_messageInfo_PromptRequest.Size(m)
}
func (m *PromptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PromptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PromptRequest proto.InternalMessageInfo

func (m *PromptRequest) GetPrompter() string {
	if m != nil {
		return m.Prompter
	}
	return ""
}

func (m *PromptRequest) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type PromptResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromptResponse) Reset()         { *m = PromptResponse{} }
func (m *PromptResponse) String() string { return proto.CompactTextString(m) }
func (*PromptResponse) ProtoMessage()    {}
func (*PromptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed1be0639533e2a, []int{1}
}

func (m *PromptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromptResponse.Unmarshal(m, b)
}
func (m *PromptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromptResponse.Marshal(b, m, deterministic)
}
func (m *PromptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromptResponse.Merge(m, src)
}
func (m *PromptResponse) XXX_Size() int {
	return xxx_messageInfo_PromptResponse.Size(m)
}
func (m *PromptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PromptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PromptResponse proto.InternalMessageInfo

func (m *PromptResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PromptRequest)(nil), "prompting.PromptRequest")
	proto.RegisterType((*PromptResponse)(nil), "prompting.PromptResponse")
}

func init() { proto.RegisterFile("service/prompting/prompting.proto", fileDescriptor_aed1be0639533e2a) }

var fileDescriptor_aed1be0639533e2a = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0xcf, 0x2d, 0x28, 0xc9, 0xcc, 0x4b, 0x47, 0xb0, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0xe1, 0x02, 0x4a, 0xce, 0x5c, 0xbc, 0x01, 0x60, 0x4e,
	0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x14, 0x17, 0x07, 0x44, 0x36, 0xb5, 0x48, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x17, 0x12, 0xe3, 0x62, 0x83, 0xb0, 0x25, 0x98, 0xc0,
	0x32, 0x50, 0x9e, 0x92, 0x0e, 0x17, 0x1f, 0xcc, 0x90, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90,
	0x29, 0x45, 0x50, 0x36, 0xcc, 0x14, 0x18, 0xdf, 0xc8, 0x87, 0x8b, 0x33, 0x00, 0x66, 0xbf, 0x90,
	0x3d, 0x17, 0x1b, 0x84, 0x23, 0x24, 0xa1, 0x87, 0x70, 0x26, 0x8a, 0x93, 0xa4, 0x24, 0xb1, 0xc8,
	0x40, 0xcc, 0x52, 0x62, 0x70, 0x32, 0x8d, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0x4f, 0xcd, 0xd3, 0xcd, 0xcc, 0x87, 0x31, 0xf5,
	0x0b, 0xb2, 0xd3, 0xf5, 0x31, 0x82, 0x24, 0x89, 0x0d, 0x1c, 0x12, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0xb8, 0xfe, 0xcd, 0x2e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PromptingClient is the client API for Prompting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PromptingClient interface {
	Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptResponse, error)
}

type promptingClient struct {
	cc grpc.ClientConnInterface
}

func NewPromptingClient(cc grpc.ClientConnInterface) PromptingClient {
	return &promptingClient{cc}
}

func (c *promptingClient) Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptResponse, error) {
	out := new(PromptResponse)
	err := c.cc.Invoke(ctx, "/prompting.Prompting/Prompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromptingServer is the server API for Prompting service.
type PromptingServer interface {
	Prompt(context.Context, *PromptRequest) (*PromptResponse, error)
}

// UnimplementedPromptingServer can be embedded to have forward compatible implementations.
type UnimplementedPromptingServer struct {
}

func (*UnimplementedPromptingServer) Prompt(ctx context.Context, req *PromptRequest) (*PromptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prompt not implemented")
}

func RegisterPromptingServer(s *grpc.Server, srv PromptingServer) {
	s.RegisterService(&_Prompting_serviceDesc, srv)
}

func _Prompting_Prompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptingServer).Prompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prompting.Prompting/Prompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptingServer).Prompt(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Prompting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prompting.Prompting",
	HandlerType: (*PromptingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prompt",
			Handler:    _Prompting_Prompt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/prompting/prompting.proto",
}
