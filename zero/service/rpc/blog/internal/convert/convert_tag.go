package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertTagPbToModel(in *blog.Tag) (out *model.Tag) {
	out = &model.Tag{
		Id:        in.Id,
		TagName:   in.TagName,
		CreatedAt: time.Unix(in.CreatedAt, 0),
		UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertTagModelToPb(in *model.Tag) (out *blog.Tag) {
	out = &blog.Tag{
		Id:        in.Id,
		TagName:   in.TagName,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
