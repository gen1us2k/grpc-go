// Code generated by protoc-gen-go.
// source: proto2_ext.proto
// DO NOT EDIT!

package grpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Extension struct {
	Baz              *int32 `protobuf:"varint,1,opt,name=baz" json:"baz,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Extension) Reset()                    { *m = Extension{} }
func (m *Extension) String() string            { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()               {}
func (*Extension) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Extension) GetBaz() int32 {
	if m != nil && m.Baz != nil {
		return *m.Baz
	}
	return 0
}

var E_Bar = &proto.ExtensionDesc{
	ExtendedType:  (*ToBeExtened)(nil),
	ExtensionType: (*int32)(nil),
	Field:         13,
	Name:          "grpc.testing.bar",
	Tag:           "varint,13,opt,name=bar",
}

var E_Baz = &proto.ExtensionDesc{
	ExtendedType:  (*ToBeExtened)(nil),
	ExtensionType: (*Extension)(nil),
	Field:         17,
	Name:          "grpc.testing.baz",
	Tag:           "bytes,17,opt,name=baz",
}

func init() {
	proto.RegisterType((*Extension)(nil), "grpc.testing.Extension")
	proto.RegisterExtension(E_Bar)
	proto.RegisterExtension(E_Baz)
}

var fileDescriptor1 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0x8a, 0x4f, 0xad, 0x28, 0xd1, 0x03, 0x33, 0x85, 0x78, 0xd2, 0x8b, 0x0a, 0x92, 0xf5,
	0x4a, 0x52, 0x8b, 0x4b, 0x32, 0xf3, 0xd2, 0xa5, 0x78, 0x20, 0xf2, 0x10, 0x39, 0x25, 0x59, 0x2e,
	0x4e, 0xd7, 0x8a, 0x92, 0xd4, 0xbc, 0xe2, 0xcc, 0xfc, 0x3c, 0x21, 0x01, 0x2e, 0xe6, 0xa4, 0xc4,
	0x2a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x10, 0xd3, 0x4a, 0x1b, 0x24, 0x52, 0x24, 0x24,
	0xa9, 0x87, 0x6c, 0x84, 0x5e, 0x48, 0xbe, 0x53, 0x2a, 0x58, 0x57, 0x6a, 0x8a, 0x04, 0x2f, 0x4c,
	0x71, 0x91, 0x95, 0x0b, 0x58, 0x3b, 0x3e, 0xc5, 0x82, 0x40, 0xc5, 0xdc, 0x46, 0xe2, 0xa8, 0x0a,
	0xe0, 0xf6, 0x83, 0xad, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x59, 0xfa, 0x16, 0xbc, 0xc0, 0x00,
	0x00, 0x00,
}
