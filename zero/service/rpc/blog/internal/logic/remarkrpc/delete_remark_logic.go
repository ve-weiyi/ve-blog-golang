package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRemarkLogic {
	return &DeleteRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除留言
func (l *DeleteRemarkLogic) DeleteRemark(in *remarkrpc.IdReq) (*remarkrpc.BatchResp, error) {
	rows, err := l.svcCtx.RemarkModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &remarkrpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
