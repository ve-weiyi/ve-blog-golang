package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertRemarkPb(in *types.Remark) (out *blogrpc.Remark) {
	return &blogrpc.Remark{
		Id:             in.Id,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		Time:           in.Time,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
	}
}

func ConvertRemarkTypes(in *blogrpc.Remark) (out *types.Remark) {
	return &types.Remark{
		Id:             in.Id,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		Time:           in.Time,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
	}
}
