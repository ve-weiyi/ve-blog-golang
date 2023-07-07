package fmtplus

import (
	"log"
	"path"
	"strings"
	"testing"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
)

func TestString(t *testing.T) {
	addImportCode := "hello  github.com/ve-weiyi/ve-blog-golang/server/api/v1/test"
	var (
		importAlias   string
		importPackage string
		packageName   string
	)
	// 删除多余的空格
	addImportCode = strings.Join(strings.Fields(addImportCode), " ")
	// 以空格划分
	importArr := strings.Split(addImportCode, " ")
	log.Println("-->", jsonconv.ObjectToJsonIndent(importArr))
	switch len(importArr) {
	case 1:
		importAlias = ""
		importPackage = importArr[0]
		packageName = path.Base(importPackage)
		break
	case 2:
		importAlias = importArr[0]
		importPackage = importArr[1]
		packageName = importAlias
		break
	default:
		break
	}

	log.Println(importAlias, importPackage, packageName)
}
