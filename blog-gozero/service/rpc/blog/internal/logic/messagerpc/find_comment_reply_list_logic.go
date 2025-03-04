package messagerpclogic

import (
	"context"
	"strings"

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

	var list []*messagerpc.CommentDetails
	for _, v := range records {
		list = append(list, convertCommentOut(v))
	}

	return &messagerpc.FindCommentReplyListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertCommentReplyQuery(in *messagerpc.FindCommentReplyListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	if in.Type != 0 {
		conditions += " type = ?"
		params = append(params, in.Type)
	}

	if in.TopicId != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " topic_id = ?"
		params = append(params, in.TopicId)
	}

	if conditions != "" {
		conditions += " and "
	}
	conditions += " parent_id = ?"
	params = append(params, in.ParentId)

	return
}
