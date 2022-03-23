// protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api.proto

package backup

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

type QueryType int32

const (
	QueryType_Phone   QueryType = 0
	QueryType_Email   QueryType = 1
	QueryType_Address QueryType = 2
)

// Enum value maps for QueryType.
var (
	QueryType_name = map[int32]string{
		0: "Phone",
		1: "Email",
		2: "Address",
	}
	QueryType_value = map[string]int32{
		"Phone":   0,
		"Email":   1,
		"Address": 2,
	}
)

func (x QueryType) Enum() *QueryType {
	p := new(QueryType)
	*p = x
	return p
}

func (x QueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (QueryType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x QueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryType.Descriptor instead.
func (QueryType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type RetrieveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type QueryType `protobuf:"varint,1,opt,name=type,proto3,enum=dtalk.backup.QueryType" json:"type,omitempty"`
	Val  string    `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *RetrieveReq) Reset() {
	*x = RetrieveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveReq) ProtoMessage() {}

func (x *RetrieveReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RetrieveReq.ProtoReflect.Descriptor instead.
func (*RetrieveReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *RetrieveReq) GetType() QueryType {
	if x != nil {
		return x.Type
	}
	return QueryType_Phone
}

func (x *RetrieveReq) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type RetrieveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Area       string `protobuf:"bytes,2,opt,name=Area,proto3" json:"Area,omitempty"`
	Phone      string `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email      string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Mnemonic   string `protobuf:"bytes,5,opt,name=Mnemonic,proto3" json:"Mnemonic,omitempty"`
	PrivateKey string `protobuf:"bytes,6,opt,name=PrivateKey,proto3" json:"PrivateKey,omitempty"`
	UpdateTime int64  `protobuf:"varint,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	CreateTime int64  `protobuf:"varint,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *RetrieveReply) Reset() {
	*x = RetrieveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveReply) ProtoMessage() {}

func (x *RetrieveReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RetrieveReply.ProtoReflect.Descriptor instead.
func (*RetrieveReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveReply) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RetrieveReply) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *RetrieveReply) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RetrieveReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RetrieveReply) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *RetrieveReply) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *RetrieveReply) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *RetrieveReply) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x4c, 0x0a, 0x0b, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12,
	0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x2a,
	0x2e, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x02, 0x32,
	0x4c, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x42, 0x0a, 0x08, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x33, 0x33, 0x2e, 0x63, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x64, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(QueryType)(0),        // 0: dtalk.backup.QueryType
	(*RetrieveReq)(nil),   // 1: dtalk.backup.RetrieveReq
	(*RetrieveReply)(nil), // 2: dtalk.backup.RetrieveReply
}
var file_api_proto_depIdxs = []int32{
	0, // 0: dtalk.backup.RetrieveReq.type:type_name -> dtalk.backup.QueryType
	1, // 1: dtalk.backup.Backup.Retrieve:input_type -> dtalk.backup.RetrieveReq
	2, // 2: dtalk.backup.Backup.Retrieve:output_type -> dtalk.backup.RetrieveReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveReq); i {
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
			switch v := v.(*RetrieveReply); i {
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
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
