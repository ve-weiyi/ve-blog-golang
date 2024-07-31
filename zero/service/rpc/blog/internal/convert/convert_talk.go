package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertTalkPbToModel(in *blog.Talk) (out *model.Talk) {
	out = &model.Talk{
		Id:        in.Id,
		UserId:    in.UserId,
		Content:   in.Content,
		Images:    in.Images,
		IsTop:     in.IsTop,
		Status:    in.Status,
		LikeCount: in.LikeCount,
		CreatedAt: time.Unix(in.CreatedAt, 0),
		UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertTalkModelToPb(in *model.Talk) (out *blog.Talk) {
	out = &blog.Talk{
		Id:        in.Id,
		UserId:    in.UserId,
		Content:   in.Content,
		Images:    in.Images,
		IsTop:     in.IsTop,
		Status:    in.Status,
		LikeCount: in.LikeCount,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
