package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
)

func ConvertUserInfoModelToPb(in *model.UserAccount) (out *blog.UserInfoResp) {
	out = &blog.UserInfoResp{
		UserId:    in.Id,
		Username:  in.Username,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Email:     in.Email,
		Phone:     in.Phone,
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

func ConvertUserDetailsModelToPb(a *model.UserAccount, roles []*model.Role) (out *blog.UserDetails) {
	var list []*blog.RoleLabel
	for _, role := range roles {
		m := &blog.RoleLabel{
			RoleId:      role.Id,
			RoleName:    role.RoleName,
			RoleComment: role.RoleComment,
		}

		list = append(list, m)
	}

	out = &blog.UserDetails{
		UserId:    a.Id,
		Username:  a.Username,
		Nickname:  a.Nickname,
		Avatar:    a.Avatar,
		Email:     a.Email,
		Phone:     a.Phone,
		Info:      a.Info,
		Status:    a.Status,
		LoginType: a.LoginType,
		IpAddress: a.IpAddress,
		IpSource:  a.IpSource,
		CreatedAt: a.CreatedAt.Unix(),
		UpdatedAt: a.UpdatedAt.Unix(),
		LoginAt:   0,
		LogoutAt:  0,
		Roles:     list,
	}

	return out
}

func ConvertUserLoginHistoryModelToPb(in *model.UserLoginHistory) (out *blog.LoginHistory) {
	out = &blog.LoginHistory{
		Id:        in.Id,
		LoginType: in.LoginType,
		Agent:     in.Agent,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		LoginTime: in.CreatedAt.String(),
	}

	return out
}
