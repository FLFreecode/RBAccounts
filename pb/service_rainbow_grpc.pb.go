// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: service_rainbow.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RainbowUser_CreateUser_FullMethodName = "/pb.rainbowUser/CreateUser"
	RainbowUser_LoginUser_FullMethodName  = "/pb.rainbowUser/LoginUser"
)

// RainbowUserClient is the client API for RainbowUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RainbowUserClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type rainbowUserClient struct {
	cc grpc.ClientConnInterface
}

func NewRainbowUserClient(cc grpc.ClientConnInterface) RainbowUserClient {
	return &rainbowUserClient{cc}
}

func (c *rainbowUserClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, RainbowUser_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rainbowUserClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, RainbowUser_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RainbowUserServer is the server API for RainbowUser service.
// All implementations must embed UnimplementedRainbowUserServer
// for forward compatibility
type RainbowUserServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedRainbowUserServer()
}

// UnimplementedRainbowUserServer must be embedded to have forward compatible implementations.
type UnimplementedRainbowUserServer struct {
}

func (UnimplementedRainbowUserServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedRainbowUserServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedRainbowUserServer) mustEmbedUnimplementedRainbowUserServer() {}

// UnsafeRainbowUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RainbowUserServer will
// result in compilation errors.
type UnsafeRainbowUserServer interface {
	mustEmbedUnimplementedRainbowUserServer()
}

func RegisterRainbowUserServer(s grpc.ServiceRegistrar, srv RainbowUserServer) {
	s.RegisterService(&RainbowUser_ServiceDesc, srv)
}

func _RainbowUser_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RainbowUserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RainbowUser_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RainbowUserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RainbowUser_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RainbowUserServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RainbowUser_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RainbowUserServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RainbowUser_ServiceDesc is the grpc.ServiceDesc for RainbowUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RainbowUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.rainbowUser",
	HandlerType: (*RainbowUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _RainbowUser_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _RainbowUser_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_rainbow.proto",
}