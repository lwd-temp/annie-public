// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: CoopRewardState.proto

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

type CoopRewardState int32

const (
	CoopRewardState_Unlock CoopRewardState = 0
	CoopRewardState_Lock   CoopRewardState = 1
	CoopRewardState_Taken  CoopRewardState = 2
)

// Enum value maps for CoopRewardState.
var (
	CoopRewardState_name = map[int32]string{
		0: "Unlock",
		1: "Lock",
		2: "Taken",
	}
	CoopRewardState_value = map[string]int32{
		"Unlock": 0,
		"Lock":   1,
		"Taken":  2,
	}
)

func (x CoopRewardState) Enum() *CoopRewardState {
	p := new(CoopRewardState)
	*p = x
	return p
}

func (x CoopRewardState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoopRewardState) Descriptor() protoreflect.EnumDescriptor {
	return file_CoopRewardState_proto_enumTypes[0].Descriptor()
}

func (CoopRewardState) Type() protoreflect.EnumType {
	return &file_CoopRewardState_proto_enumTypes[0]
}

func (x CoopRewardState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoopRewardState.Descriptor instead.
func (CoopRewardState) EnumDescriptor() ([]byte, []int) {
	return file_CoopRewardState_proto_rawDescGZIP(), []int{0}
}

var File_CoopRewardState_proto protoreflect.FileDescriptor

var file_CoopRewardState_proto_rawDesc = []byte{
	0x0a, 0x15, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x32,
	0x0a, 0x0f, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x61, 0x6b, 0x65, 0x6e,
	0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CoopRewardState_proto_rawDescOnce sync.Once
	file_CoopRewardState_proto_rawDescData = file_CoopRewardState_proto_rawDesc
)

func file_CoopRewardState_proto_rawDescGZIP() []byte {
	file_CoopRewardState_proto_rawDescOnce.Do(func() {
		file_CoopRewardState_proto_rawDescData = protoimpl.X.CompressGZIP(file_CoopRewardState_proto_rawDescData)
	})
	return file_CoopRewardState_proto_rawDescData
}

var file_CoopRewardState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CoopRewardState_proto_goTypes = []interface{}{
	(CoopRewardState)(0), // 0: proto.CoopRewardState
}
var file_CoopRewardState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CoopRewardState_proto_init() }
func file_CoopRewardState_proto_init() {
	if File_CoopRewardState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CoopRewardState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CoopRewardState_proto_goTypes,
		DependencyIndexes: file_CoopRewardState_proto_depIdxs,
		EnumInfos:         file_CoopRewardState_proto_enumTypes,
	}.Build()
	File_CoopRewardState_proto = out.File
	file_CoopRewardState_proto_rawDesc = nil
	file_CoopRewardState_proto_goTypes = nil
	file_CoopRewardState_proto_depIdxs = nil
}