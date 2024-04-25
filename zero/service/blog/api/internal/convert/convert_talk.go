package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertTalkTypes(in *blog.Talk) (out *types.TalkDetails) {
	jsonconv.ObjectMarshal(in, &out)

	jsonconv.ObjectMarshal(in.Images, &out.ImgList)
	return
}

func ConvertTalkPb(in *types.TalkDetails) (out *blog.Talk) {
	jsonconv.ObjectMarshal(in, &out)

	jsonconv.ObjectMarshal(in.ImgList, &out.Images)
	return
}
