package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/remarkrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRemarkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取留言列表
func NewFindRemarkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRemarkListLogic {
	return &FindRemarkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRemarkListLogic) FindRemarkList(req *types.RemarkQueryReq) (resp *types.PageResp, err error) {
	in := &remarkrpc.FindRemarkListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
	}
	out, err := l.svcCtx.RemarkRpc.FindRemarkList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.Remark
	for _, v := range out.List {
		list = append(list, ConvertRemarkTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
