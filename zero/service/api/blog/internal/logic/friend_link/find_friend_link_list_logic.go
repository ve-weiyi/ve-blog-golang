package friend_link

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendLinkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取友链列表
func NewFindFriendLinkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendLinkListLogic {
	return &FindFriendLinkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFriendLinkListLogic) FindFriendLinkList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.FriendLinkRpc.FindFriendLinkList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.FriendLinkRpc.FindFriendLinkCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.FriendLink
	for _, v := range out.List {
		list = append(list, convert.ConvertFriendLinkTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
