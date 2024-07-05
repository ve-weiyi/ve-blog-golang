package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertUserDetailsTypes(in *blog.User) (out *types.User) {
	jsonconv.ObjectToObject(in, &out)
	return out
}

func ConvertUserLoginHistoryTypes(in *blog.LoginHistory) (out *types.LoginHistory) {
	jsonconv.ObjectToObject(in, &out)
	return out
}

func ConvertUserMenuTypes(in *blog.MenuDetails) (out *types.UserMenu) {
	jsonconv.ObjectToObject(in, &out)

	out.Children = make([]*types.UserMenu, 0)
	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}
