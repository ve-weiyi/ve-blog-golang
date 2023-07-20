package initialize

import (
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

func RBAC() {
	permission := rbac.NewResourceEnforcer(global.DB)
	err := permission.LoadPermissions()
	if err != nil {
		global.LOG.Error("permission.LoadPermissions err:", err)
	}

	global.Permission = permission
}
