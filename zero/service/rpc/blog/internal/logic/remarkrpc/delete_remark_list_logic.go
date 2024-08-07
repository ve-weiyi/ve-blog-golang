package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRemarkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRemarkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRemarkListLogic {
	return &DeleteRemarkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除留言
func (l *DeleteRemarkListLogic) DeleteRemarkList(in *remarkrpc.IdsReq) (*remarkrpc.BatchResp, error) {
	// todo: add your logic here and delete this line

	return &remarkrpc.BatchResp{}, nil
}
