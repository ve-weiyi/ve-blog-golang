package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertOperationLogPb(in *types.OperationLog) (out *blog.OperationLog) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertOperationLogTypes(in *blog.OperationLog) (out *types.OperationLog) {
	jsonconv.ObjectToObject(in, &out)

	return
}
