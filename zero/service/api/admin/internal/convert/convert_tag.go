package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertTagPb(in *types.Tag) (out *blog.Tag) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertTagTypes(in *blog.Tag) (out *types.Tag) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertTagDetailsTypes(in *blog.Tag) (out *types.TagDetails) {
	jsonconv.ObjectToObject(in, &out)
	return
}
