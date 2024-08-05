package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertRemarkPb(in *types.Remark) (out *blogrpc.Remark) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertRemarkTypes(in *blogrpc.Remark) (out *types.Remark) {
	jsonconv.ObjectToObject(in, &out)

	return
}
