package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
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
