// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: device.proto

package device

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{0}
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ConnectId   string `protobuf:"bytes,2,opt,name=connectId,proto3" json:"connectId,omitempty"`
	DeviceUUid  string `protobuf:"bytes,3,opt,name=deviceUUid,proto3" json:"deviceUUid,omitempty"`
	DeviceType  int32  `protobuf:"varint,4,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	DeviceName  string `protobuf:"bytes,5,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	Username    string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	DeviceToken string `protobuf:"bytes,7,opt,name=deviceToken,proto3" json:"deviceToken,omitempty"`
	IsEnabled   bool   `protobuf:"varint,8,opt,name=isEnabled,proto3" json:"isEnabled,omitempty"`
	AddTime     int64  `protobuf:"varint,9,opt,name=addTime,proto3" json:"addTime,omitempty"`
	DTUid       string `protobuf:"bytes,10,opt,name=DTUid,proto3" json:"DTUid,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *DeviceInfo) GetConnectId() string {
	if x != nil {
		return x.ConnectId
	}
	return ""
}

func (x *DeviceInfo) GetDeviceUUid() string {
	if x != nil {
		return x.DeviceUUid
	}
	return ""
}

func (x *DeviceInfo) GetDeviceType() int32 {
	if x != nil {
		return x.DeviceType
	}
	return 0
}

func (x *DeviceInfo) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *DeviceInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeviceInfo) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *DeviceInfo) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *DeviceInfo) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *DeviceInfo) GetDTUid() string {
	if x != nil {
		return x.DTUid
	}
	return ""
}

type EnableThreadPushDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ConnId string `protobuf:"bytes,2,opt,name=connId,proto3" json:"connId,omitempty"`
}

func (x *EnableThreadPushDeviceRequest) Reset() {
	*x = EnableThreadPushDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableThreadPushDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableThreadPushDeviceRequest) ProtoMessage() {}

func (x *EnableThreadPushDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableThreadPushDeviceRequest.ProtoReflect.Descriptor instead.
func (*EnableThreadPushDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{2}
}

func (x *EnableThreadPushDeviceRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *EnableThreadPushDeviceRequest) GetConnId() string {
	if x != nil {
		return x.ConnId
	}
	return ""
}

type GetUserAllDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetUserAllDevicesRequest) Reset() {
	*x = GetUserAllDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAllDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAllDevicesRequest) ProtoMessage() {}

func (x *GetUserAllDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAllDevicesRequest.ProtoReflect.Descriptor instead.
func (*GetUserAllDevicesRequest) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserAllDevicesRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetUserAllDevicesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*DeviceInfo `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *GetUserAllDevicesReply) Reset() {
	*x = GetUserAllDevicesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAllDevicesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAllDevicesReply) ProtoMessage() {}

func (x *GetUserAllDevicesReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAllDevicesReply.ProtoReflect.Descriptor instead.
func (*GetUserAllDevicesReply) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserAllDevicesReply) GetDevices() []*DeviceInfo {
	if x != nil {
		return x.Devices
	}
	return nil
}

type GetDeviceByConnIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ConnID string `protobuf:"bytes,2,opt,name=connID,proto3" json:"connID,omitempty"`
}

func (x *GetDeviceByConnIdRequest) Reset() {
	*x = GetDeviceByConnIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceByConnIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceByConnIdRequest) ProtoMessage() {}

func (x *GetDeviceByConnIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceByConnIdRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceByConnIdRequest) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeviceByConnIdRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GetDeviceByConnIdRequest) GetConnID() string {
	if x != nil {
		return x.ConnID
	}
	return ""
}

type GetDeviceByConnIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *DeviceInfo `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *GetDeviceByConnIdReply) Reset() {
	*x = GetDeviceByConnIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceByConnIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceByConnIdReply) ProtoMessage() {}

func (x *GetDeviceByConnIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceByConnIdReply.ProtoReflect.Descriptor instead.
func (*GetDeviceByConnIdReply) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{6}
}

func (x *GetDeviceByConnIdReply) GetDevice() *DeviceInfo {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_device_proto protoreflect.FileDescriptor

var file_device_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0xa8, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x54, 0x55, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x54, 0x55, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x1d, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c,
	0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a,
	0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49,
	0x44, 0x22, 0x44, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79,
	0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x32, 0xaa, 0x02, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x4e, 0x0a, 0x16, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x55, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x12, 0x20,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_proto_rawDescOnce sync.Once
	file_device_proto_rawDescData = file_device_proto_rawDesc
)

func file_device_proto_rawDescGZIP() []byte {
	file_device_proto_rawDescOnce.Do(func() {
		file_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_proto_rawDescData)
	})
	return file_device_proto_rawDescData
}

var file_device_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_device_proto_goTypes = []interface{}{
	(*Empty)(nil),                         // 0: device.Empty
	(*DeviceInfo)(nil),                    // 1: device.DeviceInfo
	(*EnableThreadPushDeviceRequest)(nil), // 2: device.EnableThreadPushDeviceRequest
	(*GetUserAllDevicesRequest)(nil),      // 3: device.GetUserAllDevicesRequest
	(*GetUserAllDevicesReply)(nil),        // 4: device.GetUserAllDevicesReply
	(*GetDeviceByConnIdRequest)(nil),      // 5: device.GetDeviceByConnIdRequest
	(*GetDeviceByConnIdReply)(nil),        // 6: device.GetDeviceByConnIdReply
}
var file_device_proto_depIdxs = []int32{
	1, // 0: device.GetUserAllDevicesReply.devices:type_name -> device.DeviceInfo
	1, // 1: device.GetDeviceByConnIdReply.device:type_name -> device.DeviceInfo
	1, // 2: device.Device.AddDevice:input_type -> device.DeviceInfo
	2, // 3: device.Device.EnableThreadPushDevice:input_type -> device.EnableThreadPushDeviceRequest
	3, // 4: device.Device.GetUserAllDevices:input_type -> device.GetUserAllDevicesRequest
	5, // 5: device.Device.GetDeviceByConnId:input_type -> device.GetDeviceByConnIdRequest
	0, // 6: device.Device.AddDevice:output_type -> device.Empty
	0, // 7: device.Device.EnableThreadPushDevice:output_type -> device.Empty
	4, // 8: device.Device.GetUserAllDevices:output_type -> device.GetUserAllDevicesReply
	1, // 9: device.Device.GetDeviceByConnId:output_type -> device.DeviceInfo
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_device_proto_init() }
func file_device_proto_init() {
	if File_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
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
		file_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableThreadPushDeviceRequest); i {
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
		file_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAllDevicesRequest); i {
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
		file_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAllDevicesReply); i {
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
		file_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceByConnIdRequest); i {
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
		file_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceByConnIdReply); i {
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
			RawDescriptor: file_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_proto_goTypes,
		DependencyIndexes: file_device_proto_depIdxs,
		MessageInfos:      file_device_proto_msgTypes,
	}.Build()
	File_device_proto = out.File
	file_device_proto_rawDesc = nil
	file_device_proto_goTypes = nil
	file_device_proto_depIdxs = nil
}
