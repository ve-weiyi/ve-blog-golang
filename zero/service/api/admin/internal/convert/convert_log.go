package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertOperationLogPb(in *types.OperationLog) (out *blogrpc.OperationLog) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertOperationLogTypes(in *blogrpc.OperationLog) (out *types.OperationLog) {
	jsonconv.ObjectToObject(in, &out)

	return
}
