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
	// todo: add your logic here and delete this line

	return &friendrpc.BatchResp{}, nil
}
