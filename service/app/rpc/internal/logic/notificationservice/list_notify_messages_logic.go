package notificationservicelogic

import (
	"context"

	"github.com/ve-weiyi/vkit/adapter/gormx/queryx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/notificationrpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

type ListNotifyMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNotifyMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNotifyMessagesLogic {
	return &ListNotifyMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询通知消息列表
func (l *ListNotifyMessagesLogic) ListNotifyMessages(in *notificationrpc.ListNotifyMessagesRequest) (*notificationrpc.ListNotifyMessagesResponse, error) {
	var opts []queryx.Option
	if in.PageQuery != nil {
		opts = append(opts, queryx.WithPage(int(in.PageQuery.Page)))
		opts = append(opts, queryx.WithSize(int(in.PageQuery.PageSize)))
		opts = append(opts, queryx.WithSorts(in.PageQuery.Sorts...))
	}

	if in.Category != nil && *in.Category != "" {
		opts = append(opts, queryx.WithCondition("category = ?", *in.Category))
	}
	if in.Level != nil && *in.Level != "" {
		opts = append(opts, queryx.WithCondition("level = ?", *in.Level))
	}
	if in.Status != nil && *in.Status != "" {
		opts = append(opts, queryx.WithCondition("status = ?", *in.Status))
	}
	if in.TargetType != nil && *in.TargetType != "" {
		opts = append(opts, queryx.WithCondition("target_type = ?", *in.TargetType))
	}

	page, size, sorts, conditions, params := queryx.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TNotifyMessageModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*notificationrpc.NotifyMessage
	for _, v := range records {
		list = append(list, convertTNotifyMessageToProto(v))
	}

	return &notificationrpc.ListNotifyMessagesResponse{
		PageResult: &notificationrpc.PageResult{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
		Messages: list,
	}, nil
}
