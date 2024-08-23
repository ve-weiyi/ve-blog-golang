// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: remark.proto

// proto 包名

package remarkrpc

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
	RemarkRpc_AddRemark_FullMethodName        = "/remarkrpc.RemarkRpc/AddRemark"
	RemarkRpc_UpdateRemark_FullMethodName     = "/remarkrpc.RemarkRpc/UpdateRemark"
	RemarkRpc_DeleteRemark_FullMethodName     = "/remarkrpc.RemarkRpc/DeleteRemark"
	RemarkRpc_DeleteRemarkList_FullMethodName = "/remarkrpc.RemarkRpc/DeleteRemarkList"
	RemarkRpc_FindRemark_FullMethodName       = "/remarkrpc.RemarkRpc/FindRemark"
	RemarkRpc_FindRemarkList_FullMethodName   = "/remarkrpc.RemarkRpc/FindRemarkList"
)

// RemarkRpcClient is the client API for RemarkRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemarkRpcClient interface {
	// 创建留言
	AddRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error)
	// 更新留言
	UpdateRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error)
	// 删除留言
	DeleteRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 批量删除留言
	DeleteRemarkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询留言
	FindRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RemarkDetails, error)
	// 查询留言列表
	FindRemarkList(ctx context.Context, in *FindRemarkListReq, opts ...grpc.CallOption) (*FindRemarkListResp, error)
}

type remarkRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRemarkRpcClient(cc grpc.ClientConnInterface) RemarkRpcClient {
	return &remarkRpcClient{cc}
}

func (c *remarkRpcClient) AddRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error) {
	out := new(RemarkDetails)
	err := c.cc.Invoke(ctx, RemarkRpc_AddRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remarkRpcClient) UpdateRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error) {
	out := new(RemarkDetails)
	err := c.cc.Invoke(ctx, RemarkRpc_UpdateRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remarkRpcClient) DeleteRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, RemarkRpc_DeleteRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remarkRpcClient) DeleteRemarkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, RemarkRpc_DeleteRemarkList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remarkRpcClient) FindRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RemarkDetails, error) {
	out := new(RemarkDetails)
	err := c.cc.Invoke(ctx, RemarkRpc_FindRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remarkRpcClient) FindRemarkList(ctx context.Context, in *FindRemarkListReq, opts ...grpc.CallOption) (*FindRemarkListResp, error) {
	out := new(FindRemarkListResp)
	err := c.cc.Invoke(ctx, RemarkRpc_FindRemarkList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemarkRpcServer is the server API for RemarkRpc service.
// All implementations must embed UnimplementedRemarkRpcServer
// for forward compatibility
type RemarkRpcServer interface {
	// 创建留言
	AddRemark(context.Context, *RemarkNew) (*RemarkDetails, error)
	// 更新留言
	UpdateRemark(context.Context, *RemarkNew) (*RemarkDetails, error)
	// 删除留言
	DeleteRemark(context.Context, *IdReq) (*BatchResp, error)
	// 批量删除留言
	DeleteRemarkList(context.Context, *IdsReq) (*BatchResp, error)
	// 查询留言
	FindRemark(context.Context, *IdReq) (*RemarkDetails, error)
	// 查询留言列表
	FindRemarkList(context.Context, *FindRemarkListReq) (*FindRemarkListResp, error)
	mustEmbedUnimplementedRemarkRpcServer()
}

// UnimplementedRemarkRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRemarkRpcServer struct {
}

func (UnimplementedRemarkRpcServer) AddRemark(context.Context, *RemarkNew) (*RemarkDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRemark not implemented")
}
func (UnimplementedRemarkRpcServer) UpdateRemark(context.Context, *RemarkNew) (*RemarkDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRemark not implemented")
}
func (UnimplementedRemarkRpcServer) DeleteRemark(context.Context, *IdReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRemark not implemented")
}
func (UnimplementedRemarkRpcServer) DeleteRemarkList(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRemarkList not implemented")
}
func (UnimplementedRemarkRpcServer) FindRemark(context.Context, *IdReq) (*RemarkDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRemark not implemented")
}
func (UnimplementedRemarkRpcServer) FindRemarkList(context.Context, *FindRemarkListReq) (*FindRemarkListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRemarkList not implemented")
}
func (UnimplementedRemarkRpcServer) mustEmbedUnimplementedRemarkRpcServer() {}

// UnsafeRemarkRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemarkRpcServer will
// result in compilation errors.
type UnsafeRemarkRpcServer interface {
	mustEmbedUnimplementedRemarkRpcServer()
}

func RegisterRemarkRpcServer(s grpc.ServiceRegistrar, srv RemarkRpcServer) {
	s.RegisterService(&RemarkRpc_ServiceDesc, srv)
}

func _RemarkRpc_AddRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemarkNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarkRpcServer).AddRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemarkRpc_AddRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarkRpcServer).AddRemark(ctx, req.(*RemarkNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemarkRpc_UpdateRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemarkNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarkRpcServer).UpdateRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemarkRpc_UpdateRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarkRpcServer).UpdateRemark(ctx, req.(*RemarkNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemarkRpc_DeleteRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarkRpcServer).DeleteRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemarkRpc_DeleteRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarkRpcServer).DeleteRemark(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemarkRpc_DeleteRemarkList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarkRpcServer).DeleteRemarkList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemarkRpc_DeleteRemarkList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarkRpcServer).DeleteRemarkList(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemarkRpc_FindRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarkRpcServer).FindRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemarkRpc_FindRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarkRpcServer).FindRemark(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemarkRpc_FindRemarkList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRemarkListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarkRpcServer).FindRemarkList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemarkRpc_FindRemarkList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarkRpcServer).FindRemarkList(ctx, req.(*FindRemarkListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RemarkRpc_ServiceDesc is the grpc.ServiceDesc for RemarkRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemarkRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remarkrpc.RemarkRpc",
	HandlerType: (*RemarkRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRemark",
			Handler:    _RemarkRpc_AddRemark_Handler,
		},
		{
			MethodName: "UpdateRemark",
			Handler:    _RemarkRpc_UpdateRemark_Handler,
		},
		{
			MethodName: "DeleteRemark",
			Handler:    _RemarkRpc_DeleteRemark_Handler,
		},
		{
			MethodName: "DeleteRemarkList",
			Handler:    _RemarkRpc_DeleteRemarkList_Handler,
		},
		{
			MethodName: "FindRemark",
			Handler:    _RemarkRpc_FindRemark_Handler,
		},
		{
			MethodName: "FindRemarkList",
			Handler:    _RemarkRpc_FindRemarkList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remark.proto",
}
