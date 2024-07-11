package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertUserDetailsTypes(in *blog.User) (out *types.User) {
	jsonconv.ObjectToObject(in, &out)
	out.Id = in.UserId
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

func ConvertUserInfoTypes(in *blog.UserInfoResp) (out *types.UserInfoResp) {
	out = &types.UserInfoResp{
		UserId:      in.UserId,
		Username:    in.Username,
		Nickname:    in.Nickname,
		Avatar:      in.Avatar,
		UserInfoExt: types.UserInfoExt{},
	}

	jsonconv.JsonToObject(in.Info, &out.UserInfoExt)

	return out
}
