package svc

import (
	"log"

	"github.com/zeromicro/go-zero/core/conf"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/config"
)

func NewTestServiceContext() *ServiceContext {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	//c := NewTestConfig()

	c := config.Config{}
	conf.MustLoad("../../etc/admin-api.yaml", &c)
	return NewServiceContext(c)
}
