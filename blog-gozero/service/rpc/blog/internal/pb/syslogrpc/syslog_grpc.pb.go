// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: syslog.proto

// proto 包名

package syslogrpc

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
	SyslogRpc_AddLoginLog_FullMethodName          = "/syslogrpc.SyslogRpc/AddLoginLog"
	SyslogRpc_AddLogoutLog_FullMethodName         = "/syslogrpc.SyslogRpc/AddLogoutLog"
	SyslogRpc_DeletesLoginLog_FullMethodName      = "/syslogrpc.SyslogRpc/DeletesLoginLog"
	SyslogRpc_FindLoginLogList_FullMethodName     = "/syslogrpc.SyslogRpc/FindLoginLogList"
	SyslogRpc_AddVisitLog_FullMethodName          = "/syslogrpc.SyslogRpc/AddVisitLog"
	SyslogRpc_DeletesVisitLog_FullMethodName      = "/syslogrpc.SyslogRpc/DeletesVisitLog"
	SyslogRpc_FindVisitLogList_FullMethodName     = "/syslogrpc.SyslogRpc/FindVisitLogList"
	SyslogRpc_AddOperationLog_FullMethodName      = "/syslogrpc.SyslogRpc/AddOperationLog"
	SyslogRpc_DeletesOperationLog_FullMethodName  = "/syslogrpc.SyslogRpc/DeletesOperationLog"
	SyslogRpc_FindOperationLogList_FullMethodName = "/syslogrpc.SyslogRpc/FindOperationLogList"
	SyslogRpc_AddUploadLog_FullMethodName         = "/syslogrpc.SyslogRpc/AddUploadLog"
	SyslogRpc_DeletesUploadLog_FullMethodName     = "/syslogrpc.SyslogRpc/DeletesUploadLog"
	SyslogRpc_FindUploadLogList_FullMethodName    = "/syslogrpc.SyslogRpc/FindUploadLogList"
)

// SyslogRpcClient is the client API for SyslogRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyslogRpcClient interface {
	// 创建登录记录
	AddLoginLog(ctx context.Context, in *LoginLogNewReq, opts ...grpc.CallOption) (*EmptyResp, error)
	// 更新登录记录
	AddLogoutLog(ctx context.Context, in *AddLogoutLogReq, opts ...grpc.CallOption) (*AddLogoutLogResp, error)
	// 批量删除登录记录
	DeletesLoginLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询登录记录列表
	FindLoginLogList(ctx context.Context, in *FindLoginLogListReq, opts ...grpc.CallOption) (*FindLoginLogListResp, error)
	// 创建访问记录
	AddVisitLog(ctx context.Context, in *VisitLogNewReq, opts ...grpc.CallOption) (*EmptyResp, error)
	// 批量删除访问记录
	DeletesVisitLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询操作访问列表
	FindVisitLogList(ctx context.Context, in *FindVisitLogListReq, opts ...grpc.CallOption) (*FindVisitLogListResp, error)
	// 创建操作记录
	AddOperationLog(ctx context.Context, in *OperationLogNewReq, opts ...grpc.CallOption) (*EmptyResp, error)
	// 批量删除操作记录
	DeletesOperationLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询操作记录列表
	FindOperationLogList(ctx context.Context, in *FindOperationLogListReq, opts ...grpc.CallOption) (*FindOperationLogListResp, error)
	// 创建上传记录
	AddUploadLog(ctx context.Context, in *UploadLogNewReq, opts ...grpc.CallOption) (*UploadLogDetails, error)
	// 批量删除上传记录
	DeletesUploadLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询上传记录列表
	FindUploadLogList(ctx context.Context, in *FindUploadLogListReq, opts ...grpc.CallOption) (*FindUploadLogListResp, error)
}

type syslogRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewSyslogRpcClient(cc grpc.ClientConnInterface) SyslogRpcClient {
	return &syslogRpcClient{cc}
}

func (c *syslogRpcClient) AddLoginLog(ctx context.Context, in *LoginLogNewReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, SyslogRpc_AddLoginLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) AddLogoutLog(ctx context.Context, in *AddLogoutLogReq, opts ...grpc.CallOption) (*AddLogoutLogResp, error) {
	out := new(AddLogoutLogResp)
	err := c.cc.Invoke(ctx, SyslogRpc_AddLogoutLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) DeletesLoginLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, SyslogRpc_DeletesLoginLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) FindLoginLogList(ctx context.Context, in *FindLoginLogListReq, opts ...grpc.CallOption) (*FindLoginLogListResp, error) {
	out := new(FindLoginLogListResp)
	err := c.cc.Invoke(ctx, SyslogRpc_FindLoginLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) AddVisitLog(ctx context.Context, in *VisitLogNewReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, SyslogRpc_AddVisitLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) DeletesVisitLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, SyslogRpc_DeletesVisitLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) FindVisitLogList(ctx context.Context, in *FindVisitLogListReq, opts ...grpc.CallOption) (*FindVisitLogListResp, error) {
	out := new(FindVisitLogListResp)
	err := c.cc.Invoke(ctx, SyslogRpc_FindVisitLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) AddOperationLog(ctx context.Context, in *OperationLogNewReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, SyslogRpc_AddOperationLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) DeletesOperationLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, SyslogRpc_DeletesOperationLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) FindOperationLogList(ctx context.Context, in *FindOperationLogListReq, opts ...grpc.CallOption) (*FindOperationLogListResp, error) {
	out := new(FindOperationLogListResp)
	err := c.cc.Invoke(ctx, SyslogRpc_FindOperationLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) AddUploadLog(ctx context.Context, in *UploadLogNewReq, opts ...grpc.CallOption) (*UploadLogDetails, error) {
	out := new(UploadLogDetails)
	err := c.cc.Invoke(ctx, SyslogRpc_AddUploadLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) DeletesUploadLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, SyslogRpc_DeletesUploadLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) FindUploadLogList(ctx context.Context, in *FindUploadLogListReq, opts ...grpc.CallOption) (*FindUploadLogListResp, error) {
	out := new(FindUploadLogListResp)
	err := c.cc.Invoke(ctx, SyslogRpc_FindUploadLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyslogRpcServer is the server API for SyslogRpc service.
// All implementations must embed UnimplementedSyslogRpcServer
// for forward compatibility
type SyslogRpcServer interface {
	// 创建登录记录
	AddLoginLog(context.Context, *LoginLogNewReq) (*EmptyResp, error)
	// 更新登录记录
	AddLogoutLog(context.Context, *AddLogoutLogReq) (*AddLogoutLogResp, error)
	// 批量删除登录记录
	DeletesLoginLog(context.Context, *IdsReq) (*BatchResp, error)
	// 查询登录记录列表
	FindLoginLogList(context.Context, *FindLoginLogListReq) (*FindLoginLogListResp, error)
	// 创建访问记录
	AddVisitLog(context.Context, *VisitLogNewReq) (*EmptyResp, error)
	// 批量删除访问记录
	DeletesVisitLog(context.Context, *IdsReq) (*BatchResp, error)
	// 查询操作访问列表
	FindVisitLogList(context.Context, *FindVisitLogListReq) (*FindVisitLogListResp, error)
	// 创建操作记录
	AddOperationLog(context.Context, *OperationLogNewReq) (*EmptyResp, error)
	// 批量删除操作记录
	DeletesOperationLog(context.Context, *IdsReq) (*BatchResp, error)
	// 查询操作记录列表
	FindOperationLogList(context.Context, *FindOperationLogListReq) (*FindOperationLogListResp, error)
	// 创建上传记录
	AddUploadLog(context.Context, *UploadLogNewReq) (*UploadLogDetails, error)
	// 批量删除上传记录
	DeletesUploadLog(context.Context, *IdsReq) (*BatchResp, error)
	// 查询上传记录列表
	FindUploadLogList(context.Context, *FindUploadLogListReq) (*FindUploadLogListResp, error)
	mustEmbedUnimplementedSyslogRpcServer()
}

// UnimplementedSyslogRpcServer must be embedded to have forward compatible implementations.
type UnimplementedSyslogRpcServer struct {
}

func (UnimplementedSyslogRpcServer) AddLoginLog(context.Context, *LoginLogNewReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLoginLog not implemented")
}
func (UnimplementedSyslogRpcServer) AddLogoutLog(context.Context, *AddLogoutLogReq) (*AddLogoutLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLogoutLog not implemented")
}
func (UnimplementedSyslogRpcServer) DeletesLoginLog(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletesLoginLog not implemented")
}
func (UnimplementedSyslogRpcServer) FindLoginLogList(context.Context, *FindLoginLogListReq) (*FindLoginLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLoginLogList not implemented")
}
func (UnimplementedSyslogRpcServer) AddVisitLog(context.Context, *VisitLogNewReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVisitLog not implemented")
}
func (UnimplementedSyslogRpcServer) DeletesVisitLog(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletesVisitLog not implemented")
}
func (UnimplementedSyslogRpcServer) FindVisitLogList(context.Context, *FindVisitLogListReq) (*FindVisitLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVisitLogList not implemented")
}
func (UnimplementedSyslogRpcServer) AddOperationLog(context.Context, *OperationLogNewReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOperationLog not implemented")
}
func (UnimplementedSyslogRpcServer) DeletesOperationLog(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletesOperationLog not implemented")
}
func (UnimplementedSyslogRpcServer) FindOperationLogList(context.Context, *FindOperationLogListReq) (*FindOperationLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOperationLogList not implemented")
}
func (UnimplementedSyslogRpcServer) AddUploadLog(context.Context, *UploadLogNewReq) (*UploadLogDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUploadLog not implemented")
}
func (UnimplementedSyslogRpcServer) DeletesUploadLog(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletesUploadLog not implemented")
}
func (UnimplementedSyslogRpcServer) FindUploadLogList(context.Context, *FindUploadLogListReq) (*FindUploadLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUploadLogList not implemented")
}
func (UnimplementedSyslogRpcServer) mustEmbedUnimplementedSyslogRpcServer() {}

// UnsafeSyslogRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyslogRpcServer will
// result in compilation errors.
type UnsafeSyslogRpcServer interface {
	mustEmbedUnimplementedSyslogRpcServer()
}

func RegisterSyslogRpcServer(s grpc.ServiceRegistrar, srv SyslogRpcServer) {
	s.RegisterService(&SyslogRpc_ServiceDesc, srv)
}

func _SyslogRpc_AddLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginLogNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).AddLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_AddLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).AddLoginLog(ctx, req.(*LoginLogNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_AddLogoutLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLogoutLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).AddLogoutLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_AddLogoutLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).AddLogoutLog(ctx, req.(*AddLogoutLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_DeletesLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).DeletesLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_DeletesLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).DeletesLoginLog(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_FindLoginLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLoginLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).FindLoginLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_FindLoginLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).FindLoginLogList(ctx, req.(*FindLoginLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_AddVisitLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitLogNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).AddVisitLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_AddVisitLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).AddVisitLog(ctx, req.(*VisitLogNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_DeletesVisitLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).DeletesVisitLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_DeletesVisitLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).DeletesVisitLog(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_FindVisitLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindVisitLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).FindVisitLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_FindVisitLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).FindVisitLogList(ctx, req.(*FindVisitLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_AddOperationLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationLogNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).AddOperationLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_AddOperationLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).AddOperationLog(ctx, req.(*OperationLogNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_DeletesOperationLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).DeletesOperationLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_DeletesOperationLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).DeletesOperationLog(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_FindOperationLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOperationLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).FindOperationLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_FindOperationLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).FindOperationLogList(ctx, req.(*FindOperationLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_AddUploadLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadLogNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).AddUploadLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_AddUploadLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).AddUploadLog(ctx, req.(*UploadLogNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_DeletesUploadLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).DeletesUploadLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_DeletesUploadLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).DeletesUploadLog(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_FindUploadLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUploadLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).FindUploadLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_FindUploadLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).FindUploadLogList(ctx, req.(*FindUploadLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SyslogRpc_ServiceDesc is the grpc.ServiceDesc for SyslogRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyslogRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "syslogrpc.SyslogRpc",
	HandlerType: (*SyslogRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLoginLog",
			Handler:    _SyslogRpc_AddLoginLog_Handler,
		},
		{
			MethodName: "AddLogoutLog",
			Handler:    _SyslogRpc_AddLogoutLog_Handler,
		},
		{
			MethodName: "DeletesLoginLog",
			Handler:    _SyslogRpc_DeletesLoginLog_Handler,
		},
		{
			MethodName: "FindLoginLogList",
			Handler:    _SyslogRpc_FindLoginLogList_Handler,
		},
		{
			MethodName: "AddVisitLog",
			Handler:    _SyslogRpc_AddVisitLog_Handler,
		},
		{
			MethodName: "DeletesVisitLog",
			Handler:    _SyslogRpc_DeletesVisitLog_Handler,
		},
		{
			MethodName: "FindVisitLogList",
			Handler:    _SyslogRpc_FindVisitLogList_Handler,
		},
		{
			MethodName: "AddOperationLog",
			Handler:    _SyslogRpc_AddOperationLog_Handler,
		},
		{
			MethodName: "DeletesOperationLog",
			Handler:    _SyslogRpc_DeletesOperationLog_Handler,
		},
		{
			MethodName: "FindOperationLogList",
			Handler:    _SyslogRpc_FindOperationLogList_Handler,
		},
		{
			MethodName: "AddUploadLog",
			Handler:    _SyslogRpc_AddUploadLog_Handler,
		},
		{
			MethodName: "DeletesUploadLog",
			Handler:    _SyslogRpc_DeletesUploadLog_Handler,
		},
		{
			MethodName: "FindUploadLogList",
			Handler:    _SyslogRpc_FindUploadLogList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syslog.proto",
}
