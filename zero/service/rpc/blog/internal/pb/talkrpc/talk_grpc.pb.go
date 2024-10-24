// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: talk.proto

// proto 包名

package talkrpc

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
	TalkRpc_AddTalk_FullMethodName          = "/talkrpc.TalkRpc/AddTalk"
	TalkRpc_UpdateTalk_FullMethodName       = "/talkrpc.TalkRpc/UpdateTalk"
	TalkRpc_DeleteTalk_FullMethodName       = "/talkrpc.TalkRpc/DeleteTalk"
	TalkRpc_GetTalk_FullMethodName          = "/talkrpc.TalkRpc/GetTalk"
	TalkRpc_FindTalkList_FullMethodName     = "/talkrpc.TalkRpc/FindTalkList"
	TalkRpc_LikeTalk_FullMethodName         = "/talkrpc.TalkRpc/LikeTalk"
	TalkRpc_FindUserLikeTalk_FullMethodName = "/talkrpc.TalkRpc/FindUserLikeTalk"
)

// TalkRpcClient is the client API for TalkRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TalkRpcClient interface {
	// 创建说说
	AddTalk(ctx context.Context, in *TalkNewReq, opts ...grpc.CallOption) (*TalkDetails, error)
	// 更新说说
	UpdateTalk(ctx context.Context, in *TalkNewReq, opts ...grpc.CallOption) (*TalkDetails, error)
	// 删除说说
	DeleteTalk(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询说说
	GetTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*TalkDetails, error)
	// 查询说说列表
	FindTalkList(ctx context.Context, in *FindTalkListReq, opts ...grpc.CallOption) (*FindTalkListResp, error)
	// 点赞说说
	LikeTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error)
	// 用户点赞的说说
	FindUserLikeTalk(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindLikeTalkResp, error)
}

type talkRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewTalkRpcClient(cc grpc.ClientConnInterface) TalkRpcClient {
	return &talkRpcClient{cc}
}

func (c *talkRpcClient) AddTalk(ctx context.Context, in *TalkNewReq, opts ...grpc.CallOption) (*TalkDetails, error) {
	out := new(TalkDetails)
	err := c.cc.Invoke(ctx, TalkRpc_AddTalk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkRpcClient) UpdateTalk(ctx context.Context, in *TalkNewReq, opts ...grpc.CallOption) (*TalkDetails, error) {
	out := new(TalkDetails)
	err := c.cc.Invoke(ctx, TalkRpc_UpdateTalk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkRpcClient) DeleteTalk(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, TalkRpc_DeleteTalk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkRpcClient) GetTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*TalkDetails, error) {
	out := new(TalkDetails)
	err := c.cc.Invoke(ctx, TalkRpc_GetTalk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkRpcClient) FindTalkList(ctx context.Context, in *FindTalkListReq, opts ...grpc.CallOption) (*FindTalkListResp, error) {
	out := new(FindTalkListResp)
	err := c.cc.Invoke(ctx, TalkRpc_FindTalkList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkRpcClient) LikeTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, TalkRpc_LikeTalk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkRpcClient) FindUserLikeTalk(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindLikeTalkResp, error) {
	out := new(FindLikeTalkResp)
	err := c.cc.Invoke(ctx, TalkRpc_FindUserLikeTalk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TalkRpcServer is the server API for TalkRpc service.
// All implementations must embed UnimplementedTalkRpcServer
// for forward compatibility
type TalkRpcServer interface {
	// 创建说说
	AddTalk(context.Context, *TalkNewReq) (*TalkDetails, error)
	// 更新说说
	UpdateTalk(context.Context, *TalkNewReq) (*TalkDetails, error)
	// 删除说说
	DeleteTalk(context.Context, *IdsReq) (*BatchResp, error)
	// 查询说说
	GetTalk(context.Context, *IdReq) (*TalkDetails, error)
	// 查询说说列表
	FindTalkList(context.Context, *FindTalkListReq) (*FindTalkListResp, error)
	// 点赞说说
	LikeTalk(context.Context, *IdReq) (*EmptyResp, error)
	// 用户点赞的说说
	FindUserLikeTalk(context.Context, *UserIdReq) (*FindLikeTalkResp, error)
	mustEmbedUnimplementedTalkRpcServer()
}

// UnimplementedTalkRpcServer must be embedded to have forward compatible implementations.
type UnimplementedTalkRpcServer struct {
}

func (UnimplementedTalkRpcServer) AddTalk(context.Context, *TalkNewReq) (*TalkDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTalk not implemented")
}
func (UnimplementedTalkRpcServer) UpdateTalk(context.Context, *TalkNewReq) (*TalkDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTalk not implemented")
}
func (UnimplementedTalkRpcServer) DeleteTalk(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTalk not implemented")
}
func (UnimplementedTalkRpcServer) GetTalk(context.Context, *IdReq) (*TalkDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTalk not implemented")
}
func (UnimplementedTalkRpcServer) FindTalkList(context.Context, *FindTalkListReq) (*FindTalkListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTalkList not implemented")
}
func (UnimplementedTalkRpcServer) LikeTalk(context.Context, *IdReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeTalk not implemented")
}
func (UnimplementedTalkRpcServer) FindUserLikeTalk(context.Context, *UserIdReq) (*FindLikeTalkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserLikeTalk not implemented")
}
func (UnimplementedTalkRpcServer) mustEmbedUnimplementedTalkRpcServer() {}

// UnsafeTalkRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TalkRpcServer will
// result in compilation errors.
type UnsafeTalkRpcServer interface {
	mustEmbedUnimplementedTalkRpcServer()
}

func RegisterTalkRpcServer(s grpc.ServiceRegistrar, srv TalkRpcServer) {
	s.RegisterService(&TalkRpc_ServiceDesc, srv)
}

func _TalkRpc_AddTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalkNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkRpcServer).AddTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkRpc_AddTalk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkRpcServer).AddTalk(ctx, req.(*TalkNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TalkRpc_UpdateTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalkNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkRpcServer).UpdateTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkRpc_UpdateTalk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkRpcServer).UpdateTalk(ctx, req.(*TalkNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TalkRpc_DeleteTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkRpcServer).DeleteTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkRpc_DeleteTalk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkRpcServer).DeleteTalk(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TalkRpc_GetTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkRpcServer).GetTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkRpc_GetTalk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkRpcServer).GetTalk(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TalkRpc_FindTalkList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTalkListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkRpcServer).FindTalkList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkRpc_FindTalkList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkRpcServer).FindTalkList(ctx, req.(*FindTalkListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TalkRpc_LikeTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkRpcServer).LikeTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkRpc_LikeTalk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkRpcServer).LikeTalk(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TalkRpc_FindUserLikeTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkRpcServer).FindUserLikeTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkRpc_FindUserLikeTalk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkRpcServer).FindUserLikeTalk(ctx, req.(*UserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TalkRpc_ServiceDesc is the grpc.ServiceDesc for TalkRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TalkRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "talkrpc.TalkRpc",
	HandlerType: (*TalkRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTalk",
			Handler:    _TalkRpc_AddTalk_Handler,
		},
		{
			MethodName: "UpdateTalk",
			Handler:    _TalkRpc_UpdateTalk_Handler,
		},
		{
			MethodName: "DeleteTalk",
			Handler:    _TalkRpc_DeleteTalk_Handler,
		},
		{
			MethodName: "GetTalk",
			Handler:    _TalkRpc_GetTalk_Handler,
		},
		{
			MethodName: "FindTalkList",
			Handler:    _TalkRpc_FindTalkList_Handler,
		},
		{
			MethodName: "LikeTalk",
			Handler:    _TalkRpc_LikeTalk_Handler,
		},
		{
			MethodName: "FindUserLikeTalk",
			Handler:    _TalkRpc_FindUserLikeTalk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "talk.proto",
}
