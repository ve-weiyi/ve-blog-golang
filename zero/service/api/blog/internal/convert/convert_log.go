package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertOperationLogPb(in *types.OperationLog) (out *blog.OperationLog) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertOperationLogTypes(in *blog.OperationLog) (out *types.OperationLog) {
	jsonconv.ObjectToObject(in, &out)

	return
}
