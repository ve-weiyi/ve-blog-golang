package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertApiModelToPb(in *model.Api) (out *blog.Api) {
	out = &blog.Api{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertApiPbToModel(in *blog.Api) (out *model.Api) {
	out = &model.Api{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
		CreatedAt: time.Unix(in.CreatedAt, 0),
		UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertApiModelToDetailPb(in *model.Api) (out *blog.ApiDetails) {
	out = &blog.ApiDetails{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
