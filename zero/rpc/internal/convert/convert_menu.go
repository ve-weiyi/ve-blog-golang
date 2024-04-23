package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/repository/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

func ConvertMenuModelToPb(in *model.Menu) (out *account.Menu) {
	out = &account.Menu{
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

func ConvertMenuPbToModel(in *account.Menu) (out *model.Menu) {
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
		CreatedAt: time.Unix(in.CreatedAt, 0),
		UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertMenuModelToDetailPb(in *model.Menu) (out *account.MenuDetails) {
	out = &account.MenuDetails{
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
