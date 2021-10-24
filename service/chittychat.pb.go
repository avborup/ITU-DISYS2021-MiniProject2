// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: service/chittychat.proto

package service

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock   uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chittychat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_service_chittychat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_service_chittychat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_service_chittychat_proto protoreflect.FileDescriptor

var file_service_chittychat_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x69, 0x74,
	0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x22, 0x39, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x4b, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x12,
	0x3d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13,
	0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x1b,
	0x5a, 0x19, 0x64, 0x69, 0x73, 0x79, 0x73, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_chittychat_proto_rawDescOnce sync.Once
	file_service_chittychat_proto_rawDescData = file_service_chittychat_proto_rawDesc
)

func file_service_chittychat_proto_rawDescGZIP() []byte {
	file_service_chittychat_proto_rawDescOnce.Do(func() {
		file_service_chittychat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_chittychat_proto_rawDescData)
	})
	return file_service_chittychat_proto_rawDescData
}

var file_service_chittychat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_chittychat_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: chittychat.Message
}
var file_service_chittychat_proto_depIdxs = []int32{
	0, // 0: chittychat.chittychat.ChatSession:input_type -> chittychat.Message
	0, // 1: chittychat.chittychat.ChatSession:output_type -> chittychat.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_chittychat_proto_init() }
func file_service_chittychat_proto_init() {
	if File_service_chittychat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_chittychat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_service_chittychat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_chittychat_proto_goTypes,
		DependencyIndexes: file_service_chittychat_proto_depIdxs,
		MessageInfos:      file_service_chittychat_proto_msgTypes,
	}.Build()
	File_service_chittychat_proto = out.File
	file_service_chittychat_proto_rawDesc = nil
	file_service_chittychat_proto_goTypes = nil
	file_service_chittychat_proto_depIdxs = nil
}
