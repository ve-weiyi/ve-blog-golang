package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertPagePb(in *types.Page) (out *blogrpc.Page) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertPageTypes(in *blogrpc.Page) (out *types.Page) {
	jsonconv.ObjectToObject(in, &out)

	return
}
