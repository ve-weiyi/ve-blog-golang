package account

import (
	"context"
	"strings"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户菜单权限
func NewGetUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenusLogic {
	return &GetUserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenusLogic) GetUserMenus(req *types.EmptyReq) (resp *types.UserMenusResp, err error) {
	in := &permissionrpc.UserIdReq{
		UserId: cast.ToInt64(l.ctx.Value("uid")),
	}

	out, err := l.svcCtx.PermissionRpc.FindUserMenus(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserMenu
	for _, v := range out.List {
		list = append(list, ConvertUserMenuTypes(v))
	}

	resp = &types.UserMenusResp{}
	resp.List = list
	return
}

func ConvertUserMenuTypes(in *permissionrpc.MenuDetails) (out *types.UserMenu) {
	out = &types.UserMenu{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Title,
		Type:      in.Type,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Meta:      types.UserMenuMeta{},
		Children:  make([]*types.UserMenu, 0),
	}

	jsonconv.JsonToObject(in.Extra, &out.Meta)
	if !strings.Contains(in.Path, ":") {
		out.Meta.ShowLink = true
		out.Meta.ShowParent = true
	}

	for _, v := range in.Children {
		out.Children = append(out.Children, ConvertUserMenuTypes(v))
	}

	return
}
