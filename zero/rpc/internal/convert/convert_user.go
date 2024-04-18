package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/repository/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

func ConvertUserLoginHistoryModelToPb(in *model.UserLoginHistory) (out *account.LoginHistory) {
	out = &account.LoginHistory{
		Id:        in.Id,
		LoginType: in.LoginType,
		Agent:     in.Agent,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		LoginTime: in.CreatedAt.String(),
	}

	return out
}
