// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PlayerHomeCompInfoNotify.proto

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

type PlayerHomeCompInfoNotify_CmdId int32

const (
	PlayerHomeCompInfoNotify_NONE             PlayerHomeCompInfoNotify_CmdId = 0
	PlayerHomeCompInfoNotify_ENET_CHANNEL_ID  PlayerHomeCompInfoNotify_CmdId = 0
	PlayerHomeCompInfoNotify_ENET_IS_RELIABLE PlayerHomeCompInfoNotify_CmdId = 1
	PlayerHomeCompInfoNotify_CMD_ID           PlayerHomeCompInfoNotify_CmdId = 4563
)

// Enum value maps for PlayerHomeCompInfoNotify_CmdId.
var (
	PlayerHomeCompInfoNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		4563: "CMD_ID",
	}
	PlayerHomeCompInfoNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           4563,
	}
)

func (x PlayerHomeCompInfoNotify_CmdId) Enum() *PlayerHomeCompInfoNotify_CmdId {
	p := new(PlayerHomeCompInfoNotify_CmdId)
	*p = x
	return p
}

func (x PlayerHomeCompInfoNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerHomeCompInfoNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerHomeCompInfoNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerHomeCompInfoNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerHomeCompInfoNotify_proto_enumTypes[0]
}

func (x PlayerHomeCompInfoNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerHomeCompInfoNotify_CmdId.Descriptor instead.
func (PlayerHomeCompInfoNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerHomeCompInfoNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerHomeCompInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompInfo *PlayerHomeCompInfo `protobuf:"bytes,2,opt,name=comp_info,json=compInfo,proto3" json:"comp_info,omitempty"`
}

func (x *PlayerHomeCompInfoNotify) Reset() {
	*x = PlayerHomeCompInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerHomeCompInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerHomeCompInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerHomeCompInfoNotify) ProtoMessage() {}

func (x *PlayerHomeCompInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerHomeCompInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerHomeCompInfoNotify.ProtoReflect.Descriptor instead.
func (*PlayerHomeCompInfoNotify) Descriptor() ([]byte, []int) {
	return file_PlayerHomeCompInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerHomeCompInfoNotify) GetCompInfo() *PlayerHomeCompInfo {
	if x != nil {
		return x.CompInfo
	}
	return nil
}

var File_PlayerHomeCompInfoNotify_proto protoreflect.FileDescriptor

var file_PlayerHomeCompInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48,
	0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x36,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd3,
	0x23, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerHomeCompInfoNotify_proto_rawDescOnce sync.Once
	file_PlayerHomeCompInfoNotify_proto_rawDescData = file_PlayerHomeCompInfoNotify_proto_rawDesc
)

func file_PlayerHomeCompInfoNotify_proto_rawDescGZIP() []byte {
	file_PlayerHomeCompInfoNotify_proto_rawDescOnce.Do(func() {
		file_PlayerHomeCompInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerHomeCompInfoNotify_proto_rawDescData)
	})
	return file_PlayerHomeCompInfoNotify_proto_rawDescData
}

var file_PlayerHomeCompInfoNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerHomeCompInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerHomeCompInfoNotify_proto_goTypes = []interface{}{
	(PlayerHomeCompInfoNotify_CmdId)(0), // 0: proto.PlayerHomeCompInfoNotify.CmdId
	(*PlayerHomeCompInfoNotify)(nil),    // 1: proto.PlayerHomeCompInfoNotify
	(*PlayerHomeCompInfo)(nil),          // 2: proto.PlayerHomeCompInfo
}
var file_PlayerHomeCompInfoNotify_proto_depIdxs = []int32{
	2, // 0: proto.PlayerHomeCompInfoNotify.comp_info:type_name -> proto.PlayerHomeCompInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerHomeCompInfoNotify_proto_init() }
func file_PlayerHomeCompInfoNotify_proto_init() {
	if File_PlayerHomeCompInfoNotify_proto != nil {
		return
	}
	file_PlayerHomeCompInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerHomeCompInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerHomeCompInfoNotify); i {
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
			RawDescriptor: file_PlayerHomeCompInfoNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerHomeCompInfoNotify_proto_goTypes,
		DependencyIndexes: file_PlayerHomeCompInfoNotify_proto_depIdxs,
		EnumInfos:         file_PlayerHomeCompInfoNotify_proto_enumTypes,
		MessageInfos:      file_PlayerHomeCompInfoNotify_proto_msgTypes,
	}.Build()
	File_PlayerHomeCompInfoNotify_proto = out.File
	file_PlayerHomeCompInfoNotify_proto_rawDesc = nil
	file_PlayerHomeCompInfoNotify_proto_goTypes = nil
	file_PlayerHomeCompInfoNotify_proto_depIdxs = nil
}