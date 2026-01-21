package newsrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMessageListLogic {
	return &FindMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询留言列表
func (l *FindMessageListLogic) FindMessageList(in *newsrpc.FindMessageListReq) (*newsrpc.FindMessageListResp, error) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}

	if in.UserId != "" {
		opts = append(opts, query.WithCondition("user_id = ?", in.UserId))
	}

	if in.Status >= 0 {
		opts = append(opts, query.WithCondition("status = ?", in.Status))
	}

	page, size, sorts, conditions, params := query.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TMessageModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*newsrpc.Message
	for _, v := range records {
		list = append(list, convertMessageOut(v))
	}

	return &newsrpc.FindMessageListResp{
		List: list,
		Pagination: &newsrpc.PageResp{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
	}, nil
}

func convertMessageOut(in *model.TMessage) (out *newsrpc.Message) {
	out = &newsrpc.Message{
		Id:             in.Id,
		UserId:         in.UserId,
		TerminalId:     in.TerminalId,
		MessageContent: in.MessageContent,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt.UnixMilli(),
		UpdatedAt:      in.UpdatedAt.UnixMilli(),
	}

	return out
}
