// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: MusicGameUnknown1Enum.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MusicGameUnknown1Enum int32

const (
	MusicGameUnknown1Enum_MusicGameUnknown1Enum_NONE MusicGameUnknown1Enum = 0
	MusicGameUnknown1Enum_MusicGameUnknown1Enum_u2   MusicGameUnknown1Enum = 1
)

// Enum value maps for MusicGameUnknown1Enum.
var (
	MusicGameUnknown1Enum_name = map[int32]string{
		0: "MusicGameUnknown1Enum_NONE",
		1: "MusicGameUnknown1Enum_u2",
	}
	MusicGameUnknown1Enum_value = map[string]int32{
		"MusicGameUnknown1Enum_NONE": 0,
		"MusicGameUnknown1Enum_u2":   1,
	}
)

func (x MusicGameUnknown1Enum) Enum() *MusicGameUnknown1Enum {
	p := new(MusicGameUnknown1Enum)
	*p = x
	return p
}

func (x MusicGameUnknown1Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MusicGameUnknown1Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_MusicGameUnknown1Enum_proto_enumTypes[0].Descriptor()
}

func (MusicGameUnknown1Enum) Type() protoreflect.EnumType {
	return &file_MusicGameUnknown1Enum_proto_enumTypes[0]
}

func (x MusicGameUnknown1Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MusicGameUnknown1Enum.Descriptor instead.
func (MusicGameUnknown1Enum) EnumDescriptor() ([]byte, []int) {
	return file_MusicGameUnknown1Enum_proto_rawDescGZIP(), []int{0}
}

var File_MusicGameUnknown1Enum_proto protoreflect.FileDescriptor

var file_MusicGameUnknown1Enum_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x31, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x55, 0x0a, 0x15, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d,
	0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x31, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1e, 0x0a,
	0x1a, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x31, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x31, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x75, 0x32, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MusicGameUnknown1Enum_proto_rawDescOnce sync.Once
	file_MusicGameUnknown1Enum_proto_rawDescData = file_MusicGameUnknown1Enum_proto_rawDesc
)

func file_MusicGameUnknown1Enum_proto_rawDescGZIP() []byte {
	file_MusicGameUnknown1Enum_proto_rawDescOnce.Do(func() {
		file_MusicGameUnknown1Enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_MusicGameUnknown1Enum_proto_rawDescData)
	})
	return file_MusicGameUnknown1Enum_proto_rawDescData
}

var file_MusicGameUnknown1Enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MusicGameUnknown1Enum_proto_goTypes = []interface{}{
	(MusicGameUnknown1Enum)(0), // 0: proto.MusicGameUnknown1Enum
}
var file_MusicGameUnknown1Enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MusicGameUnknown1Enum_proto_init() }
func file_MusicGameUnknown1Enum_proto_init() {
	if File_MusicGameUnknown1Enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MusicGameUnknown1Enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MusicGameUnknown1Enum_proto_goTypes,
		DependencyIndexes: file_MusicGameUnknown1Enum_proto_depIdxs,
		EnumInfos:         file_MusicGameUnknown1Enum_proto_enumTypes,
	}.Build()
	File_MusicGameUnknown1Enum_proto = out.File
	file_MusicGameUnknown1Enum_proto_rawDesc = nil
	file_MusicGameUnknown1Enum_proto_goTypes = nil
	file_MusicGameUnknown1Enum_proto_depIdxs = nil
}