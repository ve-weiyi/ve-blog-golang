package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/zero/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleListLogic {
	return &FindRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取角色列表
func (l *FindRoleListLogic) FindRoleList(in *account.PageQuery) (*account.RolePageResp, error) {
	limit, offset, sorts, conditions, params := parsePageQuery(in)

	result, err := l.svcCtx.RoleModel.FindList(l.ctx, limit, offset, sorts, conditions, params)
	if err != nil {
		return nil, err
	}

	return &account.RolePageResp{}, nil
}

func appendRoleChildren(root *account.RoleDetailsDTO, list []*model.Role) (leafs []*account.RoleDetailsDTO) {
	for _, item := range list {
		if item.RolePid == root.Id {
			leaf := account.RoleDetailsDTO{
				Id:             item.Id,
				RolePid: root.Id,
				RoleDomain: item.
				RoleName:       "",
				RoleComment:    "",
				IsDisable:      0,
				IsDefault:      0,
				CreatedAt:      0,
				UpdatedAt:      0,
				MenuIdList:     nil,
				ResourceIdList: nil,
			}
			leaf.Children = getApiChildren(leaf, list)
			leafs = append(leafs, &leaf)
		}
	}
	return leafs
}

func getApiChildren(root account.RoleDetailsDTO, list []*entity.Api) (leafs []*account.RoleDetailsDTO) {
	for _, item := range list {
		if item.ParentID == root.ID {
			leaf := account.RoleDetailsDTO{
				Api:      *item,
				Children: nil,
			}
			leaf.Children = getApiChildren(leaf, list)
			leafs = append(leafs, &leaf)
		}
	}
	return leafs
}
