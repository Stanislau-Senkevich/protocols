// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: family/family.proto

package famv1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateFamilyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFamilyRequest) Reset() {
	*x = CreateFamilyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_family_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFamilyRequest) ProtoMessage() {}

func (x *CreateFamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_family_family_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFamilyRequest.ProtoReflect.Descriptor instead.
func (*CreateFamilyRequest) Descriptor() ([]byte, []int) {
	return file_family_family_proto_rawDescGZIP(), []int{0}
}

type CreateFamilyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyId int64 `protobuf:"varint,1,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (x *CreateFamilyResponse) Reset() {
	*x = CreateFamilyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_family_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFamilyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFamilyResponse) ProtoMessage() {}

func (x *CreateFamilyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_family_family_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFamilyResponse.ProtoReflect.Descriptor instead.
func (*CreateFamilyResponse) Descriptor() ([]byte, []int) {
	return file_family_family_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFamilyResponse) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

type LeaveFamilyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyId int64 `protobuf:"varint,1,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (x *LeaveFamilyRequest) Reset() {
	*x = LeaveFamilyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_family_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveFamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveFamilyRequest) ProtoMessage() {}

func (x *LeaveFamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_family_family_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveFamilyRequest.ProtoReflect.Descriptor instead.
func (*LeaveFamilyRequest) Descriptor() ([]byte, []int) {
	return file_family_family_proto_rawDescGZIP(), []int{2}
}

func (x *LeaveFamilyRequest) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

type LeaveFamilyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
}

func (x *LeaveFamilyResponse) Reset() {
	*x = LeaveFamilyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_family_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveFamilyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveFamilyResponse) ProtoMessage() {}

func (x *LeaveFamilyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_family_family_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveFamilyResponse.ProtoReflect.Descriptor instead.
func (*LeaveFamilyResponse) Descriptor() ([]byte, []int) {
	return file_family_family_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveFamilyResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

type GetFamilyInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyId int64 `protobuf:"varint,1,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (x *GetFamilyInfoRequest) Reset() {
	*x = GetFamilyInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_family_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFamilyInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFamilyInfoRequest) ProtoMessage() {}

func (x *GetFamilyInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_family_family_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFamilyInfoRequest.ProtoReflect.Descriptor instead.
func (*GetFamilyInfoRequest) Descriptor() ([]byte, []int) {
	return file_family_family_proto_rawDescGZIP(), []int{4}
}

func (x *GetFamilyInfoRequest) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email        string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber  string               `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Name         string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Surname      string               `protobuf:"bytes,5,opt,name=surname,proto3" json:"surname,omitempty"`
	RegisteredAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=registered_at,json=registeredAt,proto3" json:"registered_at,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_family_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_family_family_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_family_family_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UserInfo) GetRegisteredAt() *timestamp.Timestamp {
	if x != nil {
		return x.RegisteredAt
	}
	return nil
}

type GetFamilyInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*UserInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *GetFamilyInfoResponse) Reset() {
	*x = GetFamilyInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_family_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFamilyInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFamilyInfoResponse) ProtoMessage() {}

func (x *GetFamilyInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_family_family_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFamilyInfoResponse.ProtoReflect.Descriptor instead.
func (*GetFamilyInfoResponse) Descriptor() ([]byte, []int) {
	return file_family_family_proto_rawDescGZIP(), []int{6}
}

func (x *GetFamilyInfoResponse) GetInfo() []*UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_family_family_proto protoreflect.FileDescriptor

var file_family_family_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x12, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x2f, 0x0a,
	0x13, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x22, 0x33,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x49, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0xe9, 0x01, 0x0a, 0x06, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1b, 0x2e, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1a, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1c, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16,
	0x68, 0x61, 0x6b, 0x65, 0x79, 0x6e, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x76, 0x31,
	0x3b, 0x66, 0x61, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_family_family_proto_rawDescOnce sync.Once
	file_family_family_proto_rawDescData = file_family_family_proto_rawDesc
)

func file_family_family_proto_rawDescGZIP() []byte {
	file_family_family_proto_rawDescOnce.Do(func() {
		file_family_family_proto_rawDescData = protoimpl.X.CompressGZIP(file_family_family_proto_rawDescData)
	})
	return file_family_family_proto_rawDescData
}

var file_family_family_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_family_family_proto_goTypes = []interface{}{
	(*CreateFamilyRequest)(nil),   // 0: family.CreateFamilyRequest
	(*CreateFamilyResponse)(nil),  // 1: family.CreateFamilyResponse
	(*LeaveFamilyRequest)(nil),    // 2: family.LeaveFamilyRequest
	(*LeaveFamilyResponse)(nil),   // 3: family.LeaveFamilyResponse
	(*GetFamilyInfoRequest)(nil),  // 4: family.GetFamilyInfoRequest
	(*UserInfo)(nil),              // 5: family.UserInfo
	(*GetFamilyInfoResponse)(nil), // 6: family.GetFamilyInfoResponse
	(*timestamp.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_family_family_proto_depIdxs = []int32{
	7, // 0: family.UserInfo.registered_at:type_name -> google.protobuf.Timestamp
	5, // 1: family.GetFamilyInfoResponse.info:type_name -> family.UserInfo
	0, // 2: family.Family.CreateFamily:input_type -> family.CreateFamilyRequest
	2, // 3: family.Family.LeaveFamily:input_type -> family.LeaveFamilyRequest
	4, // 4: family.Family.GetFamilyInfo:input_type -> family.GetFamilyInfoRequest
	1, // 5: family.Family.CreateFamily:output_type -> family.CreateFamilyResponse
	3, // 6: family.Family.LeaveFamily:output_type -> family.LeaveFamilyResponse
	6, // 7: family.Family.GetFamilyInfo:output_type -> family.GetFamilyInfoResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_family_family_proto_init() }
func file_family_family_proto_init() {
	if File_family_family_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_family_family_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFamilyRequest); i {
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
		file_family_family_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFamilyResponse); i {
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
		file_family_family_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveFamilyRequest); i {
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
		file_family_family_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveFamilyResponse); i {
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
		file_family_family_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFamilyInfoRequest); i {
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
		file_family_family_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_family_family_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFamilyInfoResponse); i {
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
			RawDescriptor: file_family_family_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_family_family_proto_goTypes,
		DependencyIndexes: file_family_family_proto_depIdxs,
		MessageInfos:      file_family_family_proto_msgTypes,
	}.Build()
	File_family_family_proto = out.File
	file_family_family_proto_rawDesc = nil
	file_family_family_proto_goTypes = nil
	file_family_family_proto_depIdxs = nil
}
