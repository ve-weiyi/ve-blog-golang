package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertFriendLinkTypes(in *blog.FriendLink) (out *types.FriendLink) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertFriendLinkPb(in *types.FriendLink) (out *blog.FriendLink) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
