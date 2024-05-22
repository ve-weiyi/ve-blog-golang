package initest

import (
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

func TestRbacHolder(t *testing.T) {
	Init()
	r := rbac.NewPermissionHolder(global.DB, glog)
	permission, err := r.FindApiPermission("/api/v1/user", "GET")
	t.Log(err)

	t.Log(jsonconv.ObjectToJsonIndent(permission))
}
