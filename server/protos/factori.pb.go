// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.6.1
// source: factori.proto

package tutorial

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Rectangle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartX int32 `protobuf:"varint,1,opt,name=startX,proto3" json:"startX,omitempty"`
	StartY int32 `protobuf:"varint,2,opt,name=startY,proto3" json:"startY,omitempty"`
	Width  int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Rectangle) Reset() {
	*x = Rectangle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rectangle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rectangle) ProtoMessage() {}

func (x *Rectangle) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rectangle.ProtoReflect.Descriptor instead.
func (*Rectangle) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{0}
}

func (x *Rectangle) GetStartX() int32 {
	if x != nil {
		return x.StartX
	}
	return 0
}

func (x *Rectangle) GetStartY() int32 {
	if x != nil {
		return x.StartY
	}
	return 0
}

func (x *Rectangle) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Rectangle) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Interaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X               int32  `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y               int32  `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	InteractionChar string `protobuf:"bytes,3,opt,name=interactionChar,proto3" json:"interactionChar,omitempty"`
}

func (x *Interaction) Reset() {
	*x = Interaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interaction) ProtoMessage() {}

func (x *Interaction) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interaction.ProtoReflect.Descriptor instead.
func (*Interaction) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{1}
}

func (x *Interaction) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Interaction) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Interaction) GetInteractionChar() string {
	if x != nil {
		return x.InteractionChar
	}
	return ""
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int32 `protobuf:"varint,1,opt,name=typeId,proto3" json:"typeId,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{2}
}

func (x *Entity) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type Tile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities   []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	TileTypeId int32     `protobuf:"varint,2,opt,name=tileTypeId,proto3" json:"tileTypeId,omitempty"`
}

func (x *Tile) Reset() {
	*x = Tile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tile) ProtoMessage() {}

func (x *Tile) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tile.ProtoReflect.Descriptor instead.
func (*Tile) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{3}
}

func (x *Tile) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *Tile) GetTileTypeId() int32 {
	if x != nil {
		return x.TileTypeId
	}
	return 0
}

type ViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewOf *Rectangle `protobuf:"bytes,1,opt,name=viewOf,proto3" json:"viewOf,omitempty"`
	Tiles  []*Tile    `protobuf:"bytes,2,rep,name=tiles,proto3" json:"tiles,omitempty"`
}

func (x *ViewResponse) Reset() {
	*x = ViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewResponse) ProtoMessage() {}

func (x *ViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewResponse.ProtoReflect.Descriptor instead.
func (*ViewResponse) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{4}
}

func (x *ViewResponse) GetViewOf() *Rectangle {
	if x != nil {
		return x.ViewOf
	}
	return nil
}

func (x *ViewResponse) GetTiles() []*Tile {
	if x != nil {
		return x.Tiles
	}
	return nil
}

type ScreenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubViews []*ViewResponse `protobuf:"bytes,1,rep,name=subViews,proto3" json:"subViews,omitempty"`
}

func (x *ScreenResponse) Reset() {
	*x = ScreenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenResponse) ProtoMessage() {}

func (x *ScreenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenResponse.ProtoReflect.Descriptor instead.
func (*ScreenResponse) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{5}
}

func (x *ScreenResponse) GetSubViews() []*ViewResponse {
	if x != nil {
		return x.SubViews
	}
	return nil
}

type Coord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coord) Reset() {
	*x = Coord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coord) ProtoMessage() {}

func (x *Coord) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coord.ProtoReflect.Descriptor instead.
func (*Coord) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{6}
}

func (x *Coord) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coord) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type ViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullView *Rectangle `protobuf:"bytes,1,opt,name=fullView,proto3" json:"fullView,omitempty"`
	OldView  *Rectangle `protobuf:"bytes,2,opt,name=oldView,proto3" json:"oldView,omitempty"`
}

func (x *ViewRequest) Reset() {
	*x = ViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewRequest) ProtoMessage() {}

func (x *ViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewRequest.ProtoReflect.Descriptor instead.
func (*ViewRequest) Descriptor() ([]byte, []int) {
	return file_factori_proto_rawDescGZIP(), []int{7}
}

func (x *ViewRequest) GetFullView() *Rectangle {
	if x != nil {
		return x.FullView
	}
	return nil
}

func (x *ViewRequest) GetOldView() *Rectangle {
	if x != nil {
		return x.OldView
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factori_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_factori_proto_msgTypes[8]
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
	return file_factori_proto_rawDescGZIP(), []int{8}
}

var File_factori_proto protoreflect.FileDescriptor

var file_factori_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x69, 0x0a, 0x09, 0x52, 0x65, 0x63,
	0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x58,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x58, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x53, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12,
	0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x72, 0x22, 0x20, 0x0a, 0x06, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x04, 0x54,
	0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x22, 0x61, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x76, 0x69, 0x65, 0x77, 0x4f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63,
	0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x4f, 0x66, 0x12, 0x24,
	0x0a, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x74,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x0e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x56, 0x69, 0x65,
	0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x56, 0x69, 0x65, 0x77, 0x73, 0x22, 0x23, 0x0a, 0x05, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22,
	0x6d, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x74,
	0x61, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x2d, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x56, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x74,
	0x61, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x56, 0x69, 0x65, 0x77, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xcc, 0x01, 0x0a, 0x0e, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61,
	0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c,
	0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x13, 0x2e,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x1a, 0x18, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x11,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x15, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_factori_proto_rawDescOnce sync.Once
	file_factori_proto_rawDescData = file_factori_proto_rawDesc
)

func file_factori_proto_rawDescGZIP() []byte {
	file_factori_proto_rawDescOnce.Do(func() {
		file_factori_proto_rawDescData = protoimpl.X.CompressGZIP(file_factori_proto_rawDescData)
	})
	return file_factori_proto_rawDescData
}

var file_factori_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_factori_proto_goTypes = []interface{}{
	(*Rectangle)(nil),      // 0: tutorial.Rectangle
	(*Interaction)(nil),    // 1: tutorial.Interaction
	(*Entity)(nil),         // 2: tutorial.Entity
	(*Tile)(nil),           // 3: tutorial.Tile
	(*ViewResponse)(nil),   // 4: tutorial.ViewResponse
	(*ScreenResponse)(nil), // 5: tutorial.ScreenResponse
	(*Coord)(nil),          // 6: tutorial.Coord
	(*ViewRequest)(nil),    // 7: tutorial.ViewRequest
	(*Empty)(nil),          // 8: tutorial.Empty
}
var file_factori_proto_depIdxs = []int32{
	2, // 0: tutorial.Tile.entities:type_name -> tutorial.Entity
	0, // 1: tutorial.ViewResponse.viewOf:type_name -> tutorial.Rectangle
	3, // 2: tutorial.ViewResponse.tiles:type_name -> tutorial.Tile
	4, // 3: tutorial.ScreenResponse.subViews:type_name -> tutorial.ViewResponse
	0, // 4: tutorial.ViewRequest.fullView:type_name -> tutorial.Rectangle
	0, // 5: tutorial.ViewRequest.oldView:type_name -> tutorial.Rectangle
	1, // 6: tutorial.FactoryService.Interact:input_type -> tutorial.Interaction
	0, // 7: tutorial.FactoryService.RequestView:input_type -> tutorial.Rectangle
	7, // 8: tutorial.FactoryService.RequestViewStream:input_type -> tutorial.ViewRequest
	8, // 9: tutorial.FactoryService.Interact:output_type -> tutorial.Empty
	5, // 10: tutorial.FactoryService.RequestView:output_type -> tutorial.ScreenResponse
	5, // 11: tutorial.FactoryService.RequestViewStream:output_type -> tutorial.ScreenResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_factori_proto_init() }
func file_factori_proto_init() {
	if File_factori_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_factori_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rectangle); i {
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
		file_factori_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interaction); i {
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
		file_factori_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_factori_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tile); i {
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
		file_factori_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewResponse); i {
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
		file_factori_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenResponse); i {
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
		file_factori_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coord); i {
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
		file_factori_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewRequest); i {
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
		file_factori_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_factori_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_factori_proto_goTypes,
		DependencyIndexes: file_factori_proto_depIdxs,
		MessageInfos:      file_factori_proto_msgTypes,
	}.Build()
	File_factori_proto = out.File
	file_factori_proto_rawDesc = nil
	file_factori_proto_goTypes = nil
	file_factori_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FactoryServiceClient is the client API for FactoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FactoryServiceClient interface {
	Interact(ctx context.Context, in *Interaction, opts ...grpc.CallOption) (*Empty, error)
	RequestView(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (*ScreenResponse, error)
	RequestViewStream(ctx context.Context, opts ...grpc.CallOption) (FactoryService_RequestViewStreamClient, error)
}

type factoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFactoryServiceClient(cc grpc.ClientConnInterface) FactoryServiceClient {
	return &factoryServiceClient{cc}
}

func (c *factoryServiceClient) Interact(ctx context.Context, in *Interaction, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/tutorial.FactoryService/Interact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factoryServiceClient) RequestView(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (*ScreenResponse, error) {
	out := new(ScreenResponse)
	err := c.cc.Invoke(ctx, "/tutorial.FactoryService/RequestView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factoryServiceClient) RequestViewStream(ctx context.Context, opts ...grpc.CallOption) (FactoryService_RequestViewStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FactoryService_serviceDesc.Streams[0], "/tutorial.FactoryService/RequestViewStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &factoryServiceRequestViewStreamClient{stream}
	return x, nil
}

type FactoryService_RequestViewStreamClient interface {
	Send(*ViewRequest) error
	Recv() (*ScreenResponse, error)
	grpc.ClientStream
}

type factoryServiceRequestViewStreamClient struct {
	grpc.ClientStream
}

func (x *factoryServiceRequestViewStreamClient) Send(m *ViewRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *factoryServiceRequestViewStreamClient) Recv() (*ScreenResponse, error) {
	m := new(ScreenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FactoryServiceServer is the server API for FactoryService service.
type FactoryServiceServer interface {
	Interact(context.Context, *Interaction) (*Empty, error)
	RequestView(context.Context, *Rectangle) (*ScreenResponse, error)
	RequestViewStream(FactoryService_RequestViewStreamServer) error
}

// UnimplementedFactoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFactoryServiceServer struct {
}

func (*UnimplementedFactoryServiceServer) Interact(context.Context, *Interaction) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interact not implemented")
}
func (*UnimplementedFactoryServiceServer) RequestView(context.Context, *Rectangle) (*ScreenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestView not implemented")
}
func (*UnimplementedFactoryServiceServer) RequestViewStream(FactoryService_RequestViewStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestViewStream not implemented")
}

func RegisterFactoryServiceServer(s *grpc.Server, srv FactoryServiceServer) {
	s.RegisterService(&_FactoryService_serviceDesc, srv)
}

func _FactoryService_Interact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Interaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoryServiceServer).Interact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tutorial.FactoryService/Interact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoryServiceServer).Interact(ctx, req.(*Interaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FactoryService_RequestView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rectangle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoryServiceServer).RequestView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tutorial.FactoryService/RequestView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoryServiceServer).RequestView(ctx, req.(*Rectangle))
	}
	return interceptor(ctx, in, info, handler)
}

func _FactoryService_RequestViewStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FactoryServiceServer).RequestViewStream(&factoryServiceRequestViewStreamServer{stream})
}

type FactoryService_RequestViewStreamServer interface {
	Send(*ScreenResponse) error
	Recv() (*ViewRequest, error)
	grpc.ServerStream
}

type factoryServiceRequestViewStreamServer struct {
	grpc.ServerStream
}

func (x *factoryServiceRequestViewStreamServer) Send(m *ScreenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *factoryServiceRequestViewStreamServer) Recv() (*ViewRequest, error) {
	m := new(ViewRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FactoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tutorial.FactoryService",
	HandlerType: (*FactoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Interact",
			Handler:    _FactoryService_Interact_Handler,
		},
		{
			MethodName: "RequestView",
			Handler:    _FactoryService_RequestView_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestViewStream",
			Handler:       _FactoryService_RequestViewStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "factori.proto",
}
