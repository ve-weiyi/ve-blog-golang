package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertRemarkPb(in *types.Remark) (out *blog.Remark) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertRemarkTypes(in *blog.Remark) (out *types.Remark) {
	jsonconv.ObjectToObject(in, &out)

	return
}
