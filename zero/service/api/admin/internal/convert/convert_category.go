package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertCategoryPb(in *types.Category) (out *blog.Category) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertCategoryTypes(in *blog.Category) (out *types.Category) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertCategoryDetailsTypes(in *blog.Category) (out *types.CategoryDetails) {
	jsonconv.ObjectToObject(in, &out)
	return
}
