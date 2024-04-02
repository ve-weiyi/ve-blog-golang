package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLinkLogic {
	return &UpdateFriendLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFriendLinkLogic) UpdateFriendLink(req *types.FriendLink) (resp *types.FriendLink, err error) {
	// todo: add your logic here and delete this line

	return
}
