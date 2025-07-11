// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: syslog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/logic/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
)

type SyslogRpcServer struct {
	svcCtx *svc.ServiceContext
	syslogrpc.UnimplementedSyslogRpcServer
}

func NewSyslogRpcServer(svcCtx *svc.ServiceContext) *SyslogRpcServer {
	return &SyslogRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建登录记录
func (s *SyslogRpcServer) AddLoginLog(ctx context.Context, in *syslogrpc.LoginLogNewReq) (*syslogrpc.EmptyResp, error) {
	l := syslogrpclogic.NewAddLoginLogLogic(ctx, s.svcCtx)
	return l.AddLoginLog(in)
}

// 更新登录记录
func (s *SyslogRpcServer) AddLogoutLog(ctx context.Context, in *syslogrpc.AddLogoutLogReq) (*syslogrpc.AddLogoutLogResp, error) {
	l := syslogrpclogic.NewAddLogoutLogLogic(ctx, s.svcCtx)
	return l.AddLogoutLog(in)
}

// 批量删除登录记录
func (s *SyslogRpcServer) DeletesLoginLog(ctx context.Context, in *syslogrpc.IdsReq) (*syslogrpc.BatchResp, error) {
	l := syslogrpclogic.NewDeletesLoginLogLogic(ctx, s.svcCtx)
	return l.DeletesLoginLog(in)
}

// 查询登录记录列表
func (s *SyslogRpcServer) FindLoginLogList(ctx context.Context, in *syslogrpc.FindLoginLogListReq) (*syslogrpc.FindLoginLogListResp, error) {
	l := syslogrpclogic.NewFindLoginLogListLogic(ctx, s.svcCtx)
	return l.FindLoginLogList(in)
}

// 创建访问记录
func (s *SyslogRpcServer) AddVisitLog(ctx context.Context, in *syslogrpc.VisitLogNewReq) (*syslogrpc.EmptyResp, error) {
	l := syslogrpclogic.NewAddVisitLogLogic(ctx, s.svcCtx)
	return l.AddVisitLog(in)
}

// 批量删除访问记录
func (s *SyslogRpcServer) DeletesVisitLog(ctx context.Context, in *syslogrpc.IdsReq) (*syslogrpc.BatchResp, error) {
	l := syslogrpclogic.NewDeletesVisitLogLogic(ctx, s.svcCtx)
	return l.DeletesVisitLog(in)
}

// 查询操作访问列表
func (s *SyslogRpcServer) FindVisitLogList(ctx context.Context, in *syslogrpc.FindVisitLogListReq) (*syslogrpc.FindVisitLogListResp, error) {
	l := syslogrpclogic.NewFindVisitLogListLogic(ctx, s.svcCtx)
	return l.FindVisitLogList(in)
}

// 创建操作记录
func (s *SyslogRpcServer) AddOperationLog(ctx context.Context, in *syslogrpc.OperationLogNewReq) (*syslogrpc.EmptyResp, error) {
	l := syslogrpclogic.NewAddOperationLogLogic(ctx, s.svcCtx)
	return l.AddOperationLog(in)
}

// 批量删除操作记录
func (s *SyslogRpcServer) DeletesOperationLog(ctx context.Context, in *syslogrpc.IdsReq) (*syslogrpc.BatchResp, error) {
	l := syslogrpclogic.NewDeletesOperationLogLogic(ctx, s.svcCtx)
	return l.DeletesOperationLog(in)
}

// 查询操作记录列表
func (s *SyslogRpcServer) FindOperationLogList(ctx context.Context, in *syslogrpc.FindOperationLogListReq) (*syslogrpc.FindOperationLogListResp, error) {
	l := syslogrpclogic.NewFindOperationLogListLogic(ctx, s.svcCtx)
	return l.FindOperationLogList(in)
}

// 创建上传记录
func (s *SyslogRpcServer) AddUploadLog(ctx context.Context, in *syslogrpc.UploadLogNewReq) (*syslogrpc.UploadLogDetails, error) {
	l := syslogrpclogic.NewAddUploadLogLogic(ctx, s.svcCtx)
	return l.AddUploadLog(in)
}

// 批量删除上传记录
func (s *SyslogRpcServer) DeletesUploadLog(ctx context.Context, in *syslogrpc.IdsReq) (*syslogrpc.BatchResp, error) {
	l := syslogrpclogic.NewDeletesUploadLogLogic(ctx, s.svcCtx)
	return l.DeletesUploadLog(in)
}

// 查询上传记录列表
func (s *SyslogRpcServer) FindUploadLogList(ctx context.Context, in *syslogrpc.FindUploadLogListReq) (*syslogrpc.FindUploadLogListResp, error) {
	l := syslogrpclogic.NewFindUploadLogListLogic(ctx, s.svcCtx)
	return l.FindUploadLogList(in)
}
