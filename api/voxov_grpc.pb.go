// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: voxov.proto

package api

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VOxOVClient is the client API for VOxOV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VOxOVClient interface {
}

type vOxOVClient struct {
	cc grpc.ClientConnInterface
}

func NewVOxOVClient(cc grpc.ClientConnInterface) VOxOVClient {
	return &vOxOVClient{cc}
}

// VOxOVServer is the server API for VOxOV service.
// All implementations must embed UnimplementedVOxOVServer
// for forward compatibility
type VOxOVServer interface {
	mustEmbedUnimplementedVOxOVServer()
}

// UnimplementedVOxOVServer must be embedded to have forward compatible implementations.
type UnimplementedVOxOVServer struct {
}

func (UnimplementedVOxOVServer) mustEmbedUnimplementedVOxOVServer() {}

// UnsafeVOxOVServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VOxOVServer will
// result in compilation errors.
type UnsafeVOxOVServer interface {
	mustEmbedUnimplementedVOxOVServer()
}

func RegisterVOxOVServer(s grpc.ServiceRegistrar, srv VOxOVServer) {
	s.RegisterService(&VOxOV_ServiceDesc, srv)
}

// VOxOV_ServiceDesc is the grpc.ServiceDesc for VOxOV service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VOxOV_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voxov.VOxOV",
	HandlerType: (*VOxOVServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "voxov.proto",
}
