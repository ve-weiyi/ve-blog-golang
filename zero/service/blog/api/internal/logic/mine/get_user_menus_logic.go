package mine

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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

func (l *GetUserMenusLogic) GetUserMenus(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.UserMenusResp, err error) {
	in := &blog.UserReq{
		UserId: cast.ToInt64(reqCtx.HeaderXUserId),
	}

	out, err := l.svcCtx.UserRpc.FindUserMenus(l.ctx, in)
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
