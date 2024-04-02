package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLinkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFriendLinkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLinkListLogic {
	return &DeleteFriendLinkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFriendLinkListLogic) DeleteFriendLinkList(req *types.IdsReq) (resp *types.BatchResult, err error) {
	// todo: add your logic here and delete this line

	return
}
