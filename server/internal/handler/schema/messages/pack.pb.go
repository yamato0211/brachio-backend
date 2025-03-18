// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: messages/pack.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pack struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 中身
	Cards         []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pack) Reset() {
	*x = Pack{}
	mi := &file_messages_pack_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pack) ProtoMessage() {}

func (x *Pack) ProtoReflect() protoreflect.Message {
	mi := &file_messages_pack_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pack.ProtoReflect.Descriptor instead.
func (*Pack) Descriptor() ([]byte, []int) {
	return file_messages_pack_proto_rawDescGZIP(), []int{0}
}

func (x *Pack) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

var File_messages_pack_proto protoreflect.FileDescriptor

var file_messages_pack_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x1a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63,
	0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x04, 0x50, 0x61, 0x63,
	0x6b, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0xbb, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x42, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6d, 0x61,
	0x74, 0x6f, 0x30, 0x32, 0x31, 0x31, 0x2f, 0x62, 0x72, 0x61, 0x63, 0x68, 0x69, 0x6f, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x4d, 0x50, 0x58, 0xaa, 0x02, 0x0d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0xca, 0x02, 0x0d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x50, 0x61, 0x63, 0x6b, 0xe2, 0x02, 0x19,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x50, 0x61, 0x63, 0x6b, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x3a, 0x3a, 0x50, 0x61, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_messages_pack_proto_rawDescOnce sync.Once
	file_messages_pack_proto_rawDescData []byte
)

func file_messages_pack_proto_rawDescGZIP() []byte {
	file_messages_pack_proto_rawDescOnce.Do(func() {
		file_messages_pack_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_messages_pack_proto_rawDesc), len(file_messages_pack_proto_rawDesc)))
	})
	return file_messages_pack_proto_rawDescData
}

var file_messages_pack_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_pack_proto_goTypes = []any{
	(*Pack)(nil), // 0: messages.pack.Pack
	(*Card)(nil), // 1: messages.card.Card
}
var file_messages_pack_proto_depIdxs = []int32{
	1, // 0: messages.pack.Pack.cards:type_name -> messages.card.Card
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_pack_proto_init() }
func file_messages_pack_proto_init() {
	if File_messages_pack_proto != nil {
		return
	}
	file_messages_card_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_messages_pack_proto_rawDesc), len(file_messages_pack_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_pack_proto_goTypes,
		DependencyIndexes: file_messages_pack_proto_depIdxs,
		MessageInfos:      file_messages_pack_proto_msgTypes,
	}.Build()
	File_messages_pack_proto = out.File
	file_messages_pack_proto_goTypes = nil
	file_messages_pack_proto_depIdxs = nil
}
