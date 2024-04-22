package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserListLogic {
	return &FindUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserListLogic) FindUserList(req *types.PageQuery) (resp *types.PageResult, err error) {
	in := convert.ConvertPageQuery(req)
	users, err := l.svcCtx.UserRpc.FindUserList(l.ctx, in)
	if err != nil {
		return
	}

	var list []*types.UserDTO
	for _, user := range users.List {
		u := convert.ConvertUserDetailsTypes(user)
		list = append(list, u)
	}

	resp = &types.PageResult{}
	resp.Page = in.Limit.Page
	resp.PageSize = in.Limit.PageSize
	resp.Total = users.Total
	resp.List = list
	return
}
