package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertOperationLogTypes(in *blog.OperationLog) (out *types.OperationLog) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertOperationLogPb(in *types.OperationLog) (out *blog.OperationLog) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
