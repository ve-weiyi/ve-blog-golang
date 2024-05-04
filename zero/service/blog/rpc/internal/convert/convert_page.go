package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertPagePbToModel(in *blog.Page) (out *model.Page) {
	out = &model.Page{
		Id:        in.Id,
		PageName:  in.PageName,
		PageLabel: in.PageLabel,
		PageCover: in.PageCover,
		//CreatedAt: time.Unix(in.CreatedAt, 0),
		//UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertPageModelToPb(in *model.Page) (out *blog.Page) {
	out = &blog.Page{
		Id:        in.Id,
		PageName:  in.PageName,
		PageLabel: in.PageLabel,
		PageCover: in.PageCover,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
