package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新友链
func NewUpdateFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLinkLogic {
	return &UpdateFriendLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFriendLinkLogic) UpdateFriendLink(req *types.FriendLink) (resp *types.FriendLink, err error) {
	in := convert.ConvertFriendLinkPb(req)

	api, err := l.svcCtx.FriendLinkRpc.UpdateFriendLink(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertFriendLinkTypes(api), nil
}
