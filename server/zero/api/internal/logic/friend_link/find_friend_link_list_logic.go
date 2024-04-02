package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendLinkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindFriendLinkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendLinkListLogic {
	return &FindFriendLinkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFriendLinkListLogic) FindFriendLinkList(req *types.PageQuery) (resp []types.FriendLink, err error) {
	// todo: add your logic here and delete this line

	return
}
