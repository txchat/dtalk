// protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api.proto

package vip

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

type VIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *VIP) Reset() {
	*x = VIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VIP) ProtoMessage() {}

func (x *VIP) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VIP.ProtoReflect.Descriptor instead.
func (*VIP) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *VIP) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type AddVIPsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid []string `protobuf:"bytes,1,rep,name=uid,proto3" json:"uid,omitempty"`
}

func (x *AddVIPsReq) Reset() {
	*x = AddVIPsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddVIPsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVIPsReq) ProtoMessage() {}

func (x *AddVIPsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVIPsReq.ProtoReflect.Descriptor instead.
func (*AddVIPsReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *AddVIPsReq) GetUid() []string {
	if x != nil {
		return x.Uid
	}
	return nil
}

type AddVIPsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid []string `protobuf:"bytes,1,rep,name=uid,proto3" json:"uid,omitempty"`
}

func (x *AddVIPsReply) Reset() {
	*x = AddVIPsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddVIPsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVIPsReply) ProtoMessage() {}

func (x *AddVIPsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVIPsReply.ProtoReflect.Descriptor instead.
func (*AddVIPsReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *AddVIPsReply) GetUid() []string {
	if x != nil {
		return x.Uid
	}
	return nil
}

type GetVIPReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetVIPReq) Reset() {
	*x = GetVIPReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVIPReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVIPReq) ProtoMessage() {}

func (x *GetVIPReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVIPReq.ProtoReflect.Descriptor instead.
func (*GetVIPReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetVIPReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetVIPReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vip *VIP `protobuf:"bytes,1,opt,name=vip,proto3" json:"vip,omitempty"`
}

func (x *GetVIPReply) Reset() {
	*x = GetVIPReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVIPReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVIPReply) ProtoMessage() {}

func (x *GetVIPReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVIPReply.ProtoReflect.Descriptor instead.
func (*GetVIPReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetVIPReply) GetVip() *VIP {
	if x != nil {
		return x.Vip
	}
	return nil
}

type GetVIPsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetVIPsReq) Reset() {
	*x = GetVIPsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVIPsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVIPsReq) ProtoMessage() {}

func (x *GetVIPsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVIPsReq.ProtoReflect.Descriptor instead.
func (*GetVIPsReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetVIPsReq) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetVIPsReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetVIPsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vip        []*VIP `protobuf:"bytes,1,rep,name=vip,proto3" json:"vip,omitempty"`
	TotalCount int32  `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *GetVIPsReply) Reset() {
	*x = GetVIPsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVIPsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVIPsReply) ProtoMessage() {}

func (x *GetVIPsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVIPsReply.ProtoReflect.Descriptor instead.
func (*GetVIPsReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetVIPsReply) GetVip() []*VIP {
	if x != nil {
		return x.Vip
	}
	return nil
}

func (x *GetVIPsReply) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x76, 0x69, 0x70, 0x22, 0x17, 0x0a, 0x03, 0x56, 0x49, 0x50, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x1e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x56, 0x49, 0x50, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x20, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x56, 0x49, 0x50, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x1d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x56, 0x49, 0x50, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x22, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x56, 0x49, 0x50, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x20, 0x0a, 0x03, 0x76, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64,
	0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x56, 0x49, 0x50, 0x52, 0x03, 0x76, 0x69,
	0x70, 0x22, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x49, 0x50, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x50, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x56, 0x49, 0x50, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x03, 0x76,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b,
	0x2e, 0x76, 0x69, 0x70, 0x2e, 0x56, 0x49, 0x50, 0x52, 0x03, 0x76, 0x69, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb6, 0x01,
	0x0a, 0x06, 0x56, 0x49, 0x50, 0x53, 0x72, 0x76, 0x12, 0x39, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x56,
	0x49, 0x50, 0x73, 0x12, 0x15, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x41, 0x64, 0x64, 0x56, 0x49, 0x50, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x64, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x56, 0x49, 0x50, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x56, 0x49, 0x50, 0x73, 0x12, 0x15,
	0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x49,
	0x50, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x69,
	0x70, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x49, 0x50, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x56, 0x49, 0x50, 0x12, 0x14, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b,
	0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x49, 0x50, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x49,
	0x50, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x78, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x64, 0x74, 0x61, 0x6c,
	0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x69, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(*VIP)(nil),          // 0: dtalk.vip.VIP
	(*AddVIPsReq)(nil),   // 1: dtalk.vip.AddVIPsReq
	(*AddVIPsReply)(nil), // 2: dtalk.vip.AddVIPsReply
	(*GetVIPReq)(nil),    // 3: dtalk.vip.GetVIPReq
	(*GetVIPReply)(nil),  // 4: dtalk.vip.GetVIPReply
	(*GetVIPsReq)(nil),   // 5: dtalk.vip.GetVIPsReq
	(*GetVIPsReply)(nil), // 6: dtalk.vip.GetVIPsReply
}
var file_api_proto_depIdxs = []int32{
	0, // 0: dtalk.vip.GetVIPReply.vip:type_name -> dtalk.vip.VIP
	0, // 1: dtalk.vip.GetVIPsReply.vip:type_name -> dtalk.vip.VIP
	1, // 2: dtalk.vip.VIPSrv.AddVIPs:input_type -> dtalk.vip.AddVIPsReq
	5, // 3: dtalk.vip.VIPSrv.GetVIPs:input_type -> dtalk.vip.GetVIPsReq
	3, // 4: dtalk.vip.VIPSrv.GetVIP:input_type -> dtalk.vip.GetVIPReq
	2, // 5: dtalk.vip.VIPSrv.AddVIPs:output_type -> dtalk.vip.AddVIPsReply
	6, // 6: dtalk.vip.VIPSrv.GetVIPs:output_type -> dtalk.vip.GetVIPsReply
	4, // 7: dtalk.vip.VIPSrv.GetVIP:output_type -> dtalk.vip.GetVIPReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VIP); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddVIPsReq); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddVIPsReply); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVIPReq); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVIPReply); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVIPsReq); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVIPsReply); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}