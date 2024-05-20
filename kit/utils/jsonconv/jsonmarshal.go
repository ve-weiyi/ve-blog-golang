package jsonconv

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// https://www.cnblogs.com/chenqionghe/p/13067596.html
/*************************************** 下划线json ***************************************/
type JsonSnakeCase struct {
	Value interface{}
}

// 转换为json，key全部变为驼峰
func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := jjson.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

/*************************************** 驼峰json ***************************************/
type JsonCamelCase struct {
	Value interface{}
}

// 转换为json，key全部变为下划线
func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := jjson.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			if key == "id" {
				return match // 不转换"id"字段
			}
			resKey := Lcfirst(Case2Camel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// 设置json tag为下划线，需要传入指针类型。该方法无效
func SetCamelCaseJsonTag(v interface{}) {
	value := reflect.ValueOf(v)
	//log.Println("--", value.Kind(), value.Elem().Kind())
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		panic("SetCamelCaseJsonTag only accepts a pointer to a struct")
	}
	//如果是指针类型，则取指向的结构体
	//for value.Kind() == reflect.Ptr {
	//	value = value.Elem()
	//}
	t := value.Elem().Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag != "" && jsonTag != "-" {
			newJsonTag := Camel2Case(jsonTag)
			field.Tag = reflect.StructTag(strings.Replace(string(field.Tag), jsonTag, newJsonTag, 1))
			log.Println(string(field.Tag))
		} else {
			//无法在运行时添加tag
			fieldName := field.Name
			newJsonTag := Camel2Case(fieldName)
			field.Tag = reflect.StructTag(fmt.Sprintf(`json:"%v"`, newJsonTag))
			log.Println(string(field.Tag))
		}
		log.Println(field)
		log.Println(t.Field(i))
	}
}

// json解析忽略大小写
func UnmarshalJSONIgnoreCase(data []byte, obj interface{}) error {

	var tmp map[string]interface{}
	// 解析JSON数据到匿名结构体中
	if err := jjson.Unmarshal(data, &tmp); err != nil {
		return err
	}

	// 使用反射将匿名结构体中的字段值赋值给obj对象
	// Convert &&&&&p to &p using reflection
	v := reflect.ValueOf(obj)
	t := v.Elem().Type()
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("SetCamelCaseJsonTag only accepts a pointer to a struct")
	}

	checkToken := func(fieldName1, fieldName2 string) bool {
		//替换 _,转换为小写。再比较
		fieldName1 = strings.ReplaceAll(fieldName1, "_", "")
		fieldName1 = strings.ToLower(fieldName1)

		fieldName2 = strings.ReplaceAll(fieldName2, "_", "")
		fieldName2 = strings.ToLower(fieldName2)

		if fieldName1 == fieldName2 {
			return true
		}
		return false
	}

	for i := 0; i < v.Elem().NumField(); i++ {
		//目标字段名
		fieldName := t.Field(i).Name
		fieldType := v.Elem().Field(i)
		for key, value := range tmp {
			if checkToken(fieldName, key) {
				jsonValue := reflect.ValueOf(value)
				//jsonType := reflect.TypeOf(value)
				//log.Println("--", fieldName, value)
				//log.Println("2--", fieldType.Kind(), jsonType.Kind())

				switch fieldType.Kind() {
				case reflect.String:
					fieldType.SetString(jsonValue.String())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					fieldType.SetInt(int64(jsonValue.Float()))
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					fieldType.SetUint(jsonValue.Uint())
				case reflect.Float32, reflect.Float64:
					fieldType.SetFloat(jsonValue.Float())
				case reflect.Bool:
					fieldType.SetBool(jsonValue.Bool())
				case reflect.Slice:
					interfaceSlice := jsonValue.Interface().([]interface{})
					slice := make([]int, len(interfaceSlice))
					for i, v := range interfaceSlice {
						slice[i] = int(v.(float64))
					}
					// 将切片设置为字段的值
					fieldType.Set(reflect.ValueOf(slice))

				default:
					fmt.Printf("Unsupported type: %v\n", jsonValue.Kind())
				}

				break
			}
		}
	}
	return nil
}

func setFieldValue(fieldType reflect.Value, jsonValue reflect.Value) {
	log.Println("2--", fieldType.Kind(), jsonValue.Kind())

	switch fieldType.Kind() {
	case reflect.String:
		fieldType.SetString(jsonValue.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldType.SetInt(int64(jsonValue.Float()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fieldType.SetUint(jsonValue.Uint())
	case reflect.Float32, reflect.Float64:
		fieldType.SetFloat(jsonValue.Float())
	case reflect.Bool:
		fieldType.SetBool(jsonValue.Bool())
	case reflect.Slice:
		// 创建一个新的切片
		slice := reflect.MakeSlice(fieldType.Type(), 0, 0)

		// 迭代JSON数组中的每个元素
		for i := 0; i < jsonValue.Len(); i++ {
			// 获取JSON数组中的每个元素
			elementValue := jsonValue.Index(i)

			// 创建一个新的切片元素并设置其值
			newElement := reflect.New(fieldType.Type().Elem()).Elem()
			setFieldValue(newElement, elementValue)

			// 将新的切片元素添加到切片中
			slice = reflect.Append(slice, newElement)
		}

		// 将切片设置为字段的值
		fieldType.Set(slice)

	default:

		fmt.Printf("Unsupported type: %v\n", jsonValue.Kind())
		fmt.Printf("Unsupported type: %v\n", jsonValue)
		fieldType.Set(jsonValue.Convert(fieldType.Type()))

	}
}

/*************************************** 其他方法 ***************************************/
