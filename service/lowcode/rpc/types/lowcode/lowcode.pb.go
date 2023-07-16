// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: lowcode.proto

package lowcode

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

type SiteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteId int64 `protobuf:"varint,1,opt,name=siteId,proto3" json:"siteId,omitempty"`
}

func (x *SiteReq) Reset() {
	*x = SiteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lowcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiteReq) ProtoMessage() {}

func (x *SiteReq) ProtoReflect() protoreflect.Message {
	mi := &file_lowcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiteReq.ProtoReflect.Descriptor instead.
func (*SiteReq) Descriptor() ([]byte, []int) {
	return file_lowcode_proto_rawDescGZIP(), []int{0}
}

func (x *SiteReq) GetSiteId() int64 {
	if x != nil {
		return x.SiteId
	}
	return 0
}

type SiteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SiteReply) Reset() {
	*x = SiteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lowcode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiteReply) ProtoMessage() {}

func (x *SiteReply) ProtoReflect() protoreflect.Message {
	mi := &file_lowcode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiteReply.ProtoReflect.Descriptor instead.
func (*SiteReply) Descriptor() ([]byte, []int) {
	return file_lowcode_proto_rawDescGZIP(), []int{1}
}

var File_lowcode_proto protoreflect.FileDescriptor

var file_lowcode_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x21, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x73,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x3e, 0x0a, 0x07, 0x6c, 0x6f, 0x77, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6c, 0x6f,
	0x77, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lowcode_proto_rawDescOnce sync.Once
	file_lowcode_proto_rawDescData = file_lowcode_proto_rawDesc
)

func file_lowcode_proto_rawDescGZIP() []byte {
	file_lowcode_proto_rawDescOnce.Do(func() {
		file_lowcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_lowcode_proto_rawDescData)
	})
	return file_lowcode_proto_rawDescData
}

var file_lowcode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lowcode_proto_goTypes = []interface{}{
	(*SiteReq)(nil),   // 0: lowcode.siteReq
	(*SiteReply)(nil), // 1: lowcode.siteReply
}
var file_lowcode_proto_depIdxs = []int32{
	0, // 0: lowcode.lowcode.autoMigrate:input_type -> lowcode.siteReq
	1, // 1: lowcode.lowcode.autoMigrate:output_type -> lowcode.siteReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lowcode_proto_init() }
func file_lowcode_proto_init() {
	if File_lowcode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lowcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiteReq); i {
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
		file_lowcode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiteReply); i {
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
			RawDescriptor: file_lowcode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lowcode_proto_goTypes,
		DependencyIndexes: file_lowcode_proto_depIdxs,
		MessageInfos:      file_lowcode_proto_msgTypes,
	}.Build()
	File_lowcode_proto = out.File
	file_lowcode_proto_rawDesc = nil
	file_lowcode_proto_goTypes = nil
	file_lowcode_proto_depIdxs = nil
}
