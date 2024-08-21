package rolerpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
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
func (l *FindRoleListLogic) FindRoleList(in *blog.PageQuery) (*blog.FindRoleListResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.RoleModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var root blog.RoleDetails
	root.Children = appendRoleChildren(&root, result)

	out := &blog.FindRoleListResp{}
	out.Total = int64(len(root.Children))
	out.List = root.Children

	return out, nil
}

func appendRoleChildren(root *blog.RoleDetails, list []*model.Role) (leafs []*blog.RoleDetails) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convert.ConvertRoleModelToDetailPb(item)
			leaf.Children = appendRoleChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
