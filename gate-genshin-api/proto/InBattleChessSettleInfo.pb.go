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
// source: InBattleChessSettleInfo.proto

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

type InBattleChessSettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess     bool                     `protobuf:"varint,7,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	ChessExp      uint32                   `protobuf:"varint,11,opt,name=chess_exp,json=chessExp,proto3" json:"chess_exp,omitempty"`
	ChessLevel    uint32                   `protobuf:"varint,13,opt,name=chess_level,json=chessLevel,proto3" json:"chess_level,omitempty"`
	OldChessLevel uint32                   `protobuf:"varint,10,opt,name=old_chess_level,json=oldChessLevel,proto3" json:"old_chess_level,omitempty"`
	ScoreList     []*ExhibitionDisplayInfo `protobuf:"bytes,1,rep,name=score_list,json=scoreList,proto3" json:"score_list,omitempty"`
	SceneTimeMs   uint64                   `protobuf:"varint,14,opt,name=scene_time_ms,json=sceneTimeMs,proto3" json:"scene_time_ms,omitempty"`
	OldChessExp   uint32                   `protobuf:"varint,2,opt,name=old_chess_exp,json=oldChessExp,proto3" json:"old_chess_exp,omitempty"`
}

func (x *InBattleChessSettleInfo) Reset() {
	*x = InBattleChessSettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleChessSettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleChessSettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleChessSettleInfo) ProtoMessage() {}

func (x *InBattleChessSettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleChessSettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleChessSettleInfo.ProtoReflect.Descriptor instead.
func (*InBattleChessSettleInfo) Descriptor() ([]byte, []int) {
	return file_InBattleChessSettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleChessSettleInfo) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *InBattleChessSettleInfo) GetChessExp() uint32 {
	if x != nil {
		return x.ChessExp
	}
	return 0
}

func (x *InBattleChessSettleInfo) GetChessLevel() uint32 {
	if x != nil {
		return x.ChessLevel
	}
	return 0
}

func (x *InBattleChessSettleInfo) GetOldChessLevel() uint32 {
	if x != nil {
		return x.OldChessLevel
	}
	return 0
}

func (x *InBattleChessSettleInfo) GetScoreList() []*ExhibitionDisplayInfo {
	if x != nil {
		return x.ScoreList
	}
	return nil
}

func (x *InBattleChessSettleInfo) GetSceneTimeMs() uint64 {
	if x != nil {
		return x.SceneTimeMs
	}
	return 0
}

func (x *InBattleChessSettleInfo) GetOldChessExp() uint32 {
	if x != nil {
		return x.OldChessExp
	}
	return 0
}

var File_InBattleChessSettleInfo_proto protoreflect.FileDescriptor

var file_InBattleChessSettleInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a,
	0x17, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x73, 0x73,
	0x5f, 0x65, 0x78, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x68, 0x65, 0x73,
	0x73, 0x45, 0x78, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x73, 0x73,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x68, 0x65,
	0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x6f, 0x6c, 0x64, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a,
	0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x6c, 0x64, 0x5f,
	0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x6f, 0x6c, 0x64, 0x43, 0x68, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleChessSettleInfo_proto_rawDescOnce sync.Once
	file_InBattleChessSettleInfo_proto_rawDescData = file_InBattleChessSettleInfo_proto_rawDesc
)

func file_InBattleChessSettleInfo_proto_rawDescGZIP() []byte {
	file_InBattleChessSettleInfo_proto_rawDescOnce.Do(func() {
		file_InBattleChessSettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleChessSettleInfo_proto_rawDescData)
	})
	return file_InBattleChessSettleInfo_proto_rawDescData
}

var file_InBattleChessSettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InBattleChessSettleInfo_proto_goTypes = []interface{}{
	(*InBattleChessSettleInfo)(nil), // 0: InBattleChessSettleInfo
	(*ExhibitionDisplayInfo)(nil),   // 1: ExhibitionDisplayInfo
}
var file_InBattleChessSettleInfo_proto_depIdxs = []int32{
	1, // 0: InBattleChessSettleInfo.score_list:type_name -> ExhibitionDisplayInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InBattleChessSettleInfo_proto_init() }
func file_InBattleChessSettleInfo_proto_init() {
	if File_InBattleChessSettleInfo_proto != nil {
		return
	}
	file_ExhibitionDisplayInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InBattleChessSettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleChessSettleInfo); i {
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
			RawDescriptor: file_InBattleChessSettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleChessSettleInfo_proto_goTypes,
		DependencyIndexes: file_InBattleChessSettleInfo_proto_depIdxs,
		MessageInfos:      file_InBattleChessSettleInfo_proto_msgTypes,
	}.Build()
	File_InBattleChessSettleInfo_proto = out.File
	file_InBattleChessSettleInfo_proto_rawDesc = nil
	file_InBattleChessSettleInfo_proto_goTypes = nil
	file_InBattleChessSettleInfo_proto_depIdxs = nil
}