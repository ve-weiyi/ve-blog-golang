package initialize

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/logic"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

func RBAC() {
	GORM := global.DB
	if GORM == nil {
		panic("db is null")
	}
	////会自动创建数据库表并管理
	//adapter, err := gormadapter.NewAdapterByDB(GORM)
	//
	//m, err := model.NewModelFromString(SubjectDomainObjectAction)
	//if err != nil {
	//	log.Fatalln("字符串加载模型失败!", err)
	//}
	//
	//syncedCachedEnforcer, _ := casbin.NewSyncedCachedEnforcer(m, adapter)
	//syncedCachedEnforcer.SetExpireTime(60 * 60)
	//_ = syncedCachedEnforcer.LoadPolicy()

	enforcer := rbac.NewCachedEnforcer(GORM)
	var apis []*entity.Api
	ap := logic.ApiRepository{DbEngin: global.DB}
	apis, err := ap.FindAllPublicApis(context.Background())
	if err != nil {
		return
	}
	enforcer.LoadWhileList(apis)

	global.RbacEnforcer = enforcer
}
