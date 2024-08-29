// Code generated by goctl. DO NOT EDIT.
// Source: syslog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
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

// 创建操作记录
func (s *SyslogRpcServer) AddOperationLog(ctx context.Context, in *syslogrpc.OperationLog) (*syslogrpc.OperationLog, error) {
	l := syslogrpclogic.NewAddOperationLogLogic(ctx, s.svcCtx)
	return l.AddOperationLog(in)
}

// 批量删除操作记录
func (s *SyslogRpcServer) DeleteOperationLogList(ctx context.Context, in *syslogrpc.IdsReq) (*syslogrpc.BatchResp, error) {
	l := syslogrpclogic.NewDeleteOperationLogListLogic(ctx, s.svcCtx)
	return l.DeleteOperationLogList(in)
}

// 查询操作记录列表
func (s *SyslogRpcServer) FindOperationLogList(ctx context.Context, in *syslogrpc.FindOperationLogListReq) (*syslogrpc.FindOperationLogListResp, error) {
	l := syslogrpclogic.NewFindOperationLogListLogic(ctx, s.svcCtx)
	return l.FindOperationLogList(in)
}

// 上传文件
func (s *SyslogRpcServer) AddUploadLog(ctx context.Context, in *syslogrpc.UploadLogReq) (*syslogrpc.UploadLogResp, error) {
	l := syslogrpclogic.NewAddUploadLogLogic(ctx, s.svcCtx)
	return l.AddUploadLog(in)
}