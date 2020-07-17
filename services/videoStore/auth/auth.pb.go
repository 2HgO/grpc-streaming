// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

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

type UserInfo struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Res struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (m *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(m, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Res) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidationInfo struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationInfo) Reset()         { *m = ValidationInfo{} }
func (m *ValidationInfo) String() string { return proto.CompactTextString(m) }
func (*ValidationInfo) ProtoMessage()    {}
func (*ValidationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *ValidationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationInfo.Unmarshal(m, b)
}
func (m *ValidationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationInfo.Marshal(b, m, deterministic)
}
func (m *ValidationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationInfo.Merge(m, src)
}
func (m *ValidationInfo) XXX_Size() int {
	return xxx_messageInfo_ValidationInfo.Size(m)
}
func (m *ValidationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationInfo proto.InternalMessageInfo

func (m *ValidationInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidationRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationRes) Reset()         { *m = ValidationRes{} }
func (m *ValidationRes) String() string { return proto.CompactTextString(m) }
func (*ValidationRes) ProtoMessage()    {}
func (*ValidationRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *ValidationRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationRes.Unmarshal(m, b)
}
func (m *ValidationRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationRes.Marshal(b, m, deterministic)
}
func (m *ValidationRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationRes.Merge(m, src)
}
func (m *ValidationRes) XXX_Size() int {
	return xxx_messageInfo_ValidationRes.Size(m)
}
func (m *ValidationRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationRes.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationRes proto.InternalMessageInfo

func (m *ValidationRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "auth.UserInfo")
	proto.RegisterType((*Res)(nil), "auth.Res")
	proto.RegisterType((*ValidationInfo)(nil), "auth.ValidationInfo")
	proto.RegisterType((*ValidationRes)(nil), "auth.ValidationRes")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x9c, 0xb8, 0x38, 0x42, 0x8b,
	0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0x85, 0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xf2, 0x12,
	0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x41, 0x62, 0x71,
	0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x13, 0x44, 0x0e, 0xc6, 0x57, 0x32, 0xe5, 0x62, 0x0e, 0x4a,
	0x2d, 0x16, 0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0x06, 0xeb, 0xe6, 0x08,
	0x82, 0x71, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0xa0, 0x3a, 0x21, 0x1c, 0x25,
	0x35, 0x2e, 0xbe, 0xb0, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0xb0, 0x03, 0xe0,
	0xea, 0x18, 0x91, 0xd5, 0x69, 0x72, 0xf1, 0x22, 0xd4, 0xe1, 0xb5, 0xc8, 0x68, 0x2a, 0x23, 0x17,
	0x8b, 0x63, 0x69, 0x49, 0x86, 0x90, 0x0a, 0x17, 0xab, 0x4f, 0x7e, 0x7a, 0x66, 0x9e, 0x10, 0x9f,
	0x1e, 0xd8, 0xcb, 0x30, 0x3f, 0x4a, 0x71, 0x42, 0xf8, 0x41, 0xa9, 0xc5, 0x4a, 0x0c, 0x42, 0x9a,
	0x5c, 0x5c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x20, 0x69, 0xfc, 0x4a, 0x6d, 0xe0, 0x8e, 0x48,
	0x0d, 0x01, 0xb9, 0x4a, 0x48, 0x04, 0x22, 0x8b, 0xea, 0x03, 0x29, 0x61, 0x74, 0x51, 0xb0, 0x6e,
	0x27, 0xc1, 0x28, 0xfe, 0xb2, 0xcc, 0x94, 0xd4, 0xfc, 0xe0, 0x92, 0xfc, 0xa2, 0x54, 0x7d, 0x90,
	0x92, 0x24, 0x36, 0x70, 0x2c, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x97, 0xff, 0x61,
	0x93, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Res, error)
	CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Res, error)
	ValidateToken(ctx context.Context, in *ValidationInfo, opts ...grpc.CallOption) (*ValidationRes, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/auth.Auth/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *ValidationInfo, opts ...grpc.CallOption) (*ValidationRes, error) {
	out := new(ValidationRes)
	err := c.cc.Invoke(ctx, "/auth.Auth/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Login(context.Context, *UserInfo) (*Res, error)
	CreateUser(context.Context, *UserInfo) (*Res, error)
	ValidateToken(context.Context, *ValidationInfo) (*ValidationRes, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Login(ctx context.Context, req *UserInfo) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServer) CreateUser(ctx context.Context, req *UserInfo) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedAuthServer) ValidateToken(ctx context.Context, req *ValidationInfo) (*ValidationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidateToken(ctx, req.(*ValidationInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Auth_CreateUser_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _Auth_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}