package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
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
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.RoleModel.FindList(l.ctx, page, size, sorts, conditions, params)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.RoleModel.Count(l.ctx, conditions, params)
	if err != nil {
		return nil, err
	}

	var root account.RoleDetailsDTO
	root.Children = appendRoleChildren(&root, result)

	out := &account.RolePageResp{}
	out.Total = total
	out.List = root.Children

	return out, nil
}

func appendRoleChildren(root *account.RoleDetailsDTO, list []*model.Role) (leafs []*account.RoleDetailsDTO) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := &account.RoleDetailsDTO{
				Id:          item.Id,
				ParentId:    root.Id,
				RoleDomain:  item.RoleDomain,
				RoleName:    item.RoleName,
				RoleComment: item.RoleComment,
				IsDisable:   item.IsDisable,
				IsDefault:   item.IsDefault,
				CreatedAt:   item.CreatedAt.Unix(),
				UpdatedAt:   item.UpdatedAt.Unix(),
			}
			leaf.Children = appendRoleChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
