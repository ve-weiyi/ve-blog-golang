package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiListLogic {
	return &FindApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindApiListLogic) FindApiList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.ApiRpc.FindApiList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ApiDetails
	for _, role := range out.List {
		list = append(list, convert.ConvertApiDetailsTypes(role))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
