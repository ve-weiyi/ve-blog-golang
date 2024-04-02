package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendLinkLogic {
	return &FindFriendLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFriendLinkLogic) FindFriendLink(req *types.IdReq) (resp *types.FriendLink, err error) {
	// todo: add your logic here and delete this line

	return
}
