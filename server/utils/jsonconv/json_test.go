package jsonconv

import (
	"fmt"
	"log"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func TestCamel2Case(t *testing.T) {
	str := "link__intro"
	log.Println("--->", str)

	cases := Camel2Case(str)
	log.Println("--->", cases)

	camel := Case2Camel(cases)
	log.Println("--->", camel)

	cases = Camel2Case(camel)
	log.Println("--->", cases)
}

func TestJsonToObject(t *testing.T) {

}

// 序列化成字符串
func TestMarshalToString(t *testing.T) {
	order := struct {
		Id       int
		OrderNum string
		Money    float32
		PayTime  time.Time
		Extend   map[string]string
	}{
		Id:       10,
		OrderNum: "100200300",
		Money:    99.99,
		PayTime:  time.Now(),
		Extend:   map[string]string{"name": "张三"},
	}
	// 定义
	var jsonNew = jsoniter.ConfigCompatibleWithStandardLibrary
	// 设置后，没有json标签的属性，会自动转成 xx_xx
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	// 直接转成字符串
	jsonStr, _ := jsonNew.MarshalToString(order)
	fmt.Println("jsonStr:", jsonStr)

	jb, _ := jjson.MarshalIndent(order, "", " ")
	fmt.Println("jsonStr:", string(jb))

	jj := ObjectToJson(order)
	fmt.Println("jsonStr:", jj)
}

func TestUnmarshalJSONIgnoreCase(t *testing.T) {
	data := `{
			"firstname": "John",
			"last_Name": "Doe",
			"age": 30,
			"me":[1,2,3]
		}`
	type Person struct {
		FirstName string `json:"first_Name"`
		LastName  string
		Age       int
		Me        []int
	}

	var p Person
	//
	//var json = jsoniter.Config{
	//	CaseSensitive: false,
	//}.Froze()

	//data := []byte(`{"NAME":"Alice","AGE":30}`)
	//var p Person
	SetCamelCaseJsonTag(&p)
	fmt.Println("---", ObjectToJsonIndent(p))
	//err := json.Unmarshal(data, &p)

	var it interface{}
	it = &p
	if err := UnmarshalJSONIgnoreCase([]byte(data), it); err != nil {
		t.Errorf("UnmarshalJSONIgnoreCase() error = %v", err)
	}

	fmt.Println("---", data)
	fmt.Println("---", ObjectToJsonIndent(p))
}
