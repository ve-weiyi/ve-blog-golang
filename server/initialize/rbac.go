package initialize

import (
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

func RBAC() {
	global.Permission = rbac.NewPermissionHolder(global.DB, global.LOG)
}
