package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
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
		// CreatedAt:      time.Unix(in.CreatedAt, 0),
		// UpdatedAt:      time.Unix(in.UpdatedAt, 0),
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
		LikeCount:      in.LikeCount,
	}

	return out
}

func ConvertCommentUserInfoToPb(in *model.UserAccount) (out *blog.CommentUserInfo) {
	return &blog.CommentUserInfo{
		UserId:   in.Id,
		Username: in.Username,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Info:     in.Info,
	}
}
