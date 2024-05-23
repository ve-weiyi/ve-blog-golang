package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/authrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
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
