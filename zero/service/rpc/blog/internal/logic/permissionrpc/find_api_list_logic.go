package permissionrpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type FindApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiListLogic {
	return &FindApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取接口列表
func (l *FindApiListLogic) FindApiList(in *permissionrpc.FindApiListReq) (*permissionrpc.FindApiListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)

	result, err := l.svcCtx.TApiModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var root permissionrpc.ApiDetails
	root.Children = appendApiChildren(&root, result)

	out := &permissionrpc.FindApiListResp{}
	out.List = root.Children

	return out, nil
}

func appendApiChildren(root *permissionrpc.ApiDetails, list []*model.TApi) (leafs []*permissionrpc.ApiDetails) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convertApiOut(item)
			leaf.Children = appendApiChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
