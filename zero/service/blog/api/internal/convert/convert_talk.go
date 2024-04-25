package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertTalkTypes(in *blog.Talk) (out *types.Talk) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertTalkPb(in *types.Talk) (out *blog.Talk) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
