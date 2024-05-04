package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertPagePb(in *types.Page) (out *blog.Page) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertPageTypes(in *blog.Page) (out *types.Page) {
	jsonconv.ObjectToObject(in, &out)

	return
}
