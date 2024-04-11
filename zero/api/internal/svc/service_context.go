package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/apirpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/menurpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"
)

type ServiceContext struct {
	Config config.Config
	Token  *jjwt.JwtToken

	AccountRpc accountrpc.AccountRpc
	RoleRpc    rolerpc.RoleRpc
	ApiRpc     apirpc.ApiRpc
	MenuRpc    menurpc.MenuRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Token:      jjwt.NewJwtToken([]byte("ve-weiyi")),
		AccountRpc: accountrpc.NewAccountRpc(zrpc.MustNewClient(c.AccountRpcConf)),
		RoleRpc:    rolerpc.NewRoleRpc(zrpc.MustNewClient(c.RoleRpcConf)),
		ApiRpc:     apirpc.NewApiRpc(zrpc.MustNewClient(c.ApiRpcConf)),
		MenuRpc:    menurpc.NewMenuRpc(zrpc.MustNewClient(c.MenuRpcConf)),
	}
}
