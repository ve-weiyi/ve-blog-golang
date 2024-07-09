package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
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
