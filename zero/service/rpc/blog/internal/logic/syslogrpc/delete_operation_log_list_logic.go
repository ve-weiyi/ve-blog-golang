package syslogrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOperationLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOperationLogListLogic {
	return &DeleteOperationLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除操作记录
func (l *DeleteOperationLogListLogic) DeleteOperationLogList(in *syslogrpc.IdsReq) (*syslogrpc.BatchResp, error) {
	rows, err := l.svcCtx.OperationLogModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
