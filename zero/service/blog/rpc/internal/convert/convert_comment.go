package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertCommentPbToModel(in *blog.Comment) (out *model.Comment) {
	out = &model.Comment{
		Id:             in.Id,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		SessionId:      in.SessionId,
		UserId:         in.UserId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
		IsReview:       in.IsReview,
		//CreatedAt:      time.Time{},
		//UpdatedAt:      time.Time{},
	}

	return out
}

func ConvertCommentModelToPb(in *model.Comment) (out *blog.Comment) {
	out = &blog.Comment{
		Id:             in.Id,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		SessionId:      in.SessionId,
		UserId:         in.UserId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertCommentReplyPb(in *model.Comment) (out *blog.CommentReply) {
	out = &blog.CommentReply{
		Id:             in.Id,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		SessionId:      in.SessionId,
		UserId:         in.UserId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
		User:           nil,
		ReplyUser:      nil,
		LikeCount:      in.LikeCount,
	}

	return out
}
