package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建菜单
func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuLogic) AddMenu(req *types.MenuBackDTO) (resp *types.MenuBackDTO, err error) {
	in := ConvertMenuPb(req)
	out, err := l.svcCtx.PermissionRpc.AddMenu(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertMenuTypes(out)
	return resp, nil
}

func ConvertMenuPb(in *types.MenuBackDTO) (out *permissionrpc.MenuNewReq) {
	out = &permissionrpc.MenuNewReq{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Title,
		Type:      in.Type,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Extra:     "",
	}

	return
}

func ConvertMenuTypes(in *permissionrpc.MenuDetails) (out *types.MenuBackDTO) {
	out = &types.MenuBackDTO{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Title,
		Type:      in.Type,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Meta:      types.Meta{},
		Children:  nil,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	return
}
