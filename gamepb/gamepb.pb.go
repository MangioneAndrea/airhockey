// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: gamepb/gamepb.proto

package gamepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamepb_gamepb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_gamepb_gamepb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_gamepb_gamepb_proto_rawDescGZIP(), []int{0}
}

type Vector2D struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Vector2D) Reset() {
	*x = Vector2D{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamepb_gamepb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector2D) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector2D) ProtoMessage() {}

func (x *Vector2D) ProtoReflect() protoreflect.Message {
	mi := &file_gamepb_gamepb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector2D.ProtoReflect.Descriptor instead.
func (*Vector2D) Descriptor() ([]byte, []int) {
	return file_gamepb_gamepb_proto_rawDescGZIP(), []int{1}
}

func (x *Vector2D) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector2D) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerHash string `protobuf:"bytes,1,opt,name=playerHash,proto3" json:"playerHash,omitempty"`
	GameHash   string `protobuf:"bytes,2,opt,name=gameHash,proto3" json:"gameHash,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamepb_gamepb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_gamepb_gamepb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_gamepb_gamepb_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetPlayerHash() string {
	if x != nil {
		return x.PlayerHash
	}
	return ""
}

func (x *Token) GetGameHash() string {
	if x != nil {
		return x.GameHash
	}
	return ""
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameHash string `protobuf:"bytes,1,opt,name=gameHash,proto3" json:"gameHash,omitempty"`
	Token1   *Token `protobuf:"bytes,2,opt,name=token1,proto3" json:"token1,omitempty"`
	Token2   *Token `protobuf:"bytes,3,opt,name=token2,proto3" json:"token2,omitempty"`
	P1Ready  bool   `protobuf:"varint,4,opt,name=p1Ready,proto3" json:"p1Ready,omitempty"`
	P2Ready  bool   `protobuf:"varint,5,opt,name=p2Ready,proto3" json:"p2Ready,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamepb_gamepb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_gamepb_gamepb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_gamepb_gamepb_proto_rawDescGZIP(), []int{3}
}

func (x *Game) GetGameHash() string {
	if x != nil {
		return x.GameHash
	}
	return ""
}

func (x *Game) GetToken1() *Token {
	if x != nil {
		return x.Token1
	}
	return nil
}

func (x *Game) GetToken2() *Token {
	if x != nil {
		return x.Token2
	}
	return nil
}

func (x *Game) GetP1Ready() bool {
	if x != nil {
		return x.P1Ready
	}
	return false
}

func (x *Game) GetP2Ready() bool {
	if x != nil {
		return x.P2Ready
	}
	return false
}

type UserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector *Vector2D `protobuf:"bytes,1,opt,name=vector,proto3" json:"vector,omitempty"`
	Token  *Token    `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserInput) Reset() {
	*x = UserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamepb_gamepb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInput) ProtoMessage() {}

func (x *UserInput) ProtoReflect() protoreflect.Message {
	mi := &file_gamepb_gamepb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInput.ProtoReflect.Descriptor instead.
func (*UserInput) Descriptor() ([]byte, []int) {
	return file_gamepb_gamepb_proto_rawDescGZIP(), []int{4}
}

func (x *UserInput) GetVector() *Vector2D {
	if x != nil {
		return x.Vector
	}
	return nil
}

func (x *UserInput) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type GameStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player1       *Vector2D `protobuf:"bytes,1,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2       *Vector2D `protobuf:"bytes,2,opt,name=player2,proto3" json:"player2,omitempty"`
	Disk          *Vector2D `protobuf:"bytes,3,opt,name=disk,proto3" json:"disk,omitempty"`
	DiskDirection *Vector2D `protobuf:"bytes,4,opt,name=diskDirection,proto3" json:"diskDirection,omitempty"`
}

func (x *GameStatus) Reset() {
	*x = GameStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamepb_gamepb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatus) ProtoMessage() {}

func (x *GameStatus) ProtoReflect() protoreflect.Message {
	mi := &file_gamepb_gamepb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatus.ProtoReflect.Descriptor instead.
func (*GameStatus) Descriptor() ([]byte, []int) {
	return file_gamepb_gamepb_proto_rawDescGZIP(), []int{5}
}

func (x *GameStatus) GetPlayer1() *Vector2D {
	if x != nil {
		return x.Player1
	}
	return nil
}

func (x *GameStatus) GetPlayer2() *Vector2D {
	if x != nil {
		return x.Player2
	}
	return nil
}

func (x *GameStatus) GetDisk() *Vector2D {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *GameStatus) GetDiskDirection() *Vector2D {
	if x != nil {
		return x.DiskDirection
	}
	return nil
}

var File_gamepb_gamepb_proto protoreflect.FileDescriptor

var file_gamepb_gamepb_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x22, 0x06, 0x0a,
	0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x08, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32,
	0x44, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x43, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x22, 0xa4, 0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x61, 0x6d, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x61, 0x6d, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x31, 0x12, 0x25,
	0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x31, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x31, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x32, 0x52, 0x65, 0x61, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x70, 0x32, 0x52, 0x65, 0x61, 0x64, 0x79, 0x22, 0x5a, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x44, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x23, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x44, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31,
	0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x32, 0x44, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x12, 0x24, 0x0a, 0x04,
	0x64, 0x69, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x44, 0x52, 0x04, 0x64, 0x69,
	0x73, 0x6b, 0x12, 0x36, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x6b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x44, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x6b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x7d, 0x0a, 0x0f, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0c, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gamepb_gamepb_proto_rawDescOnce sync.Once
	file_gamepb_gamepb_proto_rawDescData = file_gamepb_gamepb_proto_rawDesc
)

func file_gamepb_gamepb_proto_rawDescGZIP() []byte {
	file_gamepb_gamepb_proto_rawDescOnce.Do(func() {
		file_gamepb_gamepb_proto_rawDescData = protoimpl.X.CompressGZIP(file_gamepb_gamepb_proto_rawDescData)
	})
	return file_gamepb_gamepb_proto_rawDescData
}

var file_gamepb_gamepb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gamepb_gamepb_proto_goTypes = []interface{}{
	(*Void)(nil),       // 0: gamepb.Void
	(*Vector2D)(nil),   // 1: gamepb.Vector2D
	(*Token)(nil),      // 2: gamepb.Token
	(*Game)(nil),       // 3: gamepb.Game
	(*UserInput)(nil),  // 4: gamepb.UserInput
	(*GameStatus)(nil), // 5: gamepb.GameStatus
}
var file_gamepb_gamepb_proto_depIdxs = []int32{
	2,  // 0: gamepb.Game.token1:type_name -> gamepb.Token
	2,  // 1: gamepb.Game.token2:type_name -> gamepb.Token
	1,  // 2: gamepb.UserInput.vector:type_name -> gamepb.Vector2D
	2,  // 3: gamepb.UserInput.token:type_name -> gamepb.Token
	1,  // 4: gamepb.GameStatus.player1:type_name -> gamepb.Vector2D
	1,  // 5: gamepb.GameStatus.player2:type_name -> gamepb.Vector2D
	1,  // 6: gamepb.GameStatus.disk:type_name -> gamepb.Vector2D
	1,  // 7: gamepb.GameStatus.diskDirection:type_name -> gamepb.Vector2D
	0,  // 8: gamepb.PositionService.RequestGame:input_type -> gamepb.Void
	4,  // 9: gamepb.PositionService.UpdateStatus:input_type -> gamepb.UserInput
	3,  // 10: gamepb.PositionService.RequestGame:output_type -> gamepb.Game
	5,  // 11: gamepb.PositionService.UpdateStatus:output_type -> gamepb.GameStatus
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_gamepb_gamepb_proto_init() }
func file_gamepb_gamepb_proto_init() {
	if File_gamepb_gamepb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gamepb_gamepb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_gamepb_gamepb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector2D); i {
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
		file_gamepb_gamepb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_gamepb_gamepb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_gamepb_gamepb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInput); i {
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
		file_gamepb_gamepb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStatus); i {
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
			RawDescriptor: file_gamepb_gamepb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gamepb_gamepb_proto_goTypes,
		DependencyIndexes: file_gamepb_gamepb_proto_depIdxs,
		MessageInfos:      file_gamepb_gamepb_proto_msgTypes,
	}.Build()
	File_gamepb_gamepb_proto = out.File
	file_gamepb_gamepb_proto_rawDesc = nil
	file_gamepb_gamepb_proto_goTypes = nil
	file_gamepb_gamepb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PositionServiceClient interface {
	RequestGame(ctx context.Context, in *Void, opts ...grpc.CallOption) (PositionService_RequestGameClient, error)
	UpdateStatus(ctx context.Context, opts ...grpc.CallOption) (PositionService_UpdateStatusClient, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) RequestGame(ctx context.Context, in *Void, opts ...grpc.CallOption) (PositionService_RequestGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PositionService_serviceDesc.Streams[0], "/gamepb.PositionService/RequestGame", opts...)
	if err != nil {
		return nil, err
	}
	x := &positionServiceRequestGameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PositionService_RequestGameClient interface {
	Recv() (*Game, error)
	grpc.ClientStream
}

type positionServiceRequestGameClient struct {
	grpc.ClientStream
}

func (x *positionServiceRequestGameClient) Recv() (*Game, error) {
	m := new(Game)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *positionServiceClient) UpdateStatus(ctx context.Context, opts ...grpc.CallOption) (PositionService_UpdateStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PositionService_serviceDesc.Streams[1], "/gamepb.PositionService/UpdateStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &positionServiceUpdateStatusClient{stream}
	return x, nil
}

type PositionService_UpdateStatusClient interface {
	Send(*UserInput) error
	Recv() (*GameStatus, error)
	grpc.ClientStream
}

type positionServiceUpdateStatusClient struct {
	grpc.ClientStream
}

func (x *positionServiceUpdateStatusClient) Send(m *UserInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *positionServiceUpdateStatusClient) Recv() (*GameStatus, error) {
	m := new(GameStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PositionServiceServer is the server API for PositionService service.
type PositionServiceServer interface {
	RequestGame(*Void, PositionService_RequestGameServer) error
	UpdateStatus(PositionService_UpdateStatusServer) error
}

// UnimplementedPositionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (*UnimplementedPositionServiceServer) RequestGame(*Void, PositionService_RequestGameServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestGame not implemented")
}
func (*UnimplementedPositionServiceServer) UpdateStatus(PositionService_UpdateStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}

func RegisterPositionServiceServer(s *grpc.Server, srv PositionServiceServer) {
	s.RegisterService(&_PositionService_serviceDesc, srv)
}

func _PositionService_RequestGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PositionServiceServer).RequestGame(m, &positionServiceRequestGameServer{stream})
}

type PositionService_RequestGameServer interface {
	Send(*Game) error
	grpc.ServerStream
}

type positionServiceRequestGameServer struct {
	grpc.ServerStream
}

func (x *positionServiceRequestGameServer) Send(m *Game) error {
	return x.ServerStream.SendMsg(m)
}

func _PositionService_UpdateStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PositionServiceServer).UpdateStatus(&positionServiceUpdateStatusServer{stream})
}

type PositionService_UpdateStatusServer interface {
	Send(*GameStatus) error
	Recv() (*UserInput, error)
	grpc.ServerStream
}

type positionServiceUpdateStatusServer struct {
	grpc.ServerStream
}

func (x *positionServiceUpdateStatusServer) Send(m *GameStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *positionServiceUpdateStatusServer) Recv() (*UserInput, error) {
	m := new(UserInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PositionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gamepb.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestGame",
			Handler:       _PositionService_RequestGame_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateStatus",
			Handler:       _PositionService_UpdateStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gamepb/gamepb.proto",
}
