package convert

import (
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertUserDetailsTypes(in *accountrpc.UserDetails) (out *types.User) {
	jsonconv.ObjectToObject(in, &out)
	out.Id = in.UserId
	return out
}

func ConvertUserLoginHistoryTypes(in *accountrpc.UserLoginHistory) (out *types.UserLoginHistory) {
	jsonconv.ObjectToObject(in, &out)
	return out
}

func ConvertUserMenuTypes(in *permissionrpc.MenuDetails) (out *types.UserMenu) {
	out = &types.UserMenu{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Title,
		Type:      in.Type,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Meta:      types.UserMenuMeta{},
		Children:  make([]*types.UserMenu, 0),
	}

	jsonconv.JsonToObject(in.Extra, &out.Meta)
	if !strings.Contains(in.Path, ":") {
		out.Meta.ShowLink = true
		out.Meta.ShowParent = true
	}

	for _, v := range in.Children {
		out.Children = append(out.Children, ConvertUserMenuTypes(v))
	}

	return
}

func ConvertUserInfoTypes(in *accountrpc.UserInfoResp) (out *types.UserInfoResp) {
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
