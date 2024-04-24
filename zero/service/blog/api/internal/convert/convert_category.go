package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertCategoryTypes(in *blog.Category) (out *types.Category) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertCategoryPb(in *types.Category) (out *blog.Category) {
	jsonconv.ObjectMarshal(in, &out)
	return
}

func ConvertCategoryDetailsTypes(in *blog.Category) (out *types.CategoryDetails) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
