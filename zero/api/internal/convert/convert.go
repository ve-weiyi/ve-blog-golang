package convert

import (
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

func ConvertIdReq(in *types.IdReq) (out *rolerpc.IdReq) {
	out = &rolerpc.IdReq{
		Id: in.ID,
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
		Sorts:      make([]*account.PageSort, 0),
		Conditions: make([]*account.PageCondition, 0),
	}

	for _, sort := range in.Sorts {
		out.Sorts = append(out.Sorts, &rolerpc.PageSort{
			Field: sort.Field,
			Order: sort.Order,
		})
	}

	for _, condition := range in.Conditions {
		out.Conditions = append(out.Conditions, &account.PageCondition{
			Field:    condition.Field,
			Operator: condition.Operator,
			Value:    cast.ToString(condition.Value),
		})
	}

	return
}
