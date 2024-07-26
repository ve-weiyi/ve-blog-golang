package convert

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertCommentPb(in *types.CommentNewReq) (out *blog.Comment) {
	out = &blog.Comment{
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

func ConvertCommentQueryPb(in *types.CommentQueryReq) (out *blog.PageQuery) {
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

	out = &blog.PageQuery{
		Page:       page,
		PageSize:   pageSize,
		Sorts:      sorts,
		Conditions: conditions,
		Args:       args,
	}
	return
}

func ConvertCommentTypes(in *blog.Comment) (out *types.Comment) {
	out = &types.Comment{
		Id:               in.Id,
		TopicId:          in.TopicId,
		ParentId:         in.ParentId,
		SessionId:        in.SessionId,
		UserId:           in.UserId,
		ReplyUserId:      in.ReplyUserId,
		CommentContent:   in.CommentContent,
		Type:             in.Type,
		CreatedAt:        in.CreatedAt,
		LikeCount:        in.LikeCount,
		ReplyCount:       0,
		CommentReplyList: make([]*types.CommentReply, 0),
	}

	if in.User != nil {
		out.User = types.CommentUserInfo{
			Id:       in.User.UserId,
			Nickname: in.User.Nickname,
			Avatar:   in.User.Avatar,
			Website:  in.User.Info,
		}
	}

	if in.ReplyUser != nil {
		out.ReplyUser = types.CommentUserInfo{
			Id:       in.ReplyUser.UserId,
			Nickname: in.ReplyUser.Nickname,
			Avatar:   in.ReplyUser.Avatar,
			Website:  in.ReplyUser.Info,
		}
	}

	return
}

func ConvertCommentReplyTypes(in *blog.Comment) (out *types.CommentReply) {
	out = &types.CommentReply{
		Id:             in.Id,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		SessionId:      in.SessionId,
		UserId:         in.UserId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		CreatedAt:      in.CreatedAt,
		LikeCount:      in.LikeCount,
	}

	if in.User != nil {
		out.User = types.CommentUserInfo{
			Id:       in.User.UserId,
			Nickname: in.User.Nickname,
			Avatar:   in.User.Avatar,
			Website:  in.User.Info,
		}
	}

	if in.ReplyUser != nil {
		out.ReplyUser = types.CommentUserInfo{
			Id:       in.ReplyUser.UserId,
			Nickname: in.ReplyUser.Nickname,
			Avatar:   in.ReplyUser.Avatar,
			Website:  in.ReplyUser.Info,
		}
	}

	return
}
