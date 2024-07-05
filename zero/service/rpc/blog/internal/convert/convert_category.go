package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertCategoryPbToModel(in *blog.Category) (out *model.Category) {
	out = &model.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		//CreatedAt:    time.Unix(in.CreatedAt, 0),
		//UpdatedAt:    time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertCategoryModelToPb(in *model.Category) (out *blog.Category) {
	out = &blog.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		CreatedAt:    in.CreatedAt.Unix(),
		UpdatedAt:    in.UpdatedAt.Unix(),
	}

	return out
}
