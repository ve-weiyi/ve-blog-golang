package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertUserInfoModelToPb(in *model.UserAccount) (out *blog.UserInfoResp) {
	out = &blog.UserInfoResp{
		UserId:    in.Id,
		Username:  in.Username,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Info:      in.Info,
		Status:    in.Status,
		LoginType: in.LoginType,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertUserDetailsModelToPb(a *model.UserAccount, roles []*model.Role) (out *blog.User) {
	var list []*blog.RoleLabel
	for _, role := range roles {
		m := &blog.RoleLabel{
			RoleId:      role.Id,
			RoleName:    role.RoleName,
			RoleComment: role.RoleComment,
		}

		list = append(list, m)
	}

	out = &blog.User{
		UserId:    a.Id,
		Username:  a.Username,
		Nickname:  a.Nickname,
		Avatar:    a.Avatar,
		Info:      a.Info,
		Status:    a.Status,
		LoginType: a.LoginType,
		IpAddress: a.IpAddress,
		IpSource:  a.IpSource,
		CreatedAt: a.CreatedAt.Unix(),
		UpdatedAt: a.UpdatedAt.Unix(),
		LoginAt:   a.LoginAt.Unix(),
		LogoutAt:  a.LogoutAt.Unix(),
		Roles:     list,
	}

	return out
}
