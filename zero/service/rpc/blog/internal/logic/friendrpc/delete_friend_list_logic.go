package friendrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/friendrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendListLogic {
	return &DeleteFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除友链
func (l *DeleteFriendListLogic) DeleteFriendList(in *friendrpc.IdsReq) (*friendrpc.BatchResp, error) {
	rows, err := l.svcCtx.FriendModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &friendrpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
