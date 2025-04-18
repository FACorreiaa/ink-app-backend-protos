// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: user.proto

package generated

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

type User_Role int32

const (
	User_ROLE_UNSPECIFIED User_Role = 0
	User_ADMIN            User_Role = 1
	User_STAFF            User_Role = 2
	User_USER             User_Role = 3
	User_MODERATOR        User_Role = 4
)

// Enum value maps for User_Role.
var (
	User_Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ADMIN",
		2: "STAFF",
		3: "USER",
		4: "MODERATOR",
	}
	User_Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ADMIN":            1,
		"STAFF":            2,
		"USER":             3,
		"MODERATOR":        4,
	}
)

func (x User_Role) Enum() *User_Role {
	p := new(User_Role)
	*p = x
	return p
}

func (x User_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (User_Role) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x User_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Role.Descriptor instead.
func (User_Role) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4, 0}
}

type GetUserByEmailReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Request       *BaseRequest           `protobuf:"bytes,100,opt,name=request,proto3" json:"request,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByEmailReq) Reset() {
	*x = GetUserByEmailReq{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailReq) ProtoMessage() {}

func (x *GetUserByEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailReq.ProtoReflect.Descriptor instead.
func (*GetUserByEmailReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserByEmailReq) GetRequest() *BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetUserByUsernameReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Request       *BaseRequest           `protobuf:"bytes,100,opt,name=request,proto3" json:"request,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByUsernameReq) Reset() {
	*x = GetUserByUsernameReq{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameReq) ProtoMessage() {}

func (x *GetUserByUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameReq.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByUsernameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserByUsernameReq) GetRequest() *BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetUserByUsernameRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Response      *BaseResponse          `protobuf:"bytes,100,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByUsernameRes) Reset() {
	*x = GetUserByUsernameRes{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByUsernameRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameRes) ProtoMessage() {}

func (x *GetUserByUsernameRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameRes.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByUsernameRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetUserByUsernameRes) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetUserByEmailRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Response      *BaseResponse          `protobuf:"bytes,100,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByEmailRes) Reset() {
	*x = GetUserByEmailRes{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByEmailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailRes) ProtoMessage() {}

func (x *GetUserByEmailRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailRes.ProtoReflect.Descriptor instead.
func (*GetUserByEmailRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByEmailRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetUserByEmailRes) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	PasswordHash  string                 `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	IsAdmin       bool                   `protobuf:"varint,5,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Role          User_Role              `protobuf:"varint,8,opt,name=role,proto3,enum=inkMe.user.User_Role" json:"role,omitempty"`
	StudioId      string                 `protobuf:"bytes,9,opt,name=studio_id,json=studioId,proto3" json:"studio_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
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
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *User) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *User) GetRole() User_Role {
	if x != nil {
		return x.Role
	}
	return User_ROLE_UNSPECIFIED
}

func (x *User) GetStudioId() string {
	if x != nil {
		return x.StudioId
	}
	return ""
}

// Request and response messages for new user-related services
type GetUsersReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Request       *BaseRequest           `protobuf:"bytes,100,opt,name=request,proto3" json:"request,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersReq) Reset() {
	*x = GetUsersReq{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersReq) ProtoMessage() {}

func (x *GetUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersReq.ProtoReflect.Descriptor instead.
func (*GetUsersReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersReq) GetRequest() *BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetUsersRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Users         []*User                `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Response      *BaseResponse          `protobuf:"bytes,100,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersRes) Reset() {
	*x = GetUsersRes{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRes) ProtoMessage() {}

func (x *GetUsersRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRes.ProtoReflect.Descriptor instead.
func (*GetUsersRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetUsersRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetUsersRes) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetUsersRes) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetUserByIDReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Request       *BaseRequest           `protobuf:"bytes,100,opt,name=request,proto3" json:"request,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIDReq) Reset() {
	*x = GetUserByIDReq{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDReq) ProtoMessage() {}

func (x *GetUserByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDReq.ProtoReflect.Descriptor instead.
func (*GetUserByIDReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserByIDReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserByIDReq) GetRequest() *BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetUserByIDRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Response      *BaseResponse          `protobuf:"bytes,100,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIDRes) Reset() {
	*x = GetUserByIDRes{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRes) ProtoMessage() {}

func (x *GetUserByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRes.ProtoReflect.Descriptor instead.
func (*GetUserByIDRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserByIDRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetUserByIDRes) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type DeleteUserReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Request       *BaseRequest           `protobuf:"bytes,100,opt,name=request,proto3" json:"request,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserReq) Reset() {
	*x = DeleteUserReq{}
	mi := &file_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReq) ProtoMessage() {}

func (x *DeleteUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReq.ProtoReflect.Descriptor instead.
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteUserReq) GetRequest() *BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type DeleteUserRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Response      *BaseResponse          `protobuf:"bytes,100,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRes) Reset() {
	*x = DeleteUserRes{}
	mi := &file_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRes) ProtoMessage() {}

func (x *DeleteUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRes.ProtoReflect.Descriptor instead.
func (*DeleteUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteUserRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteUserRes) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type UpdateUserReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Request       *BaseRequest           `protobuf:"bytes,100,opt,name=request,proto3" json:"request,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	mi := &file_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateUserReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateUserReq) GetRequest() *BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type UpdateUserRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Response      *BaseResponse          `protobuf:"bytes,100,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRes) Reset() {
	*x = UpdateUserRes{}
	mi := &file_user_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRes) ProtoMessage() {}

func (x *UpdateUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRes.ProtoReflect.Descriptor instead.
func (*UpdateUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateUserRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateUserRes) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type InsertUserReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Request       *BaseRequest           `protobuf:"bytes,100,opt,name=request,proto3" json:"request,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InsertUserReq) Reset() {
	*x = InsertUserReq{}
	mi := &file_user_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserReq) ProtoMessage() {}

func (x *InsertUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserReq.ProtoReflect.Descriptor instead.
func (*InsertUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{13}
}

func (x *InsertUserReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *InsertUserReq) GetRequest() *BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type InsertUserRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Response      *BaseResponse          `protobuf:"bytes,100,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InsertUserRes) Reset() {
	*x = InsertUserRes{}
	mi := &file_user_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserRes) ProtoMessage() {}

func (x *InsertUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserRes.ProtoReflect.Descriptor instead.
func (*InsertUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{14}
}

func (x *InsertUserRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *InsertUserRes) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type BaseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Downstream    string                 `protobuf:"bytes,998,opt,name=downstream,proto3" json:"downstream,omitempty"`
	RequestId     string                 `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	TraceId       string                 `protobuf:"bytes,1000,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseRequest) Reset() {
	*x = BaseRequest{}
	mi := &file_user_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseRequest) ProtoMessage() {}

func (x *BaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseRequest.ProtoReflect.Descriptor instead.
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{15}
}

func (x *BaseRequest) GetDownstream() string {
	if x != nil {
		return x.Downstream
	}
	return ""
}

func (x *BaseRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *BaseRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type BaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,997,opt,name=success,proto3" json:"success,omitempty"`
	Upstream      string                 `protobuf:"bytes,998,opt,name=upstream,proto3" json:"upstream,omitempty"`
	RequestId     string                 `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Status        string                 `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	TraceId       string                 `protobuf:"bytes,9999,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_user_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{16}
}

func (x *BaseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BaseResponse) GetUpstream() string {
	if x != nil {
		return x.Upstream
	}
	return ""
}

func (x *BaseResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *BaseResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BaseResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x12\n" +
	"inkMe.user\"\\\n" +
	"\x11GetUserByEmailReq\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x121\n" +
	"\arequest\x18d \x01(\v2\x17.inkMe.user.BaseRequestR\arequest\"e\n" +
	"\x14GetUserByUsernameReq\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x121\n" +
	"\arequest\x18d \x01(\v2\x17.inkMe.user.BaseRequestR\arequest\"r\n" +
	"\x14GetUserByUsernameRes\x12$\n" +
	"\x04user\x18\x01 \x01(\v2\x10.inkMe.user.UserR\x04user\x124\n" +
	"\bresponse\x18d \x01(\v2\x18.inkMe.user.BaseResponseR\bresponse\"o\n" +
	"\x11GetUserByEmailRes\x12$\n" +
	"\x04user\x18\x01 \x01(\v2\x10.inkMe.user.UserR\x04user\x124\n" +
	"\bresponse\x18d \x01(\v2\x18.inkMe.user.BaseResponseR\bresponse\"\xe4\x02\n" +
	"\x04User\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\busername\x18\x03 \x01(\tR\busername\x12#\n" +
	"\rpassword_hash\x18\x04 \x01(\tR\fpasswordHash\x12\x19\n" +
	"\bis_admin\x18\x05 \x01(\bR\aisAdmin\x12\x1d\n" +
	"\n" +
	"created_at\x18\x06 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\a \x01(\tR\tupdatedAt\x12)\n" +
	"\x04role\x18\b \x01(\x0e2\x15.inkMe.user.User.RoleR\x04role\x12\x1b\n" +
	"\tstudio_id\x18\t \x01(\tR\bstudioId\"K\n" +
	"\x04Role\x12\x14\n" +
	"\x10ROLE_UNSPECIFIED\x10\x00\x12\t\n" +
	"\x05ADMIN\x10\x01\x12\t\n" +
	"\x05STAFF\x10\x02\x12\b\n" +
	"\x04USER\x10\x03\x12\r\n" +
	"\tMODERATOR\x10\x04\"@\n" +
	"\vGetUsersReq\x121\n" +
	"\arequest\x18d \x01(\v2\x17.inkMe.user.BaseRequestR\arequest\"\x9f\x01\n" +
	"\vGetUsersRes\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12&\n" +
	"\x05users\x18\x03 \x03(\v2\x10.inkMe.user.UserR\x05users\x124\n" +
	"\bresponse\x18d \x01(\v2\x18.inkMe.user.BaseResponseR\bresponse\"\\\n" +
	"\x0eGetUserByIDReq\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x121\n" +
	"\arequest\x18d \x01(\v2\x17.inkMe.user.BaseRequestR\arequest\"l\n" +
	"\x0eGetUserByIDRes\x12$\n" +
	"\x04user\x18\x01 \x01(\v2\x10.inkMe.user.UserR\x04user\x124\n" +
	"\bresponse\x18d \x01(\v2\x18.inkMe.user.BaseResponseR\bresponse\"[\n" +
	"\rDeleteUserReq\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x121\n" +
	"\arequest\x18d \x01(\v2\x17.inkMe.user.BaseRequestR\arequest\"_\n" +
	"\rDeleteUserRes\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x124\n" +
	"\bresponse\x18d \x01(\v2\x18.inkMe.user.BaseResponseR\bresponse\"h\n" +
	"\rUpdateUserReq\x12$\n" +
	"\x04user\x18\x01 \x01(\v2\x10.inkMe.user.UserR\x04user\x121\n" +
	"\arequest\x18d \x01(\v2\x17.inkMe.user.BaseRequestR\arequest\"_\n" +
	"\rUpdateUserRes\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x124\n" +
	"\bresponse\x18d \x01(\v2\x18.inkMe.user.BaseResponseR\bresponse\"h\n" +
	"\rInsertUserReq\x12$\n" +
	"\x04user\x18\x01 \x01(\v2\x10.inkMe.user.UserR\x04user\x121\n" +
	"\arequest\x18d \x01(\v2\x17.inkMe.user.BaseRequestR\arequest\"_\n" +
	"\rInsertUserRes\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x124\n" +
	"\bresponse\x18d \x01(\v2\x18.inkMe.user.BaseResponseR\bresponse\"j\n" +
	"\vBaseRequest\x12\x1f\n" +
	"\n" +
	"downstream\x18\xe6\a \x01(\tR\n" +
	"downstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\x12\x1a\n" +
	"\btrace_id\x18\xe8\a \x01(\tR\atraceId\"\x9b\x01\n" +
	"\fBaseResponse\x12\x19\n" +
	"\asuccess\x18\xe5\a \x01(\bR\asuccess\x12\x1b\n" +
	"\bupstream\x18\xe6\a \x01(\tR\bupstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\x12\x17\n" +
	"\x06status\x18\xe8\a \x01(\tR\x06status\x12\x1a\n" +
	"\btrace_id\x18\x8fN \x01(\tR\atraceId2\x87\x04\n" +
	"\vUserService\x12<\n" +
	"\bGetUsers\x12\x17.inkMe.user.GetUsersReq\x1a\x17.inkMe.user.GetUsersRes\x12E\n" +
	"\vGetUserByID\x12\x1a.inkMe.user.GetUserByIDReq\x1a\x1a.inkMe.user.GetUserByIDRes\x12B\n" +
	"\n" +
	"DeleteUser\x12\x19.inkMe.user.DeleteUserReq\x1a\x19.inkMe.user.DeleteUserRes\x12B\n" +
	"\n" +
	"UpdateUser\x12\x19.inkMe.user.UpdateUserReq\x1a\x19.inkMe.user.UpdateUserRes\x12B\n" +
	"\n" +
	"InsertUser\x12\x19.inkMe.user.InsertUserReq\x1a\x19.inkMe.user.InsertUserRes\x12N\n" +
	"\x0eGetUserByEmail\x12\x1d.inkMe.user.GetUserByEmailReq\x1a\x1d.inkMe.user.GetUserByEmailRes\x12W\n" +
	"\x11GetUserByUsername\x12 .inkMe.user.GetUserByUsernameReq\x1a .inkMe.user.GetUserByUsernameResBEZCgithub.com/FACorreiaa/ink-app-backend-protos/modules/user/generatedb\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_user_proto_goTypes = []any{
	(User_Role)(0),               // 0: inkMe.user.User.Role
	(*GetUserByEmailReq)(nil),    // 1: inkMe.user.GetUserByEmailReq
	(*GetUserByUsernameReq)(nil), // 2: inkMe.user.GetUserByUsernameReq
	(*GetUserByUsernameRes)(nil), // 3: inkMe.user.GetUserByUsernameRes
	(*GetUserByEmailRes)(nil),    // 4: inkMe.user.GetUserByEmailRes
	(*User)(nil),                 // 5: inkMe.user.User
	(*GetUsersReq)(nil),          // 6: inkMe.user.GetUsersReq
	(*GetUsersRes)(nil),          // 7: inkMe.user.GetUsersRes
	(*GetUserByIDReq)(nil),       // 8: inkMe.user.GetUserByIDReq
	(*GetUserByIDRes)(nil),       // 9: inkMe.user.GetUserByIDRes
	(*DeleteUserReq)(nil),        // 10: inkMe.user.DeleteUserReq
	(*DeleteUserRes)(nil),        // 11: inkMe.user.DeleteUserRes
	(*UpdateUserReq)(nil),        // 12: inkMe.user.UpdateUserReq
	(*UpdateUserRes)(nil),        // 13: inkMe.user.UpdateUserRes
	(*InsertUserReq)(nil),        // 14: inkMe.user.InsertUserReq
	(*InsertUserRes)(nil),        // 15: inkMe.user.InsertUserRes
	(*BaseRequest)(nil),          // 16: inkMe.user.BaseRequest
	(*BaseResponse)(nil),         // 17: inkMe.user.BaseResponse
}
var file_user_proto_depIdxs = []int32{
	16, // 0: inkMe.user.GetUserByEmailReq.request:type_name -> inkMe.user.BaseRequest
	16, // 1: inkMe.user.GetUserByUsernameReq.request:type_name -> inkMe.user.BaseRequest
	5,  // 2: inkMe.user.GetUserByUsernameRes.user:type_name -> inkMe.user.User
	17, // 3: inkMe.user.GetUserByUsernameRes.response:type_name -> inkMe.user.BaseResponse
	5,  // 4: inkMe.user.GetUserByEmailRes.user:type_name -> inkMe.user.User
	17, // 5: inkMe.user.GetUserByEmailRes.response:type_name -> inkMe.user.BaseResponse
	0,  // 6: inkMe.user.User.role:type_name -> inkMe.user.User.Role
	16, // 7: inkMe.user.GetUsersReq.request:type_name -> inkMe.user.BaseRequest
	5,  // 8: inkMe.user.GetUsersRes.users:type_name -> inkMe.user.User
	17, // 9: inkMe.user.GetUsersRes.response:type_name -> inkMe.user.BaseResponse
	16, // 10: inkMe.user.GetUserByIDReq.request:type_name -> inkMe.user.BaseRequest
	5,  // 11: inkMe.user.GetUserByIDRes.user:type_name -> inkMe.user.User
	17, // 12: inkMe.user.GetUserByIDRes.response:type_name -> inkMe.user.BaseResponse
	16, // 13: inkMe.user.DeleteUserReq.request:type_name -> inkMe.user.BaseRequest
	17, // 14: inkMe.user.DeleteUserRes.response:type_name -> inkMe.user.BaseResponse
	5,  // 15: inkMe.user.UpdateUserReq.user:type_name -> inkMe.user.User
	16, // 16: inkMe.user.UpdateUserReq.request:type_name -> inkMe.user.BaseRequest
	17, // 17: inkMe.user.UpdateUserRes.response:type_name -> inkMe.user.BaseResponse
	5,  // 18: inkMe.user.InsertUserReq.user:type_name -> inkMe.user.User
	16, // 19: inkMe.user.InsertUserReq.request:type_name -> inkMe.user.BaseRequest
	17, // 20: inkMe.user.InsertUserRes.response:type_name -> inkMe.user.BaseResponse
	6,  // 21: inkMe.user.UserService.GetUsers:input_type -> inkMe.user.GetUsersReq
	8,  // 22: inkMe.user.UserService.GetUserByID:input_type -> inkMe.user.GetUserByIDReq
	10, // 23: inkMe.user.UserService.DeleteUser:input_type -> inkMe.user.DeleteUserReq
	12, // 24: inkMe.user.UserService.UpdateUser:input_type -> inkMe.user.UpdateUserReq
	14, // 25: inkMe.user.UserService.InsertUser:input_type -> inkMe.user.InsertUserReq
	1,  // 26: inkMe.user.UserService.GetUserByEmail:input_type -> inkMe.user.GetUserByEmailReq
	2,  // 27: inkMe.user.UserService.GetUserByUsername:input_type -> inkMe.user.GetUserByUsernameReq
	7,  // 28: inkMe.user.UserService.GetUsers:output_type -> inkMe.user.GetUsersRes
	9,  // 29: inkMe.user.UserService.GetUserByID:output_type -> inkMe.user.GetUserByIDRes
	11, // 30: inkMe.user.UserService.DeleteUser:output_type -> inkMe.user.DeleteUserRes
	13, // 31: inkMe.user.UserService.UpdateUser:output_type -> inkMe.user.UpdateUserRes
	15, // 32: inkMe.user.UserService.InsertUser:output_type -> inkMe.user.InsertUserRes
	4,  // 33: inkMe.user.UserService.GetUserByEmail:output_type -> inkMe.user.GetUserByEmailRes
	3,  // 34: inkMe.user.UserService.GetUserByUsername:output_type -> inkMe.user.GetUserByUsernameRes
	28, // [28:35] is the sub-list for method output_type
	21, // [21:28] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
