package convert

import (
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func EmptyReq() (out *blog.EmptyReq) {
	out = &blog.EmptyReq{}
	return
}

func ConvertIdReq(in *types.IdReq) (out *blog.IdReq) {
	out = &blog.IdReq{
		Id: in.Id,
	}
	return
}

func ConvertIdsReq(in *types.IdsReq) (out *blog.IdsReq) {
	out = &blog.IdsReq{
		Ids: in.IDS,
	}
	return
}

func ConvertPageQuery(in *types.PageQuery) (out *blog.PageQuery) {
	out = &blog.PageQuery{
		Limit: &blog.PageLimit{
			Page:     in.Limit.Page,
			PageSize: in.Limit.PageSize,
		},
		Sorts:      make([]*blog.PageSort, 0),
		Conditions: make([]*blog.PageCondition, 0),
	}

	for _, sort := range in.Sorts {
		out.Sorts = append(out.Sorts, &blog.PageSort{
			Field: sort.Field,
			Order: sort.Order,
		})
	}

	for _, condition := range in.Conditions {
		out.Conditions = append(out.Conditions, &blog.PageCondition{
			Field:    condition.Field,
			Operator: condition.Operator,
			Value:    cast.ToString(condition.Value),
		})
	}

	return
}
