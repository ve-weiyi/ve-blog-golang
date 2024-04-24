package convert

import (
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/rolerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func EmptyReq() (out *rolerpc.EmptyReq) {
	out = &rolerpc.EmptyReq{}
	return
}

func ConvertIdReq(in *types.IdReq) (out *rolerpc.IdReq) {
	out = &rolerpc.IdReq{
		Id: in.Id,
	}
	return
}

func ConvertIdsReq(in *types.IdsReq) (out *rolerpc.IdsReq) {
	out = &rolerpc.IdsReq{
		Ids: in.IDS,
	}
	return
}

func ConvertPageQuery(in *types.PageQuery) (out *rolerpc.PageQuery) {
	out = &rolerpc.PageQuery{
		Limit: &rolerpc.PageLimit{
			Page:     in.Limit.Page,
			PageSize: in.Limit.PageSize,
		},
		Sorts:      make([]*blog.PageSort, 0),
		Conditions: make([]*blog.PageCondition, 0),
	}

	for _, sort := range in.Sorts {
		out.Sorts = append(out.Sorts, &rolerpc.PageSort{
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
