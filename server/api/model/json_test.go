package model

import (
	"log"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

func TestJson(t *testing.T) {

	//log.Println("\n", jsonconv.ObjectToJsonSnakeIdent(&response.UserMenu{}))

	jstr := `{"id":1,"rolePid":0,"roleDomain":"blog","roleName":"管理员","roleComment":"admin","isDisable":1,"isDefault":false,"createdAt":"2021-03-22T14:10:21+08:00","updatedAt":"2023-05-16T20:31:16+08:00"}`

	var data entity.Role

	err := jsonconv.UnmarshalJSONIgnoreCase([]byte(jstr), &data)
	if err != nil {
		return
	}

	log.Println(jsonconv.ObjectToJsonIndent(data))
}

func TestIgnore(t *testing.T) {
	var data entity.Role
	var req interface{}
	jb := []byte(`{"id":1,"rolePid":0,"roleDomain":"blog","roleName":"管理员","roleComment":"admin","isDisable":1,"isDefault":false,"createdAt":"2021-03-22T14:10:21+08:00","updatedAt":"2023-05-16T20:31:16+08:00"}`)

	jsoniter.Unmarshal(jb, &req)
	jsonconv.SetCamelCaseJsonTag(&data)
	err := jsoniter.UnmarshalFromString(jsonconv.ObjectToJsonSnake(req), &data)
	if err != nil {
		log.Println(err)
	}

	log.Println(jsonconv.ObjectToJsonSnake(req))
	log.Println(jsoniter.MarshalToString(data))

	//var data int = 42
	//set(&data)
	//Value(&data)
}
func set1(data interface{}) {
	jsonconv.SetCamelCaseJsonTag(data)
}

func set(data interface{}) {
	Value(data)
}

func Value(v interface{}) {
	value := reflect.ValueOf(v).Elem()
	log.Println("--", reflect.TypeOf(v), value.Kind())
}
