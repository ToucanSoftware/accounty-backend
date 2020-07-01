// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf-spec/core.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// User in the system
type User struct {
	// Id of the user
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Username
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Password
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// e-mail
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// CreateUserRequest
type CreateUserRequest struct {
	// User to create
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// CreateUserResponse
type CreateUserResponse struct {
	// Created user
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// ListUsersRequest is a request for listing users
type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{3}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

// ListUsersResponse contains the list of users found
type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{4}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

// GetUserRequest is a request for getting an existing user
type GetUserRequest struct {
	// Id of the user
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{5}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// GetUserResponse contains the response with an existing user
type GetUserResponse struct {
	// Found user
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{6}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// UpdateUserRequest is a request for updating a user
type UpdateUserRequest struct {
	// Id of the user
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// User to update
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{7}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// UpdateUserResponse contains the response with the updated user
type UpdateUserResponse struct {
	// Updated user
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9af7e0fea015fe, []int{8}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "github.com.io.toucasoft.accounty.v1.User")
	proto.RegisterType((*CreateUserRequest)(nil), "github.com.io.toucasoft.accounty.v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "github.com.io.toucasoft.accounty.v1.CreateUserResponse")
	proto.RegisterType((*ListUsersRequest)(nil), "github.com.io.toucasoft.accounty.v1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "github.com.io.toucasoft.accounty.v1.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "github.com.io.toucasoft.accounty.v1.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "github.com.io.toucasoft.accounty.v1.GetUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "github.com.io.toucasoft.accounty.v1.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "github.com.io.toucasoft.accounty.v1.UpdateUserResponse")
}

func init() { proto.RegisterFile("protobuf-spec/core.proto", fileDescriptor_ea9af7e0fea015fe) }

var fileDescriptor_ea9af7e0fea015fe = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x65, 0xc7, 0x81, 0xe6, 0x55, 0x0a, 0xf5, 0x13, 0x14, 0xcb, 0xea, 0x10, 0x1d, 0x4b,
	0x19, 0xb0, 0x95, 0x16, 0x8a, 0x84, 0x84, 0x90, 0x60, 0x60, 0x01, 0x09, 0xa5, 0x74, 0x61, 0xbb,
	0xd8, 0xaf, 0xe1, 0xa4, 0xe6, 0xce, 0xf8, 0xce, 0x01, 0x84, 0x58, 0x98, 0x61, 0x42, 0x62, 0xe1,
	0x63, 0xf1, 0x15, 0x58, 0xf9, 0x0e, 0xc8, 0x67, 0x37, 0x09, 0xce, 0x62, 0xd3, 0x6c, 0x7e, 0xcf,
	0xf7, 0x7f, 0xef, 0x77, 0xfe, 0xff, 0x65, 0x08, 0xb2, 0x5c, 0x19, 0x35, 0x2d, 0xce, 0xef, 0xe9,
	0x8c, 0x92, 0x38, 0x51, 0x39, 0x45, 0xb6, 0x85, 0x77, 0x66, 0xc2, 0xbc, 0x2d, 0xa6, 0x51, 0xa2,
	0xe6, 0x91, 0x50, 0x91, 0x51, 0x45, 0xc2, 0xb5, 0x3a, 0x37, 0x11, 0x4f, 0x12, 0x55, 0x48, 0xf3,
	0x31, 0x5a, 0x8c, 0xc3, 0x83, 0x99, 0x52, 0xb3, 0x0b, 0x8a, 0x79, 0x26, 0x62, 0x2e, 0xa5, 0x32,
	0xdc, 0x08, 0x25, 0x75, 0x35, 0x82, 0x7d, 0x00, 0xef, 0x4c, 0x53, 0x8e, 0x43, 0x70, 0x45, 0x1a,
	0x38, 0x23, 0xe7, 0xb0, 0x37, 0x71, 0x45, 0x8a, 0x21, 0xec, 0x14, 0x9a, 0x72, 0xc9, 0xe7, 0x14,
	0xb8, 0x23, 0xe7, 0x70, 0x30, 0x59, 0xd6, 0x88, 0xe0, 0xd9, 0x7e, 0xcf, 0xf6, 0xed, 0x73, 0x79,
	0x3e, 0xe3, 0x5a, 0xbf, 0x57, 0x79, 0x1a, 0x78, 0xd5, 0xf9, 0xcb, 0x1a, 0x6f, 0x42, 0x9f, 0xe6,
	0x5c, 0x5c, 0x04, 0x7d, 0xfb, 0xa2, 0x2a, 0xd8, 0x04, 0xfc, 0x67, 0x39, 0x71, 0x43, 0xe5, 0xfe,
	0x09, 0xbd, 0x2b, 0x48, 0x1b, 0x7c, 0x0c, 0x5e, 0xb9, 0xc6, 0x82, 0xec, 0x1e, 0xdd, 0x8d, 0x5a,
	0x5c, 0x30, 0xb2, 0x7a, 0x2b, 0x63, 0xa7, 0x80, 0xeb, 0x33, 0x75, 0xa6, 0xa4, 0xa6, 0xab, 0x0e,
	0x45, 0xd8, 0x7b, 0x21, 0xb4, 0x29, 0x3b, 0xba, 0xe6, 0x64, 0xaf, 0xc1, 0x5f, 0xeb, 0xd5, 0x7b,
	0x9e, 0x40, 0xbf, 0x14, 0xe8, 0xc0, 0x19, 0xf5, 0xba, 0x2d, 0xaa, 0x74, 0x6c, 0x04, 0xc3, 0xe7,
	0x64, 0xd6, 0xbf, 0x47, 0xc3, 0x16, 0xf6, 0x0a, 0x6e, 0x2c, 0x4f, 0x6c, 0xe7, 0x76, 0x53, 0xf0,
	0xcf, 0xb2, 0xb4, 0x61, 0x43, 0x33, 0x0d, 0x97, 0x3b, 0xdc, 0xff, 0xb6, 0x65, 0x7d, 0xc7, 0x56,
	0xc0, 0x8f, 0xfe, 0x78, 0x70, 0xab, 0x2c, 0x5f, 0x72, 0xc9, 0x67, 0x34, 0x27, 0x69, 0x4e, 0x29,
	0x5f, 0x88, 0x84, 0xf0, 0x87, 0x03, 0xb0, 0x8a, 0x01, 0x9e, 0xb4, 0x9a, 0xbc, 0x91, 0xc5, 0xf0,
	0x61, 0x67, 0x5d, 0x75, 0x31, 0x76, 0xfb, 0xcb, 0xaf, 0xdf, 0xdf, 0x5d, 0x9f, 0x0d, 0xe2, 0xc5,
	0x38, 0xb6, 0xce, 0x3e, 0xb2, 0xc8, 0xf8, 0xcd, 0x81, 0xc1, 0x32, 0x36, 0xf8, 0xa0, 0xd5, 0xfc,
	0x66, 0xf4, 0xc2, 0x93, 0xae, 0xb2, 0x9a, 0xca, 0xb7, 0x54, 0xbb, 0xb8, 0xa2, 0xc2, 0xaf, 0x0e,
	0x5c, 0xaf, 0xe3, 0x84, 0xc7, 0xad, 0xc6, 0xfe, 0x1b, 0xcf, 0xf0, 0x7e, 0x37, 0x51, 0x4d, 0xb2,
	0x6f, 0x49, 0xf6, 0x70, 0xb8, 0x24, 0x89, 0x3f, 0x89, 0xf4, 0x33, 0xfe, 0x74, 0x00, 0x56, 0x39,
	0x69, 0xe9, 0xdb, 0x46, 0x78, 0x5b, 0xfa, 0xb6, 0x19, 0x48, 0x76, 0x60, 0xb9, 0xf6, 0xc3, 0x06,
	0x57, 0x65, 0xde, 0x53, 0x78, 0xb3, 0x53, 0xf5, 0x16, 0xe3, 0xe9, 0x35, 0xfb, 0xf3, 0x3c, 0xfe,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xa0, 0x0c, 0x1b, 0x9b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserManagementServiceClient is the client API for UserManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserManagementServiceClient interface {
	// Creates a new User in the system
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// List the users in the system
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// Get an existing users
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Updates an existing users
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type userManagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserManagementServiceClient(cc *grpc.ClientConn) UserManagementServiceClient {
	return &userManagementServiceClient{cc}
}

func (c *userManagementServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.io.toucasoft.accounty.v1.UserManagementService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/github.com.io.toucasoft.accounty.v1.UserManagementService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.io.toucasoft.accounty.v1.UserManagementService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.io.toucasoft.accounty.v1.UserManagementService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServiceServer is the server API for UserManagementService service.
type UserManagementServiceServer interface {
	// Creates a new User in the system
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// List the users in the system
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// Get an existing users
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Updates an existing users
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

func RegisterUserManagementServiceServer(s *grpc.Server, srv UserManagementServiceServer) {
	s.RegisterService(&_UserManagementService_serviceDesc, srv)
}

func _UserManagementService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.io.toucasoft.accounty.v1.UserManagementService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.io.toucasoft.accounty.v1.UserManagementService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.io.toucasoft.accounty.v1.UserManagementService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.io.toucasoft.accounty.v1.UserManagementService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.io.toucasoft.accounty.v1.UserManagementService",
	HandlerType: (*UserManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserManagementService_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserManagementService_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserManagementService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManagementService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf-spec/core.proto",
}
