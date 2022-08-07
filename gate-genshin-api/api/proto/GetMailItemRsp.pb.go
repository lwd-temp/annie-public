// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: GetMailItemRsp.proto

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

type GetMailItemRsp_CmdId int32

const (
	GetMailItemRsp_NONE             GetMailItemRsp_CmdId = 0
	GetMailItemRsp_ENET_CHANNEL_ID  GetMailItemRsp_CmdId = 0
	GetMailItemRsp_ENET_IS_RELIABLE GetMailItemRsp_CmdId = 1
	GetMailItemRsp_CMD_ID           GetMailItemRsp_CmdId = 1432
)

// Enum value maps for GetMailItemRsp_CmdId.
var (
	GetMailItemRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1432: "CMD_ID",
	}
	GetMailItemRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1432,
	}
)

func (x GetMailItemRsp_CmdId) Enum() *GetMailItemRsp_CmdId {
	p := new(GetMailItemRsp_CmdId)
	*p = x
	return p
}

func (x GetMailItemRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetMailItemRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetMailItemRsp_proto_enumTypes[0].Descriptor()
}

func (GetMailItemRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetMailItemRsp_proto_enumTypes[0]
}

func (x GetMailItemRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetMailItemRsp_CmdId.Descriptor instead.
func (GetMailItemRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetMailItemRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetMailItemRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32         `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MailIdList []uint32      `protobuf:"varint,10,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
	ItemList   []*EquipParam `protobuf:"bytes,5,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *GetMailItemRsp) Reset() {
	*x = GetMailItemRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMailItemRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMailItemRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailItemRsp) ProtoMessage() {}

func (x *GetMailItemRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMailItemRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailItemRsp.ProtoReflect.Descriptor instead.
func (*GetMailItemRsp) Descriptor() ([]byte, []int) {
	return file_GetMailItemRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMailItemRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMailItemRsp) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

func (x *GetMailItemRsp) GetItemList() []*EquipParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_GetMailItemRsp_proto protoreflect.FileDescriptor

var file_GetMailItemRsp_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcb, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0c,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x98, 0x0b, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GetMailItemRsp_proto_rawDescOnce sync.Once
	file_GetMailItemRsp_proto_rawDescData = file_GetMailItemRsp_proto_rawDesc
)

func file_GetMailItemRsp_proto_rawDescGZIP() []byte {
	file_GetMailItemRsp_proto_rawDescOnce.Do(func() {
		file_GetMailItemRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMailItemRsp_proto_rawDescData)
	})
	return file_GetMailItemRsp_proto_rawDescData
}

var file_GetMailItemRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetMailItemRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMailItemRsp_proto_goTypes = []interface{}{
	(GetMailItemRsp_CmdId)(0), // 0: proto.GetMailItemRsp.CmdId
	(*GetMailItemRsp)(nil),    // 1: proto.GetMailItemRsp
	(*EquipParam)(nil),        // 2: proto.EquipParam
}
var file_GetMailItemRsp_proto_depIdxs = []int32{
	2, // 0: proto.GetMailItemRsp.item_list:type_name -> proto.EquipParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetMailItemRsp_proto_init() }
func file_GetMailItemRsp_proto_init() {
	if File_GetMailItemRsp_proto != nil {
		return
	}
	file_EquipParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMailItemRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMailItemRsp); i {
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
			RawDescriptor: file_GetMailItemRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMailItemRsp_proto_goTypes,
		DependencyIndexes: file_GetMailItemRsp_proto_depIdxs,
		EnumInfos:         file_GetMailItemRsp_proto_enumTypes,
		MessageInfos:      file_GetMailItemRsp_proto_msgTypes,
	}.Build()
	File_GetMailItemRsp_proto = out.File
	file_GetMailItemRsp_proto_rawDesc = nil
	file_GetMailItemRsp_proto_goTypes = nil
	file_GetMailItemRsp_proto_depIdxs = nil
}