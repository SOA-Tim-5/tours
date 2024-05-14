// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: facility/facility.proto

package facility

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

type FacilityCreateDto_FacilityCategory int32

const (
	FacilityCreateDto_Draft     FacilityCreateDto_FacilityCategory = 0
	FacilityCreateDto_Published FacilityCreateDto_FacilityCategory = 1
	FacilityCreateDto_Archived  FacilityCreateDto_FacilityCategory = 2
	FacilityCreateDto_Ready     FacilityCreateDto_FacilityCategory = 3
)

// Enum value maps for FacilityCreateDto_FacilityCategory.
var (
	FacilityCreateDto_FacilityCategory_name = map[int32]string{
		0: "Draft",
		1: "Published",
		2: "Archived",
		3: "Ready",
	}
	FacilityCreateDto_FacilityCategory_value = map[string]int32{
		"Draft":     0,
		"Published": 1,
		"Archived":  2,
		"Ready":     3,
	}
)

func (x FacilityCreateDto_FacilityCategory) Enum() *FacilityCreateDto_FacilityCategory {
	p := new(FacilityCreateDto_FacilityCategory)
	*p = x
	return p
}

func (x FacilityCreateDto_FacilityCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FacilityCreateDto_FacilityCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_facility_facility_proto_enumTypes[0].Descriptor()
}

func (FacilityCreateDto_FacilityCategory) Type() protoreflect.EnumType {
	return &file_facility_facility_proto_enumTypes[0]
}

func (x FacilityCreateDto_FacilityCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FacilityCreateDto_FacilityCategory.Descriptor instead.
func (FacilityCreateDto_FacilityCategory) EnumDescriptor() ([]byte, []int) {
	return file_facility_facility_proto_rawDescGZIP(), []int{0, 0}
}

type FacilityResponseDto_FacilityCategory int32

const (
	FacilityResponseDto_Restaurant     FacilityResponseDto_FacilityCategory = 0
	FacilityResponseDto_ParkingLot     FacilityResponseDto_FacilityCategory = 1
	FacilityResponseDto_Toilet         FacilityResponseDto_FacilityCategory = 2
	FacilityResponseDto_Hospital       FacilityResponseDto_FacilityCategory = 3
	FacilityResponseDto_Cafe           FacilityResponseDto_FacilityCategory = 4
	FacilityResponseDto_Pharmacy       FacilityResponseDto_FacilityCategory = 5
	FacilityResponseDto_ExchangeOffice FacilityResponseDto_FacilityCategory = 6
	FacilityResponseDto_BusStop        FacilityResponseDto_FacilityCategory = 7
	FacilityResponseDto_Shop           FacilityResponseDto_FacilityCategory = 8
	FacilityResponseDto_Other          FacilityResponseDto_FacilityCategory = 9
)

// Enum value maps for FacilityResponseDto_FacilityCategory.
var (
	FacilityResponseDto_FacilityCategory_name = map[int32]string{
		0: "Restaurant",
		1: "ParkingLot",
		2: "Toilet",
		3: "Hospital",
		4: "Cafe",
		5: "Pharmacy",
		6: "ExchangeOffice",
		7: "BusStop",
		8: "Shop",
		9: "Other",
	}
	FacilityResponseDto_FacilityCategory_value = map[string]int32{
		"Restaurant":     0,
		"ParkingLot":     1,
		"Toilet":         2,
		"Hospital":       3,
		"Cafe":           4,
		"Pharmacy":       5,
		"ExchangeOffice": 6,
		"BusStop":        7,
		"Shop":           8,
		"Other":          9,
	}
)

func (x FacilityResponseDto_FacilityCategory) Enum() *FacilityResponseDto_FacilityCategory {
	p := new(FacilityResponseDto_FacilityCategory)
	*p = x
	return p
}

func (x FacilityResponseDto_FacilityCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FacilityResponseDto_FacilityCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_facility_facility_proto_enumTypes[1].Descriptor()
}

func (FacilityResponseDto_FacilityCategory) Type() protoreflect.EnumType {
	return &file_facility_facility_proto_enumTypes[1]
}

func (x FacilityResponseDto_FacilityCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FacilityResponseDto_FacilityCategory.Descriptor instead.
func (FacilityResponseDto_FacilityCategory) EnumDescriptor() ([]byte, []int) {
	return file_facility_facility_proto_rawDescGZIP(), []int{1, 0}
}

type FacilityCreateDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                             `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string                             `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	ImagePath   string                             `protobuf:"bytes,3,opt,name=ImagePath,proto3" json:"ImagePath,omitempty"`
	AuthorId    int64                              `protobuf:"varint,4,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	Category    FacilityCreateDto_FacilityCategory `protobuf:"varint,5,opt,name=Category,proto3,enum=FacilityCreateDto_FacilityCategory" json:"Category,omitempty"`
	Longitude   float64                            `protobuf:"fixed64,6,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Latitude    float64                            `protobuf:"fixed64,7,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
}

func (x *FacilityCreateDto) Reset() {
	*x = FacilityCreateDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facility_facility_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FacilityCreateDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FacilityCreateDto) ProtoMessage() {}

func (x *FacilityCreateDto) ProtoReflect() protoreflect.Message {
	mi := &file_facility_facility_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FacilityCreateDto.ProtoReflect.Descriptor instead.
func (*FacilityCreateDto) Descriptor() ([]byte, []int) {
	return file_facility_facility_proto_rawDescGZIP(), []int{0}
}

func (x *FacilityCreateDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FacilityCreateDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FacilityCreateDto) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *FacilityCreateDto) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *FacilityCreateDto) GetCategory() FacilityCreateDto_FacilityCategory {
	if x != nil {
		return x.Category
	}
	return FacilityCreateDto_Draft
}

func (x *FacilityCreateDto) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *FacilityCreateDto) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type FacilityResponseDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                                `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name        string                               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string                               `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	ImagePath   string                               `protobuf:"bytes,4,opt,name=ImagePath,proto3" json:"ImagePath,omitempty"`
	AuthorId    int64                                `protobuf:"varint,5,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	Category    FacilityResponseDto_FacilityCategory `protobuf:"varint,6,opt,name=Category,proto3,enum=FacilityResponseDto_FacilityCategory" json:"Category,omitempty"`
	Longitude   float64                              `protobuf:"fixed64,7,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Latitude    float64                              `protobuf:"fixed64,8,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
}

func (x *FacilityResponseDto) Reset() {
	*x = FacilityResponseDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facility_facility_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FacilityResponseDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FacilityResponseDto) ProtoMessage() {}

func (x *FacilityResponseDto) ProtoReflect() protoreflect.Message {
	mi := &file_facility_facility_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FacilityResponseDto.ProtoReflect.Descriptor instead.
func (*FacilityResponseDto) Descriptor() ([]byte, []int) {
	return file_facility_facility_proto_rawDescGZIP(), []int{1}
}

func (x *FacilityResponseDto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FacilityResponseDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FacilityResponseDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FacilityResponseDto) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *FacilityResponseDto) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *FacilityResponseDto) GetCategory() FacilityResponseDto_FacilityCategory {
	if x != nil {
		return x.Category
	}
	return FacilityResponseDto_Restaurant
}

func (x *FacilityResponseDto) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *FacilityResponseDto) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type FacilityListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FacilityResponses []*FacilityResponseDto `protobuf:"bytes,1,rep,name=facilityResponses,proto3" json:"facilityResponses,omitempty"`
}

func (x *FacilityListResponse) Reset() {
	*x = FacilityListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facility_facility_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FacilityListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FacilityListResponse) ProtoMessage() {}

func (x *FacilityListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_facility_facility_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FacilityListResponse.ProtoReflect.Descriptor instead.
func (*FacilityListResponse) Descriptor() ([]byte, []int) {
	return file_facility_facility_proto_rawDescGZIP(), []int{2}
}

func (x *FacilityListResponse) GetFacilityResponses() []*FacilityResponseDto {
	if x != nil {
		return x.FacilityResponses
	}
	return nil
}

type GetFacilityParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFacilityParams) Reset() {
	*x = GetFacilityParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facility_facility_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFacilityParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFacilityParams) ProtoMessage() {}

func (x *GetFacilityParams) ProtoReflect() protoreflect.Message {
	mi := &file_facility_facility_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFacilityParams.ProtoReflect.Descriptor instead.
func (*GetFacilityParams) Descriptor() ([]byte, []int) {
	return file_facility_facility_proto_rawDescGZIP(), []int{3}
}

func (x *GetFacilityParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_facility_facility_proto protoreflect.FileDescriptor

var file_facility_facility_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x66, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x11, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x74, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x3f, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x74, 0x6f, 0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x45, 0x0a, 0x10, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x72, 0x61, 0x66, 0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x64, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10,
	0x03, 0x22, 0xaf, 0x03, 0x0a, 0x13, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74,
	0x6f, 0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x46, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x54,
	0x6f, 0x69, 0x6c, 0x65, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x61, 0x66, 0x65, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x63, 0x79, 0x10, 0x05, 0x12, 0x12, 0x0a,
	0x0e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x10,
	0x06, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x75, 0x73, 0x53, 0x74, 0x6f, 0x70, 0x10, 0x07, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x10, 0x09, 0x22, 0x5a, 0x0a, 0x14, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x11, 0x66,
	0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x11, 0x66, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22,
	0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x8d, 0x01, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x12, 0x2e, 0x46, 0x61, 0x63,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x14,
	0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x63,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_facility_facility_proto_rawDescOnce sync.Once
	file_facility_facility_proto_rawDescData = file_facility_facility_proto_rawDesc
)

func file_facility_facility_proto_rawDescGZIP() []byte {
	file_facility_facility_proto_rawDescOnce.Do(func() {
		file_facility_facility_proto_rawDescData = protoimpl.X.CompressGZIP(file_facility_facility_proto_rawDescData)
	})
	return file_facility_facility_proto_rawDescData
}

var file_facility_facility_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_facility_facility_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_facility_facility_proto_goTypes = []interface{}{
	(FacilityCreateDto_FacilityCategory)(0),   // 0: FacilityCreateDto.FacilityCategory
	(FacilityResponseDto_FacilityCategory)(0), // 1: FacilityResponseDto.FacilityCategory
	(*FacilityCreateDto)(nil),                 // 2: FacilityCreateDto
	(*FacilityResponseDto)(nil),               // 3: FacilityResponseDto
	(*FacilityListResponse)(nil),              // 4: FacilityListResponse
	(*GetFacilityParams)(nil),                 // 5: GetFacilityParams
}
var file_facility_facility_proto_depIdxs = []int32{
	0, // 0: FacilityCreateDto.Category:type_name -> FacilityCreateDto.FacilityCategory
	1, // 1: FacilityResponseDto.Category:type_name -> FacilityResponseDto.FacilityCategory
	3, // 2: FacilityListResponse.facilityResponses:type_name -> FacilityResponseDto
	2, // 3: FacilityService.CreateFacility:input_type -> FacilityCreateDto
	5, // 4: FacilityService.GetByAuthorId:input_type -> GetFacilityParams
	3, // 5: FacilityService.CreateFacility:output_type -> FacilityResponseDto
	4, // 6: FacilityService.GetByAuthorId:output_type -> FacilityListResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_facility_facility_proto_init() }
func file_facility_facility_proto_init() {
	if File_facility_facility_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_facility_facility_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FacilityCreateDto); i {
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
		file_facility_facility_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FacilityResponseDto); i {
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
		file_facility_facility_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FacilityListResponse); i {
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
		file_facility_facility_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFacilityParams); i {
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
			RawDescriptor: file_facility_facility_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_facility_facility_proto_goTypes,
		DependencyIndexes: file_facility_facility_proto_depIdxs,
		EnumInfos:         file_facility_facility_proto_enumTypes,
		MessageInfos:      file_facility_facility_proto_msgTypes,
	}.Build()
	File_facility_facility_proto = out.File
	file_facility_facility_proto_rawDesc = nil
	file_facility_facility_proto_goTypes = nil
	file_facility_facility_proto_depIdxs = nil
}
