// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarSkillUpgradeRsp.proto

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

type AvatarSkillUpgradeRsp_CmdId int32

const (
	AvatarSkillUpgradeRsp_NONE             AvatarSkillUpgradeRsp_CmdId = 0
	AvatarSkillUpgradeRsp_ENET_CHANNEL_ID  AvatarSkillUpgradeRsp_CmdId = 0
	AvatarSkillUpgradeRsp_ENET_IS_RELIABLE AvatarSkillUpgradeRsp_CmdId = 1
	AvatarSkillUpgradeRsp_CMD_ID           AvatarSkillUpgradeRsp_CmdId = 1024
)

// Enum value maps for AvatarSkillUpgradeRsp_CmdId.
var (
	AvatarSkillUpgradeRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1024: "CMD_ID",
	}
	AvatarSkillUpgradeRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1024,
	}
)

func (x AvatarSkillUpgradeRsp_CmdId) Enum() *AvatarSkillUpgradeRsp_CmdId {
	p := new(AvatarSkillUpgradeRsp_CmdId)
	*p = x
	return p
}

func (x AvatarSkillUpgradeRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarSkillUpgradeRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarSkillUpgradeRsp_proto_enumTypes[0].Descriptor()
}

func (AvatarSkillUpgradeRsp_CmdId) Type() protoreflect.EnumType {
	return &file_AvatarSkillUpgradeRsp_proto_enumTypes[0]
}

func (x AvatarSkillUpgradeRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarSkillUpgradeRsp_CmdId.Descriptor instead.
func (AvatarSkillUpgradeRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AvatarSkillUpgradeRsp_proto_rawDescGZIP(), []int{0, 0}
}

type AvatarSkillUpgradeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AvatarGuid    uint64 `protobuf:"varint,2,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	AvatarSkillId uint32 `protobuf:"varint,3,opt,name=avatar_skill_id,json=avatarSkillId,proto3" json:"avatar_skill_id,omitempty"`
	OldLevel      uint32 `protobuf:"varint,4,opt,name=old_level,json=oldLevel,proto3" json:"old_level,omitempty"`
	CurLevel      uint32 `protobuf:"varint,5,opt,name=cur_level,json=curLevel,proto3" json:"cur_level,omitempty"`
}

func (x *AvatarSkillUpgradeRsp) Reset() {
	*x = AvatarSkillUpgradeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarSkillUpgradeRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarSkillUpgradeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarSkillUpgradeRsp) ProtoMessage() {}

func (x *AvatarSkillUpgradeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarSkillUpgradeRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarSkillUpgradeRsp.ProtoReflect.Descriptor instead.
func (*AvatarSkillUpgradeRsp) Descriptor() ([]byte, []int) {
	return file_AvatarSkillUpgradeRsp_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarSkillUpgradeRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AvatarSkillUpgradeRsp) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarSkillUpgradeRsp) GetAvatarSkillId() uint32 {
	if x != nil {
		return x.AvatarSkillId
	}
	return 0
}

func (x *AvatarSkillUpgradeRsp) GetOldLevel() uint32 {
	if x != nil {
		return x.OldLevel
	}
	return 0
}

func (x *AvatarSkillUpgradeRsp) GetCurLevel() uint32 {
	if x != nil {
		return x.CurLevel
	}
	return 0
}

var File_AvatarSkillUpgradeRsp_proto protoreflect.FileDescriptor

var file_AvatarSkillUpgradeRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x4d, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0x80, 0x08, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarSkillUpgradeRsp_proto_rawDescOnce sync.Once
	file_AvatarSkillUpgradeRsp_proto_rawDescData = file_AvatarSkillUpgradeRsp_proto_rawDesc
)

func file_AvatarSkillUpgradeRsp_proto_rawDescGZIP() []byte {
	file_AvatarSkillUpgradeRsp_proto_rawDescOnce.Do(func() {
		file_AvatarSkillUpgradeRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarSkillUpgradeRsp_proto_rawDescData)
	})
	return file_AvatarSkillUpgradeRsp_proto_rawDescData
}

var file_AvatarSkillUpgradeRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarSkillUpgradeRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarSkillUpgradeRsp_proto_goTypes = []interface{}{
	(AvatarSkillUpgradeRsp_CmdId)(0), // 0: proto.AvatarSkillUpgradeRsp.CmdId
	(*AvatarSkillUpgradeRsp)(nil),    // 1: proto.AvatarSkillUpgradeRsp
}
var file_AvatarSkillUpgradeRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarSkillUpgradeRsp_proto_init() }
func file_AvatarSkillUpgradeRsp_proto_init() {
	if File_AvatarSkillUpgradeRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarSkillUpgradeRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarSkillUpgradeRsp); i {
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
			RawDescriptor: file_AvatarSkillUpgradeRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarSkillUpgradeRsp_proto_goTypes,
		DependencyIndexes: file_AvatarSkillUpgradeRsp_proto_depIdxs,
		EnumInfos:         file_AvatarSkillUpgradeRsp_proto_enumTypes,
		MessageInfos:      file_AvatarSkillUpgradeRsp_proto_msgTypes,
	}.Build()
	File_AvatarSkillUpgradeRsp_proto = out.File
	file_AvatarSkillUpgradeRsp_proto_rawDesc = nil
	file_AvatarSkillUpgradeRsp_proto_goTypes = nil
	file_AvatarSkillUpgradeRsp_proto_depIdxs = nil
}