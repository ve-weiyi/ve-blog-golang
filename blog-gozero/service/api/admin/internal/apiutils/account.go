package apiutils

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/accountrpc"
)

func GetUserInfos(ctx context.Context, svcCtx *svc.ServiceContext, uids []string) (map[string]*accountrpc.User, error) {
	users, err := svcCtx.AccountRpc.FindUserList(ctx, &accountrpc.FindUserListReq{
		UserIds: uids,
	})
	if err != nil {
		return nil, err
	}

	usm := make(map[string]*accountrpc.User)
	for _, v := range users.List {
		usm[v.UserId] = v
	}

	return usm, err
}
