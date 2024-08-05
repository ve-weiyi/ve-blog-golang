package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
)

func ConvertPageTypes(in *blogrpc.Page) (out *types.Page) {
	jsonconv.ObjectToObject(in, &out)

	return
}
