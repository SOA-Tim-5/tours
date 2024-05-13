// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc2
// source: preference/preference.proto

package preference

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

type PreferenceCreateDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	DifficultyLevel int32    `protobuf:"varint,2,opt,name=DifficultyLevel,proto3" json:"DifficultyLevel,omitempty"`
	WalkingRating   int32    `protobuf:"varint,3,opt,name=WalkingRating,proto3" json:"WalkingRating,omitempty"`
	CyclingRating   int32    `protobuf:"varint,4,opt,name=CyclingRating,proto3" json:"CyclingRating,omitempty"`
	CarRating       int32    `protobuf:"varint,5,opt,name=CarRating,proto3" json:"CarRating,omitempty"`
	BoatRating      int32    `protobuf:"varint,6,opt,name=BoatRating,proto3" json:"BoatRating,omitempty"`
	SelectedTags    []string `protobuf:"bytes,7,rep,name=SelectedTags,proto3" json:"SelectedTags,omitempty"`
}

func (x *PreferenceCreateDto) Reset() {
	*x = PreferenceCreateDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preference_preference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreferenceCreateDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreferenceCreateDto) ProtoMessage() {}

func (x *PreferenceCreateDto) ProtoReflect() protoreflect.Message {
	mi := &file_preference_preference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreferenceCreateDto.ProtoReflect.Descriptor instead.
func (*PreferenceCreateDto) Descriptor() ([]byte, []int) {
	return file_preference_preference_proto_rawDescGZIP(), []int{0}
}

func (x *PreferenceCreateDto) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PreferenceCreateDto) GetDifficultyLevel() int32 {
	if x != nil {
		return x.DifficultyLevel
	}
	return 0
}

func (x *PreferenceCreateDto) GetWalkingRating() int32 {
	if x != nil {
		return x.WalkingRating
	}
	return 0
}

func (x *PreferenceCreateDto) GetCyclingRating() int32 {
	if x != nil {
		return x.CyclingRating
	}
	return 0
}

func (x *PreferenceCreateDto) GetCarRating() int32 {
	if x != nil {
		return x.CarRating
	}
	return 0
}

func (x *PreferenceCreateDto) GetBoatRating() int32 {
	if x != nil {
		return x.BoatRating
	}
	return 0
}

func (x *PreferenceCreateDto) GetSelectedTags() []string {
	if x != nil {
		return x.SelectedTags
	}
	return nil
}

type PreferenceResponseDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId          int64    `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	DifficultyLevel int32    `protobuf:"varint,3,opt,name=DifficultyLevel,proto3" json:"DifficultyLevel,omitempty"`
	WalkingRating   int32    `protobuf:"varint,4,opt,name=WalkingRating,proto3" json:"WalkingRating,omitempty"`
	CyclingRating   int32    `protobuf:"varint,5,opt,name=CyclingRating,proto3" json:"CyclingRating,omitempty"`
	CarRating       int32    `protobuf:"varint,6,opt,name=CarRating,proto3" json:"CarRating,omitempty"`
	BoatRating      int32    `protobuf:"varint,7,opt,name=BoatRating,proto3" json:"BoatRating,omitempty"`
	SelectedTags    []string `protobuf:"bytes,8,rep,name=SelectedTags,proto3" json:"SelectedTags,omitempty"`
}

func (x *PreferenceResponseDto) Reset() {
	*x = PreferenceResponseDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preference_preference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreferenceResponseDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreferenceResponseDto) ProtoMessage() {}

func (x *PreferenceResponseDto) ProtoReflect() protoreflect.Message {
	mi := &file_preference_preference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreferenceResponseDto.ProtoReflect.Descriptor instead.
func (*PreferenceResponseDto) Descriptor() ([]byte, []int) {
	return file_preference_preference_proto_rawDescGZIP(), []int{1}
}

func (x *PreferenceResponseDto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PreferenceResponseDto) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PreferenceResponseDto) GetDifficultyLevel() int32 {
	if x != nil {
		return x.DifficultyLevel
	}
	return 0
}

func (x *PreferenceResponseDto) GetWalkingRating() int32 {
	if x != nil {
		return x.WalkingRating
	}
	return 0
}

func (x *PreferenceResponseDto) GetCyclingRating() int32 {
	if x != nil {
		return x.CyclingRating
	}
	return 0
}

func (x *PreferenceResponseDto) GetCarRating() int32 {
	if x != nil {
		return x.CarRating
	}
	return 0
}

func (x *PreferenceResponseDto) GetBoatRating() int32 {
	if x != nil {
		return x.BoatRating
	}
	return 0
}

func (x *PreferenceResponseDto) GetSelectedTags() []string {
	if x != nil {
		return x.SelectedTags
	}
	return nil
}

type PreferenceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreferenceResponses []*PreferenceResponseDto `protobuf:"bytes,1,rep,name=preferenceResponses,proto3" json:"preferenceResponses,omitempty"`
}

func (x *PreferenceListResponse) Reset() {
	*x = PreferenceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preference_preference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreferenceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreferenceListResponse) ProtoMessage() {}

func (x *PreferenceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_preference_preference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreferenceListResponse.ProtoReflect.Descriptor instead.
func (*PreferenceListResponse) Descriptor() ([]byte, []int) {
	return file_preference_preference_proto_rawDescGZIP(), []int{2}
}

func (x *PreferenceListResponse) GetPreferenceResponses() []*PreferenceResponseDto {
	if x != nil {
		return x.PreferenceResponses
	}
	return nil
}

type GetPreferenceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPreferenceParams) Reset() {
	*x = GetPreferenceParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preference_preference_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPreferenceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPreferenceParams) ProtoMessage() {}

func (x *GetPreferenceParams) ProtoReflect() protoreflect.Message {
	mi := &file_preference_preference_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPreferenceParams.ProtoReflect.Descriptor instead.
func (*GetPreferenceParams) Descriptor() ([]byte, []int) {
	return file_preference_preference_proto_rawDescGZIP(), []int{3}
}

func (x *GetPreferenceParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_preference_preference_proto protoreflect.FileDescriptor

var file_preference_preference_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02,
	0x0a, 0x13, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a,
	0x0d, 0x43, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x43, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x61, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x6f, 0x61, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x42, 0x6f, 0x61, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x61, 0x67,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x54, 0x61, 0x67, 0x73, 0x22, 0x97, 0x02, 0x0a, 0x15, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x79, 0x63, 0x6c, 0x69,
	0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x43, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x61, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x43, 0x61, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x42,
	0x6f, 0x61, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x42, 0x6f, 0x61, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x61, 0x67, 0x73, 0x22,
	0x62, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x13,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8f, 0x01, 0x0a, 0x11, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x50, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x74, 0x6f,
	0x1a, 0x16, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_preference_preference_proto_rawDescOnce sync.Once
	file_preference_preference_proto_rawDescData = file_preference_preference_proto_rawDesc
)

func file_preference_preference_proto_rawDescGZIP() []byte {
	file_preference_preference_proto_rawDescOnce.Do(func() {
		file_preference_preference_proto_rawDescData = protoimpl.X.CompressGZIP(file_preference_preference_proto_rawDescData)
	})
	return file_preference_preference_proto_rawDescData
}

var file_preference_preference_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_preference_preference_proto_goTypes = []interface{}{
	(*PreferenceCreateDto)(nil),    // 0: PreferenceCreateDto
	(*PreferenceResponseDto)(nil),  // 1: PreferenceResponseDto
	(*PreferenceListResponse)(nil), // 2: PreferenceListResponse
	(*GetPreferenceParams)(nil),    // 3: GetPreferenceParams
}
var file_preference_preference_proto_depIdxs = []int32{
	1, // 0: PreferenceListResponse.preferenceResponses:type_name -> PreferenceResponseDto
	0, // 1: PreferenceService.Create:input_type -> PreferenceCreateDto
	3, // 2: PreferenceService.GetByTouristId:input_type -> GetPreferenceParams
	1, // 3: PreferenceService.Create:output_type -> PreferenceResponseDto
	1, // 4: PreferenceService.GetByTouristId:output_type -> PreferenceResponseDto
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_preference_preference_proto_init() }
func file_preference_preference_proto_init() {
	if File_preference_preference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_preference_preference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreferenceCreateDto); i {
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
		file_preference_preference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreferenceResponseDto); i {
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
		file_preference_preference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreferenceListResponse); i {
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
		file_preference_preference_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPreferenceParams); i {
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
			RawDescriptor: file_preference_preference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_preference_preference_proto_goTypes,
		DependencyIndexes: file_preference_preference_proto_depIdxs,
		MessageInfos:      file_preference_preference_proto_msgTypes,
	}.Build()
	File_preference_preference_proto = out.File
	file_preference_preference_proto_rawDesc = nil
	file_preference_preference_proto_goTypes = nil
	file_preference_preference_proto_depIdxs = nil
}
