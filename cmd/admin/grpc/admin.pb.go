// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmd/admin/grpc/admin.proto

package grpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/elojah/trax/pkg/gogoproto"
	pbtypes "github.com/elojah/trax/pkg/pbtypes"
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

func init() { proto.RegisterFile("cmd/admin/grpc/admin.proto", fileDescriptor_6b45207160ef9c5f) }
func init() { golang_proto.RegisterFile("cmd/admin/grpc/admin.proto", fileDescriptor_6b45207160ef9c5f) }

var fileDescriptor_6b45207160ef9c5f = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x21, 0x4f, 0xc4, 0x40,
	0x10, 0x85, 0x77, 0x92, 0x03, 0x51, 0x41, 0x48, 0x0d, 0xa4, 0xe2, 0x39, 0x90, 0xdd, 0xe4, 0xf8,
	0x05, 0x40, 0x50, 0x18, 0x82, 0xc4, 0xb5, 0xbd, 0xcd, 0x5e, 0x39, 0x7a, 0xb3, 0xe9, 0x2d, 0x09,
	0x75, 0xfc, 0x04, 0x7e, 0x06, 0x3f, 0x01, 0x79, 0xf2, 0x64, 0xe5, 0x49, 0x76, 0x6a, 0x90, 0x27,
	0x91, 0xa4, 0x2d, 0x06, 0xcc, 0xb9, 0x37, 0xef, 0x9b, 0x37, 0xc9, 0xbc, 0x28, 0x29, 0xaa, 0x99,
	0xce, 0x66, 0x55, 0xb9, 0xd4, 0xb6, 0x76, 0xc5, 0x28, 0x53, 0x57, 0xb3, 0xe7, 0x78, 0xd2, 0x3b,
	0xc9, 0xa9, 0x5b, 0x58, 0x6d, 0xd9, 0xf2, 0x60, 0x0e, 0x6a, 0xe4, 0xc9, 0x49, 0x4f, 0x5c, 0xee,
	0x1b, 0x67, 0x56, 0xda, 0x54, 0xce, 0x37, 0x23, 0x98, 0x16, 0xd1, 0xc1, 0x65, 0x7f, 0x27, 0x9e,
	0x46, 0xc7, 0xf7, 0xec, 0x33, 0x6f, 0xae, 0x99, 0x17, 0xa5, 0xb9, 0x35, 0xcd, 0x2a, 0x3e, 0x4a,
	0x7f, 0x23, 0xe9, 0x4d, 0x1f, 0x49, 0xfe, 0xcd, 0xf1, 0x79, 0x34, 0xb9, 0x2b, 0x97, 0x76, 0xdf,
	0xde, 0x15, 0x6f, 0x02, 0x54, 0x1b, 0xa0, 0xb6, 0x01, 0x6a, 0x17, 0x40, 0xdf, 0x01, 0xf4, 0x2a,
	0xa0, 0x77, 0x01, 0x7d, 0x08, 0x68, 0x2d, 0xa0, 0x8d, 0x80, 0x5a, 0x01, 0x7d, 0x0a, 0xe8, 0x4b,
	0xa0, 0x76, 0x02, 0x7a, 0xeb, 0xa0, 0xd6, 0x1d, 0xa8, 0xed, 0xa0, 0xb6, 0x1d, 0xd4, 0xc3, 0x99,
	0x2d, 0xfd, 0xfc, 0x39, 0x4f, 0x0b, 0xae, 0xb4, 0x79, 0xe2, 0xc7, 0x6c, 0xae, 0x7d, 0x9d, 0xbd,
	0xe8, 0xbf, 0xd5, 0xe4, 0x87, 0xc3, 0x73, 0x17, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x81,
	0x7d, 0x49, 0x33, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	// Cookie secure management
	RotateCookieKeys(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error)
	// Ping
	Ping(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) RotateCookieKeys(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error) {
	out := new(pbtypes.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/RotateCookieKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Ping(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error) {
	out := new(pbtypes.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	// Cookie secure management
	RotateCookieKeys(context.Context, *pbtypes.Empty) (*pbtypes.Empty, error)
	// Ping
	Ping(context.Context, *pbtypes.Empty) (*pbtypes.Empty, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) RotateCookieKeys(ctx context.Context, req *pbtypes.Empty) (*pbtypes.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateCookieKeys not implemented")
}
func (*UnimplementedAdminServer) Ping(ctx context.Context, req *pbtypes.Empty) (*pbtypes.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_RotateCookieKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RotateCookieKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/RotateCookieKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RotateCookieKeys(ctx, req.(*pbtypes.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Ping(ctx, req.(*pbtypes.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RotateCookieKeys",
			Handler:    _Admin_RotateCookieKeys_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Admin_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/admin/grpc/admin.proto",
}
