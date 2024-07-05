package parsex

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	model2 "github.com/ve-weiyi/ve-blog-golang/zero/tools/quickstart/gin/cmd/model"
)

func Test_ParseTableFormDsn(t *testing.T) {
	const dsn = "root:mysql7914@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	fromDsn, err := model2.ParseTableFromDsn(dsn)
	t.Log(err)
	//t.Log(jsonconv.ObjectToJsonIndent(fromDsn))

	const sql = "test.sql"
	fromSql, err := model2.ParseTableFromSql(sql)
	t.Log(err)
	//t.Log(jsonconv.ObjectToJsonIndent(fromSql))

	assert.Equal(t, jsonconv.ObjectToJsonIndent(fromDsn), jsonconv.ObjectToJsonIndent(fromSql))
}

func Test_ParseTableFromSql(t *testing.T) {
	const sql = "test.sql"
	fromSql, err := model2.ParseTableFromSql(sql)
	t.Log(err)
	t.Log(jsonconv.ObjectToJsonIndent(fromSql))
}

func Test_ParseAPI(t *testing.T) {
	//const api = "../test.api"
	const api = "/Users/weiyi/Github/ve-blog-golang/zero/service/api/blog/proto/module/category.api"
	fromApi, err := ParseAPI(api)
	t.Log(err)
	//t.Log(jsonconv.ObjectToJsonIndent(fromApi))

	for _, v := range fromApi.Service.Groups {
		t.Log("--->", v.Annotation.Properties["group"])
		for _, v2 := range v.Routes {
			t.Log(v2.Path)
		}
		//t.Log(jsonconv.ObjectToJsonIndent(v))
	}

	//for _, v := range fromApi.Types {
	//	t.Log(jsonconv.ObjectToJsonIndent(v))
	//}

}
