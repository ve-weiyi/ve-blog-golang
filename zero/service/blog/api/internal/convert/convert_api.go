package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertApiPb(in *types.Api) (out *blog.Api) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertApiTypes(in *blog.Api) (out *types.Api) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertApiDetailsTypes(in *blog.ApiDetails) (out *types.ApiDetails) {
	jsonconv.ObjectToObject(in, &out)
	return
}
