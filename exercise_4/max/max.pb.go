// Code generated by protoc-gen-go.
// source: max.proto
// DO NOT EDIT!

/*
Package maxpb is a generated protocol buffer package.

It is generated from these files:
	max.proto

It has these top-level messages:
	Max
	MaxRequest
	MaxResponse
*/
package maxpb

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

type Max struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *Max) Reset()                    { *m = Max{} }
func (m *Max) String() string            { return proto.CompactTextString(m) }
func (*Max) ProtoMessage()               {}
func (*Max) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Max) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type MaxRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *MaxRequest) Reset()                    { *m = MaxRequest{} }
func (m *MaxRequest) String() string            { return proto.CompactTextString(m) }
func (*MaxRequest) ProtoMessage()               {}
func (*MaxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MaxRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type MaxResponse struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *MaxResponse) Reset()                    { *m = MaxResponse{} }
func (m *MaxResponse) String() string            { return proto.CompactTextString(m) }
func (*MaxResponse) ProtoMessage()               {}
func (*MaxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MaxResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Max)(nil), "max.Max")
	proto.RegisterType((*MaxRequest)(nil), "max.MaxRequest")
	proto.RegisterType((*MaxResponse)(nil), "max.MaxResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MaxService service

type MaxServiceClient interface {
	// BiDi Streaming
	FindMax(ctx context.Context, opts ...grpc.CallOption) (MaxService_FindMaxClient, error)
}

type maxServiceClient struct {
	cc *grpc.ClientConn
}

func NewMaxServiceClient(cc *grpc.ClientConn) MaxServiceClient {
	return &maxServiceClient{cc}
}

func (c *maxServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (MaxService_FindMaxClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MaxService_serviceDesc.Streams[0], c.cc, "/max.MaxService/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &maxServiceFindMaxClient{stream}
	return x, nil
}

type MaxService_FindMaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type maxServiceFindMaxClient struct {
	grpc.ClientStream
}

func (x *maxServiceFindMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maxServiceFindMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MaxService service

type MaxServiceServer interface {
	// BiDi Streaming
	FindMax(MaxService_FindMaxServer) error
}

func RegisterMaxServiceServer(s *grpc.Server, srv MaxServiceServer) {
	s.RegisterService(&_MaxService_serviceDesc, srv)
}

func _MaxService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaxServiceServer).FindMax(&maxServiceFindMaxServer{stream})
}

type MaxService_FindMaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type maxServiceFindMaxServer struct {
	grpc.ServerStream
}

func (x *maxServiceFindMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maxServiceFindMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MaxService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "max.MaxService",
	HandlerType: (*MaxServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMax",
			Handler:       _MaxService_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "max.proto",
}

func init() { proto.RegisterFile("max.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4d, 0xac, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x32, 0x95, 0x64, 0xb9, 0x98, 0x7d, 0x13, 0x2b,
	0x84, 0xc4, 0xb8, 0xd8, 0xf2, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x83, 0xa0, 0x3c, 0x25, 0x15, 0x2e, 0x2e, 0xa0, 0x74, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09,
	0x4e, 0x55, 0xaa, 0x5c, 0xdc, 0x60, 0x55, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x65, 0x45,
	0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x30, 0x65, 0x10, 0x9e, 0x91, 0x03, 0xd8, 0xb0, 0xe0, 0xd4, 0xa2,
	0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x23, 0x2e, 0x76, 0xb7, 0xcc, 0xbc, 0x14, 0x90, 0xed, 0xfc, 0x7a,
	0x20, 0x57, 0x21, 0x2c, 0x92, 0x12, 0x40, 0x08, 0x40, 0xcc, 0x54, 0x62, 0xd0, 0x60, 0x34, 0x60,
	0x74, 0x62, 0x8f, 0x62, 0x05, 0x4a, 0x14, 0x24, 0x25, 0xb1, 0x81, 0xbd, 0x60, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x2f, 0x58, 0x45, 0xcf, 0x00, 0x00, 0x00,
}
