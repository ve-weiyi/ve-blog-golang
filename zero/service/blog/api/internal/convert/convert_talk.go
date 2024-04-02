package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertTalkPb(in *types.TalkDetails) (out *blog.Talk) {
	jsonconv.ObjectToObject(in, &out)

	out.Images = jsonconv.ObjectToJson(in.ImgList)
	return
}

func ConvertTalkTypes(in *blog.Talk) (out *types.TalkDetails) {
	jsonconv.ObjectToObject(in, &out)

	jsonconv.JsonToObject(in.Images, &out.ImgList)
	return
}
