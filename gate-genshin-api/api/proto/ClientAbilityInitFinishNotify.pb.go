// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ClientAbilityInitFinishNotify.proto

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

type ClientAbilityInitFinishNotify_CmdId int32

const (
	ClientAbilityInitFinishNotify_ENET_CHANNEL_ID  ClientAbilityInitFinishNotify_CmdId = 0
	ClientAbilityInitFinishNotify_NONE             ClientAbilityInitFinishNotify_CmdId = 0
	ClientAbilityInitFinishNotify_ENET_IS_RELIABLE ClientAbilityInitFinishNotify_CmdId = 1
	ClientAbilityInitFinishNotify_IS_ALLOW_CLIENT  ClientAbilityInitFinishNotify_CmdId = 1
	ClientAbilityInitFinishNotify_CMD_ID           ClientAbilityInitFinishNotify_CmdId = 1184
)

// Enum value maps for ClientAbilityInitFinishNotify_CmdId.
var (
	ClientAbilityInitFinishNotify_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		1184: "CMD_ID",
	}
	ClientAbilityInitFinishNotify_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           1184,
	}
)

func (x ClientAbilityInitFinishNotify_CmdId) Enum() *ClientAbilityInitFinishNotify_CmdId {
	p := new(ClientAbilityInitFinishNotify_CmdId)
	*p = x
	return p
}

func (x ClientAbilityInitFinishNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientAbilityInitFinishNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ClientAbilityInitFinishNotify_proto_enumTypes[0].Descriptor()
}

func (ClientAbilityInitFinishNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ClientAbilityInitFinishNotify_proto_enumTypes[0]
}

func (x ClientAbilityInitFinishNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientAbilityInitFinishNotify_CmdId.Descriptor instead.
func (ClientAbilityInitFinishNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ClientAbilityInitFinishNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ClientAbilityInitFinishNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId uint32                `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Invokes  []*AbilityInvokeEntry `protobuf:"bytes,2,rep,name=invokes,proto3" json:"invokes,omitempty"`
}

func (x *ClientAbilityInitFinishNotify) Reset() {
	*x = ClientAbilityInitFinishNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientAbilityInitFinishNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAbilityInitFinishNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAbilityInitFinishNotify) ProtoMessage() {}

func (x *ClientAbilityInitFinishNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ClientAbilityInitFinishNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAbilityInitFinishNotify.ProtoReflect.Descriptor instead.
func (*ClientAbilityInitFinishNotify) Descriptor() ([]byte, []int) {
	return file_ClientAbilityInitFinishNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ClientAbilityInitFinishNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *ClientAbilityInitFinishNotify) GetInvokes() []*AbilityInvokeEntry {
	if x != nil {
		return x.Invokes
	}
	return nil
}

var File_ClientAbilityInitFinishNotify_proto protoreflect.FileDescriptor

var file_ClientAbilityInitFinishNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x69, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x1d, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d,
	0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xa0, 0x09, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ClientAbilityInitFinishNotify_proto_rawDescOnce sync.Once
	file_ClientAbilityInitFinishNotify_proto_rawDescData = file_ClientAbilityInitFinishNotify_proto_rawDesc
)

func file_ClientAbilityInitFinishNotify_proto_rawDescGZIP() []byte {
	file_ClientAbilityInitFinishNotify_proto_rawDescOnce.Do(func() {
		file_ClientAbilityInitFinishNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientAbilityInitFinishNotify_proto_rawDescData)
	})
	return file_ClientAbilityInitFinishNotify_proto_rawDescData
}

var file_ClientAbilityInitFinishNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ClientAbilityInitFinishNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientAbilityInitFinishNotify_proto_goTypes = []interface{}{
	(ClientAbilityInitFinishNotify_CmdId)(0), // 0: proto.ClientAbilityInitFinishNotify.CmdId
	(*ClientAbilityInitFinishNotify)(nil),    // 1: proto.ClientAbilityInitFinishNotify
	(*AbilityInvokeEntry)(nil),               // 2: proto.AbilityInvokeEntry
}
var file_ClientAbilityInitFinishNotify_proto_depIdxs = []int32{
	2, // 0: proto.ClientAbilityInitFinishNotify.invokes:type_name -> proto.AbilityInvokeEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ClientAbilityInitFinishNotify_proto_init() }
func file_ClientAbilityInitFinishNotify_proto_init() {
	if File_ClientAbilityInitFinishNotify_proto != nil {
		return
	}
	file_AbilityInvokeEntry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClientAbilityInitFinishNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAbilityInitFinishNotify); i {
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
			RawDescriptor: file_ClientAbilityInitFinishNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientAbilityInitFinishNotify_proto_goTypes,
		DependencyIndexes: file_ClientAbilityInitFinishNotify_proto_depIdxs,
		EnumInfos:         file_ClientAbilityInitFinishNotify_proto_enumTypes,
		MessageInfos:      file_ClientAbilityInitFinishNotify_proto_msgTypes,
	}.Build()
	File_ClientAbilityInitFinishNotify_proto = out.File
	file_ClientAbilityInitFinishNotify_proto_rawDesc = nil
	file_ClientAbilityInitFinishNotify_proto_goTypes = nil
	file_ClientAbilityInitFinishNotify_proto_depIdxs = nil
}