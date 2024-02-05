// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/cmd/api/grpc/api.proto

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
	proto.RegisterFile("github.com/elojah/trax/cmd/api/grpc/api.proto", fileDescriptor_e44be87dd4a64048)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/cmd/api/grpc/api.proto", fileDescriptor_e44be87dd4a64048)
}

var fileDescriptor_e44be87dd4a64048 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8d, 0xa1, 0x4e, 0xc4, 0x40,
	0x14, 0x45, 0xe7, 0x85, 0x0d, 0xa2, 0x02, 0x51, 0x59, 0x71, 0x25, 0x6e, 0xe7, 0x25, 0xf0, 0x05,
	0x90, 0x90, 0x80, 0xdb, 0x5f, 0x68, 0x4b, 0x33, 0x5b, 0xa0, 0xcc, 0xa4, 0x0c, 0x09, 0xeb, 0xf8,
	0x04, 0x3e, 0x83, 0x4f, 0x40, 0xae, 0xac, 0xac, 0xac, 0x64, 0xde, 0x18, 0x64, 0x25, 0x92, 0xb4,
	0x45, 0x91, 0xb0, 0xee, 0x9e, 0xe4, 0x9c, 0xdc, 0x64, 0x6d, 0x6a, 0xbf, 0x7d, 0x2e, 0x74, 0x69,
	0x1b, 0xae, 0x1e, 0xec, 0x5d, 0xbe, 0x65, 0xdf, 0xe6, 0x2f, 0x5c, 0x36, 0xb7, 0x9c, 0xbb, 0x9a,
	0x4d, 0xeb, 0xca, 0x69, 0x68, 0xd7, 0x5a, 0x6f, 0xd3, 0xd5, 0xc4, 0x19, 0xff, 0x13, 0xb9, 0x7b,
	0xc3, 0xc6, 0x1a, 0x3b, 0xbb, 0xf3, 0x5a, 0xb2, 0x4c, 0x1f, 0x08, 0x5c, 0xe1, 0x77, 0xae, 0x7a,
	0xe2, 0xaa, 0x71, 0x7e, 0xb7, 0xf8, 0x67, 0xeb, 0xe4, 0xe8, 0x62, 0x73, 0x93, 0x9e, 0x26, 0xab,
	0x4d, 0xfd, 0x68, 0xd2, 0x13, 0xfd, 0x2b, 0xe9, 0xab, 0x49, 0xca, 0xfe, 0xf0, 0xe5, 0x75, 0x17,
	0xa0, 0xfa, 0x00, 0x35, 0x04, 0xa8, 0x31, 0x80, 0xbe, 0x03, 0xe8, 0x55, 0x40, 0xef, 0x02, 0xfa,
	0x10, 0xd0, 0x5e, 0x40, 0x9d, 0x80, 0x7a, 0x01, 0x7d, 0x0a, 0xe8, 0x4b, 0xa0, 0x46, 0x01, 0xbd,
	0x45, 0xa8, 0x7d, 0x04, 0xf5, 0x11, 0x6a, 0x88, 0x50, 0xc5, 0xf1, 0xfc, 0x7f, 0xfe, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x59, 0x07, 0xd3, 0x2a, 0x17, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	// Ping
	Ping(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Ping(ctx context.Context, in *pbtypes.Empty, opts ...grpc.CallOption) (*pbtypes.Empty, error) {
	out := new(pbtypes.Empty)
	err := c.cc.Invoke(ctx, "/grpc.API/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	// Ping
	Ping(context.Context, *pbtypes.Empty) (*pbtypes.Empty, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) Ping(ctx context.Context, req *pbtypes.Empty) (*pbtypes.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Ping(ctx, req.(*pbtypes.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _API_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/elojah/trax/cmd/api/grpc/api.proto",
}
