// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: TakePlayerLevelRewardRsp.proto

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

type TakePlayerLevelRewardRsp_CmdId int32

const (
	TakePlayerLevelRewardRsp_NONE             TakePlayerLevelRewardRsp_CmdId = 0
	TakePlayerLevelRewardRsp_ENET_CHANNEL_ID  TakePlayerLevelRewardRsp_CmdId = 0
	TakePlayerLevelRewardRsp_ENET_IS_RELIABLE TakePlayerLevelRewardRsp_CmdId = 1
	TakePlayerLevelRewardRsp_CMD_ID           TakePlayerLevelRewardRsp_CmdId = 102
)

// Enum value maps for TakePlayerLevelRewardRsp_CmdId.
var (
	TakePlayerLevelRewardRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		102: "CMD_ID",
	}
	TakePlayerLevelRewardRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           102,
	}
)

func (x TakePlayerLevelRewardRsp_CmdId) Enum() *TakePlayerLevelRewardRsp_CmdId {
	p := new(TakePlayerLevelRewardRsp_CmdId)
	*p = x
	return p
}

func (x TakePlayerLevelRewardRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TakePlayerLevelRewardRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_TakePlayerLevelRewardRsp_proto_enumTypes[0].Descriptor()
}

func (TakePlayerLevelRewardRsp_CmdId) Type() protoreflect.EnumType {
	return &file_TakePlayerLevelRewardRsp_proto_enumTypes[0]
}

func (x TakePlayerLevelRewardRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TakePlayerLevelRewardRsp_CmdId.Descriptor instead.
func (TakePlayerLevelRewardRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_TakePlayerLevelRewardRsp_proto_rawDescGZIP(), []int{0, 0}
}

type TakePlayerLevelRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Level    uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	RewardId uint32 `protobuf:"varint,3,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
}

func (x *TakePlayerLevelRewardRsp) Reset() {
	*x = TakePlayerLevelRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakePlayerLevelRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakePlayerLevelRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakePlayerLevelRewardRsp) ProtoMessage() {}

func (x *TakePlayerLevelRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakePlayerLevelRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakePlayerLevelRewardRsp.ProtoReflect.Descriptor instead.
func (*TakePlayerLevelRewardRsp) Descriptor() ([]byte, []int) {
	return file_TakePlayerLevelRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakePlayerLevelRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *TakePlayerLevelRewardRsp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *TakePlayerLevelRewardRsp) GetRewardId() uint32 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

var File_TakePlayerLevelRewardRsp_proto protoreflect.FileDescriptor

var file_TakePlayerLevelRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x18, 0x54, 0x61, 0x6b, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x22, 0x4c, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x66, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_TakePlayerLevelRewardRsp_proto_rawDescOnce sync.Once
	file_TakePlayerLevelRewardRsp_proto_rawDescData = file_TakePlayerLevelRewardRsp_proto_rawDesc
)

func file_TakePlayerLevelRewardRsp_proto_rawDescGZIP() []byte {
	file_TakePlayerLevelRewardRsp_proto_rawDescOnce.Do(func() {
		file_TakePlayerLevelRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakePlayerLevelRewardRsp_proto_rawDescData)
	})
	return file_TakePlayerLevelRewardRsp_proto_rawDescData
}

var file_TakePlayerLevelRewardRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TakePlayerLevelRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakePlayerLevelRewardRsp_proto_goTypes = []interface{}{
	(TakePlayerLevelRewardRsp_CmdId)(0), // 0: proto.TakePlayerLevelRewardRsp.CmdId
	(*TakePlayerLevelRewardRsp)(nil),    // 1: proto.TakePlayerLevelRewardRsp
}
var file_TakePlayerLevelRewardRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TakePlayerLevelRewardRsp_proto_init() }
func file_TakePlayerLevelRewardRsp_proto_init() {
	if File_TakePlayerLevelRewardRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TakePlayerLevelRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakePlayerLevelRewardRsp); i {
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
			RawDescriptor: file_TakePlayerLevelRewardRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakePlayerLevelRewardRsp_proto_goTypes,
		DependencyIndexes: file_TakePlayerLevelRewardRsp_proto_depIdxs,
		EnumInfos:         file_TakePlayerLevelRewardRsp_proto_enumTypes,
		MessageInfos:      file_TakePlayerLevelRewardRsp_proto_msgTypes,
	}.Build()
	File_TakePlayerLevelRewardRsp_proto = out.File
	file_TakePlayerLevelRewardRsp_proto_rawDesc = nil
	file_TakePlayerLevelRewardRsp_proto_goTypes = nil
	file_TakePlayerLevelRewardRsp_proto_depIdxs = nil
}