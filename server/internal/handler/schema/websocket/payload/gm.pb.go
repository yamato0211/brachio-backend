// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: websocket/payload/gm.proto

package payload

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

// ルーム作成イベント
type CreateRoomEventPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ルームの合言葉
	Password      string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoomEventPayload) Reset() {
	*x = CreateRoomEventPayload{}
	mi := &file_websocket_payload_gm_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomEventPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomEventPayload) ProtoMessage() {}

func (x *CreateRoomEventPayload) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_payload_gm_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomEventPayload.ProtoReflect.Descriptor instead.
func (*CreateRoomEventPayload) Descriptor() ([]byte, []int) {
	return file_websocket_payload_gm_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomEventPayload) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// ルーム入室イベント
type EnterRoomEventPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ルームの合言葉
	Password      string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	DeckId        string `protobuf:"bytes,2,opt,name=deckId,proto3" json:"deckId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnterRoomEventPayload) Reset() {
	*x = EnterRoomEventPayload{}
	mi := &file_websocket_payload_gm_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnterRoomEventPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterRoomEventPayload) ProtoMessage() {}

func (x *EnterRoomEventPayload) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_payload_gm_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterRoomEventPayload.ProtoReflect.Descriptor instead.
func (*EnterRoomEventPayload) Descriptor() ([]byte, []int) {
	return file_websocket_payload_gm_proto_rawDescGZIP(), []int{1}
}

func (x *EnterRoomEventPayload) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *EnterRoomEventPayload) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

// マッチング完了イベント
type MatchingCompleteEventPayload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OpponentId    string                 `protobuf:"bytes,1,opt,name=opponentId,proto3" json:"opponentId,omitempty"`
	BattleId      string                 `protobuf:"bytes,2,opt,name=battleId,proto3" json:"battleId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchingCompleteEventPayload) Reset() {
	*x = MatchingCompleteEventPayload{}
	mi := &file_websocket_payload_gm_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchingCompleteEventPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingCompleteEventPayload) ProtoMessage() {}

func (x *MatchingCompleteEventPayload) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_payload_gm_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingCompleteEventPayload.ProtoReflect.Descriptor instead.
func (*MatchingCompleteEventPayload) Descriptor() ([]byte, []int) {
	return file_websocket_payload_gm_proto_rawDescGZIP(), []int{2}
}

func (x *MatchingCompleteEventPayload) GetOpponentId() string {
	if x != nil {
		return x.OpponentId
	}
	return ""
}

func (x *MatchingCompleteEventPayload) GetBattleId() string {
	if x != nil {
		return x.BattleId
	}
	return ""
}

// 先攻後攻決定イベント
type DecideOrderEventPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 先攻のユーザーID
	FirstUserId string `protobuf:"bytes,1,opt,name=firstUserId,proto3" json:"firstUserId,omitempty"`
	// 後攻のユーザーID
	SecondUserId  string `protobuf:"bytes,2,opt,name=secondUserId,proto3" json:"secondUserId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecideOrderEventPayload) Reset() {
	*x = DecideOrderEventPayload{}
	mi := &file_websocket_payload_gm_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecideOrderEventPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecideOrderEventPayload) ProtoMessage() {}

func (x *DecideOrderEventPayload) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_payload_gm_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecideOrderEventPayload.ProtoReflect.Descriptor instead.
func (*DecideOrderEventPayload) Descriptor() ([]byte, []int) {
	return file_websocket_payload_gm_proto_rawDescGZIP(), []int{3}
}

func (x *DecideOrderEventPayload) GetFirstUserId() string {
	if x != nil {
		return x.FirstUserId
	}
	return ""
}

func (x *DecideOrderEventPayload) GetSecondUserId() string {
	if x != nil {
		return x.SecondUserId
	}
	return ""
}

var File_websocket_payload_gm_proto protoreflect.FileDescriptor

var file_websocket_payload_gm_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2f, 0x67, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x77, 0x65,
	0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x67, 0x6d, 0x22, 0x34, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x1c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x5f, 0x0a, 0x17, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x42, 0xe6, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x67, 0x6d, 0x42,
	0x07, 0x47, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6d, 0x61, 0x74, 0x6f, 0x30, 0x32, 0x31,
	0x31, 0x2f, 0x62, 0x72, 0x61, 0x63, 0x68, 0x69, 0x6f, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0xa2, 0x02, 0x03, 0x57, 0x50,
	0x47, 0xaa, 0x02, 0x14, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x47, 0x6d, 0xca, 0x02, 0x14, 0x57, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x5c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5c, 0x47, 0x6d, 0xe2,
	0x02, 0x20, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x5c, 0x47, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x16, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x3a,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x3a, 0x3a, 0x47, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_websocket_payload_gm_proto_rawDescOnce sync.Once
	file_websocket_payload_gm_proto_rawDescData []byte
)

func file_websocket_payload_gm_proto_rawDescGZIP() []byte {
	file_websocket_payload_gm_proto_rawDescOnce.Do(func() {
		file_websocket_payload_gm_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_websocket_payload_gm_proto_rawDesc), len(file_websocket_payload_gm_proto_rawDesc)))
	})
	return file_websocket_payload_gm_proto_rawDescData
}

var file_websocket_payload_gm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_websocket_payload_gm_proto_goTypes = []any{
	(*CreateRoomEventPayload)(nil),       // 0: websocket.payload.gm.CreateRoomEventPayload
	(*EnterRoomEventPayload)(nil),        // 1: websocket.payload.gm.EnterRoomEventPayload
	(*MatchingCompleteEventPayload)(nil), // 2: websocket.payload.gm.MatchingCompleteEventPayload
	(*DecideOrderEventPayload)(nil),      // 3: websocket.payload.gm.DecideOrderEventPayload
}
var file_websocket_payload_gm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_websocket_payload_gm_proto_init() }
func file_websocket_payload_gm_proto_init() {
	if File_websocket_payload_gm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_websocket_payload_gm_proto_rawDesc), len(file_websocket_payload_gm_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_websocket_payload_gm_proto_goTypes,
		DependencyIndexes: file_websocket_payload_gm_proto_depIdxs,
		MessageInfos:      file_websocket_payload_gm_proto_msgTypes,
	}.Build()
	File_websocket_payload_gm_proto = out.File
	file_websocket_payload_gm_proto_goTypes = nil
	file_websocket_payload_gm_proto_depIdxs = nil
}
