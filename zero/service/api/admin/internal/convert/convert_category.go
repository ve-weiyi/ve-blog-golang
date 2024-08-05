package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertCategoryPb(in *types.Category) (out *blogrpc.Category) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertCategoryTypes(in *blogrpc.Category) (out *types.Category) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertCategoryDetailsTypes(in *blogrpc.Category) (out *types.CategoryDetails) {
	jsonconv.ObjectToObject(in, &out)
	return
}
