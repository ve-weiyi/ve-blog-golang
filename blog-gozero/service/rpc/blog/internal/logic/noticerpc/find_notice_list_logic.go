package noticerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/noticerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindNoticeListLogic {
	return &FindNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询通知列表
func (l *FindNoticeListLogic) FindNoticeList(in *noticerpc.FindNoticeListReq) (*noticerpc.FindNoticeListResp, error) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}

	if in.Type != "" {
		opts = append(opts, query.WithCondition("type = ?", in.Type))
	}

	if in.Level != "" {
		opts = append(opts, query.WithCondition("level = ?", in.Level))
	}

	if in.PublishStatus != 0 {
		opts = append(opts, query.WithCondition("publish_status = ?", in.PublishStatus))
	}

	if in.AppName != "" {
		opts = append(opts, query.WithCondition("app_name = ?", in.AppName))
	}

	page, size, sorts, conditions, params := query.NewQueryBuilder(opts...).Build()

	records, total, err := l.svcCtx.TSystemNoticeModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*noticerpc.Notice
	for _, v := range records {
		list = append(list, convertNoticeOut(v))
	}

	return &noticerpc.FindNoticeListResp{
		List: list,
		Pagination: &noticerpc.PageResp{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
	}, nil
}
