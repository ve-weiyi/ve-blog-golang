package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertTalkPb(in *types.TalkDetails) (out *blogrpc.Talk) {
	jsonconv.ObjectToObject(in, &out)

	out.Images = jsonconv.ObjectToJson(in.ImgList)
	return
}

func ConvertTalkTypes(in *blogrpc.Talk) (out *types.TalkDetails) {
	jsonconv.ObjectToObject(in, &out)

	jsonconv.JsonToObject(in.Images, &out.ImgList)
	return
}
