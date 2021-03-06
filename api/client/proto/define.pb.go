// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/client/proto/define.proto

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	api/client/proto/define.proto

It has these top-level messages:
*/
package client

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

type SRV int32

const (
	SRV_client_auth     SRV = 0
	SRV_client_register SRV = 1
)

var SRV_name = map[int32]string{
	0: "client_auth",
	1: "client_register",
}
var SRV_value = map[string]int32{
	"client_auth":     0,
	"client_register": 1,
}

func (x SRV) String() string {
	return proto.EnumName(SRV_name, int32(x))
}
func (SRV) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("client.SRV", SRV_name, SRV_value)
}

func init() { proto.RegisterFile("api/client/proto/define.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x49, 0x4d,
	0xcb, 0xcc, 0x4b, 0xd5, 0x03, 0x73, 0x84, 0xd8, 0x20, 0x52, 0x5a, 0xda, 0x5c, 0xcc, 0xc1, 0x41,
	0x61, 0x42, 0xfc, 0x5c, 0xdc, 0x10, 0x81, 0xf8, 0xc4, 0xd2, 0x92, 0x0c, 0x01, 0x06, 0x21, 0x61,
	0x2e, 0x7e, 0xa8, 0x40, 0x51, 0x6a, 0x7a, 0x66, 0x71, 0x49, 0x6a, 0x91, 0x00, 0x63, 0x12, 0x1b,
	0x58, 0xaf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x35, 0x4b, 0x09, 0x5c, 0x00, 0x00, 0x00,
}
