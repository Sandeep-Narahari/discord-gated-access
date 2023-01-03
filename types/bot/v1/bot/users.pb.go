// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: v1/bot/users.proto

package bot

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DiscordName    string `protobuf:"bytes,2,opt,name=discord_name,json=discordName,proto3" json:"discord_name,omitempty"`
	DiscordId      string `protobuf:"bytes,3,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
	AccountAddress string `protobuf:"bytes,4,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetDiscordName() string {
	if x != nil {
		return x.DiscordName
	}
	return ""
}

func (x *User) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

func (x *User) GetAccountAddress() string {
	if x != nil {
		return x.AccountAddress
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscordId       string `protobuf:"bytes,1,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
	AccountAddress  string `protobuf:"bytes,2,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	DiscordUserName string `protobuf:"bytes,3,opt,name=discord_user_name,json=discordUserName,proto3" json:"discord_user_name,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

func (x *CreateUserRequest) GetAccountAddress() string {
	if x != nil {
		return x.AccountAddress
	}
	return ""
}

func (x *CreateUserRequest) GetDiscordUserName() string {
	if x != nil {
		return x.DiscordUserName
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{3}
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserByDiscordIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscordId string `protobuf:"bytes,1,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
}

func (x *GetUserByDiscordIDRequest) Reset() {
	*x = GetUserByDiscordIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByDiscordIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByDiscordIDRequest) ProtoMessage() {}

func (x *GetUserByDiscordIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByDiscordIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserByDiscordIDRequest) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserByDiscordIDRequest) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

type GetUserByAccountAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetUserByAccountAddressRequest) Reset() {
	*x = GetUserByAccountAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByAccountAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByAccountAddressRequest) ProtoMessage() {}

func (x *GetUserByAccountAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByAccountAddressRequest.ProtoReflect.Descriptor instead.
func (*GetUserByAccountAddressRequest) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserByAccountAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bot_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bot_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_v1_bot_users_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_v1_bot_users_proto protoreflect.FileDescriptor

var file_v1_bot_users_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x87,
	0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x11,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xa3, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x4d, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x6d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x74, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26,
	0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x75, 0x74, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bot_users_proto_rawDescOnce sync.Once
	file_v1_bot_users_proto_rawDescData = file_v1_bot_users_proto_rawDesc
)

func file_v1_bot_users_proto_rawDescGZIP() []byte {
	file_v1_bot_users_proto_rawDescOnce.Do(func() {
		file_v1_bot_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bot_users_proto_rawDescData)
	})
	return file_v1_bot_users_proto_rawDescData
}

var file_v1_bot_users_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_bot_users_proto_goTypes = []interface{}{
	(*User)(nil),                           // 0: v1.bot.User
	(*CreateUserRequest)(nil),              // 1: v1.bot.CreateUserRequest
	(*CreateUserResponse)(nil),             // 2: v1.bot.CreateUserResponse
	(*GetUsersRequest)(nil),                // 3: v1.bot.GetUsersRequest
	(*GetUsersResponse)(nil),               // 4: v1.bot.GetUsersResponse
	(*GetUserByDiscordIDRequest)(nil),      // 5: v1.bot.GetUserByDiscordIDRequest
	(*GetUserByAccountAddressRequest)(nil), // 6: v1.bot.GetUserByAccountAddressRequest
	(*GetUserResponse)(nil),                // 7: v1.bot.GetUserResponse
}
var file_v1_bot_users_proto_depIdxs = []int32{
	0, // 0: v1.bot.GetUsersResponse.users:type_name -> v1.bot.User
	0, // 1: v1.bot.GetUserResponse.user:type_name -> v1.bot.User
	1, // 2: v1.bot.UserService.CreateUser:input_type -> v1.bot.CreateUserRequest
	3, // 3: v1.bot.UserService.Users:input_type -> v1.bot.GetUsersRequest
	5, // 4: v1.bot.UserService.UserByDiscordID:input_type -> v1.bot.GetUserByDiscordIDRequest
	6, // 5: v1.bot.UserService.UserByAccountAddress:input_type -> v1.bot.GetUserByAccountAddressRequest
	2, // 6: v1.bot.UserService.CreateUser:output_type -> v1.bot.CreateUserResponse
	4, // 7: v1.bot.UserService.Users:output_type -> v1.bot.GetUsersResponse
	7, // 8: v1.bot.UserService.UserByDiscordID:output_type -> v1.bot.GetUserResponse
	7, // 9: v1.bot.UserService.UserByAccountAddress:output_type -> v1.bot.GetUserResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_bot_users_proto_init() }
func file_v1_bot_users_proto_init() {
	if File_v1_bot_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_bot_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_v1_bot_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_v1_bot_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_v1_bot_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_v1_bot_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
		file_v1_bot_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByDiscordIDRequest); i {
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
		file_v1_bot_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByAccountAddressRequest); i {
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
		file_v1_bot_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
			RawDescriptor: file_v1_bot_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bot_users_proto_goTypes,
		DependencyIndexes: file_v1_bot_users_proto_depIdxs,
		MessageInfos:      file_v1_bot_users_proto_msgTypes,
	}.Build()
	File_v1_bot_users_proto = out.File
	file_v1_bot_users_proto_rawDesc = nil
	file_v1_bot_users_proto_goTypes = nil
	file_v1_bot_users_proto_depIdxs = nil
}
