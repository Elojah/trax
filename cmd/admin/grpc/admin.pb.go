// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/cmd/admin/grpc/admin.proto

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

func init() {
	proto.RegisterFile("github.com/elojah/trax/cmd/admin/grpc/admin.proto", fileDescriptor_7a310f739c32c4bd)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/cmd/admin/grpc/admin.proto", fileDescriptor_7a310f739c32c4bd)
}

var fileDescriptor_7a310f739c32c4bd = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0x3f, 0x4a, 0x3c, 0x31,
	0x14, 0xc0, 0xf1, 0x3c, 0xd8, 0xdf, 0x0f, 0x9c, 0x42, 0x25, 0xe5, 0x14, 0xaf, 0xb4, 0x92, 0x04,
	0xd7, 0x13, 0xa8, 0x08, 0x82, 0x08, 0xa2, 0x78, 0x80, 0xcc, 0x6c, 0xc8, 0xc6, 0x75, 0x26, 0x21,
	0x13, 0xc1, 0xe9, 0x3c, 0x82, 0xde, 0xc2, 0x23, 0x58, 0x6e, 0xb9, 0xe5, 0x94, 0x5b, 0x9a, 0x4c,
	0x63, 0xb9, 0xa5, 0xa5, 0xcc, 0xcc, 0x62, 0xa1, 0xf8, 0xa7, 0x7b, 0x3c, 0x3e, 0x5f, 0x92, 0x97,
	0xec, 0x29, 0xed, 0xa7, 0xb7, 0x19, 0xcb, 0x4d, 0xc1, 0xe5, 0x8d, 0xb9, 0x16, 0x53, 0xee, 0x9d,
	0xb8, 0xe3, 0x79, 0x31, 0xe1, 0x62, 0x52, 0xe8, 0x92, 0x2b, 0x67, 0xf3, 0x61, 0x64, 0xd6, 0x19,
	0x6f, 0xe8, 0xa8, 0xdb, 0xa4, 0xfc, 0x9b, 0xd0, 0xce, 0x14, 0x57, 0x46, 0x99, 0xde, 0xf6, 0xd3,
	0x90, 0xa5, 0xec, 0x87, 0xc0, 0x66, 0xbe, 0xb6, 0xb2, 0xe2, 0xb2, 0xb0, 0xbe, 0x5e, 0x7b, 0xfe,
	0x07, 0x5f, 0x79, 0xa7, 0x4b, 0x35, 0x04, 0xe3, 0x47, 0x48, 0xfe, 0x1d, 0x74, 0xff, 0xa4, 0xbb,
	0xc9, 0xc6, 0x99, 0x56, 0x4e, 0x78, 0x79, 0x65, 0xe9, 0x16, 0x5b, 0x6b, 0x76, 0xd9, 0xeb, 0x74,
	0xf3, 0x63, 0x71, 0xdc, 0x3d, 0x47, 0xc7, 0xc9, 0xf6, 0x85, 0xf1, 0xc2, 0xcb, 0x23, 0x63, 0x66,
	0x5a, 0x9e, 0xca, 0xba, 0xa2, 0x9f, 0xcc, 0x97, 0x66, 0x27, 0x19, 0x9d, 0xeb, 0x52, 0xfd, 0xe6,
	0x0e, 0x4f, 0x16, 0x01, 0x49, 0x13, 0x90, 0x2c, 0x03, 0x92, 0x55, 0x40, 0x78, 0x0b, 0x08, 0xf7,
	0x11, 0xe1, 0x29, 0x22, 0x3c, 0x47, 0x84, 0x79, 0x44, 0x58, 0x44, 0x84, 0x26, 0x22, 0xbc, 0x44,
	0x84, 0xd7, 0x88, 0x64, 0x15, 0x11, 0x1e, 0x5a, 0x24, 0xf3, 0x16, 0xa1, 0x69, 0x91, 0x2c, 0x5b,
	0x24, 0xd9, 0xff, 0xfe, 0xc8, 0xfd, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xd4, 0xd7, 0x93,
	0xb1, 0x01, 0x00, 0x00,
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
	// DB migrations
	MigrateUp(ctx context.Context, in *pbtypes.String, opts ...grpc.CallOption) (*pbtypes.Empty, error)
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

func (c *adminClient) MigrateUp(ctx context.Context, in *pbtypes.String, opts ...grpc.CallOption) (*pbtypes.Empty, error) {
	out := new(pbtypes.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/MigrateUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	// DB migrations
	MigrateUp(context.Context, *pbtypes.String) (*pbtypes.Empty, error)
	// Cookie secure management
	RotateCookieKeys(context.Context, *pbtypes.Empty) (*pbtypes.Empty, error)
	// Ping
	Ping(context.Context, *pbtypes.Empty) (*pbtypes.Empty, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) MigrateUp(ctx context.Context, req *pbtypes.String) (*pbtypes.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateUp not implemented")
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

func _Admin_MigrateUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MigrateUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/MigrateUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MigrateUp(ctx, req.(*pbtypes.String))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "MigrateUp",
			Handler:    _Admin_MigrateUp_Handler,
		},
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
	Metadata: "github.com/elojah/trax/cmd/admin/grpc/admin.proto",
}
