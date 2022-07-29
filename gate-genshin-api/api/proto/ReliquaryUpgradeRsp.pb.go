// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ReliquaryUpgradeRsp.proto

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

type ReliquaryUpgradeRsp_CmdId int32

const (
	ReliquaryUpgradeRsp_NONE             ReliquaryUpgradeRsp_CmdId = 0
	ReliquaryUpgradeRsp_ENET_CHANNEL_ID  ReliquaryUpgradeRsp_CmdId = 0
	ReliquaryUpgradeRsp_ENET_IS_RELIABLE ReliquaryUpgradeRsp_CmdId = 1
	ReliquaryUpgradeRsp_CMD_ID           ReliquaryUpgradeRsp_CmdId = 696
)

// Enum value maps for ReliquaryUpgradeRsp_CmdId.
var (
	ReliquaryUpgradeRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		696: "CMD_ID",
	}
	ReliquaryUpgradeRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           696,
	}
)

func (x ReliquaryUpgradeRsp_CmdId) Enum() *ReliquaryUpgradeRsp_CmdId {
	p := new(ReliquaryUpgradeRsp_CmdId)
	*p = x
	return p
}

func (x ReliquaryUpgradeRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReliquaryUpgradeRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ReliquaryUpgradeRsp_proto_enumTypes[0].Descriptor()
}

func (ReliquaryUpgradeRsp_CmdId) Type() protoreflect.EnumType {
	return &file_ReliquaryUpgradeRsp_proto_enumTypes[0]
}

func (x ReliquaryUpgradeRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReliquaryUpgradeRsp_CmdId.Descriptor instead.
func (ReliquaryUpgradeRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ReliquaryUpgradeRsp_proto_rawDescGZIP(), []int{0, 0}
}

type ReliquaryUpgradeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode             int32    `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	TargetReliquaryGuid uint64   `protobuf:"varint,2,opt,name=target_reliquary_guid,json=targetReliquaryGuid,proto3" json:"target_reliquary_guid,omitempty"`
	OldLevel            uint32   `protobuf:"varint,3,opt,name=old_level,json=oldLevel,proto3" json:"old_level,omitempty"`
	CurLevel            uint32   `protobuf:"varint,4,opt,name=cur_level,json=curLevel,proto3" json:"cur_level,omitempty"`
	PowerUpRate         uint32   `protobuf:"varint,5,opt,name=power_up_rate,json=powerUpRate,proto3" json:"power_up_rate,omitempty"`
	OldAppendPropList   []uint32 `protobuf:"varint,6,rep,packed,name=old_append_prop_list,json=oldAppendPropList,proto3" json:"old_append_prop_list,omitempty"`
	CurAppendPropList   []uint32 `protobuf:"varint,7,rep,packed,name=cur_append_prop_list,json=curAppendPropList,proto3" json:"cur_append_prop_list,omitempty"`
}

func (x *ReliquaryUpgradeRsp) Reset() {
	*x = ReliquaryUpgradeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReliquaryUpgradeRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReliquaryUpgradeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReliquaryUpgradeRsp) ProtoMessage() {}

func (x *ReliquaryUpgradeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ReliquaryUpgradeRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReliquaryUpgradeRsp.ProtoReflect.Descriptor instead.
func (*ReliquaryUpgradeRsp) Descriptor() ([]byte, []int) {
	return file_ReliquaryUpgradeRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ReliquaryUpgradeRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ReliquaryUpgradeRsp) GetTargetReliquaryGuid() uint64 {
	if x != nil {
		return x.TargetReliquaryGuid
	}
	return 0
}

func (x *ReliquaryUpgradeRsp) GetOldLevel() uint32 {
	if x != nil {
		return x.OldLevel
	}
	return 0
}

func (x *ReliquaryUpgradeRsp) GetCurLevel() uint32 {
	if x != nil {
		return x.CurLevel
	}
	return 0
}

func (x *ReliquaryUpgradeRsp) GetPowerUpRate() uint32 {
	if x != nil {
		return x.PowerUpRate
	}
	return 0
}

func (x *ReliquaryUpgradeRsp) GetOldAppendPropList() []uint32 {
	if x != nil {
		return x.OldAppendPropList
	}
	return nil
}

func (x *ReliquaryUpgradeRsp) GetCurAppendPropList() []uint32 {
	if x != nil {
		return x.CurAppendPropList
	}
	return nil
}

var File_ReliquaryUpgradeRsp_proto protoreflect.FileDescriptor

var file_ReliquaryUpgradeRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf2, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72,
	0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x69, 0x71,
	0x75, 0x61, 0x72, 0x79, 0x47, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x6c, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x55, 0x70, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x6f, 0x6c, 0x64, 0x5f, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x6f, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x5f, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x75, 0x72, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xb8, 0x05, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReliquaryUpgradeRsp_proto_rawDescOnce sync.Once
	file_ReliquaryUpgradeRsp_proto_rawDescData = file_ReliquaryUpgradeRsp_proto_rawDesc
)

func file_ReliquaryUpgradeRsp_proto_rawDescGZIP() []byte {
	file_ReliquaryUpgradeRsp_proto_rawDescOnce.Do(func() {
		file_ReliquaryUpgradeRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReliquaryUpgradeRsp_proto_rawDescData)
	})
	return file_ReliquaryUpgradeRsp_proto_rawDescData
}

var file_ReliquaryUpgradeRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ReliquaryUpgradeRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReliquaryUpgradeRsp_proto_goTypes = []interface{}{
	(ReliquaryUpgradeRsp_CmdId)(0), // 0: proto.ReliquaryUpgradeRsp.CmdId
	(*ReliquaryUpgradeRsp)(nil),    // 1: proto.ReliquaryUpgradeRsp
}
var file_ReliquaryUpgradeRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReliquaryUpgradeRsp_proto_init() }
func file_ReliquaryUpgradeRsp_proto_init() {
	if File_ReliquaryUpgradeRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ReliquaryUpgradeRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReliquaryUpgradeRsp); i {
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
			RawDescriptor: file_ReliquaryUpgradeRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReliquaryUpgradeRsp_proto_goTypes,
		DependencyIndexes: file_ReliquaryUpgradeRsp_proto_depIdxs,
		EnumInfos:         file_ReliquaryUpgradeRsp_proto_enumTypes,
		MessageInfos:      file_ReliquaryUpgradeRsp_proto_msgTypes,
	}.Build()
	File_ReliquaryUpgradeRsp_proto = out.File
	file_ReliquaryUpgradeRsp_proto_rawDesc = nil
	file_ReliquaryUpgradeRsp_proto_goTypes = nil
	file_ReliquaryUpgradeRsp_proto_depIdxs = nil
}