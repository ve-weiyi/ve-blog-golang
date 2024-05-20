package convert

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
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

func ConvertCommentQueryTypes(in *types.CommentQueryReq) (out *blog.PageQuery) {
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

func ConvertCommentTypes(in *blog.Comment) (out *types.CommentNewReq) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertCommentDTOTypes(in *blog.CommentReply) (out *types.CommentDTO) {
	out = &types.CommentDTO{
		Id:               in.Id,
		TopicId:          in.TopicId,
		ParentId:         in.ParentId,
		SessionId:        in.SessionId,
		UserId:           0,
		Nickname:         "",
		Avatar:           "",
		Website:          "",
		ReplyUserId:      0,
		ReplyNickname:    "",
		ReplyAvatar:      "",
		ReplyWebsite:     "",
		CommentContent:   in.CommentContent,
		Type:             in.Type,
		CreatedAt:        in.CreatedAt,
		LikeCount:        in.LikeCount,
		ReplyCount:       0,
		CommentReplyList: make([]*types.CommentReply, 0),
	}

	if in.User != nil {
		out.UserId = in.User.UserId
		out.Avatar = in.User.Avatar
		out.Nickname = in.User.Nickname
	}

	if in.ReplyUser != nil {
		out.ReplyUserId = in.ReplyUser.UserId
		out.ReplyAvatar = in.ReplyUser.Avatar
		out.ReplyNickname = in.ReplyUser.Nickname
	}

	return
}

func ConvertCommentReplyTypes(in *blog.CommentReply) (out *types.CommentReply) {
	out = &types.CommentReply{
		Id:             in.Id,
		ParentId:       in.ParentId,
		TopicId:        in.TopicId,
		UserId:         0,
		Nickname:       "",
		Avatar:         "",
		Website:        "",
		ReplyUserId:    0,
		ReplyNickname:  "",
		ReplyAvatar:    "",
		ReplyWebsite:   "",
		CommentContent: in.CommentContent,
		Type:           in.Type,
		CreatedAt:      in.CreatedAt,
		LikeCount:      in.LikeCount,
	}

	if in.User != nil {
		out.UserId = in.User.UserId
		out.Avatar = in.User.Avatar
		out.Nickname = in.User.Nickname
	}

	if in.ReplyUser != nil {
		out.ReplyUserId = in.ReplyUser.UserId
		out.ReplyAvatar = in.ReplyUser.Avatar
		out.ReplyNickname = in.ReplyUser.Nickname
	}

	return
}

func ConvertCommentBackTypes(in *blog.Comment) (out *types.CommentBackDTO) {
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
