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
// source: AbilityMixinRecoverInfo.proto

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

type AbilityMixinRecoverInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalId              uint32                 `protobuf:"varint,3,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	DataList             []uint32               `protobuf:"varint,4,rep,packed,name=data_list,json=dataList,proto3" json:"data_list,omitempty"`
	IsServerbuffModifier bool                   `protobuf:"varint,5,opt,name=is_serverbuff_modifier,json=isServerbuffModifier,proto3" json:"is_serverbuff_modifier,omitempty"`
	MassivePropList      []*MassivePropSyncInfo `protobuf:"bytes,6,rep,name=massive_prop_list,json=massivePropList,proto3" json:"massive_prop_list,omitempty"`
	// Types that are assignable to Source:
	//	*AbilityMixinRecoverInfo_InstancedAbilityId
	//	*AbilityMixinRecoverInfo_InstancedModifierId
	Source isAbilityMixinRecoverInfo_Source `protobuf_oneof:"source"`
}

func (x *AbilityMixinRecoverInfo) Reset() {
	*x = AbilityMixinRecoverInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityMixinRecoverInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMixinRecoverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMixinRecoverInfo) ProtoMessage() {}

func (x *AbilityMixinRecoverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityMixinRecoverInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMixinRecoverInfo.ProtoReflect.Descriptor instead.
func (*AbilityMixinRecoverInfo) Descriptor() ([]byte, []int) {
	return file_AbilityMixinRecoverInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityMixinRecoverInfo) GetLocalId() uint32 {
	if x != nil {
		return x.LocalId
	}
	return 0
}

func (x *AbilityMixinRecoverInfo) GetDataList() []uint32 {
	if x != nil {
		return x.DataList
	}
	return nil
}

func (x *AbilityMixinRecoverInfo) GetIsServerbuffModifier() bool {
	if x != nil {
		return x.IsServerbuffModifier
	}
	return false
}

func (x *AbilityMixinRecoverInfo) GetMassivePropList() []*MassivePropSyncInfo {
	if x != nil {
		return x.MassivePropList
	}
	return nil
}

func (m *AbilityMixinRecoverInfo) GetSource() isAbilityMixinRecoverInfo_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *AbilityMixinRecoverInfo) GetInstancedAbilityId() uint32 {
	if x, ok := x.GetSource().(*AbilityMixinRecoverInfo_InstancedAbilityId); ok {
		return x.InstancedAbilityId
	}
	return 0
}

func (x *AbilityMixinRecoverInfo) GetInstancedModifierId() uint32 {
	if x, ok := x.GetSource().(*AbilityMixinRecoverInfo_InstancedModifierId); ok {
		return x.InstancedModifierId
	}
	return 0
}

type isAbilityMixinRecoverInfo_Source interface {
	isAbilityMixinRecoverInfo_Source()
}

type AbilityMixinRecoverInfo_InstancedAbilityId struct {
	InstancedAbilityId uint32 `protobuf:"varint,1,opt,name=instanced_ability_id,json=instancedAbilityId,proto3,oneof"`
}

type AbilityMixinRecoverInfo_InstancedModifierId struct {
	InstancedModifierId uint32 `protobuf:"varint,2,opt,name=instanced_modifier_id,json=instancedModifierId,proto3,oneof"`
}

func (*AbilityMixinRecoverInfo_InstancedAbilityId) isAbilityMixinRecoverInfo_Source() {}

func (*AbilityMixinRecoverInfo_InstancedModifierId) isAbilityMixinRecoverInfo_Source() {}

var File_AbilityMixinRecoverInfo_proto protoreflect.FileDescriptor

var file_AbilityMixinRecoverInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x79, 0x6e, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x02, 0x0a, 0x17, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x16, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x75, 0x66, 0x66, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x11, 0x6d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x79, 0x6e,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x6d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x13, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityMixinRecoverInfo_proto_rawDescOnce sync.Once
	file_AbilityMixinRecoverInfo_proto_rawDescData = file_AbilityMixinRecoverInfo_proto_rawDesc
)

func file_AbilityMixinRecoverInfo_proto_rawDescGZIP() []byte {
	file_AbilityMixinRecoverInfo_proto_rawDescOnce.Do(func() {
		file_AbilityMixinRecoverInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityMixinRecoverInfo_proto_rawDescData)
	})
	return file_AbilityMixinRecoverInfo_proto_rawDescData
}

var file_AbilityMixinRecoverInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityMixinRecoverInfo_proto_goTypes = []interface{}{
	(*AbilityMixinRecoverInfo)(nil), // 0: AbilityMixinRecoverInfo
	(*MassivePropSyncInfo)(nil),     // 1: MassivePropSyncInfo
}
var file_AbilityMixinRecoverInfo_proto_depIdxs = []int32{
	1, // 0: AbilityMixinRecoverInfo.massive_prop_list:type_name -> MassivePropSyncInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AbilityMixinRecoverInfo_proto_init() }
func file_AbilityMixinRecoverInfo_proto_init() {
	if File_AbilityMixinRecoverInfo_proto != nil {
		return
	}
	file_MassivePropSyncInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityMixinRecoverInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMixinRecoverInfo); i {
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
	file_AbilityMixinRecoverInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AbilityMixinRecoverInfo_InstancedAbilityId)(nil),
		(*AbilityMixinRecoverInfo_InstancedModifierId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AbilityMixinRecoverInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityMixinRecoverInfo_proto_goTypes,
		DependencyIndexes: file_AbilityMixinRecoverInfo_proto_depIdxs,
		MessageInfos:      file_AbilityMixinRecoverInfo_proto_msgTypes,
	}.Build()
	File_AbilityMixinRecoverInfo_proto = out.File
	file_AbilityMixinRecoverInfo_proto_rawDesc = nil
	file_AbilityMixinRecoverInfo_proto_goTypes = nil
	file_AbilityMixinRecoverInfo_proto_depIdxs = nil
}