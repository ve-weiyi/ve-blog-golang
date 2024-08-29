package permissionrpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type FindMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuListLogic {
	return &FindMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取菜单列表
func (l *FindMenuListLogic) FindMenuList(in *permissionrpc.FindMenuListReq) (*permissionrpc.FindMenuListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)

	result, err := l.svcCtx.MenuModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var root permissionrpc.MenuDetails
	root.Children = appendMenuChildren(&root, result)

	out := &permissionrpc.FindMenuListResp{}
	out.List = root.Children

	return out, nil
}

func appendMenuChildren(root *permissionrpc.MenuDetails, list []*model.Menu) (leafs []*permissionrpc.MenuDetails) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convertMenuOut(item)
			leaf.Children = appendMenuChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
