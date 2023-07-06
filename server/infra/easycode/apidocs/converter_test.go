package apidocs

import (
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"testing"
)

func TestSwagger(t *testing.T) {
	converter := &SwaggerConverter{}

	converter.ReadSwagJSON(global.GetRuntimeRoot() + "/server/docs/swagger.json")

	converter.toTypeScriptApis("./preview")
}
