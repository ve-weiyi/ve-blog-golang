package user

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenusLogic {
	return &GetUserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenusLogic) GetUserMenus(req *types.EmptyReq) (resp *types.UserMenusResp, err error) {
	in := convert.EmptyReq()
	out, err := l.svcCtx.UserRpc.GetUserMenus(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserMenu
	for _, v := range out.List {
		list = append(list, convert.ConvertUserMenuTypes(v))
	}

	resp = &types.UserMenusResp{}
	resp.List = list
	return
}
