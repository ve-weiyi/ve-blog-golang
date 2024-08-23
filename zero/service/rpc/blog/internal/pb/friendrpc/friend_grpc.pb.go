// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: friend.proto

// proto 包名

package friendrpc

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
	FriendRpc_AddFriend_FullMethodName        = "/friendrpc.FriendRpc/AddFriend"
	FriendRpc_UpdateFriend_FullMethodName     = "/friendrpc.FriendRpc/UpdateFriend"
	FriendRpc_DeleteFriend_FullMethodName     = "/friendrpc.FriendRpc/DeleteFriend"
	FriendRpc_DeleteFriendList_FullMethodName = "/friendrpc.FriendRpc/DeleteFriendList"
	FriendRpc_FindFriendList_FullMethodName   = "/friendrpc.FriendRpc/FindFriendList"
)

// FriendRpcClient is the client API for FriendRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendRpcClient interface {
	// 创建友链
	AddFriend(ctx context.Context, in *FriendNew, opts ...grpc.CallOption) (*FriendDetails, error)
	// 更新友链
	UpdateFriend(ctx context.Context, in *FriendNew, opts ...grpc.CallOption) (*FriendDetails, error)
	// 删除友链
	DeleteFriend(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 批量删除友链
	DeleteFriendList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询友链列表
	FindFriendList(ctx context.Context, in *FindFriendListReq, opts ...grpc.CallOption) (*FindFriendListResp, error)
}

type friendRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendRpcClient(cc grpc.ClientConnInterface) FriendRpcClient {
	return &friendRpcClient{cc}
}

func (c *friendRpcClient) AddFriend(ctx context.Context, in *FriendNew, opts ...grpc.CallOption) (*FriendDetails, error) {
	out := new(FriendDetails)
	err := c.cc.Invoke(ctx, FriendRpc_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendRpcClient) UpdateFriend(ctx context.Context, in *FriendNew, opts ...grpc.CallOption) (*FriendDetails, error) {
	out := new(FriendDetails)
	err := c.cc.Invoke(ctx, FriendRpc_UpdateFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendRpcClient) DeleteFriend(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, FriendRpc_DeleteFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendRpcClient) DeleteFriendList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, FriendRpc_DeleteFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendRpcClient) FindFriendList(ctx context.Context, in *FindFriendListReq, opts ...grpc.CallOption) (*FindFriendListResp, error) {
	out := new(FindFriendListResp)
	err := c.cc.Invoke(ctx, FriendRpc_FindFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendRpcServer is the server API for FriendRpc service.
// All implementations must embed UnimplementedFriendRpcServer
// for forward compatibility
type FriendRpcServer interface {
	// 创建友链
	AddFriend(context.Context, *FriendNew) (*FriendDetails, error)
	// 更新友链
	UpdateFriend(context.Context, *FriendNew) (*FriendDetails, error)
	// 删除友链
	DeleteFriend(context.Context, *IdReq) (*BatchResp, error)
	// 批量删除友链
	DeleteFriendList(context.Context, *IdsReq) (*BatchResp, error)
	// 查询友链列表
	FindFriendList(context.Context, *FindFriendListReq) (*FindFriendListResp, error)
	mustEmbedUnimplementedFriendRpcServer()
}

// UnimplementedFriendRpcServer must be embedded to have forward compatible implementations.
type UnimplementedFriendRpcServer struct {
}

func (UnimplementedFriendRpcServer) AddFriend(context.Context, *FriendNew) (*FriendDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendRpcServer) UpdateFriend(context.Context, *FriendNew) (*FriendDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriend not implemented")
}
func (UnimplementedFriendRpcServer) DeleteFriend(context.Context, *IdReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendRpcServer) DeleteFriendList(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriendList not implemented")
}
func (UnimplementedFriendRpcServer) FindFriendList(context.Context, *FindFriendListReq) (*FindFriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFriendList not implemented")
}
func (UnimplementedFriendRpcServer) mustEmbedUnimplementedFriendRpcServer() {}

// UnsafeFriendRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendRpcServer will
// result in compilation errors.
type UnsafeFriendRpcServer interface {
	mustEmbedUnimplementedFriendRpcServer()
}

func RegisterFriendRpcServer(s grpc.ServiceRegistrar, srv FriendRpcServer) {
	s.RegisterService(&FriendRpc_ServiceDesc, srv)
}

func _FriendRpc_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendRpcServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendRpc_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendRpcServer).AddFriend(ctx, req.(*FriendNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendRpc_UpdateFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendRpcServer).UpdateFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendRpc_UpdateFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendRpcServer).UpdateFriend(ctx, req.(*FriendNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendRpc_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendRpcServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendRpc_DeleteFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendRpcServer).DeleteFriend(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendRpc_DeleteFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendRpcServer).DeleteFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendRpc_DeleteFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendRpcServer).DeleteFriendList(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendRpc_FindFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendRpcServer).FindFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendRpc_FindFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendRpcServer).FindFriendList(ctx, req.(*FindFriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendRpc_ServiceDesc is the grpc.ServiceDesc for FriendRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "friendrpc.FriendRpc",
	HandlerType: (*FriendRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _FriendRpc_AddFriend_Handler,
		},
		{
			MethodName: "UpdateFriend",
			Handler:    _FriendRpc_UpdateFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _FriendRpc_DeleteFriend_Handler,
		},
		{
			MethodName: "DeleteFriendList",
			Handler:    _FriendRpc_DeleteFriendList_Handler,
		},
		{
			MethodName: "FindFriendList",
			Handler:    _FriendRpc_FindFriendList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend.proto",
}
