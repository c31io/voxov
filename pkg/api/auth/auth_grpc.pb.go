// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auth.proto

package api

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// sms auth
	WaitMsg(ctx context.Context, in *WaitMsgRequest, opts ...grpc.CallOption) (*WaitMsgReply, error)
	CheckMsg(ctx context.Context, in *CheckMsgRequest, opts ...grpc.CallOption) (*CheckMsgReply, error)
	// device crud
	NewDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
	GetDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
	EditDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
	RemoveDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
	AllDevice(ctx context.Context, in *AllDeviceRequest, opts ...grpc.CallOption) (*AllDeviceReply, error)
	// device auth
	CheckDevice(ctx context.Context, in *CheckDeviceRequest, opts ...grpc.CallOption) (*CheckDeviceReply, error)
	// person read
	GetPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) WaitMsg(ctx context.Context, in *WaitMsgRequest, opts ...grpc.CallOption) (*WaitMsgReply, error) {
	out := new(WaitMsgReply)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/WaitMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckMsg(ctx context.Context, in *CheckMsgRequest, opts ...grpc.CallOption) (*CheckMsgReply, error) {
	out := new(CheckMsgReply)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/CheckMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) NewDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/NewDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) EditDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/EditDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RemoveDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/RemoveDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AllDevice(ctx context.Context, in *AllDeviceRequest, opts ...grpc.CallOption) (*AllDeviceReply, error) {
	out := new(AllDeviceReply)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/AllDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckDevice(ctx context.Context, in *CheckDeviceRequest, opts ...grpc.CallOption) (*CheckDeviceReply, error) {
	out := new(CheckDeviceReply)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/CheckDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/voxov_auth.Auth/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// sms auth
	WaitMsg(context.Context, *WaitMsgRequest) (*WaitMsgReply, error)
	CheckMsg(context.Context, *CheckMsgRequest) (*CheckMsgReply, error)
	// device crud
	NewDevice(context.Context, *Device) (*Device, error)
	GetDevice(context.Context, *Device) (*Device, error)
	EditDevice(context.Context, *Device) (*Device, error)
	RemoveDevice(context.Context, *Device) (*Device, error)
	AllDevice(context.Context, *AllDeviceRequest) (*AllDeviceReply, error)
	// device auth
	CheckDevice(context.Context, *CheckDeviceRequest) (*CheckDeviceReply, error)
	// person read
	GetPerson(context.Context, *Person) (*Person, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) WaitMsg(context.Context, *WaitMsgRequest) (*WaitMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitMsg not implemented")
}
func (UnimplementedAuthServer) CheckMsg(context.Context, *CheckMsgRequest) (*CheckMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckMsg not implemented")
}
func (UnimplementedAuthServer) NewDevice(context.Context, *Device) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewDevice not implemented")
}
func (UnimplementedAuthServer) GetDevice(context.Context, *Device) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedAuthServer) EditDevice(context.Context, *Device) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDevice not implemented")
}
func (UnimplementedAuthServer) RemoveDevice(context.Context, *Device) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevice not implemented")
}
func (UnimplementedAuthServer) AllDevice(context.Context, *AllDeviceRequest) (*AllDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllDevice not implemented")
}
func (UnimplementedAuthServer) CheckDevice(context.Context, *CheckDeviceRequest) (*CheckDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDevice not implemented")
}
func (UnimplementedAuthServer) GetPerson(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_WaitMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).WaitMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/WaitMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).WaitMsg(ctx, req.(*WaitMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/CheckMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckMsg(ctx, req.(*CheckMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_NewDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).NewDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/NewDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).NewDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_EditDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).EditDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/EditDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).EditDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RemoveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RemoveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/RemoveDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RemoveDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AllDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AllDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/AllDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AllDevice(ctx, req.(*AllDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/CheckDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckDevice(ctx, req.(*CheckDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voxov_auth.Auth/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetPerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voxov_auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WaitMsg",
			Handler:    _Auth_WaitMsg_Handler,
		},
		{
			MethodName: "CheckMsg",
			Handler:    _Auth_CheckMsg_Handler,
		},
		{
			MethodName: "NewDevice",
			Handler:    _Auth_NewDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _Auth_GetDevice_Handler,
		},
		{
			MethodName: "EditDevice",
			Handler:    _Auth_EditDevice_Handler,
		},
		{
			MethodName: "RemoveDevice",
			Handler:    _Auth_RemoveDevice_Handler,
		},
		{
			MethodName: "AllDevice",
			Handler:    _Auth_AllDevice_Handler,
		},
		{
			MethodName: "CheckDevice",
			Handler:    _Auth_CheckDevice_Handler,
		},
		{
			MethodName: "GetPerson",
			Handler:    _Auth_GetPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
