package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertTalkPb(in *types.TalkDetails) (out *blogrpc.Talk) {
	out = &blogrpc.Talk{
		Id:        in.Id,
		UserId:    in.UserId,
		Content:   in.Content,
		Images:    jsonconv.ObjectToJson(in.ImgList),
		IsTop:     in.IsTop,
		Status:    in.Status,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		LikeCount: in.LikeCount,
	}

	return
}

func ConvertTalkTypes(in *blogrpc.Talk) (out *types.TalkDetails) {
	out = &types.TalkDetails{
		Id:        in.Id,
		UserId:    in.UserId,
		Content:   in.Content,
		IsTop:     in.IsTop,
		Status:    in.Status,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		LikeCount: in.LikeCount,
	}
	jsonconv.JsonToObject(in.Images, &out.ImgList)

	return
}
