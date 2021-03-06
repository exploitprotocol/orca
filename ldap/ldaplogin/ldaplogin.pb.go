// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ldaplogin/ldaplogin.proto

package ldaplogin

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

type AuthReply_AuthStatus int32

const (
	AuthReply_SERVER_ERROR AuthReply_AuthStatus = 0
	AuthReply_OK           AuthReply_AuthStatus = 1
	AuthReply_FAILED       AuthReply_AuthStatus = 2
)

var AuthReply_AuthStatus_name = map[int32]string{
	0: "SERVER_ERROR",
	1: "OK",
	2: "FAILED",
}

var AuthReply_AuthStatus_value = map[string]int32{
	"SERVER_ERROR": 0,
	"OK":           1,
	"FAILED":       2,
}

func (x AuthReply_AuthStatus) String() string {
	return proto.EnumName(AuthReply_AuthStatus_name, int32(x))
}

func (AuthReply_AuthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60a91fa331746fa7, []int{2, 0}
}

type PasswdAuthRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswdAuthRequest) Reset()         { *m = PasswdAuthRequest{} }
func (m *PasswdAuthRequest) String() string { return proto.CompactTextString(m) }
func (*PasswdAuthRequest) ProtoMessage()    {}
func (*PasswdAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60a91fa331746fa7, []int{0}
}

func (m *PasswdAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswdAuthRequest.Unmarshal(m, b)
}
func (m *PasswdAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswdAuthRequest.Marshal(b, m, deterministic)
}
func (m *PasswdAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswdAuthRequest.Merge(m, src)
}
func (m *PasswdAuthRequest) XXX_Size() int {
	return xxx_messageInfo_PasswdAuthRequest.Size(m)
}
func (m *PasswdAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswdAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PasswdAuthRequest proto.InternalMessageInfo

func (m *PasswdAuthRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *PasswdAuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type KeyAuthRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyAuthRequest) Reset()         { *m = KeyAuthRequest{} }
func (m *KeyAuthRequest) String() string { return proto.CompactTextString(m) }
func (*KeyAuthRequest) ProtoMessage()    {}
func (*KeyAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60a91fa331746fa7, []int{1}
}

func (m *KeyAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyAuthRequest.Unmarshal(m, b)
}
func (m *KeyAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyAuthRequest.Marshal(b, m, deterministic)
}
func (m *KeyAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyAuthRequest.Merge(m, src)
}
func (m *KeyAuthRequest) XXX_Size() int {
	return xxx_messageInfo_KeyAuthRequest.Size(m)
}
func (m *KeyAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyAuthRequest proto.InternalMessageInfo

func (m *KeyAuthRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *KeyAuthRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type AuthReply struct {
	Status               AuthReply_AuthStatus `protobuf:"varint,1,opt,name=status,proto3,enum=AuthReply_AuthStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthReply) Reset()         { *m = AuthReply{} }
func (m *AuthReply) String() string { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()    {}
func (*AuthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_60a91fa331746fa7, []int{2}
}

func (m *AuthReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReply.Unmarshal(m, b)
}
func (m *AuthReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReply.Marshal(b, m, deterministic)
}
func (m *AuthReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReply.Merge(m, src)
}
func (m *AuthReply) XXX_Size() int {
	return xxx_messageInfo_AuthReply.Size(m)
}
func (m *AuthReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReply proto.InternalMessageInfo

func (m *AuthReply) GetStatus() AuthReply_AuthStatus {
	if m != nil {
		return m.Status
	}
	return AuthReply_SERVER_ERROR
}

func init() {
	proto.RegisterEnum("AuthReply_AuthStatus", AuthReply_AuthStatus_name, AuthReply_AuthStatus_value)
	proto.RegisterType((*PasswdAuthRequest)(nil), "PasswdAuthRequest")
	proto.RegisterType((*KeyAuthRequest)(nil), "KeyAuthRequest")
	proto.RegisterType((*AuthReply)(nil), "AuthReply")
}

func init() { proto.RegisterFile("ldaplogin/ldaplogin.proto", fileDescriptor_60a91fa331746fa7) }

var fileDescriptor_60a91fa331746fa7 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9b, 0x80, 0xd1, 0x3c, 0x4a, 0x8d, 0x83, 0x42, 0x2d, 0x1e, 0x24, 0xa7, 0x5e, 0x5c,
	0x21, 0xfe, 0x82, 0x40, 0x56, 0x90, 0x04, 0x5a, 0xa6, 0xe0, 0x55, 0x52, 0xb3, 0x68, 0x21, 0x34,
	0x6b, 0x77, 0x83, 0xe4, 0xdf, 0x4b, 0x37, 0x92, 0x50, 0x3c, 0x78, 0x9b, 0x37, 0xef, 0xf1, 0x31,
	0xf3, 0x70, 0x5b, 0x57, 0xa5, 0xae, 0x9b, 0x8f, 0xdd, 0xfe, 0x71, 0x98, 0x84, 0x3e, 0x34, 0xb6,
	0x89, 0x25, 0xae, 0xd6, 0xa5, 0x31, 0xdf, 0x55, 0xda, 0xda, 0x4f, 0x56, 0x5f, 0xad, 0x32, 0x96,
	0xae, 0x71, 0xe6, 0x32, 0x73, 0xef, 0xde, 0x5b, 0x86, 0xdc, 0x0b, 0x5a, 0xe0, 0x42, 0x1f, 0xa3,
	0xcd, 0xa1, 0x9a, 0xfb, 0xce, 0x18, 0x74, 0x9c, 0x61, 0x96, 0xab, 0xee, 0x7f, 0xc6, 0x1d, 0x42,
	0xdd, 0x6e, 0xeb, 0xdd, 0x7b, 0xae, 0x3a, 0x07, 0x99, 0xf2, 0xb8, 0x88, 0xf7, 0x08, 0x7b, 0x84,
	0xae, 0x3b, 0x7a, 0x40, 0x60, 0x6c, 0x69, 0x5b, 0xe3, 0x08, 0xb3, 0xe4, 0x46, 0x0c, 0x9e, 0x9b,
	0x36, 0xce, 0xe4, 0xdf, 0x50, 0x9c, 0x00, 0xe3, 0x96, 0x22, 0x4c, 0x37, 0x92, 0x5f, 0x25, 0xbf,
	0x49, 0xe6, 0x15, 0x47, 0x13, 0x0a, 0xe0, 0xaf, 0xf2, 0xc8, 0x23, 0x20, 0x78, 0x4e, 0x5f, 0x0a,
	0x99, 0x45, 0x7e, 0xa2, 0x10, 0x16, 0x59, 0xba, 0x2e, 0xdc, 0x69, 0xa2, 0x07, 0xf4, 0x6d, 0x10,
	0x89, 0x3f, 0xb5, 0x2c, 0x30, 0x5e, 0x10, 0x4f, 0x68, 0x89, 0xf3, 0xa3, 0xcc, 0x55, 0x47, 0x97,
	0xe2, 0xf4, 0xf9, 0xd3, 0xe4, 0x36, 0x70, 0x55, 0x3f, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x93,
	0x1d, 0x22, 0xf2, 0x87, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LDAPLoginClient is the client API for LDAPLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LDAPLoginClient interface {
	AuthPasswd(ctx context.Context, in *PasswdAuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	AuthKey(ctx context.Context, in *KeyAuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
}

type lDAPLoginClient struct {
	cc *grpc.ClientConn
}

func NewLDAPLoginClient(cc *grpc.ClientConn) LDAPLoginClient {
	return &lDAPLoginClient{cc}
}

func (c *lDAPLoginClient) AuthPasswd(ctx context.Context, in *PasswdAuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/LDAPLogin/AuthPasswd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lDAPLoginClient) AuthKey(ctx context.Context, in *KeyAuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/LDAPLogin/AuthKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LDAPLoginServer is the server API for LDAPLogin service.
type LDAPLoginServer interface {
	AuthPasswd(context.Context, *PasswdAuthRequest) (*AuthReply, error)
	AuthKey(context.Context, *KeyAuthRequest) (*AuthReply, error)
}

// UnimplementedLDAPLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLDAPLoginServer struct {
}

func (*UnimplementedLDAPLoginServer) AuthPasswd(ctx context.Context, req *PasswdAuthRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthPasswd not implemented")
}
func (*UnimplementedLDAPLoginServer) AuthKey(ctx context.Context, req *KeyAuthRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthKey not implemented")
}

func RegisterLDAPLoginServer(s *grpc.Server, srv LDAPLoginServer) {
	s.RegisterService(&_LDAPLogin_serviceDesc, srv)
}

func _LDAPLogin_AuthPasswd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswdAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LDAPLoginServer).AuthPasswd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LDAPLogin/AuthPasswd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LDAPLoginServer).AuthPasswd(ctx, req.(*PasswdAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LDAPLogin_AuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LDAPLoginServer).AuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LDAPLogin/AuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LDAPLoginServer).AuthKey(ctx, req.(*KeyAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LDAPLogin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LDAPLogin",
	HandlerType: (*LDAPLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthPasswd",
			Handler:    _LDAPLogin_AuthPasswd_Handler,
		},
		{
			MethodName: "AuthKey",
			Handler:    _LDAPLogin_AuthKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ldaplogin/ldaplogin.proto",
}
