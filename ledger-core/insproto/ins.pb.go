// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: insproto/ins.proto

package insproto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *FieldMap) Reset()         { *m = FieldMap{} }
func (m *FieldMap) String() string { return proto.CompactTextString(m) }
func (*FieldMap) ProtoMessage()    {}
func (*FieldMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e08ba0a4ad7c7f, []int{0}
}
func (m *FieldMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldMap.Unmarshal(m, b)
}
func (m *FieldMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldMap.Marshal(b, m, deterministic)
}
func (m *FieldMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMap.Merge(m, src)
}
func (m *FieldMap) XXX_Size() int {
	return xxx_messageInfo_FieldMap.Size(m)
}
func (m *FieldMap) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMap.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMap proto.InternalMessageInfo

var E_NotationAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63901,
	Name:          "insproto.notation_all",
	Tag:           "varint,63901,opt,name=notation_all",
	Filename:      "insproto/ins.proto",
}

var E_ZeroableAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63902,
	Name:          "insproto.zeroable_all",
	Tag:           "varint,63902,opt,name=zeroable_all",
	Filename:      "insproto/ins.proto",
}

var E_MessageCtxApplyAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         63903,
	Name:          "insproto.message_ctx_apply_all",
	Tag:           "bytes,63903,opt,name=message_ctx_apply_all",
	Filename:      "insproto/ins.proto",
}

var E_ProjectionNames = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         63905,
	Name:          "insproto.projection_names",
	Tag:           "bytes,63905,opt,name=projection_names",
	Filename:      "insproto/ins.proto",
}

var E_ContextMethodAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         63906,
	Name:          "insproto.context_method_all",
	Tag:           "bytes,63906,opt,name=context_method_all",
	Filename:      "insproto/ins.proto",
}

var E_FieldsMappingAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63907,
	Name:          "insproto.fields_mapping_all",
	Tag:           "varint,63907,opt,name=fields_mapping_all",
	Filename:      "insproto/ins.proto",
}

var E_MessageMappingAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63908,
	Name:          "insproto.message_mapping_all",
	Tag:           "varint,63908,opt,name=message_mapping_all",
	Filename:      "insproto/ins.proto",
}

var E_RegisterAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         63909,
	Name:          "insproto.register_all",
	Tag:           "bytes,63909,opt,name=register_all",
	Filename:      "insproto/ins.proto",
}

var E_ContextAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         63910,
	Name:          "insproto.context_all",
	Tag:           "bytes,63910,opt,name=context_all",
	Filename:      "insproto/ins.proto",
}

var E_NullableAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63912,
	Name:          "insproto.nullable_all",
	Tag:           "varint,63912,opt,name=nullable_all",
	Filename:      "insproto/ins.proto",
}

var E_Notation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64901,
	Name:          "insproto.notation",
	Tag:           "varint,64901,opt,name=notation",
	Filename:      "insproto/ins.proto",
}

var E_FieldsZeroable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64902,
	Name:          "insproto.fields_zeroable",
	Tag:           "varint,64902,opt,name=fields_zeroable",
	Filename:      "insproto/ins.proto",
}

var E_Id = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         64904,
	Name:          "insproto.id",
	Tag:           "varint,64904,opt,name=id",
	Filename:      "insproto/ins.proto",
}

var E_MessageCtxApply = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         64903,
	Name:          "insproto.message_ctx_apply",
	Tag:           "bytes,64903,opt,name=message_ctx_apply",
	Filename:      "insproto/ins.proto",
}

var E_Projection = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64905,
	Name:          "insproto.projection",
	Tag:           "varint,64905,opt,name=projection",
	Filename:      "insproto/ins.proto",
}

var E_ContextMethod = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         64906,
	Name:          "insproto.context_method",
	Tag:           "bytes,64906,opt,name=context_method",
	Filename:      "insproto/ins.proto",
}

var E_FieldsMapping = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64907,
	Name:          "insproto.fields_mapping",
	Tag:           "varint,64907,opt,name=fields_mapping",
	Filename:      "insproto/ins.proto",
}

var E_MessageMapping = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64908,
	Name:          "insproto.message_mapping",
	Tag:           "varint,64908,opt,name=message_mapping",
	Filename:      "insproto/ins.proto",
}

var E_Register = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         64909,
	Name:          "insproto.register",
	Tag:           "bytes,64909,opt,name=register",
	Filename:      "insproto/ins.proto",
}

var E_Context = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         64910,
	Name:          "insproto.context",
	Tag:           "bytes,64910,opt,name=context",
	Filename:      "insproto/ins.proto",
}

var E_ProjectionId = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         64911,
	Name:          "insproto.projection_id",
	Tag:           "varint,64911,opt,name=projection_id",
	Filename:      "insproto/ins.proto",
}

var E_Zeroable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65902,
	Name:          "insproto.zeroable",
	Tag:           "varint,65902,opt,name=zeroable",
	Filename:      "insproto/ins.proto",
}

var E_CtxApply = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65903,
	Name:          "insproto.ctx_apply",
	Tag:           "bytes,65903,opt,name=ctx_apply",
	Filename:      "insproto/ins.proto",
}

var E_RawBytes = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65904,
	Name:          "insproto.raw_bytes",
	Tag:           "varint,65904,opt,name=raw_bytes",
	Filename:      "insproto/ins.proto",
}

var E_Mapping = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65907,
	Name:          "insproto.mapping",
	Tag:           "varint,65907,opt,name=mapping",
	Filename:      "insproto/ins.proto",
}

func init() {
	proto.RegisterType((*FieldMap)(nil), "insproto.FieldMap")
	proto.RegisterExtension(E_NotationAll)
	proto.RegisterExtension(E_ZeroableAll)
	proto.RegisterExtension(E_MessageCtxApplyAll)
	proto.RegisterExtension(E_ProjectionNames)
	proto.RegisterExtension(E_ContextMethodAll)
	proto.RegisterExtension(E_FieldsMappingAll)
	proto.RegisterExtension(E_MessageMappingAll)
	proto.RegisterExtension(E_RegisterAll)
	proto.RegisterExtension(E_ContextAll)
	proto.RegisterExtension(E_NullableAll)
	proto.RegisterExtension(E_Notation)
	proto.RegisterExtension(E_FieldsZeroable)
	proto.RegisterExtension(E_Id)
	proto.RegisterExtension(E_MessageCtxApply)
	proto.RegisterExtension(E_Projection)
	proto.RegisterExtension(E_ContextMethod)
	proto.RegisterExtension(E_FieldsMapping)
	proto.RegisterExtension(E_MessageMapping)
	proto.RegisterExtension(E_Register)
	proto.RegisterExtension(E_Context)
	proto.RegisterExtension(E_ProjectionId)
	proto.RegisterExtension(E_Zeroable)
	proto.RegisterExtension(E_CtxApply)
	proto.RegisterExtension(E_RawBytes)
	proto.RegisterExtension(E_Mapping)
}

func init() { proto.RegisterFile("insproto/ins.proto", fileDescriptor_72e08ba0a4ad7c7f) }

var fileDescriptor_72e08ba0a4ad7c7f = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x4f, 0xd4, 0x40,
	0x14, 0xc7, 0xbb, 0x48, 0xc2, 0xee, 0xe3, 0x77, 0xd5, 0x84, 0x10, 0x6d, 0x89, 0x27, 0x2f, 0xee,
	0xc6, 0xa3, 0x03, 0xc6, 0x2c, 0x26, 0x44, 0x0c, 0x8b, 0x71, 0x8f, 0x5e, 0x9a, 0x6e, 0x3b, 0x94,
	0x9a, 0xd9, 0x4e, 0xd3, 0x0e, 0x01, 0x3c, 0x9a, 0xa8, 0xf8, 0x9b, 0x7f, 0xc0, 0x1f, 0xf5, 0x57,
	0xfc, 0x53, 0x3c, 0x72, 0xf4, 0x48, 0xe0, 0xae, 0x24, 0x1c, 0xb9, 0x98, 0x69, 0xdf, 0xec, 0xb2,
	0x60, 0x32, 0xbd, 0xb5, 0x9b, 0xf9, 0x7c, 0xf6, 0xbd, 0x79, 0xdf, 0x57, 0x30, 0xc3, 0x28, 0x8d,
	0x13, 0x2e, 0x78, 0x23, 0x8c, 0xd2, 0x7a, 0xfe, 0x64, 0x56, 0xd5, 0x6f, 0xb3, 0x73, 0x01, 0xe7,
	0x01, 0xa3, 0x8d, 0xfc, 0xad, 0xb3, 0xb1, 0xd6, 0xf0, 0x69, 0xea, 0x25, 0x61, 0x2c, 0x78, 0x52,
	0x9c, 0x9d, 0xbd, 0x11, 0x84, 0x62, 0x7d, 0xa3, 0x53, 0xf7, 0x78, 0xb7, 0x11, 0xf0, 0x80, 0xf7,
	0x8f, 0xca, 0xb7, 0xc2, 0x2c, 0x9f, 0x8a, 0xe3, 0xd7, 0xe6, 0xa0, 0xba, 0x14, 0x52, 0xe6, 0xb7,
	0xdc, 0x98, 0x5c, 0xda, 0xc9, 0x6c, 0x63, 0x37, 0xb3, 0x8d, 0xcf, 0x99, 0x6d, 0xec, 0x67, 0x76,
	0xe5, 0x28, 0xb3, 0x0d, 0xd2, 0x84, 0xb1, 0x88, 0x0b, 0x57, 0x84, 0x3c, 0x72, 0x5c, 0xc6, 0xcc,
	0x2b, 0xf5, 0xa2, 0x86, 0xba, 0x12, 0xd7, 0x97, 0x42, 0x46, 0x1f, 0xc4, 0xf2, 0x40, 0x3a, 0xf3,
	0xe1, 0xf8, 0xc2, 0x5c, 0xe5, 0x7a, 0xb5, 0x3d, 0xaa, 0x98, 0x26, 0x63, 0x52, 0xf1, 0x84, 0x26,
	0xdc, 0xed, 0x30, 0x5a, 0x42, 0xf1, 0x51, 0x29, 0x14, 0x23, 0x15, 0x0f, 0xe1, 0x72, 0x97, 0xa6,
	0xa9, 0x1b, 0x50, 0xc7, 0x13, 0x5b, 0x8e, 0x1b, 0xc7, 0x6c, 0xbb, 0x84, 0xeb, 0x53, 0xee, 0xaa,
	0xb5, 0x4d, 0x84, 0xef, 0x8a, 0xad, 0xa6, 0x44, 0xa5, 0x72, 0x19, 0xa6, 0xe2, 0x84, 0x3f, 0xa6,
	0x5e, 0xde, 0x5a, 0xe4, 0x76, 0x69, 0xaa, 0xb1, 0x65, 0x68, 0x9b, 0xec, 0x73, 0xab, 0x12, 0x23,
	0x2b, 0x60, 0x7a, 0x3c, 0x12, 0x74, 0x4b, 0x38, 0x5d, 0x2a, 0xd6, 0xb9, 0x5f, 0xa2, 0xb4, 0x2f,
	0x28, 0x9b, 0x42, 0xb2, 0x95, 0x83, 0xb2, 0xb0, 0x15, 0x30, 0xd7, 0xe4, 0x4c, 0x52, 0xa7, 0xeb,
	0xc6, 0x71, 0x18, 0x05, 0x25, 0x6c, 0x5f, 0xf1, 0xd2, 0xa6, 0x0a, 0xb2, 0x55, 0x80, 0xd2, 0xb6,
	0x0a, 0x17, 0xd5, 0xcd, 0x95, 0xd7, 0x7d, 0x43, 0xdd, 0x34, 0xa2, 0xa7, 0x7c, 0x4d, 0x18, 0x4b,
	0x68, 0x10, 0xa6, 0x82, 0x26, 0x25, 0x44, 0xdf, 0xb1, 0xcb, 0x51, 0xc5, 0x48, 0xc5, 0x1d, 0x18,
	0x55, 0xd7, 0xa5, 0x37, 0xfc, 0x40, 0x03, 0x20, 0x82, 0x35, 0x44, 0x1b, 0x8c, 0x95, 0x0c, 0xd4,
	0xcf, 0x5e, 0x26, 0x91, 0x91, 0x8a, 0xdb, 0x50, 0x55, 0x11, 0x35, 0xed, 0x73, 0x78, 0xab, 0x68,
	0x5b, 0x19, 0x9e, 0x9d, 0x14, 0x86, 0x1e, 0x42, 0xee, 0xc3, 0x24, 0xce, 0x48, 0xa5, 0x54, 0x6f,
	0x79, 0x8e, 0x96, 0x89, 0x82, 0x7c, 0x84, 0x20, 0xb9, 0x09, 0x43, 0xa1, 0xaf, 0xc7, 0x77, 0x72,
	0x7c, 0xb8, 0x3d, 0x14, 0xfa, 0xa4, 0x05, 0xd3, 0xe7, 0xd6, 0x41, 0x6f, 0x78, 0x71, 0x82, 0xf9,
	0x3d, 0xb3, 0x0d, 0xa4, 0x09, 0xd0, 0x8f, 0xb4, 0xde, 0xf3, 0x12, 0x1b, 0x39, 0x05, 0x91, 0x7b,
	0x30, 0x31, 0xb8, 0x02, 0x7a, 0xcd, 0x2b, 0x2c, 0x67, 0x7c, 0x60, 0x03, 0xa4, 0x69, 0x30, 0xfe,
	0x7a, 0xd3, 0x6b, 0x2c, 0x68, 0x7c, 0x20, 0xfd, 0x72, 0x48, 0x67, 0xa2, 0xaf, 0x57, 0xbd, 0x51,
	0x43, 0x1a, 0x4c, 0xbe, 0xcc, 0x8b, 0x8a, 0xb0, 0x5e, 0xf2, 0x16, 0x3b, 0xeb, 0x21, 0x64, 0x1e,
	0x46, 0xb0, 0x4b, 0x3d, 0xfd, 0x0e, 0x69, 0x45, 0x90, 0x25, 0x18, 0x3f, 0xf5, 0xa5, 0x2a, 0x93,
	0x95, 0xf7, 0x98, 0x95, 0xb1, 0x3e, 0xb7, 0xec, 0x93, 0x79, 0xa8, 0xf6, 0xd2, 0x7a, 0xf5, 0x3f,
	0x2b, 0x43, 0x99, 0xaf, 0x04, 0x7f, 0x9e, 0x0e, 0x17, 0x89, 0x57, 0x00, 0x59, 0x80, 0x5a, 0x3f,
	0x6a, 0x1a, 0xfa, 0x6f, 0x4e, 0xd7, 0xda, 0x55, 0x4f, 0x25, 0x6c, 0x01, 0x6a, 0x89, 0xbb, 0xe9,
	0x74, 0xb6, 0x05, 0x4d, 0x75, 0xf4, 0x91, 0xfa, 0xef, 0xc4, 0xdd, 0x5c, 0x94, 0x00, 0xb9, 0x05,
	0x23, 0x6a, 0x80, 0x1a, 0xf6, 0x18, 0x59, 0x75, 0x7e, 0x71, 0xe6, 0xd7, 0x81, 0x55, 0xd9, 0x3b,
	0xb0, 0x2a, 0xfb, 0x07, 0x56, 0x65, 0xf7, 0xd0, 0x32, 0xf6, 0x0e, 0x2d, 0xe3, 0xf7, 0xa1, 0x65,
	0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x27, 0x19, 0x7b, 0x6a, 0x07, 0x00, 0x00,
}

func (m *FieldMap) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovIns(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIns(x uint64) (n int) {
	return sovIns(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
