package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertUserDetailsTypes(in *blog.User) (out *types.User) {
	jsonconv.ObjectMarshal(in, &out)
	return out
}

func ConvertUserLoginHistoryTypes(in *blog.LoginHistory) (out *types.LoginHistory) {
	jsonconv.ObjectMarshal(in, &out)
	return out
}

func ConvertUserMenuTypes(in *blog.MenuDetails) (out *types.UserMenu) {
	jsonconv.ObjectMarshal(in, &out)

	out.Children = make([]*types.UserMenu, 0)
	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}
