// Code generated by protoc-gen-go.
// source: greet.proto
// DO NOT EDIT!

/*
Package greetpb is a generated protocol buffer package.

It is generated from these files:
	greet.proto

It has these top-level messages:
	Greeting
	GreetRequest
	GreetResponse
*/
package greetpb

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

type Greeting struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
}

func (m *Greeting) Reset()                    { *m = Greeting{} }
func (m *Greeting) String() string            { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()               {}
func (*Greeting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetRequest) Reset()                    { *m = GreetRequest{} }
func (m *GreetRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()               {}
func (*GreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *GreetResponse) Reset()                    { *m = GreetResponse{} }
func (m *GreetResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()               {}
func (*GreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GreetService service

type GreetServiceClient interface {
	// Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := grpc.Invoke(ctx, "/greet.GreetService/Greet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GreetService service

type GreetServiceServer interface {
	// Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xdc, 0xb8, 0x38, 0xdc,
	0x41, 0x8c, 0xcc, 0xbc, 0x74, 0x21, 0x59, 0x2e, 0xae, 0xb4, 0xcc, 0xa2, 0xe2, 0x92, 0xf8, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0, 0x88, 0x1f, 0x50, 0x40,
	0x48, 0x9a, 0x8b, 0x33, 0x27, 0x11, 0x26, 0xcb, 0x04, 0x96, 0xe5, 0x00, 0x09, 0x80, 0x24, 0x95,
	0xac, 0xb9, 0x78, 0xc0, 0xe6, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x69, 0x73, 0x71,
	0xa4, 0x43, 0xcd, 0x05, 0x9b, 0xc4, 0x6d, 0xc4, 0xaf, 0x07, 0xb1, 0x1e, 0x66, 0x5d, 0x10, 0x5c,
	0x81, 0x92, 0x3a, 0x17, 0x2f, 0x54, 0x73, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17,
	0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0xd4, 0x15, 0x50, 0x9e, 0x91, 0x0b, 0xd4, 0x96, 0xe0,
	0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x13, 0x2e, 0x56, 0x30, 0x5f, 0x48, 0x18, 0xd9, 0x70,
	0xa8, 0x1b, 0xa4, 0x44, 0x50, 0x05, 0x21, 0x66, 0x2b, 0x31, 0x38, 0x71, 0x46, 0xb1, 0x83, 0x25,
	0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x81, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xea, 0x89,
	0x4f, 0x1b, 0x01, 0x00, 0x00,
}