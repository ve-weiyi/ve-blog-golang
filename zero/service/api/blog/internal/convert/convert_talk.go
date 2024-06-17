package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertTalkTypes(in *blogrpc.Talk) (out *types.Talk) {
	out = &types.Talk{
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
