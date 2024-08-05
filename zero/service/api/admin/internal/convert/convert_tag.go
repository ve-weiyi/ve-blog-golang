package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertTagPb(in *types.Tag) (out *blogrpc.Tag) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertTagTypes(in *blogrpc.Tag) (out *types.Tag) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertTagDetailsTypes(in *blogrpc.Tag) (out *types.TagDetails) {
	jsonconv.ObjectToObject(in, &out)
	return
}
