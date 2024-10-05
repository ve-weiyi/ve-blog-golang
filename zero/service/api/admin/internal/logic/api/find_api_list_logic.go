package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取api路由列表
func NewFindApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiListLogic {
	return &FindApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindApiListLogic) FindApiList(req *types.ApiQuery) (resp *types.PageResp, err error) {
	in := &permissionrpc.FindApiListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Path:     req.Path,
		Method:   req.Method,
	}

	out, err := l.svcCtx.PermissionRpc.FindApiList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ApiBackDTO
	for _, v := range out.List {
		m := convertApiTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = int64(len(list))
	resp.List = list
	return resp, nil
}
