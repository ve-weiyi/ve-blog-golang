package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFriendLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建友链
func NewCreateFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFriendLinkLogic {
	return &CreateFriendLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFriendLinkLogic) CreateFriendLink(req *types.FriendLink) (resp *types.FriendLink, err error) {
	in := convert.ConvertFriendLinkPb(req)
	out, err := l.svcCtx.FriendLinkRpc.CreateFriendLink(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convert.ConvertFriendLinkTypes(out)
	return resp, nil
}
