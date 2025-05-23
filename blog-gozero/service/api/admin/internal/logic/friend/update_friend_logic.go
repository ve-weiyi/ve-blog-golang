package friend

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新友链
func NewUpdateFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLogic {
	return &UpdateFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFriendLogic) UpdateFriend(req *types.FriendNewReq) (resp *types.FriendBackVO, err error) {
	in := ConvertFriendPb(req)
	out, err := l.svcCtx.WebsiteRpc.UpdateFriend(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertFriendTypes(out)
	return resp, nil
}
