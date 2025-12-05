package messagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentReplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentReplyListLogic {
	return &FindCommentReplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询评论回复列表
func (l *FindCommentReplyListLogic) FindCommentReplyList(in *messagerpc.FindCommentReplyListReq) (*messagerpc.FindCommentReplyListResp, error) {
	page, size, sorts, conditions, params := convertCommentReplyQuery(in)

	records, total, err := l.svcCtx.TCommentModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*messagerpc.CommentDetailsResp
	for _, v := range records {
		list = append(list, convertCommentOut(v))
	}

	return &messagerpc.FindCommentReplyListResp{
		List: list,
		Pagination: &messagerpc.PageResp{
			Page:     in.Paginate.Page,
			PageSize: in.Paginate.PageSize,
			Total:    total,
		},
	}, nil
}

func convertCommentReplyQuery(in *messagerpc.FindCommentReplyListReq) (page int, size int, sorts string, conditions string, params []any) {
	opts := []query.Option{
		query.WithPage(int(in.Paginate.Page)),
		query.WithSize(int(in.Paginate.PageSize)),
		query.WithSorts(in.Paginate.Sorts...),
		query.WithCondition("parent_id = ?", in.ParentId),
	}

	if in.Type != 0 {
		opts = append(opts, query.WithCondition("type = ?", in.Type))
	}

	if in.TopicId != 0 {
		opts = append(opts, query.WithCondition("topic_id = ?", in.TopicId))
	}

	return query.NewQueryBuilder(opts...).Build()
}
