package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertPagePb(in *types.Page) (out *blog.Page) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertPageTypes(in *blog.Page) (out *types.Page) {
	jsonconv.ObjectToObject(in, &out)

	return
}
