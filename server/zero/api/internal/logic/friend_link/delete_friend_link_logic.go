package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLinkLogic {
	return &DeleteFriendLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFriendLinkLogic) DeleteFriendLink(req *types.IdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
