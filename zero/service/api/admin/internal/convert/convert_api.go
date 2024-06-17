package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertApiPb(in *types.Api) (out *blogrpc.Api) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertApiTypes(in *blogrpc.Api) (out *types.Api) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertApiDetailsTypes(in *blogrpc.ApiDetails) (out *types.ApiDetails) {
	jsonconv.ObjectToObject(in, &out)
	return
}
