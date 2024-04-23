package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOnlineUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOnlineUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOnlineUserListLogic {
	return &FindOnlineUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOnlineUserListLogic) FindOnlineUserList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	users, err := l.svcCtx.UserRpc.FindUserList(l.ctx, in)
	if err != nil {
		return
	}

	var list []*types.User
	for _, user := range users.List {
		u := convert.ConvertUserDetailsTypes(user)
		list = append(list, u)
	}

	resp = &types.PageResp{}
	resp.Page = in.Limit.Page
	resp.PageSize = in.Limit.PageSize
	resp.Total = users.Total
	resp.List = list
	return
}
