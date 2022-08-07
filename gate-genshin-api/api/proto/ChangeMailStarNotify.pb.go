// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ChangeMailStarNotify.proto

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

type ChangeMailStarNotify_CmdId int32

const (
	ChangeMailStarNotify_NONE             ChangeMailStarNotify_CmdId = 0
	ChangeMailStarNotify_ENET_CHANNEL_ID  ChangeMailStarNotify_CmdId = 0
	ChangeMailStarNotify_ENET_IS_RELIABLE ChangeMailStarNotify_CmdId = 1
	ChangeMailStarNotify_IS_ALLOW_CLIENT  ChangeMailStarNotify_CmdId = 1
	ChangeMailStarNotify_CMD_ID           ChangeMailStarNotify_CmdId = 1407
)

// Enum value maps for ChangeMailStarNotify_CmdId.
var (
	ChangeMailStarNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		1407: "CMD_ID",
	}
	ChangeMailStarNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           1407,
	}
)

func (x ChangeMailStarNotify_CmdId) Enum() *ChangeMailStarNotify_CmdId {
	p := new(ChangeMailStarNotify_CmdId)
	*p = x
	return p
}

func (x ChangeMailStarNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeMailStarNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ChangeMailStarNotify_proto_enumTypes[0].Descriptor()
}

func (ChangeMailStarNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ChangeMailStarNotify_proto_enumTypes[0]
}

func (x ChangeMailStarNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeMailStarNotify_CmdId.Descriptor instead.
func (ChangeMailStarNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ChangeMailStarNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ChangeMailStarNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailIdList []uint32 `protobuf:"varint,5,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
	IsStar     bool     `protobuf:"varint,9,opt,name=is_star,json=isStar,proto3" json:"is_star,omitempty"`
}

func (x *ChangeMailStarNotify) Reset() {
	*x = ChangeMailStarNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChangeMailStarNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeMailStarNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeMailStarNotify) ProtoMessage() {}

func (x *ChangeMailStarNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChangeMailStarNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeMailStarNotify.ProtoReflect.Descriptor instead.
func (*ChangeMailStarNotify) Descriptor() ([]byte, []int) {
	return file_ChangeMailStarNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeMailStarNotify) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

func (x *ChangeMailStarNotify) GetIsStar() bool {
	if x != nil {
		return x.IsStar
	}
	return false
}

var File_ChangeMailStarNotify_proto protoreflect.FileDescriptor

var file_ChangeMailStarNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61,
	0x69, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0c,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x53, 0x74, 0x61, 0x72, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0xff, 0x0a, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChangeMailStarNotify_proto_rawDescOnce sync.Once
	file_ChangeMailStarNotify_proto_rawDescData = file_ChangeMailStarNotify_proto_rawDesc
)

func file_ChangeMailStarNotify_proto_rawDescGZIP() []byte {
	file_ChangeMailStarNotify_proto_rawDescOnce.Do(func() {
		file_ChangeMailStarNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChangeMailStarNotify_proto_rawDescData)
	})
	return file_ChangeMailStarNotify_proto_rawDescData
}

var file_ChangeMailStarNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChangeMailStarNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChangeMailStarNotify_proto_goTypes = []interface{}{
	(ChangeMailStarNotify_CmdId)(0), // 0: proto.ChangeMailStarNotify.CmdId
	(*ChangeMailStarNotify)(nil),    // 1: proto.ChangeMailStarNotify
}
var file_ChangeMailStarNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChangeMailStarNotify_proto_init() }
func file_ChangeMailStarNotify_proto_init() {
	if File_ChangeMailStarNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChangeMailStarNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeMailStarNotify); i {
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
			RawDescriptor: file_ChangeMailStarNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChangeMailStarNotify_proto_goTypes,
		DependencyIndexes: file_ChangeMailStarNotify_proto_depIdxs,
		EnumInfos:         file_ChangeMailStarNotify_proto_enumTypes,
		MessageInfos:      file_ChangeMailStarNotify_proto_msgTypes,
	}.Build()
	File_ChangeMailStarNotify_proto = out.File
	file_ChangeMailStarNotify_proto_rawDesc = nil
	file_ChangeMailStarNotify_proto_goTypes = nil
	file_ChangeMailStarNotify_proto_depIdxs = nil
}