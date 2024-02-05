// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/cmd/auth/grpc/auth.proto

package grpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/elojah/trax/pkg/gogoproto"
	pbtypes "github.com/elojah/trax/pkg/pbtypes"
	dto "github.com/elojah/trax/pkg/user/dto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("github.com/elojah/trax/cmd/auth/grpc/auth.proto", fileDescriptor_caddd9128ba9588f)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/cmd/auth/grpc/auth.proto", fileDescriptor_caddd9128ba9588f)
}

var fileDescriptor_caddd9128ba9588f = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xfd, 0xa4, 0x8a, 0x21, 0x42, 0x54, 0xca, 0x98, 0xe1, 0x8d, 0x8c, 0x36, 0x82, 0x13,
	0x80, 0x84, 0x60, 0x44, 0x2d, 0x17, 0x68, 0x12, 0xe3, 0x84, 0x36, 0x79, 0x96, 0xe3, 0x48, 0x74,
	0xe3, 0x08, 0x88, 0x53, 0x70, 0x04, 0xc6, 0x8e, 0x1d, 0x33, 0x76, 0x24, 0xce, 0xc2, 0xd8, 0x91,
	0x11, 0xc5, 0xa9, 0x18, 0x90, 0x40, 0xd9, 0x7e, 0x5b, 0xdf, 0xf7, 0x3f, 0xfd, 0x81, 0x50, 0xb9,
	0xcd, 0xea, 0x98, 0x27, 0x54, 0x08, 0xb9, 0xa2, 0xc7, 0x45, 0x26, 0xac, 0x59, 0x3c, 0x89, 0xa4,
	0x48, 0xc5, 0xa2, 0xb6, 0x99, 0x50, 0x46, 0x27, 0x3e, 0x71, 0x6d, 0xc8, 0x52, 0x38, 0xe9, 0x3f,
	0xa2, 0xbf, 0x34, 0xbd, 0x54, 0x42, 0x91, 0x22, 0xcf, 0xfa, 0x34, 0x68, 0x11, 0xff, 0x47, 0xd0,
	0xb1, 0x5d, 0x6b, 0x59, 0x09, 0x59, 0x68, 0xbb, 0x3e, 0xf0, 0x62, 0x04, 0x5f, 0x59, 0x93, 0x97,
	0x6a, 0xc4, 0x81, 0xba, 0x92, 0x46, 0xa4, 0x96, 0x7c, 0x18, 0xf8, 0xf3, 0x57, 0x08, 0x26, 0x97,
	0xb5, 0xcd, 0xc2, 0xb3, 0xe0, 0x78, 0x9e, 0xab, 0x32, 0x2f, 0x6f, 0x88, 0xd4, 0x4a, 0x86, 0x53,
	0x7e, 0xe8, 0xe7, 0x73, 0xdf, 0x1f, 0x4d, 0x79, 0x6a, 0x89, 0x0f, 0xcc, 0x4c, 0x56, 0xba, 0x37,
	0x66, 0xf2, 0xc1, 0xc8, 0x2a, 0xbb, 0xa7, 0xa5, 0x2c, 0x47, 0x18, 0xa7, 0xc1, 0xe4, 0x2e, 0x2f,
	0x55, 0x78, 0xf2, 0x43, 0x5e, 0xf7, 0x5b, 0xa3, 0x5f, 0xef, 0xab, 0xdb, 0x6d, 0x8b, 0xac, 0x69,
	0x91, 0xed, 0x5a, 0x64, 0xfb, 0x16, 0xe1, 0xab, 0x45, 0x78, 0x76, 0x08, 0x6f, 0x0e, 0xe1, 0xdd,
	0x21, 0x6c, 0x1c, 0xc2, 0xd6, 0x21, 0x34, 0x0e, 0xe1, 0xc3, 0x21, 0x7c, 0x3a, 0x64, 0x7b, 0x87,
	0xf0, 0xd2, 0x21, 0xdb, 0x74, 0x08, 0x4d, 0x87, 0x6c, 0xd7, 0x21, 0x8b, 0x8f, 0xfc, 0xca, 0x8b,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x1a, 0xe7, 0x1e, 0xe0, 0x01, 0x00, 0x00,
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
	// Signin Google
	SigninGoogle(ctx context.Context, in *pbtypes.String, opts ...grpc.CallOption) (*dto.SigninResp, error)
	// Refresh token
	RefreshToken(ctx context.Context, in *pbtypes.String, opts ...grpc.CallOption) (*dto.SigninResp, error)
	// Ping
	Ping(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) SigninGoogle(ctx context.Context, in *pbtypes.String, opts ...grpc.CallOption) (*dto.SigninResp, error) {
	out := new(dto.SigninResp)
	err := c.cc.Invoke(ctx, "/grpc.Auth/SigninGoogle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshToken(ctx context.Context, in *pbtypes.String, opts ...grpc.CallOption) (*dto.SigninResp, error) {
	out := new(dto.SigninResp)
	err := c.cc.Invoke(ctx, "/grpc.Auth/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Ping(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error) {
	out := new(pbtypes.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Auth/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// Signin Google
	SigninGoogle(context.Context, *pbtypes.String) (*dto.SigninResp, error)
	// Refresh token
	RefreshToken(context.Context, *pbtypes.String) (*dto.SigninResp, error)
	// Ping
	Ping(context.Context, *pbtypes.Empty) (*pbtypes.Empty, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) SigninGoogle(ctx context.Context, req *pbtypes.String) (*dto.SigninResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigninGoogle not implemented")
}
func (*UnimplementedAuthServer) RefreshToken(ctx context.Context, req *pbtypes.String) (*dto.SigninResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (*UnimplementedAuthServer) Ping(ctx context.Context, req *pbtypes.Empty) (*pbtypes.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_SigninGoogle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SigninGoogle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/SigninGoogle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SigninGoogle(ctx, req.(*pbtypes.String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshToken(ctx, req.(*pbtypes.String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Ping(ctx, req.(*pbtypes.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SigninGoogle",
			Handler:    _Auth_SigninGoogle_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Auth_RefreshToken_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Auth_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/elojah/trax/cmd/auth/grpc/auth.proto",
}
