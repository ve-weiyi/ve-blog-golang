package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertFriendTypes(in *blog.FriendLink) (out *types.Friend) {
	jsonconv.ObjectToObject(in, &out)

	return
}
