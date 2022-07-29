// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: CombatInvocationsNotify.proto

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

type CombatInvocationsNotify_CmdId int32

const (
	CombatInvocationsNotify_ENET_CHANNEL_ID  CombatInvocationsNotify_CmdId = 0
	CombatInvocationsNotify_NONE             CombatInvocationsNotify_CmdId = 0
	CombatInvocationsNotify_ENET_IS_RELIABLE CombatInvocationsNotify_CmdId = 1
	CombatInvocationsNotify_IS_ALLOW_CLIENT  CombatInvocationsNotify_CmdId = 1
	CombatInvocationsNotify_CMD_ID           CombatInvocationsNotify_CmdId = 359
)

// Enum value maps for CombatInvocationsNotify_CmdId.
var (
	CombatInvocationsNotify_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		359: "CMD_ID",
	}
	CombatInvocationsNotify_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           359,
	}
)

func (x CombatInvocationsNotify_CmdId) Enum() *CombatInvocationsNotify_CmdId {
	p := new(CombatInvocationsNotify_CmdId)
	*p = x
	return p
}

func (x CombatInvocationsNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CombatInvocationsNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_CombatInvocationsNotify_proto_enumTypes[0].Descriptor()
}

func (CombatInvocationsNotify_CmdId) Type() protoreflect.EnumType {
	return &file_CombatInvocationsNotify_proto_enumTypes[0]
}

func (x CombatInvocationsNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CombatInvocationsNotify_CmdId.Descriptor instead.
func (CombatInvocationsNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_CombatInvocationsNotify_proto_rawDescGZIP(), []int{0, 0}
}

type CombatInvocationsNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvokeList []*CombatInvokeEntry `protobuf:"bytes,1,rep,name=invoke_list,json=invokeList,proto3" json:"invoke_list,omitempty"`
}

func (x *CombatInvocationsNotify) Reset() {
	*x = CombatInvocationsNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CombatInvocationsNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombatInvocationsNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombatInvocationsNotify) ProtoMessage() {}

func (x *CombatInvocationsNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CombatInvocationsNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombatInvocationsNotify.ProtoReflect.Descriptor instead.
func (*CombatInvocationsNotify) Descriptor() ([]byte, []int) {
	return file_CombatInvocationsNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CombatInvocationsNotify) GetInvokeList() []*CombatInvokeEntry {
	if x != nil {
		return x.InvokeList
	}
	return nil
}

var File_CombatInvocationsNotify_proto protoreflect.FileDescriptor

var file_CombatInvocationsNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb8, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x39, 0x0a, 0x0b, 0x69,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xe7, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CombatInvocationsNotify_proto_rawDescOnce sync.Once
	file_CombatInvocationsNotify_proto_rawDescData = file_CombatInvocationsNotify_proto_rawDesc
)

func file_CombatInvocationsNotify_proto_rawDescGZIP() []byte {
	file_CombatInvocationsNotify_proto_rawDescOnce.Do(func() {
		file_CombatInvocationsNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CombatInvocationsNotify_proto_rawDescData)
	})
	return file_CombatInvocationsNotify_proto_rawDescData
}

var file_CombatInvocationsNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CombatInvocationsNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CombatInvocationsNotify_proto_goTypes = []interface{}{
	(CombatInvocationsNotify_CmdId)(0), // 0: proto.CombatInvocationsNotify.CmdId
	(*CombatInvocationsNotify)(nil),    // 1: proto.CombatInvocationsNotify
	(*CombatInvokeEntry)(nil),          // 2: proto.CombatInvokeEntry
}
var file_CombatInvocationsNotify_proto_depIdxs = []int32{
	2, // 0: proto.CombatInvocationsNotify.invoke_list:type_name -> proto.CombatInvokeEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CombatInvocationsNotify_proto_init() }
func file_CombatInvocationsNotify_proto_init() {
	if File_CombatInvocationsNotify_proto != nil {
		return
	}
	file_CombatInvokeEntry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CombatInvocationsNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombatInvocationsNotify); i {
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
			RawDescriptor: file_CombatInvocationsNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CombatInvocationsNotify_proto_goTypes,
		DependencyIndexes: file_CombatInvocationsNotify_proto_depIdxs,
		EnumInfos:         file_CombatInvocationsNotify_proto_enumTypes,
		MessageInfos:      file_CombatInvocationsNotify_proto_msgTypes,
	}.Build()
	File_CombatInvocationsNotify_proto = out.File
	file_CombatInvocationsNotify_proto_rawDesc = nil
	file_CombatInvocationsNotify_proto_goTypes = nil
	file_CombatInvocationsNotify_proto_depIdxs = nil
}