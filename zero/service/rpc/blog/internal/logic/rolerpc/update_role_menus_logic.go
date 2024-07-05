package rolerpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

type UpdateRoleMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenusLogic {
	return &UpdateRoleMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色菜单
func (l *UpdateRoleMenusLogic) UpdateRoleMenus(in *blog.UpdateRoleMenusReq) (*blog.EmptyResp, error) {
	// 删除角色菜单
	_, err := l.svcCtx.RoleMenuModel.DeleteBatch(l.ctx, "role_id in (?)", in.RoleId)
	if err != nil {
		return nil, err
	}

	var roleMenus []*model.RoleMenu
	for _, menuId := range in.MenuIds {
		m := &model.RoleMenu{
			RoleId: in.RoleId,
			MenuId: menuId,
		}

		roleMenus = append(roleMenus, m)
	}

	// 添加角色菜单
	_, err = l.svcCtx.RoleMenuModel.InsertBatch(l.ctx, roleMenus...)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
