package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/rolerpc"
)

func ConvertApiTypes(in *rolerpc.Api) (out *types.Api) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertApiPb(in *types.Api) (out *rolerpc.Api) {
	jsonconv.ObjectMarshal(in, &out)
	return
}

func ConvertApiDetailsTypes(in *rolerpc.ApiDetails) (out *types.ApiDetails) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
