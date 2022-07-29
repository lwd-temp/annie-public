// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: GetPlayerTokenRsp.proto

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

type GetPlayerTokenRsp_CmdId int32

const (
	GetPlayerTokenRsp_NONE             GetPlayerTokenRsp_CmdId = 0
	GetPlayerTokenRsp_ENET_CHANNEL_ID  GetPlayerTokenRsp_CmdId = 0
	GetPlayerTokenRsp_ENET_IS_RELIABLE GetPlayerTokenRsp_CmdId = 1
	GetPlayerTokenRsp_CMD_ID           GetPlayerTokenRsp_CmdId = 118
)

// Enum value maps for GetPlayerTokenRsp_CmdId.
var (
	GetPlayerTokenRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		118: "CMD_ID",
	}
	GetPlayerTokenRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           118,
	}
)

func (x GetPlayerTokenRsp_CmdId) Enum() *GetPlayerTokenRsp_CmdId {
	p := new(GetPlayerTokenRsp_CmdId)
	*p = x
	return p
}

func (x GetPlayerTokenRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPlayerTokenRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetPlayerTokenRsp_proto_enumTypes[0].Descriptor()
}

func (GetPlayerTokenRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetPlayerTokenRsp_proto_enumTypes[0]
}

func (x GetPlayerTokenRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPlayerTokenRsp_CmdId.Descriptor instead.
func (GetPlayerTokenRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetPlayerTokenRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetPlayerTokenRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Msg                    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Uid                    uint32 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Token                  string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	BlackUidEndTime        uint32 `protobuf:"varint,5,opt,name=black_uid_end_time,json=blackUidEndTime,proto3" json:"black_uid_end_time,omitempty"`
	AccountType            uint32 `protobuf:"varint,6,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	AccountUid             string `protobuf:"bytes,7,opt,name=account_uid,json=accountUid,proto3" json:"account_uid,omitempty"`
	IsProficientPlayer     bool   `protobuf:"varint,8,opt,name=is_proficient_player,json=isProficientPlayer,proto3" json:"is_proficient_player,omitempty"`
	SecretKey              string `protobuf:"bytes,9,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	GmUid                  uint32 `protobuf:"varint,10,opt,name=gm_uid,json=gmUid,proto3" json:"gm_uid,omitempty"`
	SecretKeySeed          uint64 `protobuf:"varint,11,opt,name=secret_key_seed,json=secretKeySeed,proto3" json:"secret_key_seed,omitempty"`
	SecurityCmdBuffer      []byte `protobuf:"bytes,12,opt,name=security_cmd_buffer,json=securityCmdBuffer,proto3" json:"security_cmd_buffer,omitempty"`
	PlatformType           uint32 `protobuf:"varint,13,opt,name=platform_type,json=platformType,proto3" json:"platform_type,omitempty"`
	ExtraBinData           []byte `protobuf:"bytes,14,opt,name=extra_bin_data,json=extraBinData,proto3" json:"extra_bin_data,omitempty"`
	IsGuest                bool   `protobuf:"varint,15,opt,name=is_guest,json=isGuest,proto3" json:"is_guest,omitempty"`
	ChannelId              uint32 `protobuf:"varint,16,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	SubChannelId           uint32 `protobuf:"varint,17,opt,name=sub_channel_id,json=subChannelId,proto3" json:"sub_channel_id,omitempty"`
	Tag                    uint32 `protobuf:"varint,18,opt,name=tag,proto3" json:"tag,omitempty"`
	CountryCode            string `protobuf:"bytes,19,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	IsLoginWhiteList       bool   `protobuf:"varint,20,opt,name=is_login_white_list,json=isLoginWhiteList,proto3" json:"is_login_white_list,omitempty"`
	PsnId                  string `protobuf:"bytes,21,opt,name=psn_id,json=psnId,proto3" json:"psn_id,omitempty"`
	ClientVersionRandomKey string `protobuf:"bytes,22,opt,name=client_version_random_key,json=clientVersionRandomKey,proto3" json:"client_version_random_key,omitempty"`
	RegPlatform            uint32 `protobuf:"varint,23,opt,name=reg_platform,json=regPlatform,proto3" json:"reg_platform,omitempty"`
	ClientIpStr            string `protobuf:"bytes,24,opt,name=client_ip_str,json=clientIpStr,proto3" json:"client_ip_str,omitempty"`
}

func (x *GetPlayerTokenRsp) Reset() {
	*x = GetPlayerTokenRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPlayerTokenRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerTokenRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerTokenRsp) ProtoMessage() {}

func (x *GetPlayerTokenRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPlayerTokenRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerTokenRsp.ProtoReflect.Descriptor instead.
func (*GetPlayerTokenRsp) Descriptor() ([]byte, []int) {
	return file_GetPlayerTokenRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerTokenRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetBlackUidEndTime() uint32 {
	if x != nil {
		return x.BlackUidEndTime
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetAccountType() uint32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetAccountUid() string {
	if x != nil {
		return x.AccountUid
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetIsProficientPlayer() bool {
	if x != nil {
		return x.IsProficientPlayer
	}
	return false
}

func (x *GetPlayerTokenRsp) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetGmUid() uint32 {
	if x != nil {
		return x.GmUid
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetSecretKeySeed() uint64 {
	if x != nil {
		return x.SecretKeySeed
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetSecurityCmdBuffer() []byte {
	if x != nil {
		return x.SecurityCmdBuffer
	}
	return nil
}

func (x *GetPlayerTokenRsp) GetPlatformType() uint32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetExtraBinData() []byte {
	if x != nil {
		return x.ExtraBinData
	}
	return nil
}

func (x *GetPlayerTokenRsp) GetIsGuest() bool {
	if x != nil {
		return x.IsGuest
	}
	return false
}

func (x *GetPlayerTokenRsp) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetSubChannelId() uint32 {
	if x != nil {
		return x.SubChannelId
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetTag() uint32 {
	if x != nil {
		return x.Tag
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetIsLoginWhiteList() bool {
	if x != nil {
		return x.IsLoginWhiteList
	}
	return false
}

func (x *GetPlayerTokenRsp) GetPsnId() string {
	if x != nil {
		return x.PsnId
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetClientVersionRandomKey() string {
	if x != nil {
		return x.ClientVersionRandomKey
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetRegPlatform() uint32 {
	if x != nil {
		return x.RegPlatform
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetClientIpStr() string {
	if x != nil {
		return x.ClientIpStr
	}
	return ""
}

var File_GetPlayerTokenRsp_proto protoreflect.FileDescriptor

var file_GetPlayerTokenRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8e, 0x07, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x12, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x5f, 0x75, 0x69, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x69, 0x64,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06,
	0x67, 0x6d, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x6d,
	0x55, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6d, 0x64, 0x5f, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x6d, 0x64, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x69, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x42,
	0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x69,
	0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x73,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x73, 0x6e, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x72,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70,
	0x53, 0x74, 0x72, 0x22, 0x4c, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x76, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPlayerTokenRsp_proto_rawDescOnce sync.Once
	file_GetPlayerTokenRsp_proto_rawDescData = file_GetPlayerTokenRsp_proto_rawDesc
)

func file_GetPlayerTokenRsp_proto_rawDescGZIP() []byte {
	file_GetPlayerTokenRsp_proto_rawDescOnce.Do(func() {
		file_GetPlayerTokenRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPlayerTokenRsp_proto_rawDescData)
	})
	return file_GetPlayerTokenRsp_proto_rawDescData
}

var file_GetPlayerTokenRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetPlayerTokenRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPlayerTokenRsp_proto_goTypes = []interface{}{
	(GetPlayerTokenRsp_CmdId)(0), // 0: proto.GetPlayerTokenRsp.CmdId
	(*GetPlayerTokenRsp)(nil),    // 1: proto.GetPlayerTokenRsp
}
var file_GetPlayerTokenRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetPlayerTokenRsp_proto_init() }
func file_GetPlayerTokenRsp_proto_init() {
	if File_GetPlayerTokenRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetPlayerTokenRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerTokenRsp); i {
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
			RawDescriptor: file_GetPlayerTokenRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPlayerTokenRsp_proto_goTypes,
		DependencyIndexes: file_GetPlayerTokenRsp_proto_depIdxs,
		EnumInfos:         file_GetPlayerTokenRsp_proto_enumTypes,
		MessageInfos:      file_GetPlayerTokenRsp_proto_msgTypes,
	}.Build()
	File_GetPlayerTokenRsp_proto = out.File
	file_GetPlayerTokenRsp_proto_rawDesc = nil
	file_GetPlayerTokenRsp_proto_goTypes = nil
	file_GetPlayerTokenRsp_proto_depIdxs = nil
}