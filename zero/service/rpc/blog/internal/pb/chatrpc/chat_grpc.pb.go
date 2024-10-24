// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: chat.proto

// proto 包名

package chatrpc

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
	ChatRpc_AddChatRecord_FullMethodName      = "/chatrpc.ChatRpc/AddChatRecord"
	ChatRpc_UpdateChatRecord_FullMethodName   = "/chatrpc.ChatRpc/UpdateChatRecord"
	ChatRpc_DeleteChatRecord_FullMethodName   = "/chatrpc.ChatRpc/DeleteChatRecord"
	ChatRpc_GetChatRecord_FullMethodName      = "/chatrpc.ChatRpc/GetChatRecord"
	ChatRpc_FindChatRecordList_FullMethodName = "/chatrpc.ChatRpc/FindChatRecordList"
)

// ChatRpcClient is the client API for ChatRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatRpcClient interface {
	// 创建聊天记录
	AddChatRecord(ctx context.Context, in *ChatRecordNewReq, opts ...grpc.CallOption) (*ChatRecordDetails, error)
	// 更新聊天记录
	UpdateChatRecord(ctx context.Context, in *ChatRecordNewReq, opts ...grpc.CallOption) (*ChatRecordDetails, error)
	// 删除聊天记录
	DeleteChatRecord(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询聊天记录
	GetChatRecord(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*ChatRecordDetails, error)
	// 查询聊天记录列表
	FindChatRecordList(ctx context.Context, in *FindChatRecordListReq, opts ...grpc.CallOption) (*FindChatRecordListResp, error)
}

type chatRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewChatRpcClient(cc grpc.ClientConnInterface) ChatRpcClient {
	return &chatRpcClient{cc}
}

func (c *chatRpcClient) AddChatRecord(ctx context.Context, in *ChatRecordNewReq, opts ...grpc.CallOption) (*ChatRecordDetails, error) {
	out := new(ChatRecordDetails)
	err := c.cc.Invoke(ctx, ChatRpc_AddChatRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRpcClient) UpdateChatRecord(ctx context.Context, in *ChatRecordNewReq, opts ...grpc.CallOption) (*ChatRecordDetails, error) {
	out := new(ChatRecordDetails)
	err := c.cc.Invoke(ctx, ChatRpc_UpdateChatRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRpcClient) DeleteChatRecord(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, ChatRpc_DeleteChatRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRpcClient) GetChatRecord(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*ChatRecordDetails, error) {
	out := new(ChatRecordDetails)
	err := c.cc.Invoke(ctx, ChatRpc_GetChatRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRpcClient) FindChatRecordList(ctx context.Context, in *FindChatRecordListReq, opts ...grpc.CallOption) (*FindChatRecordListResp, error) {
	out := new(FindChatRecordListResp)
	err := c.cc.Invoke(ctx, ChatRpc_FindChatRecordList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatRpcServer is the server API for ChatRpc service.
// All implementations must embed UnimplementedChatRpcServer
// for forward compatibility
type ChatRpcServer interface {
	// 创建聊天记录
	AddChatRecord(context.Context, *ChatRecordNewReq) (*ChatRecordDetails, error)
	// 更新聊天记录
	UpdateChatRecord(context.Context, *ChatRecordNewReq) (*ChatRecordDetails, error)
	// 删除聊天记录
	DeleteChatRecord(context.Context, *IdsReq) (*BatchResp, error)
	// 查询聊天记录
	GetChatRecord(context.Context, *IdReq) (*ChatRecordDetails, error)
	// 查询聊天记录列表
	FindChatRecordList(context.Context, *FindChatRecordListReq) (*FindChatRecordListResp, error)
	mustEmbedUnimplementedChatRpcServer()
}

// UnimplementedChatRpcServer must be embedded to have forward compatible implementations.
type UnimplementedChatRpcServer struct {
}

func (UnimplementedChatRpcServer) AddChatRecord(context.Context, *ChatRecordNewReq) (*ChatRecordDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChatRecord not implemented")
}
func (UnimplementedChatRpcServer) UpdateChatRecord(context.Context, *ChatRecordNewReq) (*ChatRecordDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatRecord not implemented")
}
func (UnimplementedChatRpcServer) DeleteChatRecord(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChatRecord not implemented")
}
func (UnimplementedChatRpcServer) GetChatRecord(context.Context, *IdReq) (*ChatRecordDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatRecord not implemented")
}
func (UnimplementedChatRpcServer) FindChatRecordList(context.Context, *FindChatRecordListReq) (*FindChatRecordListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindChatRecordList not implemented")
}
func (UnimplementedChatRpcServer) mustEmbedUnimplementedChatRpcServer() {}

// UnsafeChatRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatRpcServer will
// result in compilation errors.
type UnsafeChatRpcServer interface {
	mustEmbedUnimplementedChatRpcServer()
}

func RegisterChatRpcServer(s grpc.ServiceRegistrar, srv ChatRpcServer) {
	s.RegisterService(&ChatRpc_ServiceDesc, srv)
}

func _ChatRpc_AddChatRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRecordNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRpcServer).AddChatRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatRpc_AddChatRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRpcServer).AddChatRecord(ctx, req.(*ChatRecordNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRpc_UpdateChatRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRecordNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRpcServer).UpdateChatRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatRpc_UpdateChatRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRpcServer).UpdateChatRecord(ctx, req.(*ChatRecordNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRpc_DeleteChatRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRpcServer).DeleteChatRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatRpc_DeleteChatRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRpcServer).DeleteChatRecord(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRpc_GetChatRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRpcServer).GetChatRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatRpc_GetChatRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRpcServer).GetChatRecord(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRpc_FindChatRecordList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindChatRecordListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRpcServer).FindChatRecordList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatRpc_FindChatRecordList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRpcServer).FindChatRecordList(ctx, req.(*FindChatRecordListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatRpc_ServiceDesc is the grpc.ServiceDesc for ChatRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatrpc.ChatRpc",
	HandlerType: (*ChatRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddChatRecord",
			Handler:    _ChatRpc_AddChatRecord_Handler,
		},
		{
			MethodName: "UpdateChatRecord",
			Handler:    _ChatRpc_UpdateChatRecord_Handler,
		},
		{
			MethodName: "DeleteChatRecord",
			Handler:    _ChatRpc_DeleteChatRecord_Handler,
		},
		{
			MethodName: "GetChatRecord",
			Handler:    _ChatRpc_GetChatRecord_Handler,
		},
		{
			MethodName: "FindChatRecordList",
			Handler:    _ChatRpc_FindChatRecordList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
