// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ChooseCurAvatarTeamReq.proto

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

type ChooseCurAvatarTeamReq_CmdId int32

const (
	ChooseCurAvatarTeamReq_ENET_CHANNEL_ID  ChooseCurAvatarTeamReq_CmdId = 0
	ChooseCurAvatarTeamReq_NONE             ChooseCurAvatarTeamReq_CmdId = 0
	ChooseCurAvatarTeamReq_ENET_IS_RELIABLE ChooseCurAvatarTeamReq_CmdId = 1
	ChooseCurAvatarTeamReq_IS_ALLOW_CLIENT  ChooseCurAvatarTeamReq_CmdId = 1
	ChooseCurAvatarTeamReq_CMD_ID           ChooseCurAvatarTeamReq_CmdId = 1663
)

// Enum value maps for ChooseCurAvatarTeamReq_CmdId.
var (
	ChooseCurAvatarTeamReq_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		1663: "CMD_ID",
	}
	ChooseCurAvatarTeamReq_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           1663,
	}
)

func (x ChooseCurAvatarTeamReq_CmdId) Enum() *ChooseCurAvatarTeamReq_CmdId {
	p := new(ChooseCurAvatarTeamReq_CmdId)
	*p = x
	return p
}

func (x ChooseCurAvatarTeamReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChooseCurAvatarTeamReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ChooseCurAvatarTeamReq_proto_enumTypes[0].Descriptor()
}

func (ChooseCurAvatarTeamReq_CmdId) Type() protoreflect.EnumType {
	return &file_ChooseCurAvatarTeamReq_proto_enumTypes[0]
}

func (x ChooseCurAvatarTeamReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChooseCurAvatarTeamReq_CmdId.Descriptor instead.
func (ChooseCurAvatarTeamReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ChooseCurAvatarTeamReq_proto_rawDescGZIP(), []int{0, 0}
}

type ChooseCurAvatarTeamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId uint32 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *ChooseCurAvatarTeamReq) Reset() {
	*x = ChooseCurAvatarTeamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChooseCurAvatarTeamReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChooseCurAvatarTeamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChooseCurAvatarTeamReq) ProtoMessage() {}

func (x *ChooseCurAvatarTeamReq) ProtoReflect() protoreflect.Message {
	mi := &file_ChooseCurAvatarTeamReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChooseCurAvatarTeamReq.ProtoReflect.Descriptor instead.
func (*ChooseCurAvatarTeamReq) Descriptor() ([]byte, []int) {
	return file_ChooseCurAvatarTeamReq_proto_rawDescGZIP(), []int{0}
}

func (x *ChooseCurAvatarTeamReq) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

var File_ChooseCurAvatarTeamReq_proto protoreflect.FileDescriptor

var file_ChooseCurAvatarTeamReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43, 0x75, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65,
	0x43, 0x75, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64,
	0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xff, 0x0c, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ChooseCurAvatarTeamReq_proto_rawDescOnce sync.Once
	file_ChooseCurAvatarTeamReq_proto_rawDescData = file_ChooseCurAvatarTeamReq_proto_rawDesc
)

func file_ChooseCurAvatarTeamReq_proto_rawDescGZIP() []byte {
	file_ChooseCurAvatarTeamReq_proto_rawDescOnce.Do(func() {
		file_ChooseCurAvatarTeamReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChooseCurAvatarTeamReq_proto_rawDescData)
	})
	return file_ChooseCurAvatarTeamReq_proto_rawDescData
}

var file_ChooseCurAvatarTeamReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChooseCurAvatarTeamReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChooseCurAvatarTeamReq_proto_goTypes = []interface{}{
	(ChooseCurAvatarTeamReq_CmdId)(0), // 0: proto.ChooseCurAvatarTeamReq.CmdId
	(*ChooseCurAvatarTeamReq)(nil),    // 1: proto.ChooseCurAvatarTeamReq
}
var file_ChooseCurAvatarTeamReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChooseCurAvatarTeamReq_proto_init() }
func file_ChooseCurAvatarTeamReq_proto_init() {
	if File_ChooseCurAvatarTeamReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChooseCurAvatarTeamReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChooseCurAvatarTeamReq); i {
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
			RawDescriptor: file_ChooseCurAvatarTeamReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChooseCurAvatarTeamReq_proto_goTypes,
		DependencyIndexes: file_ChooseCurAvatarTeamReq_proto_depIdxs,
		EnumInfos:         file_ChooseCurAvatarTeamReq_proto_enumTypes,
		MessageInfos:      file_ChooseCurAvatarTeamReq_proto_msgTypes,
	}.Build()
	File_ChooseCurAvatarTeamReq_proto = out.File
	file_ChooseCurAvatarTeamReq_proto_rawDesc = nil
	file_ChooseCurAvatarTeamReq_proto_goTypes = nil
	file_ChooseCurAvatarTeamReq_proto_depIdxs = nil
}