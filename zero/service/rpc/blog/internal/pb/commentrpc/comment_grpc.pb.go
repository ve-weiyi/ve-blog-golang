// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: comment.proto

// proto 包名

package commentrpc

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
	CommentRpc_AddComment_FullMethodName           = "/commentrpc.CommentRpc/AddComment"
	CommentRpc_UpdateComment_FullMethodName        = "/commentrpc.CommentRpc/UpdateComment"
	CommentRpc_DeleteComment_FullMethodName        = "/commentrpc.CommentRpc/DeleteComment"
	CommentRpc_DeleteCommentList_FullMethodName    = "/commentrpc.CommentRpc/DeleteCommentList"
	CommentRpc_GetComment_FullMethodName           = "/commentrpc.CommentRpc/GetComment"
	CommentRpc_FindCommentList_FullMethodName      = "/commentrpc.CommentRpc/FindCommentList"
	CommentRpc_FindCommentReplyList_FullMethodName = "/commentrpc.CommentRpc/FindCommentReplyList"
	CommentRpc_LikeComment_FullMethodName          = "/commentrpc.CommentRpc/LikeComment"
	CommentRpc_FindUserLikeComment_FullMethodName  = "/commentrpc.CommentRpc/FindUserLikeComment"
)

// CommentRpcClient is the client API for CommentRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentRpcClient interface {
	// 创建评论
	AddComment(ctx context.Context, in *CommentNew, opts ...grpc.CallOption) (*CommentDetails, error)
	// 更新评论
	UpdateComment(ctx context.Context, in *CommentNew, opts ...grpc.CallOption) (*CommentDetails, error)
	// 删除评论
	DeleteComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 批量删除评论
	DeleteCommentList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询评论
	GetComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CommentDetails, error)
	// 查询评论列表
	FindCommentList(ctx context.Context, in *FindCommentListReq, opts ...grpc.CallOption) (*FindCommentListResp, error)
	// 查询评论回复列表
	FindCommentReplyList(ctx context.Context, in *FindCommentReplyListReq, opts ...grpc.CallOption) (*FindCommentReplyListResp, error)
	// 点赞评论
	LikeComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error)
	// 用户点赞的评论
	FindUserLikeComment(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindLikeCommentResp, error)
}

type commentRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentRpcClient(cc grpc.ClientConnInterface) CommentRpcClient {
	return &commentRpcClient{cc}
}

func (c *commentRpcClient) AddComment(ctx context.Context, in *CommentNew, opts ...grpc.CallOption) (*CommentDetails, error) {
	out := new(CommentDetails)
	err := c.cc.Invoke(ctx, CommentRpc_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) UpdateComment(ctx context.Context, in *CommentNew, opts ...grpc.CallOption) (*CommentDetails, error) {
	out := new(CommentDetails)
	err := c.cc.Invoke(ctx, CommentRpc_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) DeleteComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, CommentRpc_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) DeleteCommentList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, CommentRpc_DeleteCommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) GetComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CommentDetails, error) {
	out := new(CommentDetails)
	err := c.cc.Invoke(ctx, CommentRpc_GetComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) FindCommentList(ctx context.Context, in *FindCommentListReq, opts ...grpc.CallOption) (*FindCommentListResp, error) {
	out := new(FindCommentListResp)
	err := c.cc.Invoke(ctx, CommentRpc_FindCommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) FindCommentReplyList(ctx context.Context, in *FindCommentReplyListReq, opts ...grpc.CallOption) (*FindCommentReplyListResp, error) {
	out := new(FindCommentReplyListResp)
	err := c.cc.Invoke(ctx, CommentRpc_FindCommentReplyList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) LikeComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, CommentRpc_LikeComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) FindUserLikeComment(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindLikeCommentResp, error) {
	out := new(FindLikeCommentResp)
	err := c.cc.Invoke(ctx, CommentRpc_FindUserLikeComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentRpcServer is the server API for CommentRpc service.
// All implementations must embed UnimplementedCommentRpcServer
// for forward compatibility
type CommentRpcServer interface {
	// 创建评论
	AddComment(context.Context, *CommentNew) (*CommentDetails, error)
	// 更新评论
	UpdateComment(context.Context, *CommentNew) (*CommentDetails, error)
	// 删除评论
	DeleteComment(context.Context, *IdReq) (*BatchResp, error)
	// 批量删除评论
	DeleteCommentList(context.Context, *IdsReq) (*BatchResp, error)
	// 查询评论
	GetComment(context.Context, *IdReq) (*CommentDetails, error)
	// 查询评论列表
	FindCommentList(context.Context, *FindCommentListReq) (*FindCommentListResp, error)
	// 查询评论回复列表
	FindCommentReplyList(context.Context, *FindCommentReplyListReq) (*FindCommentReplyListResp, error)
	// 点赞评论
	LikeComment(context.Context, *IdReq) (*EmptyResp, error)
	// 用户点赞的评论
	FindUserLikeComment(context.Context, *UserIdReq) (*FindLikeCommentResp, error)
	mustEmbedUnimplementedCommentRpcServer()
}

// UnimplementedCommentRpcServer must be embedded to have forward compatible implementations.
type UnimplementedCommentRpcServer struct {
}

func (UnimplementedCommentRpcServer) AddComment(context.Context, *CommentNew) (*CommentDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCommentRpcServer) UpdateComment(context.Context, *CommentNew) (*CommentDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentRpcServer) DeleteComment(context.Context, *IdReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentRpcServer) DeleteCommentList(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommentList not implemented")
}
func (UnimplementedCommentRpcServer) GetComment(context.Context, *IdReq) (*CommentDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentRpcServer) FindCommentList(context.Context, *FindCommentListReq) (*FindCommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCommentList not implemented")
}
func (UnimplementedCommentRpcServer) FindCommentReplyList(context.Context, *FindCommentReplyListReq) (*FindCommentReplyListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCommentReplyList not implemented")
}
func (UnimplementedCommentRpcServer) LikeComment(context.Context, *IdReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeComment not implemented")
}
func (UnimplementedCommentRpcServer) FindUserLikeComment(context.Context, *UserIdReq) (*FindLikeCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserLikeComment not implemented")
}
func (UnimplementedCommentRpcServer) mustEmbedUnimplementedCommentRpcServer() {}

// UnsafeCommentRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentRpcServer will
// result in compilation errors.
type UnsafeCommentRpcServer interface {
	mustEmbedUnimplementedCommentRpcServer()
}

func RegisterCommentRpcServer(s grpc.ServiceRegistrar, srv CommentRpcServer) {
	s.RegisterService(&CommentRpc_ServiceDesc, srv)
}

func _CommentRpc_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).AddComment(ctx, req.(*CommentNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).UpdateComment(ctx, req.(*CommentNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).DeleteComment(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_DeleteCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).DeleteCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_DeleteCommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).DeleteCommentList(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).GetComment(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_FindCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).FindCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_FindCommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).FindCommentList(ctx, req.(*FindCommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_FindCommentReplyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCommentReplyListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).FindCommentReplyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_FindCommentReplyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).FindCommentReplyList(ctx, req.(*FindCommentReplyListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_LikeComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).LikeComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_LikeComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).LikeComment(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_FindUserLikeComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).FindUserLikeComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentRpc_FindUserLikeComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).FindUserLikeComment(ctx, req.(*UserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentRpc_ServiceDesc is the grpc.ServiceDesc for CommentRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commentrpc.CommentRpc",
	HandlerType: (*CommentRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _CommentRpc_AddComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _CommentRpc_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentRpc_DeleteComment_Handler,
		},
		{
			MethodName: "DeleteCommentList",
			Handler:    _CommentRpc_DeleteCommentList_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _CommentRpc_GetComment_Handler,
		},
		{
			MethodName: "FindCommentList",
			Handler:    _CommentRpc_FindCommentList_Handler,
		},
		{
			MethodName: "FindCommentReplyList",
			Handler:    _CommentRpc_FindCommentReplyList_Handler,
		},
		{
			MethodName: "LikeComment",
			Handler:    _CommentRpc_LikeComment_Handler,
		},
		{
			MethodName: "FindUserLikeComment",
			Handler:    _CommentRpc_FindUserLikeComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
