package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

func Test_ParseTableFormDsn(t *testing.T) {
	const dsn = "root:mysql7914@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	fromDsn, err := ParseTableFromDsn(dsn)
	t.Log(err)
	//t.Log(jsonconv.ObjectToJsonIndent(fromDsn))

	const sql = "test.sql"
	fromSql, err := ParseTableFromSql(sql)
	t.Log(err)
	//t.Log(jsonconv.ObjectToJsonIndent(fromSql))

	assert.Equal(t, jsonconv.ObjectToJsonIndent(fromDsn), jsonconv.ObjectToJsonIndent(fromSql))
}
