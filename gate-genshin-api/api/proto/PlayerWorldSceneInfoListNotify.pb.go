// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PlayerWorldSceneInfoListNotify.proto

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

type PlayerWorldSceneInfoListNotify_CmdId int32

const (
	PlayerWorldSceneInfoListNotify_NONE             PlayerWorldSceneInfoListNotify_CmdId = 0
	PlayerWorldSceneInfoListNotify_ENET_CHANNEL_ID  PlayerWorldSceneInfoListNotify_CmdId = 0
	PlayerWorldSceneInfoListNotify_ENET_IS_RELIABLE PlayerWorldSceneInfoListNotify_CmdId = 1
	PlayerWorldSceneInfoListNotify_CMD_ID           PlayerWorldSceneInfoListNotify_CmdId = 3172
)

// Enum value maps for PlayerWorldSceneInfoListNotify_CmdId.
var (
	PlayerWorldSceneInfoListNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		3172: "CMD_ID",
	}
	PlayerWorldSceneInfoListNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           3172,
	}
)

func (x PlayerWorldSceneInfoListNotify_CmdId) Enum() *PlayerWorldSceneInfoListNotify_CmdId {
	p := new(PlayerWorldSceneInfoListNotify_CmdId)
	*p = x
	return p
}

func (x PlayerWorldSceneInfoListNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerWorldSceneInfoListNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerWorldSceneInfoListNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerWorldSceneInfoListNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerWorldSceneInfoListNotify_proto_enumTypes[0]
}

func (x PlayerWorldSceneInfoListNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerWorldSceneInfoListNotify_CmdId.Descriptor instead.
func (PlayerWorldSceneInfoListNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerWorldSceneInfoListNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerWorldSceneInfoListNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoList []*PlayerWorldSceneInfo `protobuf:"bytes,5,rep,name=info_list,json=infoList,proto3" json:"info_list,omitempty"`
}

func (x *PlayerWorldSceneInfoListNotify) Reset() {
	*x = PlayerWorldSceneInfoListNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerWorldSceneInfoListNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerWorldSceneInfoListNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerWorldSceneInfoListNotify) ProtoMessage() {}

func (x *PlayerWorldSceneInfoListNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerWorldSceneInfoListNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerWorldSceneInfoListNotify.ProtoReflect.Descriptor instead.
func (*PlayerWorldSceneInfoListNotify) Descriptor() ([]byte, []int) {
	return file_PlayerWorldSceneInfoListNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerWorldSceneInfoListNotify) GetInfoList() []*PlayerWorldSceneInfo {
	if x != nil {
		return x.InfoList
	}
	return nil
}

var File_PlayerWorldSceneInfoListNotify_proto protoreflect.FileDescriptor

var file_PlayerWorldSceneInfoListNotify_proto_rawDesc = []byte{
	0x0a, 0x24, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x1e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x38, 0x0a, 0x09,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xe4,
	0x18, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerWorldSceneInfoListNotify_proto_rawDescOnce sync.Once
	file_PlayerWorldSceneInfoListNotify_proto_rawDescData = file_PlayerWorldSceneInfoListNotify_proto_rawDesc
)

func file_PlayerWorldSceneInfoListNotify_proto_rawDescGZIP() []byte {
	file_PlayerWorldSceneInfoListNotify_proto_rawDescOnce.Do(func() {
		file_PlayerWorldSceneInfoListNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerWorldSceneInfoListNotify_proto_rawDescData)
	})
	return file_PlayerWorldSceneInfoListNotify_proto_rawDescData
}

var file_PlayerWorldSceneInfoListNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerWorldSceneInfoListNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerWorldSceneInfoListNotify_proto_goTypes = []interface{}{
	(PlayerWorldSceneInfoListNotify_CmdId)(0), // 0: proto.PlayerWorldSceneInfoListNotify.CmdId
	(*PlayerWorldSceneInfoListNotify)(nil),    // 1: proto.PlayerWorldSceneInfoListNotify
	(*PlayerWorldSceneInfo)(nil),              // 2: proto.PlayerWorldSceneInfo
}
var file_PlayerWorldSceneInfoListNotify_proto_depIdxs = []int32{
	2, // 0: proto.PlayerWorldSceneInfoListNotify.info_list:type_name -> proto.PlayerWorldSceneInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerWorldSceneInfoListNotify_proto_init() }
func file_PlayerWorldSceneInfoListNotify_proto_init() {
	if File_PlayerWorldSceneInfoListNotify_proto != nil {
		return
	}
	file_PlayerWorldSceneInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerWorldSceneInfoListNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerWorldSceneInfoListNotify); i {
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
			RawDescriptor: file_PlayerWorldSceneInfoListNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerWorldSceneInfoListNotify_proto_goTypes,
		DependencyIndexes: file_PlayerWorldSceneInfoListNotify_proto_depIdxs,
		EnumInfos:         file_PlayerWorldSceneInfoListNotify_proto_enumTypes,
		MessageInfos:      file_PlayerWorldSceneInfoListNotify_proto_msgTypes,
	}.Build()
	File_PlayerWorldSceneInfoListNotify_proto = out.File
	file_PlayerWorldSceneInfoListNotify_proto_rawDesc = nil
	file_PlayerWorldSceneInfoListNotify_proto_goTypes = nil
	file_PlayerWorldSceneInfoListNotify_proto_depIdxs = nil
}