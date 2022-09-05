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
// source: AchievementAllDataNotify.proto

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

// CmdId: 2676
// EnetChannelId: 0
// EnetIsReliable: true
type AchievementAllDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AchievementList       []*Achievement `protobuf:"bytes,4,rep,name=achievement_list,json=achievementList,proto3" json:"achievement_list,omitempty"`
	RewardTakenGoalIdList []uint32       `protobuf:"varint,2,rep,packed,name=reward_taken_goal_id_list,json=rewardTakenGoalIdList,proto3" json:"reward_taken_goal_id_list,omitempty"`
}

func (x *AchievementAllDataNotify) Reset() {
	*x = AchievementAllDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AchievementAllDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AchievementAllDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AchievementAllDataNotify) ProtoMessage() {}

func (x *AchievementAllDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AchievementAllDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AchievementAllDataNotify.ProtoReflect.Descriptor instead.
func (*AchievementAllDataNotify) Descriptor() ([]byte, []int) {
	return file_AchievementAllDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AchievementAllDataNotify) GetAchievementList() []*Achievement {
	if x != nil {
		return x.AchievementList
	}
	return nil
}

func (x *AchievementAllDataNotify) GetRewardTakenGoalIdList() []uint32 {
	if x != nil {
		return x.RewardTakenGoalIdList
	}
	return nil
}

var File_AchievementAllDataNotify_proto protoreflect.FileDescriptor

var file_AchievementAllDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x37, 0x0a, 0x10, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x19, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x15, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AchievementAllDataNotify_proto_rawDescOnce sync.Once
	file_AchievementAllDataNotify_proto_rawDescData = file_AchievementAllDataNotify_proto_rawDesc
)

func file_AchievementAllDataNotify_proto_rawDescGZIP() []byte {
	file_AchievementAllDataNotify_proto_rawDescOnce.Do(func() {
		file_AchievementAllDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AchievementAllDataNotify_proto_rawDescData)
	})
	return file_AchievementAllDataNotify_proto_rawDescData
}

var file_AchievementAllDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AchievementAllDataNotify_proto_goTypes = []interface{}{
	(*AchievementAllDataNotify)(nil), // 0: AchievementAllDataNotify
	(*Achievement)(nil),              // 1: Achievement
}
var file_AchievementAllDataNotify_proto_depIdxs = []int32{
	1, // 0: AchievementAllDataNotify.achievement_list:type_name -> Achievement
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AchievementAllDataNotify_proto_init() }
func file_AchievementAllDataNotify_proto_init() {
	if File_AchievementAllDataNotify_proto != nil {
		return
	}
	file_Achievement_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AchievementAllDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AchievementAllDataNotify); i {
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
			RawDescriptor: file_AchievementAllDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AchievementAllDataNotify_proto_goTypes,
		DependencyIndexes: file_AchievementAllDataNotify_proto_depIdxs,
		MessageInfos:      file_AchievementAllDataNotify_proto_msgTypes,
	}.Build()
	File_AchievementAllDataNotify_proto = out.File
	file_AchievementAllDataNotify_proto_rawDesc = nil
	file_AchievementAllDataNotify_proto_goTypes = nil
	file_AchievementAllDataNotify_proto_depIdxs = nil
}