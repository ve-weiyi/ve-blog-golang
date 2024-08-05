package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertFriendLinkPb(in *types.FriendLink) (out *blogrpc.FriendLink) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertFriendLinkTypes(in *blogrpc.FriendLink) (out *types.FriendLink) {
	jsonconv.ObjectToObject(in, &out)

	return
}
