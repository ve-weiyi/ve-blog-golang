package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

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

func (l *FindUserListLogic) FindUserList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := &blogrpc.FindUserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
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
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = users.Total
	resp.List = list
	return
}
