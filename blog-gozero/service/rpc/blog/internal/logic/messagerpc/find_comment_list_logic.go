package messagerpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentListLogic {
	return &FindCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取评论列表
func (l *FindCommentListLogic) FindCommentList(in *messagerpc.FindCommentListReq) (*messagerpc.FindCommentListResp, error) {
	page, size, sorts, conditions, params := convertCommentQuery(in)

	records, total, err := l.svcCtx.TCommentModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*messagerpc.CommentDetails
	for _, v := range records {
		m := convertCommentOut(v)
		list = append(list, m)
	}

	return &messagerpc.FindCommentListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertCommentQuery(in *messagerpc.FindCommentListReq) (page int, size int, sorts string, conditions string, params []any) {
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

func convertCommentOut(in *model.TComment) (out *messagerpc.CommentDetails) {
	out = &messagerpc.CommentDetails{
		Id:             in.Id,
		UserId:         in.UserId,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		ReplyMsgId:     in.ReplyMsgId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
		LikeCount:      in.LikeCount,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
	}

	return out
}
