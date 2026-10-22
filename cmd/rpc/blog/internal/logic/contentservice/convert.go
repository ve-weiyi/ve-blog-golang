package contentservicelogic

import (
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/model"
)

func convertCategoryOut(in *model.TCategory) *contentrpc.Category {
	return &contentrpc.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		CreatedAt:    in.CreatedAt.UnixMilli(),
		UpdatedAt:    in.UpdatedAt.UnixMilli(),
	}
}

func convertTagOut(in *model.TTag) *contentrpc.Tag {
	return &contentrpc.Tag{
		Id:        in.Id,
		TagName:   in.TagName,
		CreatedAt: in.CreatedAt.UnixMilli(),
		UpdatedAt: in.UpdatedAt.UnixMilli(),
	}
}
