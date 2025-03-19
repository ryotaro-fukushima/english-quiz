// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/testeditions/testeditions_opaque/test_extension2.opaque.proto

package testeditions_opaque

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	unsafe "unsafe"
)

type OtherRepeatedFieldEncoding struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OtherRepeatedFieldEncoding) Reset() {
	*x = OtherRepeatedFieldEncoding{}
	mi := &file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OtherRepeatedFieldEncoding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherRepeatedFieldEncoding) ProtoMessage() {}

func (x *OtherRepeatedFieldEncoding) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type OtherRepeatedFieldEncoding_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 OtherRepeatedFieldEncoding_builder) Build() *OtherRepeatedFieldEncoding {
	m0 := &OtherRepeatedFieldEncoding{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*TestFeatureResolution)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         6,
		Name:          "opaque.goproto.proto.testeditions.other_file_global_expanded_extension_overriden",
		Tag:           "varint,6,rep,name=other_file_global_expanded_extension_overriden",
		Filename:      "internal/testprotos/testeditions/testeditions_opaque/test_extension2.opaque.proto",
	},
	{
		ExtendedType:  (*TestFeatureResolution)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         7,
		Name:          "opaque.goproto.proto.testeditions.other_file_global_packed_extension",
		Tag:           "varint,7,rep,packed,name=other_file_global_packed_extension",
		Filename:      "internal/testprotos/testeditions/testeditions_opaque/test_extension2.opaque.proto",
	},
	{
		ExtendedType:  (*TestFeatureResolution)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         8,
		Name:          "opaque.goproto.proto.testeditions.OtherRepeatedFieldEncoding.other_file_message_expanded_extension_overriden",
		Tag:           "varint,8,rep,name=other_file_message_expanded_extension_overriden",
		Filename:      "internal/testprotos/testeditions/testeditions_opaque/test_extension2.opaque.proto",
	},
	{
		ExtendedType:  (*TestFeatureResolution)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         9,
		Name:          "opaque.goproto.proto.testeditions.OtherRepeatedFieldEncoding.other_file_message_packed_extension",
		Tag:           "varint,9,rep,packed,name=other_file_message_packed_extension",
		Filename:      "internal/testprotos/testeditions/testeditions_opaque/test_extension2.opaque.proto",
	},
}

// Extension fields to TestFeatureResolution.
var (
	// repeated int32 other_file_global_expanded_extension_overriden = 6;
	E_OtherFileGlobalExpandedExtensionOverriden = &file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_extTypes[0]
	// repeated int32 other_file_global_packed_extension = 7;
	E_OtherFileGlobalPackedExtension = &file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_extTypes[1]
	// repeated int32 other_file_message_expanded_extension_overriden = 8;
	E_OtherRepeatedFieldEncoding_OtherFileMessageExpandedExtensionOverriden = &file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_extTypes[2]
	// repeated int32 other_file_message_packed_extension = 9;
	E_OtherRepeatedFieldEncoding_OtherFileMessagePackedExtension = &file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_extTypes[3]
)

var File_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto protoreflect.FileDescriptor

var file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_rawDesc = string([]byte{
	0x0a, 0x51, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x1a,
	0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x32, 0xa4, 0x01, 0x0a, 0x2f, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x6e, 0x12, 0x38,
	0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x42, 0x05,
	0xaa, 0x01, 0x02, 0x18, 0x02, 0x52, 0x2a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x6e, 0x32, 0x86, 0x01, 0x0a, 0x23, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x1f, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0xa2, 0x01, 0x0a, 0x2e, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x6e, 0x12, 0x38, 0x2e,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x42, 0x05, 0xaa,
	0x01, 0x02, 0x18, 0x02, 0x52, 0x29, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x6e, 0x3a,
	0x84, 0x01, 0x0a, 0x22, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x1e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x5b, 0x5a, 0x4f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x92, 0x03, 0x07, 0xd2, 0x3e, 0x02, 0x10,
	0x03, 0x18, 0x01, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_goTypes = []any{
	(*OtherRepeatedFieldEncoding)(nil), // 0: opaque.goproto.proto.testeditions.OtherRepeatedFieldEncoding
	(*TestFeatureResolution)(nil),      // 1: opaque.goproto.proto.testeditions.TestFeatureResolution
}
var file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_depIdxs = []int32{
	1, // 0: opaque.goproto.proto.testeditions.other_file_global_expanded_extension_overriden:extendee -> opaque.goproto.proto.testeditions.TestFeatureResolution
	1, // 1: opaque.goproto.proto.testeditions.other_file_global_packed_extension:extendee -> opaque.goproto.proto.testeditions.TestFeatureResolution
	1, // 2: opaque.goproto.proto.testeditions.OtherRepeatedFieldEncoding.other_file_message_expanded_extension_overriden:extendee -> opaque.goproto.proto.testeditions.TestFeatureResolution
	1, // 3: opaque.goproto.proto.testeditions.OtherRepeatedFieldEncoding.other_file_message_packed_extension:extendee -> opaque.goproto.proto.testeditions.TestFeatureResolution
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_init()
}
func file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_init() {
	if File_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto != nil {
		return
	}
	file_internal_testprotos_testeditions_testeditions_opaque_test_extension_opaque_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_rawDesc), len(file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_extTypes,
	}.Build()
	File_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto = out.File
	file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_goTypes = nil
	file_internal_testprotos_testeditions_testeditions_opaque_test_extension2_opaque_proto_depIdxs = nil
}
