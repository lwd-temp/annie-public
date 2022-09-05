// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarExpeditionCallBackRsp.proto

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

// CmdId: 1726
// EnetChannelId: 0
// EnetIsReliable: true
type AvatarExpeditionCallBackRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpeditionInfoMap map[uint64]*AvatarExpeditionInfo `protobuf:"bytes,9,rep,name=expedition_info_map,json=expeditionInfoMap,proto3" json:"expedition_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Retcode           int32                            `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *AvatarExpeditionCallBackRsp) Reset() {
	*x = AvatarExpeditionCallBackRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarExpeditionCallBackRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarExpeditionCallBackRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarExpeditionCallBackRsp) ProtoMessage() {}

func (x *AvatarExpeditionCallBackRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarExpeditionCallBackRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarExpeditionCallBackRsp.ProtoReflect.Descriptor instead.
func (*AvatarExpeditionCallBackRsp) Descriptor() ([]byte, []int) {
	return file_AvatarExpeditionCallBackRsp_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarExpeditionCallBackRsp) GetExpeditionInfoMap() map[uint64]*AvatarExpeditionInfo {
	if x != nil {
		return x.ExpeditionInfoMap
	}
	return nil
}

func (x *AvatarExpeditionCallBackRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_AvatarExpeditionCallBackRsp_proto protoreflect.FileDescriptor

var file_AvatarExpeditionCallBackRsp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf9, 0x01, 0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x12,
	0x63, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x11, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x4d, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x5b,
	0x0a, 0x16, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarExpeditionCallBackRsp_proto_rawDescOnce sync.Once
	file_AvatarExpeditionCallBackRsp_proto_rawDescData = file_AvatarExpeditionCallBackRsp_proto_rawDesc
)

func file_AvatarExpeditionCallBackRsp_proto_rawDescGZIP() []byte {
	file_AvatarExpeditionCallBackRsp_proto_rawDescOnce.Do(func() {
		file_AvatarExpeditionCallBackRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarExpeditionCallBackRsp_proto_rawDescData)
	})
	return file_AvatarExpeditionCallBackRsp_proto_rawDescData
}

var file_AvatarExpeditionCallBackRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AvatarExpeditionCallBackRsp_proto_goTypes = []interface{}{
	(*AvatarExpeditionCallBackRsp)(nil), // 0: AvatarExpeditionCallBackRsp
	nil,                                 // 1: AvatarExpeditionCallBackRsp.ExpeditionInfoMapEntry
	(*AvatarExpeditionInfo)(nil),        // 2: AvatarExpeditionInfo
}
var file_AvatarExpeditionCallBackRsp_proto_depIdxs = []int32{
	1, // 0: AvatarExpeditionCallBackRsp.expedition_info_map:type_name -> AvatarExpeditionCallBackRsp.ExpeditionInfoMapEntry
	2, // 1: AvatarExpeditionCallBackRsp.ExpeditionInfoMapEntry.value:type_name -> AvatarExpeditionInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AvatarExpeditionCallBackRsp_proto_init() }
func file_AvatarExpeditionCallBackRsp_proto_init() {
	if File_AvatarExpeditionCallBackRsp_proto != nil {
		return
	}
	file_AvatarExpeditionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarExpeditionCallBackRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarExpeditionCallBackRsp); i {
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
			RawDescriptor: file_AvatarExpeditionCallBackRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarExpeditionCallBackRsp_proto_goTypes,
		DependencyIndexes: file_AvatarExpeditionCallBackRsp_proto_depIdxs,
		MessageInfos:      file_AvatarExpeditionCallBackRsp_proto_msgTypes,
	}.Build()
	File_AvatarExpeditionCallBackRsp_proto = out.File
	file_AvatarExpeditionCallBackRsp_proto_rawDesc = nil
	file_AvatarExpeditionCallBackRsp_proto_goTypes = nil
	file_AvatarExpeditionCallBackRsp_proto_depIdxs = nil
}