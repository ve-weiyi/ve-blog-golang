package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/authrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertUserInfo(in *authrpc.LoginResp) (out *types.UserInfo) {
	out = &types.UserInfo{
		UserId:   in.UserId,
		Username: in.Username,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Intro:    in.Intro,
		Website:  in.Website,
		Email:    in.Email,
	}
	return
}
