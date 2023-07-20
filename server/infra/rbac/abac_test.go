package rbac

import (
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/testinit"
)

func TestResourceEnforcer(t *testing.T) {
	testinit.Init()

	abac := NewResourceEnforcer()
	err := abac.LoadPermissions()
	t.Log(err)

	ok, err := abac.VerifyUserPermissions(10, "/api/v1/version", "GET")
	t.Log(ok, err)
}
