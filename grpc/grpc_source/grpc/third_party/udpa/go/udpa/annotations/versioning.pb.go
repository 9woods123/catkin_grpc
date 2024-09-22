// Code generated by protoc-gen-go. DO NOT EDIT.
// source: udpa/annotations/versioning.proto

package udpa_annotations

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type VersioningAnnotation struct {
	PreviousMessageType  string   `protobuf:"bytes,1,opt,name=previous_message_type,json=previousMessageType,proto3" json:"previous_message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersioningAnnotation) Reset()         { *m = VersioningAnnotation{} }
func (m *VersioningAnnotation) String() string { return proto.CompactTextString(m) }
func (*VersioningAnnotation) ProtoMessage()    {}
func (*VersioningAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bc0544382e16cfc, []int{0}
}

func (m *VersioningAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersioningAnnotation.Unmarshal(m, b)
}
func (m *VersioningAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersioningAnnotation.Marshal(b, m, deterministic)
}
func (m *VersioningAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersioningAnnotation.Merge(m, src)
}
func (m *VersioningAnnotation) XXX_Size() int {
	return xxx_messageInfo_VersioningAnnotation.Size(m)
}
func (m *VersioningAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_VersioningAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_VersioningAnnotation proto.InternalMessageInfo

func (m *VersioningAnnotation) GetPreviousMessageType() string {
	if m != nil {
		return m.PreviousMessageType
	}
	return ""
}

var E_Versioning = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*VersioningAnnotation)(nil),
	Field:         7881811,
	Name:          "udpa.annotations.versioning",
	Tag:           "bytes,7881811,opt,name=versioning",
	Filename:      "udpa/annotations/versioning.proto",
}

func init() {
	proto.RegisterType((*VersioningAnnotation)(nil), "udpa.annotations.VersioningAnnotation")
	proto.RegisterExtension(E_Versioning)
}

func init() { proto.RegisterFile("udpa/annotations/versioning.proto", fileDescriptor_5bc0544382e16cfc) }

var fileDescriptor_5bc0544382e16cfc = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4d, 0x29, 0x48,
	0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x4b, 0x2d,
	0x2a, 0xce, 0xcc, 0xcf, 0xcb, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x00,
	0x29, 0xd1, 0x43, 0x52, 0x22, 0xa5, 0x90, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x0f, 0x96, 0x4f,
	0x2a, 0x4d, 0xd3, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x82, 0xe8, 0x51,
	0xf2, 0xe2, 0x12, 0x09, 0x83, 0x9b, 0xe3, 0x08, 0xd7, 0x2a, 0x64, 0xc4, 0x25, 0x5a, 0x50, 0x94,
	0x5a, 0x96, 0x99, 0x5f, 0x5a, 0x1c, 0x9f, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x1a, 0x5f, 0x52,
	0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x24, 0x0c, 0x93, 0xf4, 0x85, 0xc8, 0x85,
	0x54, 0x16, 0xa4, 0x5a, 0x65, 0x71, 0x71, 0x21, 0xdc, 0x24, 0x24, 0xaf, 0x07, 0xb1, 0x5c, 0x0f,
	0x66, 0xb9, 0x1e, 0x54, 0xad, 0x7f, 0x01, 0xd8, 0x71, 0x12, 0x97, 0x3b, 0x1e, 0x32, 0x2b, 0x30,
	0x6a, 0x70, 0x1b, 0xa9, 0xe9, 0xa1, 0x3b, 0x5c, 0x0f, 0x9b, 0x9b, 0x82, 0x90, 0x4c, 0x4f, 0x62,
	0x03, 0x9b, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x9c, 0xb8, 0x85, 0x17, 0x01, 0x00,
	0x00,
}
