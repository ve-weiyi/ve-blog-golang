package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertMenuPbToModel(in *blog.Menu) (out *model.Menu) {
	out = &model.Menu{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Name,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Type:      in.Type,
		Extra:     in.Extra,
		// CreatedAt: time.Unix(in.CreatedAt, 0),
		// UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertMenuModelToPb(in *model.Menu) (out *blog.Menu) {
	out = &blog.Menu{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Name,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Type:      in.Type,
		Extra:     in.Extra,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}
	return out
}

func ConvertMenuModelToDetailPb(in *model.Menu) (out *blog.MenuDetails) {
	out = &blog.MenuDetails{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Name,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Type:      in.Type,
		Extra:     in.Extra,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
