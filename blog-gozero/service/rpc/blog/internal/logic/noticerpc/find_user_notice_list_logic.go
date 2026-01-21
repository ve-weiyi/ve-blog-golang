package noticerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/enums"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rpcutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/noticerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserNoticeListLogic {
	return &FindUserNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户可见通知列表
func (l *FindUserNoticeListLogic) FindUserNoticeList(in *noticerpc.FindUserNoticeListReq) (*noticerpc.FindUserNoticeListResp, error) {
	appName, _ := rpcutils.GetAppNameFromCtx(l.ctx)

	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}

	opts = append(opts, query.WithCondition("publish_status = ?", enums.NoticeStatusPublished))

	if appName != "" {
		opts = append(opts, query.WithCondition("app_name = ?", appName))
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

	return &noticerpc.FindUserNoticeListResp{
		List: list,
		Pagination: &noticerpc.PageResp{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
	}, nil
}

func convertNoticeOut(in *model.TSystemNotice) *noticerpc.Notice {
	out := &noticerpc.Notice{
		Id:            in.Id,
		Title:         in.Title,
		Content:       in.Content,
		Type:          in.Type,
		Level:         in.Level,
		PublishStatus: in.PublishStatus,
		AppName:       in.AppName,
		PublisherId:   in.PublisherId,
		CreatedAt:     in.CreatedAt.UnixMilli(),
		UpdatedAt:     in.UpdatedAt.UnixMilli(),
	}

	if in.PublishTime.Valid {
		out.PublishTime = in.PublishTime.Time.UnixMilli()
	}

	if in.RevokeTime.Valid {
		out.RevokeTime = in.RevokeTime.Time.UnixMilli()
	}

	return out
}
