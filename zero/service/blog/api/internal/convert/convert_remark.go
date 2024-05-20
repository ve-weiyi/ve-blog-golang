package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertRemarkPb(in *types.Remark) (out *blog.Remark) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertRemarkTypes(in *blog.Remark) (out *types.Remark) {
	jsonconv.ObjectToObject(in, &out)

	return
}
