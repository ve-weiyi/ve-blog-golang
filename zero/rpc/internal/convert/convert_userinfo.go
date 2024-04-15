package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

func ConvertUserInfoModelToPb(in *model.UserInformation) (out *account.UserInfoResp) {
	out = &account.UserInfoResp{
		Id:        in.Id,
		UserId:    in.UserId,
		Email:     in.Email,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Phone:     in.Phone,
		Intro:     in.Intro,
		Website:   in.Website,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertUserDetailsModelToPb(a *model.UserAccount, info *model.UserInformation, roles []*model.Role) (out *account.UserDTO) {
	var list []*account.RoleLabelDTO
	for _, role := range roles {
		m := &account.RoleLabelDTO{
			RoleName:    role.RoleName,
			RoleComment: role.RoleComment,
		}

		list = append(list, m)
	}

	out = &account.UserDTO{
		Id:           a.Id,
		Username:     a.Username,
		Email:        info.Email,
		Nickname:     info.Nickname,
		Avatar:       info.Avatar,
		Phone:        info.Phone,
		Intro:        info.Intro,
		Website:      info.Website,
		Status:       a.Status,
		RegisterType: a.RegisterType,
		IpAddress:    a.IpAddress,
		IpSource:     a.IpSource,
		Roles:        list,
		CreatedAt:    info.CreatedAt.Unix(),
		UpdatedAt:    info.UpdatedAt.Unix(),
	}

	return out
}
