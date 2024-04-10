package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/apirpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/menurpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/userrpc"
)

type ServiceContext struct {
	Config config.Config
	Token  *jjwt.JwtToken

	AccountRpc accountrpc.AccountRpc
	ApiRpc     apirpc.ApiRpc
	MenuRpc    menurpc.MenuRpc
	RoleRpc    rolerpc.RoleRpc
	UserRpc    userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Token:      jjwt.NewJwtToken([]byte("ve-weiyi")),
		AccountRpc: accountrpc.NewAccountRpc(zrpc.MustNewClient(c.AccountRpcConf)),
		ApiRpc:     apirpc.NewApiRpc(zrpc.MustNewClient(c.ApiRpcConf)),
		MenuRpc:    menurpc.NewMenuRpc(zrpc.MustNewClient(c.MenuRpcConf)),
		RoleRpc:    rolerpc.NewRoleRpc(zrpc.MustNewClient(c.RoleRpcConf)),
		UserRpc:    userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
