// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Procedures
type UserGetByIDRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetByIDRequest) Reset()         { *m = UserGetByIDRequest{} }
func (m *UserGetByIDRequest) String() string { return proto.CompactTextString(m) }
func (*UserGetByIDRequest) ProtoMessage()    {}
func (*UserGetByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *UserGetByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetByIDRequest.Unmarshal(m, b)
}
func (m *UserGetByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetByIDRequest.Marshal(b, m, deterministic)
}
func (m *UserGetByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetByIDRequest.Merge(m, src)
}
func (m *UserGetByIDRequest) XXX_Size() int {
	return xxx_messageInfo_UserGetByIDRequest.Size(m)
}
func (m *UserGetByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetByIDRequest proto.InternalMessageInfo

func (m *UserGetByIDRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserGetByIDResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetByIDResponse) Reset()         { *m = UserGetByIDResponse{} }
func (m *UserGetByIDResponse) String() string { return proto.CompactTextString(m) }
func (*UserGetByIDResponse) ProtoMessage()    {}
func (*UserGetByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *UserGetByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetByIDResponse.Unmarshal(m, b)
}
func (m *UserGetByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetByIDResponse.Marshal(b, m, deterministic)
}
func (m *UserGetByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetByIDResponse.Merge(m, src)
}
func (m *UserGetByIDResponse) XXX_Size() int {
	return xxx_messageInfo_UserGetByIDResponse.Size(m)
}
func (m *UserGetByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetByIDResponse proto.InternalMessageInfo

func (m *UserGetByIDResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Models
type User struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
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

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*UserGetByIDRequest)(nil), "pb.UserGetByIDRequest")
	proto.RegisterType((*UserGetByIDResponse)(nil), "pb.UserGetByIDResponse")
	proto.RegisterType((*User)(nil), "pb.User")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe1, 0x12, 0x0a,
	0x2d, 0x4e, 0x2d, 0x72, 0x4f, 0x2d, 0x71, 0xaa, 0xf4, 0x74, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca,
	0x4c, 0x51, 0x32, 0xe6, 0x12, 0x46, 0x51, 0x55, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc3,
	0xc5, 0x02, 0x32, 0x0f, 0xac, 0x90, 0xdb, 0x88, 0x43, 0xaf, 0x20, 0x49, 0x0f, 0xa4, 0x2c, 0x08,
	0x2c, 0xaa, 0x14, 0xc6, 0xc5, 0x02, 0xe2, 0xa1, 0x1b, 0x26, 0x24, 0xc2, 0xc5, 0x9a, 0x93, 0x9f,
	0x9e, 0x99, 0x27, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x80, 0x44, 0x53, 0x73, 0x13,
	0x33, 0x73, 0x24, 0x98, 0x21, 0xa2, 0x60, 0x8e, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa,
	0x04, 0x0b, 0x58, 0x10, 0xcc, 0x36, 0x0a, 0xe0, 0xe2, 0x01, 0x99, 0x5b, 0x1c, 0x9c, 0x5a, 0x54,
	0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc0, 0xc5, 0x8d, 0xe4, 0x38, 0x21, 0x31, 0x98, 0x33, 0x50, 0xfd,
	0x24, 0x25, 0x8e, 0x21, 0x0e, 0xf1, 0x85, 0x12, 0x43, 0x12, 0x1b, 0x38, 0x3c, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0x89, 0x2b, 0x4e, 0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersServiceClient interface {
	UserGetByID(ctx context.Context, in *UserGetByIDRequest, opts ...grpc.CallOption) (*UserGetByIDResponse, error)
}

type usersServiceClient struct {
	cc *grpc.ClientConn
}

func NewUsersServiceClient(cc *grpc.ClientConn) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) UserGetByID(ctx context.Context, in *UserGetByIDRequest, opts ...grpc.CallOption) (*UserGetByIDResponse, error) {
	out := new(UserGetByIDResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/UserGetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
type UsersServiceServer interface {
	UserGetByID(context.Context, *UserGetByIDRequest) (*UserGetByIDResponse, error)
}

// UnimplementedUsersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (*UnimplementedUsersServiceServer) UserGetByID(ctx context.Context, req *UserGetByIDRequest) (*UserGetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetByID not implemented")
}

func RegisterUsersServiceServer(s *grpc.Server, srv UsersServiceServer) {
	s.RegisterService(&_UsersService_serviceDesc, srv)
}

func _UsersService_UserGetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UserGetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/UserGetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UserGetByID(ctx, req.(*UserGetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserGetByID",
			Handler:    _UsersService_UserGetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
