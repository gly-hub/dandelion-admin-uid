// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server.proto

package base

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdd, 0x4a, 0x23, 0x31,
	0x18, 0x86, 0xb7, 0x07, 0xfb, 0x43, 0x76, 0xb7, 0xcb, 0xe6, 0x64, 0xb7, 0xa3, 0x8c, 0xd5, 0x0b,
	0x18, 0x41, 0x0f, 0xc5, 0x83, 0xb6, 0x62, 0x11, 0x2a, 0x94, 0xfe, 0x20, 0x78, 0x52, 0xa6, 0xfa,
	0x59, 0x07, 0xda, 0x49, 0xcc, 0x37, 0x23, 0xd4, 0x4b, 0xf0, 0x0a, 0xbc, 0x24, 0x0f, 0xbd, 0x04,
	0xa9, 0x37, 0x22, 0x49, 0xa6, 0x9d, 0xa4, 0x99, 0x69, 0x4f, 0x0a, 0x79, 0x9f, 0x37, 0xcf, 0xf0,
	0xa5, 0x09, 0xf9, 0x85, 0x20, 0x1e, 0x41, 0x04, 0x5c, 0xb0, 0x84, 0x79, 0x7f, 0xc2, 0x34, 0xb9,
	0x67, 0x22, 0x7a, 0x82, 0x2c, 0xa8, 0xe2, 0x1c, 0x47, 0x33, 0x88, 0x53, 0x73, 0x9d, 0xe2, 0x6a,
	0x83, 0x5a, 0x0b, 0x36, 0xcd, 0xfa, 0x47, 0xcf, 0x84, 0xfc, 0x6d, 0x86, 0x08, 0x7d, 0x65, 0x95,
	0xbf, 0xd1, 0x0d, 0xd0, 0x06, 0xf9, 0xda, 0x61, 0x93, 0x28, 0xa6, 0xbb, 0xc1, 0xea, 0x03, 0x23,
	0xb5, 0x21, 0x50, 0x79, 0x37, 0x14, 0xe1, 0x0c, 0x3d, 0xaf, 0x98, 0xf6, 0x00, 0x39, 0xed, 0x92,
	0x9f, 0x6d, 0x48, 0x86, 0x08, 0xe2, 0x12, 0xe2, 0x94, 0xee, 0x39, 0x55, 0x83, 0xf6, 0xe0, 0xc1,
	0xab, 0x6f, 0x2e, 0x58, 0xc6, 0x8b, 0xf8, 0x8e, 0x95, 0x1b, 0x25, 0xdd, 0x68, 0xd4, 0x05, 0xe4,
	0xf4, 0x5c, 0x19, 0xe5, 0x07, 0x06, 0x02, 0x80, 0x7a, 0x81, 0x3c, 0xb8, 0xbc, 0xbb, 0x04, 0x52,
	0xb6, 0x53, 0xca, 0x90, 0xd3, 0x16, 0x21, 0x2d, 0x01, 0x61, 0x02, 0x6a, 0xd4, 0x9a, 0x59, 0xcd,
	0x73, 0x69, 0xf1, 0xca, 0x90, 0x96, 0x0c, 0xf9, 0x6d, 0xa1, 0x24, 0xcf, 0x1d, 0x89, 0x89, 0xb4,
	0xe4, 0x0c, 0xa6, 0x50, 0x24, 0xc9, 0x73, 0x47, 0x62, 0x22, 0xe4, 0xf4, 0x94, 0xfc, 0xe8, 0x33,
	0xa1, 0x46, 0xa4, 0xff, 0xcc, 0xde, 0x32, 0x95, 0x82, 0xff, 0xc5, 0x00, 0x39, 0x1d, 0x90, 0xdf,
	0x7a, 0xb4, 0xfe, 0x1c, 0xe5, 0x71, 0xd3, 0x7a, 0xb0, 0xbc, 0x84, 0xd6, 0xe4, 0x19, 0x96, 0xb2,
	0xfd, 0x2d, 0x0d, 0x6d, 0xd5, 0xb3, 0x96, 0x5a, 0x2d, 0x5c, 0x68, 0x5d, 0x6b, 0x68, 0xab, 0x1e,
	0xbe, 0xd4, 0x6a, 0xe1, 0x42, 0xeb, 0x5a, 0x03, 0x39, 0xbd, 0x22, 0xd5, 0x36, 0x24, 0x59, 0xd2,
	0x89, 0x30, 0xa1, 0xce, 0x26, 0x9b, 0x4b, 0xef, 0xc1, 0xb6, 0x8a, 0x79, 0xd1, 0x7a, 0x6c, 0x0a,
	0xb4, 0x16, 0xc8, 0x87, 0x6c, 0x9d, 0x98, 0xcc, 0xf5, 0xdf, 0x5b, 0x82, 0xcc, 0x8b, 0xe6, 0x4a,
	0xf2, 0xdc, 0x91, 0x98, 0xc8, 0xbc, 0x68, 0xae, 0x24, 0xcf, 0x1d, 0x89, 0x89, 0x56, 0xef, 0x4f,
	0x2e, 0xd5, 0x21, 0x59, 0x55, 0x03, 0xe8, 0xf7, 0x57, 0xc6, 0xb4, 0xa7, 0x81, 0x18, 0x4d, 0x62,
	0x79, 0x58, 0x68, 0x7b, 0x0c, 0xe0, 0x78, 0x2c, 0x86, 0xbc, 0x59, 0x7b, 0x5d, 0xf8, 0x95, 0xb7,
	0x85, 0x5f, 0x79, 0x5f, 0xf8, 0x95, 0x97, 0x0f, 0xff, 0xcb, 0xf5, 0xf7, 0xe0, 0xf0, 0x64, 0x1c,
	0x22, 0x8c, 0xbf, 0xa9, 0x1d, 0xc7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0x86, 0x58, 0x8a,
	0x7f, 0x05, 0x00, 0x00,
}