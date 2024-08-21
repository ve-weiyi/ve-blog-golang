package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

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
	in := &blogrpc.IdReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.FriendLinkRpc.FindFriendLink(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertFriendLinkTypes(out), nil
}
