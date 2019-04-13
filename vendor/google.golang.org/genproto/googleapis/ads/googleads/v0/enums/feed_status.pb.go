// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_status.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Possible statuses of a feed.
type FeedStatusEnum_FeedStatus int32

const (
	// Not specified.
	FeedStatusEnum_UNSPECIFIED FeedStatusEnum_FeedStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedStatusEnum_UNKNOWN FeedStatusEnum_FeedStatus = 1
	// Feed is enabled.
	FeedStatusEnum_ENABLED FeedStatusEnum_FeedStatus = 2
	// Feed has been removed.
	FeedStatusEnum_REMOVED FeedStatusEnum_FeedStatus = 3
)

var FeedStatusEnum_FeedStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}
var FeedStatusEnum_FeedStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedStatusEnum_FeedStatus) String() string {
	return proto.EnumName(FeedStatusEnum_FeedStatus_name, int32(x))
}
func (FeedStatusEnum_FeedStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_status_4a6d02cc16d1f2a8, []int{0, 0}
}

// Container for enum describing possible statuses of a feed.
type FeedStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedStatusEnum) Reset()         { *m = FeedStatusEnum{} }
func (m *FeedStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedStatusEnum) ProtoMessage()    {}
func (*FeedStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_status_4a6d02cc16d1f2a8, []int{0}
}
func (m *FeedStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedStatusEnum.Unmarshal(m, b)
}
func (m *FeedStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedStatusEnum.Marshal(b, m, deterministic)
}
func (dst *FeedStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedStatusEnum.Merge(dst, src)
}
func (m *FeedStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedStatusEnum.Size(m)
}
func (m *FeedStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedStatusEnum)(nil), "google.ads.googleads.v0.enums.FeedStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedStatusEnum_FeedStatus", FeedStatusEnum_FeedStatus_name, FeedStatusEnum_FeedStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_status.proto", fileDescriptor_feed_status_4a6d02cc16d1f2a8)
}

var fileDescriptor_feed_status_4a6d02cc16d1f2a8 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x4f, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0xb5, 0x1d, 0x28, 0x64, 0xe0, 0x4a, 0xdf, 0xf7, 0xb0, 0x7d, 0x40, 0x1a, 0xf0, 0x2d, 0x3e,
	0xa5, 0x36, 0x1b, 0x43, 0xcd, 0x8a, 0x63, 0x11, 0xa4, 0x20, 0xd5, 0xc4, 0x20, 0xac, 0xcd, 0xd8,
	0x6d, 0xf7, 0x41, 0x3e, 0xfa, 0x29, 0x7e, 0x87, 0x4f, 0x7e, 0x85, 0x24, 0xdd, 0xba, 0x27, 0x7d,
	0x09, 0xe7, 0xde, 0x73, 0x4e, 0xee, 0x39, 0x28, 0x31, 0xd6, 0x9a, 0x8d, 0x4e, 0x4a, 0x05, 0x07,
	0xe8, 0xd0, 0x9e, 0x24, 0xba, 0x6e, 0x2b, 0x48, 0xde, 0xb4, 0x56, 0xcf, 0xd0, 0x94, 0x4d, 0x0b,
	0x78, 0xbb, 0xb3, 0x8d, 0x8d, 0xc7, 0x9d, 0x0a, 0x97, 0x0a, 0x70, 0x6f, 0xc0, 0x7b, 0x82, 0xbd,
	0x61, 0x2a, 0xd1, 0xe5, 0x4c, 0x6b, 0xb5, 0xf2, 0x16, 0x5e, 0xb7, 0xd5, 0x34, 0x43, 0xe8, 0xb4,
	0x89, 0x47, 0x68, 0xb8, 0x16, 0xab, 0x9c, 0xdf, 0x2c, 0x66, 0x0b, 0x9e, 0x45, 0x67, 0xf1, 0x10,
	0x5d, 0xac, 0xc5, 0xad, 0x58, 0x3e, 0x8a, 0x28, 0x70, 0x03, 0x17, 0x2c, 0xbd, 0xe3, 0x59, 0x14,
	0xba, 0xe1, 0x81, 0xdf, 0x2f, 0x25, 0xcf, 0xa2, 0x41, 0xfa, 0x1d, 0xa0, 0xc9, 0xab, 0xad, 0xf0,
	0xbf, 0xd7, 0xd3, 0xd1, 0xe9, 0x52, 0xee, 0xd2, 0xe6, 0xc1, 0x53, 0x7a, 0x70, 0x18, 0xbb, 0x29,
	0x6b, 0x83, 0xed, 0xce, 0x24, 0x46, 0xd7, 0xbe, 0xcb, 0xb1, 0xf0, 0xf6, 0x1d, 0xfe, 0xe8, 0x7f,
	0xed, 0xdf, 0x8f, 0x70, 0x30, 0x67, 0xec, 0x33, 0x1c, 0xcf, 0xbb, 0xaf, 0x98, 0x02, 0xdc, 0x41,
	0x87, 0x24, 0xc1, 0xae, 0x27, 0x7c, 0x1d, 0xf9, 0x82, 0x29, 0x28, 0x7a, 0xbe, 0x90, 0xa4, 0xf0,
	0xfc, 0x4f, 0x38, 0xe9, 0x96, 0x94, 0x32, 0x05, 0x94, 0xf6, 0x0a, 0x4a, 0x25, 0xa1, 0xd4, 0x6b,
	0x5e, 0xce, 0x7d, 0xb0, 0xab, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x32, 0x93, 0xf5, 0x97,
	0x01, 0x00, 0x00,
}