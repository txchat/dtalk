// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pusher.proto

package pusher

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

type PushGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App  string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Gid  string `protobuf:"bytes,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Body []byte `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PushGroupReq) Reset() {
	*x = PushGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pusher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushGroupReq) ProtoMessage() {}

func (x *PushGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_pusher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushGroupReq.ProtoReflect.Descriptor instead.
func (*PushGroupReq) Descriptor() ([]byte, []int) {
	return file_pusher_proto_rawDescGZIP(), []int{0}
}

func (x *PushGroupReq) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *PushGroupReq) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *PushGroupReq) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type PushGroupResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushGroupResp) Reset() {
	*x = PushGroupResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pusher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushGroupResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushGroupResp) ProtoMessage() {}

func (x *PushGroupResp) ProtoReflect() protoreflect.Message {
	mi := &file_pusher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushGroupResp.ProtoReflect.Descriptor instead.
func (*PushGroupResp) Descriptor() ([]byte, []int) {
	return file_pusher_proto_rawDescGZIP(), []int{1}
}

type PushListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	From string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Uid  []string `protobuf:"bytes,3,rep,name=uid,proto3" json:"uid,omitempty"`
	Body []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PushListReq) Reset() {
	*x = PushListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pusher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushListReq) ProtoMessage() {}

func (x *PushListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pusher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushListReq.ProtoReflect.Descriptor instead.
func (*PushListReq) Descriptor() ([]byte, []int) {
	return file_pusher_proto_rawDescGZIP(), []int{2}
}

func (x *PushListReq) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *PushListReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PushListReq) GetUid() []string {
	if x != nil {
		return x.Uid
	}
	return nil
}

func (x *PushListReq) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type PushListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushListResp) Reset() {
	*x = PushListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pusher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushListResp) ProtoMessage() {}

func (x *PushListResp) ProtoReflect() protoreflect.Message {
	mi := &file_pusher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushListResp.ProtoReflect.Descriptor instead.
func (*PushListResp) Descriptor() ([]byte, []int) {
	return file_pusher_proto_rawDescGZIP(), []int{3}
}

var File_pusher_proto protoreflect.FileDescriptor

var file_pusher_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x75, 0x73, 0x68, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x0f,
	0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x59, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75,
	0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0x79, 0x0a, 0x06, 0x50, 0x75,
	0x73, 0x68, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x50, 0x75, 0x73, 0x68, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x14, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x65, 0x72,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35,
	0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x14, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pusher_proto_rawDescOnce sync.Once
	file_pusher_proto_rawDescData = file_pusher_proto_rawDesc
)

func file_pusher_proto_rawDescGZIP() []byte {
	file_pusher_proto_rawDescOnce.Do(func() {
		file_pusher_proto_rawDescData = protoimpl.X.CompressGZIP(file_pusher_proto_rawDescData)
	})
	return file_pusher_proto_rawDescData
}

var file_pusher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pusher_proto_goTypes = []interface{}{
	(*PushGroupReq)(nil),  // 0: pusher.PushGroupReq
	(*PushGroupResp)(nil), // 1: pusher.PushGroupResp
	(*PushListReq)(nil),   // 2: pusher.PushListReq
	(*PushListResp)(nil),  // 3: pusher.PushListResp
}
var file_pusher_proto_depIdxs = []int32{
	0, // 0: pusher.Pusher.PushGroup:input_type -> pusher.PushGroupReq
	2, // 1: pusher.Pusher.PushList:input_type -> pusher.PushListReq
	1, // 2: pusher.Pusher.PushGroup:output_type -> pusher.PushGroupResp
	3, // 3: pusher.Pusher.PushList:output_type -> pusher.PushListResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pusher_proto_init() }
func file_pusher_proto_init() {
	if File_pusher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pusher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushGroupReq); i {
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
		file_pusher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushGroupResp); i {
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
		file_pusher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushListReq); i {
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
		file_pusher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushListResp); i {
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
			RawDescriptor: file_pusher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pusher_proto_goTypes,
		DependencyIndexes: file_pusher_proto_depIdxs,
		MessageInfos:      file_pusher_proto_msgTypes,
	}.Build()
	File_pusher_proto = out.File
	file_pusher_proto_rawDesc = nil
	file_pusher_proto_goTypes = nil
	file_pusher_proto_depIdxs = nil
}
