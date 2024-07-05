package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询友链
func NewFindFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendLinkLogic {
	return &FindFriendLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFriendLinkLogic) FindFriendLink(req *types.IdReq) (resp *types.FriendLink, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.FriendLinkRpc.FindFriendLink(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertFriendLinkTypes(out), nil
}
