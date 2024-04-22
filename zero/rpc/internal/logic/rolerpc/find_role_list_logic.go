package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/repository/model"
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
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.RoleModel.FindList(l.ctx, limit, offset, sorts, conditions, params)
	if err != nil {
		return nil, err
	}

	var root account.RoleDetailsDTO
	root.Children = appendRoleChildren(&root, result)

	out := &account.RolePageResp{}
	out.Total = int64(len(root.Children))
	out.List = root.Children

	return out, nil
}

func appendRoleChildren(root *account.RoleDetailsDTO, list []*model.Role) (leafs []*account.RoleDetailsDTO) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convert.ConvertRoleModelToDetailPb(item)
			leaf.Children = appendRoleChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
