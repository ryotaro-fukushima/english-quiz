// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/imports/test_a_1/m1.proto

package test_a_1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type E1 int32

const (
	E1_E1_ZERO E1 = 0
)

// Enum value maps for E1.
var (
	E1_name = map[int32]string{
		0: "E1_ZERO",
	}
	E1_value = map[string]int32{
		"E1_ZERO": 0,
	}
)

func (x E1) Enum() *E1 {
	p := new(E1)
	*p = x
	return p
}

func (x E1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (E1) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_enumTypes[0].Descriptor()
}

func (E1) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_enumTypes[0]
}

func (x E1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use E1.Descriptor instead.
func (E1) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescGZIP(), []int{0}
}

type M1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *M1) Reset() {
	*x = M1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M1) ProtoMessage() {}

func (x *M1) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M1.ProtoReflect.Descriptor instead.
func (*M1) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescGZIP(), []int{0}
}

type M1_1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M1 *M1 `protobuf:"bytes,1,opt,name=m1,proto3" json:"m1,omitempty"`
}

func (x *M1_1) Reset() {
	*x = M1_1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M1_1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M1_1) ProtoMessage() {}

func (x *M1_1) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M1_1.ProtoReflect.Descriptor instead.
func (*M1_1) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescGZIP(), []int{1}
}

func (x *M1_1) GetM1() *M1 {
	if x != nil {
		return x.M1
	}
	return nil
}

var File_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x2f, 0x6d, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x22, 0x04,
	0x0a, 0x02, 0x4d, 0x31, 0x22, 0x22, 0x0a, 0x04, 0x4d, 0x31, 0x5f, 0x31, 0x12, 0x1a, 0x0a, 0x02,
	0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x61, 0x2e, 0x4d, 0x31, 0x52, 0x02, 0x6d, 0x31, 0x2a, 0x11, 0x0a, 0x02, 0x45, 0x31, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x31, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x61, 0x5f, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescData = file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_goTypes = []interface{}{
	(E1)(0),      // 0: test.a.E1
	(*M1)(nil),   // 1: test.a.M1
	(*M1_1)(nil), // 2: test.a.M1_1
}
var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_depIdxs = []int32{
	1, // 0: test.a.M1_1.m1:type_name -> test.a.M1
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_init() }
func file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_init() {
	if File_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M1_1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto = out.File
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m1_proto_depIdxs = nil
}
