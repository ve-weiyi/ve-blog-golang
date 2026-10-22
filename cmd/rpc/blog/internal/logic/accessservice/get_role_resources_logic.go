package accessservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/accessrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type GetRoleResourcesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleResourcesLogic {
	return &GetRoleResourcesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量查询角色资源，避免按角色逐个回源。
func (l *GetRoleResourcesLogic) GetRoleResources(in *accessrpc.GetRoleResourcesRequest) (*accessrpc.GetRoleResourcesResponse, error) {
	if len(in.RoleIds) == 0 {
		return &accessrpc.GetRoleResourcesResponse{}, nil
	}

	var apiRows []struct {
		RoleId int64 `gorm:"column:role_id"`
		ApiId  int64 `gorm:"column:api_id"`
	}
	if err := l.svcCtx.GormDB.Table("`t_role_api`").
		Select("role_id, api_id").
		Where("role_id IN ?", in.RoleIds).
		Scan(&apiRows).Error; err != nil {
		return nil, err
	}

	var menuRows []struct {
		RoleId int64 `gorm:"column:role_id"`
		MenuId int64 `gorm:"column:menu_id"`
	}
	if err := l.svcCtx.GormDB.Table("`t_role_menu`").
		Select("role_id, menu_id").
		Where("role_id IN ?", in.RoleIds).
		Scan(&menuRows).Error; err != nil {
		return nil, err
	}

	apiIds := make(map[int64][]int64, len(in.RoleIds))
	for _, row := range apiRows {
		apiIds[row.RoleId] = append(apiIds[row.RoleId], row.ApiId)
	}
	menuIds := make(map[int64][]int64, len(in.RoleIds))
	for _, row := range menuRows {
		menuIds[row.RoleId] = append(menuIds[row.RoleId], row.MenuId)
	}

	list := make([]*accessrpc.RoleResource, 0, len(in.RoleIds))
	for _, roleId := range in.RoleIds {
		list = append(list, &accessrpc.RoleResource{
			RoleId:  roleId,
			ApiIds:  apiIds[roleId],
			MenuIds: menuIds[roleId],
		})
	}

	return &accessrpc.GetRoleResourcesResponse{List: list}, nil
}
