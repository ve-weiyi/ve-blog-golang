package notice

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户通知列表
func NewFindUserNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserNoticeListLogic {
	return &FindUserNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserNoticeListLogic) FindUserNoticeList(req *types.QueryUserNoticeReq) (resp *types.PageResp, err error) {
	in := &noticerpc.FindUserNoticeListReq{
		Paginate: &noticerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
	}

	out, err := l.svcCtx.NoticeRpc.FindUserNoticeList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.NoticeBackVO
	for _, v := range out.List {
		list = append(list, convertNoticeOut(v))
	}

	resp = &types.PageResp{}
	resp.Page = out.Pagination.Page
	resp.PageSize = out.Pagination.PageSize
	resp.Total = out.Pagination.Total
	resp.List = list
	return resp, nil
}
