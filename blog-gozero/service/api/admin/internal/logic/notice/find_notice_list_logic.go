package notice

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取通知列表
func NewFindNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindNoticeListLogic {
	return &FindNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindNoticeListLogic) FindNoticeList(req *types.QueryNoticeReq) (resp *types.PageResp, err error) {
	in := &noticerpc.FindNoticeListReq{
		Paginate: &noticerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		Type:          req.Type,
		Level:         req.Level,
		PublishStatus: req.PublishStatus,
		AppName:       req.AppName,
	}

	out, err := l.svcCtx.NoticeRpc.FindNoticeList(l.ctx, in)
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

func convertNoticeOut(in *noticerpc.Notice) *types.NoticeBackVO {
	return &types.NoticeBackVO{
		Id:            in.Id,
		Title:         in.Title,
		Content:       in.Content,
		Type:          in.Type,
		Level:         in.Level,
		PublishStatus: in.PublishStatus,
		AppName:       in.AppName,
		PublisherId:   in.PublisherId,
		PublishTime:   in.PublishTime,
		RevokeTime:    in.RevokeTime,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
	}
}
