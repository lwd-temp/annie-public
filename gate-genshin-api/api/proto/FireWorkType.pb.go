// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: FireWorkType.proto

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

type FireWorkType int32

const (
	FireWorkType_FireWorkType_ODJKANKMPPJ FireWorkType = 0
	FireWorkType_FireWorkType_EFGLHEIODFN FireWorkType = 1
	FireWorkType_FireWorkType_JPBBBCFGHAK FireWorkType = 2
	FireWorkType_FireWorkType_IDCMGHBHBFH FireWorkType = 3
	FireWorkType_FireWorkType_ODDBNNDFMBO FireWorkType = 4
	FireWorkType_FireWorkType_AGIDMOGJOBD FireWorkType = 5
)

// Enum value maps for FireWorkType.
var (
	FireWorkType_name = map[int32]string{
		0: "FireWorkType_ODJKANKMPPJ",
		1: "FireWorkType_EFGLHEIODFN",
		2: "FireWorkType_JPBBBCFGHAK",
		3: "FireWorkType_IDCMGHBHBFH",
		4: "FireWorkType_ODDBNNDFMBO",
		5: "FireWorkType_AGIDMOGJOBD",
	}
	FireWorkType_value = map[string]int32{
		"FireWorkType_ODJKANKMPPJ": 0,
		"FireWorkType_EFGLHEIODFN": 1,
		"FireWorkType_JPBBBCFGHAK": 2,
		"FireWorkType_IDCMGHBHBFH": 3,
		"FireWorkType_ODDBNNDFMBO": 4,
		"FireWorkType_AGIDMOGJOBD": 5,
	}
)

func (x FireWorkType) Enum() *FireWorkType {
	p := new(FireWorkType)
	*p = x
	return p
}

func (x FireWorkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FireWorkType) Descriptor() protoreflect.EnumDescriptor {
	return file_FireWorkType_proto_enumTypes[0].Descriptor()
}

func (FireWorkType) Type() protoreflect.EnumType {
	return &file_FireWorkType_proto_enumTypes[0]
}

func (x FireWorkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FireWorkType.Descriptor instead.
func (FireWorkType) EnumDescriptor() ([]byte, []int) {
	return file_FireWorkType_proto_rawDescGZIP(), []int{0}
}

var File_FireWorkType_proto protoreflect.FileDescriptor

var file_FireWorkType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc2, 0x01, 0x0a, 0x0c,
	0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18,
	0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4f, 0x44, 0x4a,
	0x4b, 0x41, 0x4e, 0x4b, 0x4d, 0x50, 0x50, 0x4a, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x69,
	0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x45, 0x46, 0x47, 0x4c, 0x48,
	0x45, 0x49, 0x4f, 0x44, 0x46, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4a, 0x50, 0x42, 0x42, 0x42, 0x43, 0x46,
	0x47, 0x48, 0x41, 0x4b, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x49, 0x44, 0x43, 0x4d, 0x47, 0x48, 0x42, 0x48, 0x42,
	0x46, 0x48, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x4f, 0x44, 0x44, 0x42, 0x4e, 0x4e, 0x44, 0x46, 0x4d, 0x42, 0x4f,
	0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x41, 0x47, 0x49, 0x44, 0x4d, 0x4f, 0x47, 0x4a, 0x4f, 0x42, 0x44, 0x10, 0x05,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FireWorkType_proto_rawDescOnce sync.Once
	file_FireWorkType_proto_rawDescData = file_FireWorkType_proto_rawDesc
)

func file_FireWorkType_proto_rawDescGZIP() []byte {
	file_FireWorkType_proto_rawDescOnce.Do(func() {
		file_FireWorkType_proto_rawDescData = protoimpl.X.CompressGZIP(file_FireWorkType_proto_rawDescData)
	})
	return file_FireWorkType_proto_rawDescData
}

var file_FireWorkType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FireWorkType_proto_goTypes = []interface{}{
	(FireWorkType)(0), // 0: proto.FireWorkType
}
var file_FireWorkType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FireWorkType_proto_init() }
func file_FireWorkType_proto_init() {
	if File_FireWorkType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FireWorkType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FireWorkType_proto_goTypes,
		DependencyIndexes: file_FireWorkType_proto_depIdxs,
		EnumInfos:         file_FireWorkType_proto_enumTypes,
	}.Build()
	File_FireWorkType_proto = out.File
	file_FireWorkType_proto_rawDesc = nil
	file_FireWorkType_proto_goTypes = nil
	file_FireWorkType_proto_depIdxs = nil
}