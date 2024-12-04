// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: website.proto

// proto 包名

package websiterpc

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
	WebsiteRpc_GetUserDailyVisit_FullMethodName = "/websiterpc.WebsiteRpc/GetUserDailyVisit"
	WebsiteRpc_GetUserTotalVisit_FullMethodName = "/websiterpc.WebsiteRpc/GetUserTotalVisit"
	WebsiteRpc_AddPage_FullMethodName           = "/websiterpc.WebsiteRpc/AddPage"
	WebsiteRpc_UpdatePage_FullMethodName        = "/websiterpc.WebsiteRpc/UpdatePage"
	WebsiteRpc_DeletePage_FullMethodName        = "/websiterpc.WebsiteRpc/DeletePage"
	WebsiteRpc_FindPageList_FullMethodName      = "/websiterpc.WebsiteRpc/FindPageList"
	WebsiteRpc_AddFriend_FullMethodName         = "/websiterpc.WebsiteRpc/AddFriend"
	WebsiteRpc_UpdateFriend_FullMethodName      = "/websiterpc.WebsiteRpc/UpdateFriend"
	WebsiteRpc_DeleteFriend_FullMethodName      = "/websiterpc.WebsiteRpc/DeleteFriend"
	WebsiteRpc_FindFriendList_FullMethodName    = "/websiterpc.WebsiteRpc/FindFriendList"
)

// WebsiteRpcClient is the client API for WebsiteRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebsiteRpcClient interface {
	// 用户日浏览量分析
	GetUserDailyVisit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*UserDailyVisitRsp, error)
	// 用户总流量数
	GetUserTotalVisit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*CountResp, error)
	// 创建页面
	AddPage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error)
	// 更新页面
	UpdatePage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error)
	// 删除页面
	DeletePage(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询页面列表
	FindPageList(ctx context.Context, in *FindPageListReq, opts ...grpc.CallOption) (*FindPageListResp, error)
	// 创建友链
	AddFriend(ctx context.Context, in *FriendNewReq, opts ...grpc.CallOption) (*FriendDetails, error)
	// 更新友链
	UpdateFriend(ctx context.Context, in *FriendNewReq, opts ...grpc.CallOption) (*FriendDetails, error)
	// 删除友链
	DeleteFriend(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询友链列表
	FindFriendList(ctx context.Context, in *FindFriendListReq, opts ...grpc.CallOption) (*FindFriendListResp, error)
}

type websiteRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewWebsiteRpcClient(cc grpc.ClientConnInterface) WebsiteRpcClient {
	return &websiteRpcClient{cc}
}

func (c *websiteRpcClient) GetUserDailyVisit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*UserDailyVisitRsp, error) {
	out := new(UserDailyVisitRsp)
	err := c.cc.Invoke(ctx, WebsiteRpc_GetUserDailyVisit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) GetUserTotalVisit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*CountResp, error) {
	out := new(CountResp)
	err := c.cc.Invoke(ctx, WebsiteRpc_GetUserTotalVisit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) AddPage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error) {
	out := new(PageDetails)
	err := c.cc.Invoke(ctx, WebsiteRpc_AddPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) UpdatePage(ctx context.Context, in *PageNewReq, opts ...grpc.CallOption) (*PageDetails, error) {
	out := new(PageDetails)
	err := c.cc.Invoke(ctx, WebsiteRpc_UpdatePage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) DeletePage(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, WebsiteRpc_DeletePage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) FindPageList(ctx context.Context, in *FindPageListReq, opts ...grpc.CallOption) (*FindPageListResp, error) {
	out := new(FindPageListResp)
	err := c.cc.Invoke(ctx, WebsiteRpc_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) AddFriend(ctx context.Context, in *FriendNewReq, opts ...grpc.CallOption) (*FriendDetails, error) {
	out := new(FriendDetails)
	err := c.cc.Invoke(ctx, WebsiteRpc_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) UpdateFriend(ctx context.Context, in *FriendNewReq, opts ...grpc.CallOption) (*FriendDetails, error) {
	out := new(FriendDetails)
	err := c.cc.Invoke(ctx, WebsiteRpc_UpdateFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) DeleteFriend(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, WebsiteRpc_DeleteFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteRpcClient) FindFriendList(ctx context.Context, in *FindFriendListReq, opts ...grpc.CallOption) (*FindFriendListResp, error) {
	out := new(FindFriendListResp)
	err := c.cc.Invoke(ctx, WebsiteRpc_FindFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebsiteRpcServer is the server API for WebsiteRpc service.
// All implementations must embed UnimplementedWebsiteRpcServer
// for forward compatibility
type WebsiteRpcServer interface {
	// 用户日浏览量分析
	GetUserDailyVisit(context.Context, *EmptyReq) (*UserDailyVisitRsp, error)
	// 用户总流量数
	GetUserTotalVisit(context.Context, *EmptyReq) (*CountResp, error)
	// 创建页面
	AddPage(context.Context, *PageNewReq) (*PageDetails, error)
	// 更新页面
	UpdatePage(context.Context, *PageNewReq) (*PageDetails, error)
	// 删除页面
	DeletePage(context.Context, *IdsReq) (*BatchResp, error)
	// 查询页面列表
	FindPageList(context.Context, *FindPageListReq) (*FindPageListResp, error)
	// 创建友链
	AddFriend(context.Context, *FriendNewReq) (*FriendDetails, error)
	// 更新友链
	UpdateFriend(context.Context, *FriendNewReq) (*FriendDetails, error)
	// 删除友链
	DeleteFriend(context.Context, *IdsReq) (*BatchResp, error)
	// 查询友链列表
	FindFriendList(context.Context, *FindFriendListReq) (*FindFriendListResp, error)
	mustEmbedUnimplementedWebsiteRpcServer()
}

// UnimplementedWebsiteRpcServer must be embedded to have forward compatible implementations.
type UnimplementedWebsiteRpcServer struct {
}

func (UnimplementedWebsiteRpcServer) GetUserDailyVisit(context.Context, *EmptyReq) (*UserDailyVisitRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDailyVisit not implemented")
}
func (UnimplementedWebsiteRpcServer) GetUserTotalVisit(context.Context, *EmptyReq) (*CountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTotalVisit not implemented")
}
func (UnimplementedWebsiteRpcServer) AddPage(context.Context, *PageNewReq) (*PageDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPage not implemented")
}
func (UnimplementedWebsiteRpcServer) UpdatePage(context.Context, *PageNewReq) (*PageDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePage not implemented")
}
func (UnimplementedWebsiteRpcServer) DeletePage(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePage not implemented")
}
func (UnimplementedWebsiteRpcServer) FindPageList(context.Context, *FindPageListReq) (*FindPageListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedWebsiteRpcServer) AddFriend(context.Context, *FriendNewReq) (*FriendDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedWebsiteRpcServer) UpdateFriend(context.Context, *FriendNewReq) (*FriendDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriend not implemented")
}
func (UnimplementedWebsiteRpcServer) DeleteFriend(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedWebsiteRpcServer) FindFriendList(context.Context, *FindFriendListReq) (*FindFriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFriendList not implemented")
}
func (UnimplementedWebsiteRpcServer) mustEmbedUnimplementedWebsiteRpcServer() {}

// UnsafeWebsiteRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebsiteRpcServer will
// result in compilation errors.
type UnsafeWebsiteRpcServer interface {
	mustEmbedUnimplementedWebsiteRpcServer()
}

func RegisterWebsiteRpcServer(s grpc.ServiceRegistrar, srv WebsiteRpcServer) {
	s.RegisterService(&WebsiteRpc_ServiceDesc, srv)
}

func _WebsiteRpc_GetUserDailyVisit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).GetUserDailyVisit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_GetUserDailyVisit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).GetUserDailyVisit(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_GetUserTotalVisit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).GetUserTotalVisit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_GetUserTotalVisit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).GetUserTotalVisit(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_AddPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).AddPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_AddPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).AddPage(ctx, req.(*PageNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_UpdatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).UpdatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_UpdatePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).UpdatePage(ctx, req.(*PageNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_DeletePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).DeletePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_DeletePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).DeletePage(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPageListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).FindPageList(ctx, req.(*FindPageListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).AddFriend(ctx, req.(*FriendNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_UpdateFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).UpdateFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_UpdateFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).UpdateFriend(ctx, req.(*FriendNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_DeleteFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).DeleteFriend(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteRpc_FindFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteRpcServer).FindFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteRpc_FindFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteRpcServer).FindFriendList(ctx, req.(*FindFriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WebsiteRpc_ServiceDesc is the grpc.ServiceDesc for WebsiteRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebsiteRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "websiterpc.WebsiteRpc",
	HandlerType: (*WebsiteRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserDailyVisit",
			Handler:    _WebsiteRpc_GetUserDailyVisit_Handler,
		},
		{
			MethodName: "GetUserTotalVisit",
			Handler:    _WebsiteRpc_GetUserTotalVisit_Handler,
		},
		{
			MethodName: "AddPage",
			Handler:    _WebsiteRpc_AddPage_Handler,
		},
		{
			MethodName: "UpdatePage",
			Handler:    _WebsiteRpc_UpdatePage_Handler,
		},
		{
			MethodName: "DeletePage",
			Handler:    _WebsiteRpc_DeletePage_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _WebsiteRpc_FindPageList_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _WebsiteRpc_AddFriend_Handler,
		},
		{
			MethodName: "UpdateFriend",
			Handler:    _WebsiteRpc_UpdateFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _WebsiteRpc_DeleteFriend_Handler,
		},
		{
			MethodName: "FindFriendList",
			Handler:    _WebsiteRpc_FindFriendList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "website.proto",
}
