package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertFriendLinkPb(in *types.FriendLink) (out *blog.FriendLink) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertFriendLinkTypes(in *blog.FriendLink) (out *types.FriendLink) {
	jsonconv.ObjectToObject(in, &out)

	return
}
