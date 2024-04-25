package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertRemarkTypes(in *blog.Remark) (out *types.Remark) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertRemarkPb(in *types.Remark) (out *blog.Remark) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
