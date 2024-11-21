// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: page.proto

// proto 包名

package pagerpc

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
	PageRpc_AddPage_FullMethodName      = "/pagerpc.PageRpc/AddPage"
	PageRpc_UpdatePage_FullMethodName   = "/pagerpc.PageRpc/UpdatePage"
	PageRpc_DeletePage_FullMethodName   = "/pagerpc.PageRpc/DeletePage"
	PageRpc_FindPageList_FullMethodName = "/pagerpc.PageRpc/FindPageList"
)

// PageRpcClient is the client API for PageRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PageRpcClient interface {
	// 创建页面
	AddPage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error)
	// 更新页面
	UpdatePage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error)
	// 删除页面
	DeletePage(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询页面列表
	FindPageList(ctx context.Context, in *FindPageListReq, opts ...grpc.CallOption) (*FindPageListResp, error)
}

type pageRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewPageRpcClient(cc grpc.ClientConnInterface) PageRpcClient {
	return &pageRpcClient{cc}
}

func (c *pageRpcClient) AddPage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error) {
	out := new(PageDetails)
	err := c.cc.Invoke(ctx, PageRpc_AddPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageRpcClient) UpdatePage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error) {
	out := new(PageDetails)
	err := c.cc.Invoke(ctx, PageRpc_UpdatePage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageRpcClient) DeletePage(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, PageRpc_DeletePage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageRpcClient) FindPageList(ctx context.Context, in *FindPageListReq, opts ...grpc.CallOption) (*FindPageListResp, error) {
	out := new(FindPageListResp)
	err := c.cc.Invoke(ctx, PageRpc_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PageRpcServer is the server API for PageRpc service.
// All implementations must embed UnimplementedPageRpcServer
// for forward compatibility
type PageRpcServer interface {
	// 创建页面
	AddPage(context.Context, *PageNewReq) (*PageDetails, error)
	// 更新页面
	UpdatePage(context.Context, *PageNewReq) (*PageDetails, error)
	// 删除页面
	DeletePage(context.Context, *IdsReq) (*BatchResp, error)
	// 查询页面列表
	FindPageList(context.Context, *FindPageListReq) (*FindPageListResp, error)
	mustEmbedUnimplementedPageRpcServer()
}

// UnimplementedPageRpcServer must be embedded to have forward compatible implementations.
type UnimplementedPageRpcServer struct {
}

func (UnimplementedPageRpcServer) AddPage(context.Context, *PageNewReq) (*PageDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPage not implemented")
}
func (UnimplementedPageRpcServer) UpdatePage(context.Context, *PageNewReq) (*PageDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePage not implemented")
}
func (UnimplementedPageRpcServer) DeletePage(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePage not implemented")
}
func (UnimplementedPageRpcServer) FindPageList(context.Context, *FindPageListReq) (*FindPageListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedPageRpcServer) mustEmbedUnimplementedPageRpcServer() {}

// UnsafePageRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PageRpcServer will
// result in compilation errors.
type UnsafePageRpcServer interface {
	mustEmbedUnimplementedPageRpcServer()
}

func RegisterPageRpcServer(s grpc.ServiceRegistrar, srv PageRpcServer) {
	s.RegisterService(&PageRpc_ServiceDesc, srv)
}

func _PageRpc_AddPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageRpcServer).AddPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageRpc_AddPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageRpcServer).AddPage(ctx, req.(*PageNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageRpc_UpdatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageRpcServer).UpdatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageRpc_UpdatePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageRpcServer).UpdatePage(ctx, req.(*PageNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageRpc_DeletePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageRpcServer).DeletePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageRpc_DeletePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageRpcServer).DeletePage(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageRpc_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPageListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageRpcServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageRpc_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageRpcServer).FindPageList(ctx, req.(*FindPageListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PageRpc_ServiceDesc is the grpc.ServiceDesc for PageRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PageRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pagerpc.PageRpc",
	HandlerType: (*PageRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPage",
			Handler:    _PageRpc_AddPage_Handler,
		},
		{
			MethodName: "UpdatePage",
			Handler:    _PageRpc_UpdatePage_Handler,
		},
		{
			MethodName: "DeletePage",
			Handler:    _PageRpc_DeletePage_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _PageRpc_FindPageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "page.proto",
}
