package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertTalkModelToPb(in *model.Talk) (out *blog.Talk) {
	out = &blog.Talk{
		Id:        in.Id,
		UserId:    in.UserId,
		Content:   in.Content,
		Images:    in.Images,
		IsTop:     in.IsTop,
		Status:    in.Status,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertTalkPbToModel(in *blog.Talk) (out *model.Talk) {
	out = &model.Talk{
		Id:        in.Id,
		UserId:    in.UserId,
		Content:   in.Content,
		Images:    in.Images,
		IsTop:     in.IsTop,
		Status:    in.Status,
		CreatedAt: time.Unix(in.CreatedAt, 0),
		UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}
