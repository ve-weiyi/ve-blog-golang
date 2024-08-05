package convert

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertCommentPb(in *types.CommentNewReq) (out *blogrpc.Comment) {
	out = &blogrpc.Comment{
		Id:             0,
		ParentId:       in.ParentId,
		TopicId:        in.TopicId,
		SessionId:      in.SessionId,
		UserId:         0,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         0,
		IsReview:       0,
		CreatedAt:      0,
		UpdatedAt:      0,
	}
	return
}

func ConvertCommentQueryTypes(in *types.CommentQueryReq) (out *blogrpc.PageQuery) {
	var page, pageSize int64
	var sorts, conditions string
	var args []string

	page = in.Page
	pageSize = in.PageSize

	if in.OrderBy != "" {
		sorts = fmt.Sprintf("`%s` desc", in.OrderBy)
	}

	if in.TopicId >= 0 {
		conditions = "topic_id = ? "
		args = append(args, cast.ToString(in.TopicId))
	}

	if in.ParentId >= 0 {
		conditions = conditions + "and "
		conditions = conditions + "parent_id = ? "
		args = append(args, cast.ToString(in.ParentId))
	}

	if in.Type >= 0 {
		conditions = conditions + "and "
		conditions = conditions + "type = ? "
		args = append(args, cast.ToString(in.Type))
	}

	out = &blogrpc.PageQuery{
		Page:       page,
		PageSize:   pageSize,
		Sorts:      sorts,
		Conditions: conditions,
		Args:       args,
	}
	return
}

func ConvertCommentTypes(in *blogrpc.Comment) (out *types.CommentNewReq) {
	return &types.CommentNewReq{
		ParentId:       in.ParentId,
		TopicId:        in.TopicId,
		SessionId:      in.SessionId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
	}
}

func ConvertCommentBackTypes(in *blogrpc.Comment) (out *types.CommentBackDTO) {
	out = &types.CommentBackDTO{
		Id:             in.Id,
		Type:           in.Type,
		TopicTitle:     "",
		Avatar:         "",
		Nickname:       "",
		CommentContent: in.CommentContent,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt,
	}

	return
}
