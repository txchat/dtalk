// protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: github.com/txchat/dtalk/api/proto/chat.proto

package chat

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// event define
type Chat_Type int32

const (
	Chat_message Chat_Type = 0
	Chat_signal  Chat_Type = 1
)

// Enum value maps for Chat_Type.
var (
	Chat_Type_name = map[int32]string{
		0: "message",
		1: "signal",
	}
	Chat_Type_value = map[string]int32{
		"message": 0,
		"signal":  1,
	}
)

func (x Chat_Type) Enum() *Chat_Type {
	p := new(Chat_Type)
	*p = x
	return p
}

func (x Chat_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Chat_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_txchat_dtalk_api_proto_chat_proto_enumTypes[0].Descriptor()
}

func (Chat_Type) Type() protoreflect.EnumType {
	return &file_github_com_txchat_dtalk_api_proto_chat_proto_enumTypes[0]
}

func (x Chat_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Chat_Type.Descriptor instead.
func (Chat_Type) EnumDescriptor() ([]byte, []int) {
	return file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescGZIP(), []int{0, 0}
}

type SendMessageReply_FailedType int32

const (
	SendMessageReply_IsOK                   SendMessageReply_FailedType = 0
	SendMessageReply_InnerError             SendMessageReply_FailedType = 1
	SendMessageReply_UnSupportedMessageType SendMessageReply_FailedType = 2
	SendMessageReply_InsufficientPermission SendMessageReply_FailedType = 3
	SendMessageReply_IllegalFormat          SendMessageReply_FailedType = 4
	SendMessageReply_OutdatedFormat         SendMessageReply_FailedType = 5
)

// Enum value maps for SendMessageReply_FailedType.
var (
	SendMessageReply_FailedType_name = map[int32]string{
		0: "IsOK",
		1: "InnerError",
		2: "UnSupportedMessageType",
		3: "InsufficientPermission",
		4: "IllegalFormat",
		5: "OutdatedFormat",
	}
	SendMessageReply_FailedType_value = map[string]int32{
		"IsOK":                   0,
		"InnerError":             1,
		"UnSupportedMessageType": 2,
		"InsufficientPermission": 3,
		"IllegalFormat":          4,
		"OutdatedFormat":         5,
	}
)

func (x SendMessageReply_FailedType) Enum() *SendMessageReply_FailedType {
	p := new(SendMessageReply_FailedType)
	*p = x
	return p
}

func (x SendMessageReply_FailedType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMessageReply_FailedType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_txchat_dtalk_api_proto_chat_proto_enumTypes[1].Descriptor()
}

func (SendMessageReply_FailedType) Type() protoreflect.EnumType {
	return &file_github_com_txchat_dtalk_api_proto_chat_proto_enumTypes[1]
}

func (x SendMessageReply_FailedType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMessageReply_FailedType.Descriptor instead.
func (SendMessageReply_FailedType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescGZIP(), []int{1, 0}
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Chat_Type `protobuf:"varint,1,opt,name=type,proto3,enum=dtalk.api.proto.Chat_Type" json:"type,omitempty"`
	Seq  int64     `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Body []byte    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetType() Chat_Type {
	if x != nil {
		return x.Type
	}
	return Chat_message
}

func (x *Chat) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Chat) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type SendMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     SendMessageReply_FailedType `protobuf:"varint,1,opt,name=code,proto3,enum=dtalk.api.proto.SendMessageReply_FailedType" json:"code,omitempty"`
	Mid      string                      `protobuf:"bytes,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Datetime int64                       `protobuf:"varint,3,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Repeat   bool                        `protobuf:"varint,4,opt,name=repeat,proto3" json:"repeat,omitempty"`
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageReply) GetCode() SendMessageReply_FailedType {
	if x != nil {
		return x.Code
	}
	return SendMessageReply_IsOK
}

func (x *SendMessageReply) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *SendMessageReply) GetDatetime() int64 {
	if x != nil {
		return x.Datetime
	}
	return 0
}

func (x *SendMessageReply) GetRepeat() bool {
	if x != nil {
		return x.Repeat
	}
	return false
}

var File_github_com_txchat_dtalk_api_proto_chat_proto protoreflect.FileDescriptor

var file_github_com_txchat_dtalk_api_proto_chat_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x78, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7d, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x1f, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x10, 0x01, 0x22, 0xa2,
	0x02, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x0a,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x73,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x02,
	0x12, 0x1a, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x10, 0x04, 0x12,
	0x12, 0x0a, 0x0e, 0x4f, 0x75, 0x74, 0x64, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x10, 0x05, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x78, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x3b, 0x63, 0x68,
	0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescOnce sync.Once
	file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescData = file_github_com_txchat_dtalk_api_proto_chat_proto_rawDesc
)

func file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescGZIP() []byte {
	file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescOnce.Do(func() {
		file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescData)
	})
	return file_github_com_txchat_dtalk_api_proto_chat_proto_rawDescData
}

var file_github_com_txchat_dtalk_api_proto_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_txchat_dtalk_api_proto_chat_proto_goTypes = []interface{}{
	(Chat_Type)(0),                   // 0: dtalk.api.proto.Chat.Type
	(SendMessageReply_FailedType)(0), // 1: dtalk.api.proto.SendMessageReply.FailedType
	(*Chat)(nil),                     // 2: dtalk.api.proto.Chat
	(*SendMessageReply)(nil),         // 3: dtalk.api.proto.SendMessageReply
}
var file_github_com_txchat_dtalk_api_proto_chat_proto_depIdxs = []int32{
	0, // 0: dtalk.api.proto.Chat.type:type_name -> dtalk.api.proto.Chat.Type
	1, // 1: dtalk.api.proto.SendMessageReply.code:type_name -> dtalk.api.proto.SendMessageReply.FailedType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_txchat_dtalk_api_proto_chat_proto_init() }
func file_github_com_txchat_dtalk_api_proto_chat_proto_init() {
	if File_github_com_txchat_dtalk_api_proto_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReply); i {
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
			RawDescriptor: file_github_com_txchat_dtalk_api_proto_chat_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_txchat_dtalk_api_proto_chat_proto_goTypes,
		DependencyIndexes: file_github_com_txchat_dtalk_api_proto_chat_proto_depIdxs,
		EnumInfos:         file_github_com_txchat_dtalk_api_proto_chat_proto_enumTypes,
		MessageInfos:      file_github_com_txchat_dtalk_api_proto_chat_proto_msgTypes,
	}.Build()
	File_github_com_txchat_dtalk_api_proto_chat_proto = out.File
	file_github_com_txchat_dtalk_api_proto_chat_proto_rawDesc = nil
	file_github_com_txchat_dtalk_api_proto_chat_proto_goTypes = nil
	file_github_com_txchat_dtalk_api_proto_chat_proto_depIdxs = nil
}
