package newsrpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/infra/queryx"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindChatListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindChatListLogic {
	return &FindChatListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询聊天记录列表
func (l *FindChatListLogic) FindChatList(in *newsrpc.FindChatListReq) (*newsrpc.FindChatListResp, error) {
	page, size, sorts, conditions, params := convertChatQuery(in)

	records, total, err := l.svcCtx.TChatModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*newsrpc.Chat
	for _, v := range records {
		list = append(list, convertChatOut(v))
	}

	return &newsrpc.FindChatListResp{
		List: list,
		Pagination: &newsrpc.PageResp{
			Total: total,
		},
	}, nil
}

func convertChatQuery(in *newsrpc.FindChatListReq) (page int, size int, sorts string, conditions string, params []any) {
	opts := []queryx.Option{
		queryx.WithPage(1),
		queryx.WithSize(int(in.Limit)),
		queryx.WithSorts("created_at desc"),
	}

	if in.After != 0 {
		opts = append(opts, queryx.WithCondition("created_at >= ?", time.UnixMilli(in.After)))
	}

	if in.Before != 0 {
		opts = append(opts, queryx.WithCondition("created_at <= ?", time.UnixMilli(in.Before)))
	}

	if in.UserId != "" {
		opts = append(opts, queryx.WithCondition("user_id = ?", in.UserId))
	}

	if in.Type != "" {
		opts = append(opts, queryx.WithCondition("type = ?", in.Type))
	}

	if in.Status != 0 {
		opts = append(opts, queryx.WithCondition("status = ?", in.Status))
	}

	return queryx.NewQueryBuilder(opts...).Build()
}
