package initialize

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

func RBAC() {
	permission := rbac.NewPermissionHolder(global.DB, glog.Default())

	global.Permission = permission
}
