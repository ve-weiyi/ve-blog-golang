package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertCommentModelToPb(in *model.Comment) (out *blog.Comment) {
	out = &blog.Comment{
		Id:             in.Id,
		UserId:         in.UserId,
		TopicId:        in.TopicId,
		CommentContent: in.CommentContent,
		ReplyUserId:    in.ReplyUserId,
		ParentId:       in.ParentId,
		Type:           in.Type,
		IsDelete:       in.IsDelete,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertCommentPbToModel(in *blog.Comment) (out *model.Comment) {
	out = &model.Comment{
		Id:             in.Id,
		UserId:         in.UserId,
		TopicId:        in.TopicId,
		CommentContent: in.CommentContent,
		ReplyUserId:    in.ReplyUserId,
		ParentId:       in.ParentId,
		Type:           in.Type,
		IsDelete:       in.IsDelete,
		IsReview:       in.IsReview,
		CreatedAt:      time.Unix(in.CreatedAt, 0),
		UpdatedAt:      time.Unix(in.UpdatedAt, 0),
	}

	return out
}
