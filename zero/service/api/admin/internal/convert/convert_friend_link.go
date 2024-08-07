package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertFriendLinkPb(in *types.Friend) (out *blogrpc.FriendLink) {
	return &blogrpc.FriendLink{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}

func ConvertFriendLinkTypes(in *blogrpc.FriendLink) (out *types.Friend) {

	return &types.Friend{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}
