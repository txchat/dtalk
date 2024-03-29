// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: call.proto

package call

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

type RTCType int32

const (
	RTCType_Undefined RTCType = 0
	RTCType_Audio     RTCType = 1
	RTCType_Video     RTCType = 2
)

// Enum value maps for RTCType.
var (
	RTCType_name = map[int32]string{
		0: "Undefined",
		1: "Audio",
		2: "Video",
	}
	RTCType_value = map[string]int32{
		"Undefined": 0,
		"Audio":     1,
		"Video":     2,
	}
)

func (x RTCType) Enum() *RTCType {
	p := new(RTCType)
	*p = x
	return p
}

func (x RTCType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RTCType) Descriptor() protoreflect.EnumDescriptor {
	return file_call_proto_enumTypes[0].Descriptor()
}

func (RTCType) Type() protoreflect.EnumType {
	return &file_call_proto_enumTypes[0]
}

func (x RTCType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RTCType.Descriptor instead.
func (RTCType) EnumDescriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{0}
}

type RejectType int32

const (
	RejectType_Reject   RejectType = 0
	RejectType_Occupied RejectType = 1
)

// Enum value maps for RejectType.
var (
	RejectType_name = map[int32]string{
		0: "Reject",
		1: "Occupied",
	}
	RejectType_value = map[string]int32{
		"Reject":   0,
		"Occupied": 1,
	}
)

func (x RejectType) Enum() *RejectType {
	p := new(RejectType)
	*p = x
	return p
}

func (x RejectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RejectType) Descriptor() protoreflect.EnumDescriptor {
	return file_call_proto_enumTypes[1].Descriptor()
}

func (RejectType) Type() protoreflect.EnumType {
	return &file_call_proto_enumTypes[1]
}

func (x RejectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RejectType.Descriptor instead.
func (RejectType) EnumDescriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{1}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId int64 `protobuf:"varint,1,opt,name=TraceId,proto3" json:"TraceId,omitempty"`
	RoomId  int64 `protobuf:"varint,2,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	RTCType int32 `protobuf:"varint,3,opt,name=RTCType,proto3" json:"RTCType,omitempty"`
	// Deadline 超出后对方未接就结束通话
	Deadline int64 `protobuf:"varint,4,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
	// Status 0=对方未接通, 1=双方正在通话中, 2=通话结束
	Status     int32    `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Invitees   []string `protobuf:"bytes,6,rep,name=Invitees,proto3" json:"Invitees,omitempty"`
	Caller     string   `protobuf:"bytes,7,opt,name=Caller,proto3" json:"Caller,omitempty"`
	Timeout    int32    `protobuf:"varint,8,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	CreateTime int64    `protobuf:"varint,9,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// GroupId 0=私聊，^0=群id
	GroupId int64 `protobuf:"varint,10,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetTraceId() int64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *Session) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *Session) GetRTCType() int32 {
	if x != nil {
		return x.RTCType
	}
	return 0
}

func (x *Session) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *Session) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Session) GetInvitees() []string {
	if x != nil {
		return x.Invitees
	}
	return nil
}

func (x *Session) GetCaller() string {
	if x != nil {
		return x.Caller
	}
	return ""
}

func (x *Session) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Session) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Session) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type PrivateOfferReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator string  `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Invitee  string  `protobuf:"bytes,2,opt,name=invitee,proto3" json:"invitee,omitempty"`
	RTCType  RTCType `protobuf:"varint,3,opt,name=RTCType,proto3,enum=call.RTCType" json:"RTCType,omitempty"`
}

func (x *PrivateOfferReq) Reset() {
	*x = PrivateOfferReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateOfferReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateOfferReq) ProtoMessage() {}

func (x *PrivateOfferReq) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateOfferReq.ProtoReflect.Descriptor instead.
func (*PrivateOfferReq) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{1}
}

func (x *PrivateOfferReq) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *PrivateOfferReq) GetInvitee() string {
	if x != nil {
		return x.Invitee
	}
	return ""
}

func (x *PrivateOfferReq) GetRTCType() RTCType {
	if x != nil {
		return x.RTCType
	}
	return RTCType_Undefined
}

type PrivateOfferResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *PrivateOfferResp) Reset() {
	*x = PrivateOfferResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateOfferResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateOfferResp) ProtoMessage() {}

func (x *PrivateOfferResp) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateOfferResp.ProtoReflect.Descriptor instead.
func (*PrivateOfferResp) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateOfferResp) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type GroupOfferReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator string   `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Invitees []string `protobuf:"bytes,2,rep,name=invitees,proto3" json:"invitees,omitempty"`
	RTCType  RTCType  `protobuf:"varint,3,opt,name=RTCType,proto3,enum=call.RTCType" json:"RTCType,omitempty"`
	GroupID  int64    `protobuf:"varint,4,opt,name=groupID,proto3" json:"groupID,omitempty"`
}

func (x *GroupOfferReq) Reset() {
	*x = GroupOfferReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupOfferReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupOfferReq) ProtoMessage() {}

func (x *GroupOfferReq) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupOfferReq.ProtoReflect.Descriptor instead.
func (*GroupOfferReq) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{3}
}

func (x *GroupOfferReq) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *GroupOfferReq) GetInvitees() []string {
	if x != nil {
		return x.Invitees
	}
	return nil
}

func (x *GroupOfferReq) GetRTCType() RTCType {
	if x != nil {
		return x.RTCType
	}
	return RTCType_Undefined
}

func (x *GroupOfferReq) GetGroupID() int64 {
	if x != nil {
		return x.GroupID
	}
	return 0
}

type GroupOfferResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *GroupOfferResp) Reset() {
	*x = GroupOfferResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupOfferResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupOfferResp) ProtoMessage() {}

func (x *GroupOfferResp) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupOfferResp.ProtoReflect.Descriptor instead.
func (*GroupOfferResp) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{4}
}

func (x *GroupOfferResp) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type AcceptReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	TraceId  int64  `protobuf:"varint,2,opt,name=traceId,proto3" json:"traceId,omitempty"`
}

func (x *AcceptReq) Reset() {
	*x = AcceptReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptReq) ProtoMessage() {}

func (x *AcceptReq) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptReq.ProtoReflect.Descriptor instead.
func (*AcceptReq) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{5}
}

func (x *AcceptReq) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *AcceptReq) GetTraceId() int64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

type AcceptResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId        int64  `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	UserSign      string `protobuf:"bytes,2,opt,name=UserSign,proto3" json:"UserSign,omitempty"`
	PrivateMapKey string `protobuf:"bytes,3,opt,name=PrivateMapKey,proto3" json:"PrivateMapKey,omitempty"`
	SDKAppID      int32  `protobuf:"varint,4,opt,name=SDKAppID,proto3" json:"SDKAppID,omitempty"`
}

func (x *AcceptResp) Reset() {
	*x = AcceptResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptResp) ProtoMessage() {}

func (x *AcceptResp) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptResp.ProtoReflect.Descriptor instead.
func (*AcceptResp) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptResp) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *AcceptResp) GetUserSign() string {
	if x != nil {
		return x.UserSign
	}
	return ""
}

func (x *AcceptResp) GetPrivateMapKey() string {
	if x != nil {
		return x.PrivateMapKey
	}
	return ""
}

func (x *AcceptResp) GetSDKAppID() int32 {
	if x != nil {
		return x.SDKAppID
	}
	return 0
}

type RejectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator   string     `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	TraceId    int64      `protobuf:"varint,2,opt,name=traceId,proto3" json:"traceId,omitempty"`
	RejectType RejectType `protobuf:"varint,3,opt,name=rejectType,proto3,enum=call.RejectType" json:"rejectType,omitempty"`
}

func (x *RejectReq) Reset() {
	*x = RejectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectReq) ProtoMessage() {}

func (x *RejectReq) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectReq.ProtoReflect.Descriptor instead.
func (*RejectReq) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{7}
}

func (x *RejectReq) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *RejectReq) GetTraceId() int64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *RejectReq) GetRejectType() RejectType {
	if x != nil {
		return x.RejectType
	}
	return RejectType_Reject
}

type RejectResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RejectResp) Reset() {
	*x = RejectResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectResp) ProtoMessage() {}

func (x *RejectResp) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectResp.ProtoReflect.Descriptor instead.
func (*RejectResp) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{8}
}

type CheckTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	TraceId  int64  `protobuf:"varint,2,opt,name=traceId,proto3" json:"traceId,omitempty"`
}

func (x *CheckTaskReq) Reset() {
	*x = CheckTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTaskReq) ProtoMessage() {}

func (x *CheckTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTaskReq.ProtoReflect.Descriptor instead.
func (*CheckTaskReq) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{9}
}

func (x *CheckTaskReq) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *CheckTaskReq) GetTraceId() int64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

type CheckTaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *CheckTaskResp) Reset() {
	*x = CheckTaskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTaskResp) ProtoMessage() {}

func (x *CheckTaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_call_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTaskResp.ProtoReflect.Descriptor instead.
func (*CheckTaskResp) Descriptor() ([]byte, []int) {
	return file_call_proto_rawDescGZIP(), []int{10}
}

func (x *CheckTaskResp) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_call_proto protoreflect.FileDescriptor

var file_call_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x6c, 0x6c, 0x22, 0x91, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x12,
	0x27, 0x0a, 0x07, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3b, 0x0a, 0x10, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x07, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x52, 0x54, 0x43, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x22, 0x39, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x82, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x70, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x44, 0x4b,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x44, 0x4b,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x22, 0x73, 0x0a, 0x09, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63,
	0x61, 0x6c, 0x6c, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x44, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x38,
	0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x27, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x2e, 0x0a, 0x07, 0x52, 0x54, 0x43, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x02, 0x2a, 0x26, 0x0a, 0x0a, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x10, 0x01,
	0x32, 0x8e, 0x02, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x3d, 0x0a, 0x0c, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x6c, 0x6c,
	0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x63, 0x61,
	0x6c, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2b, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x61,
	0x6c, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63,
	0x61, 0x6c, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b,
	0x0a, 0x06, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x61, 0x6c, 0x6c,
	0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x09, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63,
	0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_call_proto_rawDescOnce sync.Once
	file_call_proto_rawDescData = file_call_proto_rawDesc
)

func file_call_proto_rawDescGZIP() []byte {
	file_call_proto_rawDescOnce.Do(func() {
		file_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_call_proto_rawDescData)
	})
	return file_call_proto_rawDescData
}

var file_call_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_call_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_call_proto_goTypes = []interface{}{
	(RTCType)(0),             // 0: call.RTCType
	(RejectType)(0),          // 1: call.RejectType
	(*Session)(nil),          // 2: call.Session
	(*PrivateOfferReq)(nil),  // 3: call.PrivateOfferReq
	(*PrivateOfferResp)(nil), // 4: call.PrivateOfferResp
	(*GroupOfferReq)(nil),    // 5: call.GroupOfferReq
	(*GroupOfferResp)(nil),   // 6: call.GroupOfferResp
	(*AcceptReq)(nil),        // 7: call.AcceptReq
	(*AcceptResp)(nil),       // 8: call.AcceptResp
	(*RejectReq)(nil),        // 9: call.RejectReq
	(*RejectResp)(nil),       // 10: call.RejectResp
	(*CheckTaskReq)(nil),     // 11: call.CheckTaskReq
	(*CheckTaskResp)(nil),    // 12: call.CheckTaskResp
}
var file_call_proto_depIdxs = []int32{
	0,  // 0: call.PrivateOfferReq.RTCType:type_name -> call.RTCType
	2,  // 1: call.PrivateOfferResp.session:type_name -> call.Session
	0,  // 2: call.GroupOfferReq.RTCType:type_name -> call.RTCType
	2,  // 3: call.GroupOfferResp.session:type_name -> call.Session
	1,  // 4: call.RejectReq.rejectType:type_name -> call.RejectType
	2,  // 5: call.CheckTaskResp.session:type_name -> call.Session
	3,  // 6: call.Call.PrivateOffer:input_type -> call.PrivateOfferReq
	5,  // 7: call.Call.GroupOffer:input_type -> call.GroupOfferReq
	7,  // 8: call.Call.Accept:input_type -> call.AcceptReq
	9,  // 9: call.Call.Reject:input_type -> call.RejectReq
	11, // 10: call.Call.CheckTask:input_type -> call.CheckTaskReq
	4,  // 11: call.Call.PrivateOffer:output_type -> call.PrivateOfferResp
	6,  // 12: call.Call.GroupOffer:output_type -> call.GroupOfferResp
	8,  // 13: call.Call.Accept:output_type -> call.AcceptResp
	10, // 14: call.Call.Reject:output_type -> call.RejectResp
	12, // 15: call.Call.CheckTask:output_type -> call.CheckTaskResp
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_call_proto_init() }
func file_call_proto_init() {
	if File_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_call_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateOfferReq); i {
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
		file_call_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateOfferResp); i {
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
		file_call_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupOfferReq); i {
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
		file_call_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupOfferResp); i {
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
		file_call_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptReq); i {
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
		file_call_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptResp); i {
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
		file_call_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectReq); i {
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
		file_call_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectResp); i {
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
		file_call_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTaskReq); i {
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
		file_call_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTaskResp); i {
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
			RawDescriptor: file_call_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_call_proto_goTypes,
		DependencyIndexes: file_call_proto_depIdxs,
		EnumInfos:         file_call_proto_enumTypes,
		MessageInfos:      file_call_proto_msgTypes,
	}.Build()
	File_call_proto = out.File
	file_call_proto_rawDesc = nil
	file_call_proto_goTypes = nil
	file_call_proto_depIdxs = nil
}
