// Code generated by protoc-gen-go.
// source: simple.proto
// DO NOT EDIT!

/*
Package simplepb is a generated protocol buffer package.

It is generated from these files:
	simple.proto

It has these top-level messages:
	SimpleMessage
*/
package simplepb

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

type SimpleMessage struct {
	Id         int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	IsSimple   bool    `protobuf:"varint,2,opt,name=is_simple,json=isSimple" json:"is_simple,omitempty"`
	Name       string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	SampleList []int32 `protobuf:"varint,4,rep,packed,name=sample_list,json=sampleList" json:"sample_list,omitempty"`
}

func (m *SimpleMessage) Reset()                    { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string            { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()               {}
func (*SimpleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SimpleMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SimpleMessage) GetIsSimple() bool {
	if m != nil {
		return m.IsSimple
	}
	return false
}

func (m *SimpleMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleMessage) GetSampleList() []int32 {
	if m != nil {
		return m.SampleList
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "example.simple.SimpleMessage")
}

func init() { proto.RegisterFile("simple.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0xad, 0x48, 0x04, 0x73, 0x21,
	0xa2, 0x4a, 0x85, 0x5c, 0xbc, 0xc1, 0x60, 0x96, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x10,
	0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x90, 0x25, 0x24, 0xcd,
	0xc5, 0x99, 0x59, 0x1c, 0x0f, 0x51, 0x2d, 0xc1, 0x04, 0x14, 0xe6, 0x08, 0xe2, 0xc8, 0x2c, 0x86,
	0xe8, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x06, 0x8a, 0x73, 0x06, 0x81,
	0xd9, 0x42, 0xf2, 0x5c, 0xdc, 0xc5, 0x60, 0x2b, 0xe2, 0x73, 0x32, 0x8b, 0x4b, 0x24, 0x58, 0x14,
	0x98, 0x81, 0x26, 0x71, 0x41, 0x84, 0x7c, 0x80, 0x22, 0x4e, 0x5c, 0x51, 0x1c, 0x10, 0xe3, 0x0a,
	0x92, 0x92, 0xd8, 0xc0, 0xae, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xa7, 0x7d, 0xbc,
	0xa5, 0x00, 0x00, 0x00,
}
