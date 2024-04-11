package jsonconv

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

/*
*
ConfigDefault(默认API行为)、
ConfigCompatibleWithStandardLibrary(支持标准库的行为，比如encoding/jjson)、
ConfigFast(通过忽略float类型数据的精度保证最高效的性能)
*/
var jjson jsoniter.API

func init() {
	//- `IndentionStep`：设置 JSON 缩进的空格数，默认为 0，表示不缩进。
	//- `MarshalFloatWith6Digits`：设置是否将浮点数序列化为 6 位小数，默认为 false。
	//- `EscapeHTML`：设置是否将 HTML 字符转义，默认为 true。
	//- `SortMapKeys`：设置是否按照键名排序序列化 map，默认为 false。
	//- `UseNumber`：设置是否将数字解码为 `json.Number` 类型，默认为 false。
	//- `DisallowUnknownFields`：设置是否在解码时禁止未知字段，默认为 false。
	//- `TagKey`：设置结构体 tag 的键名，默认为 `json`。
	//- `OnlyTaggedField`：设置是否只序列化带有 tag 的结构体字段，默认为 false。
	//- `ValidateJsonRawMessage`：设置是否验证 `json.RawMessage` 类型的字段，默认为 false。
	//- `ObjectFieldMustBeSimpleString`：设置是否要求对象类型的字段必须是简单字符串，默认为 false。
	//- `CaseSensitive`：设置是否大小写敏感，默认为 true。
	jjson = jsoniter.Config{
		CaseSensitive: false,
	}.Froze()
}

// 默认json
func ObjectMarshal(data any, obj any) (err error) {
	bytes, err := jjson.Marshal(data)
	if err != nil {
		return err
	}

	err = jjson.Unmarshal(bytes, obj)
	if err != nil {
		return err
	}

	return nil
}

// 调用 JsonToObject(jsonStr , &obj)
func JsonToObject(jsonStr string, obj any) error {
	err := jjson.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		//log.Println("error:format", "jjson", jsonStr, "obj", obj)
		return err
	}

	return nil
}

// 默认json
func ObjectToJson(data any) string {
	bytes, err := jjson.Marshal(data)
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 转换行结构json
func ObjectToJsonIndent(data any) string {
	bytes, err := jjson.MarshalIndent(data, "", " ")
	if err != nil {
		return ""
	}
	if string(bytes) == "{}" {
		return fmt.Sprintf("%+v", data)
	}
	return string(bytes)
}

// 转下划线json
func ObjectToJsonSnake(data any) string {
	bytes, err := jjson.Marshal(JsonSnakeCase{Value: data})
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 转下划线json
func ObjectToJsonSnakeIdent(data any) string {
	bytes, err := jjson.MarshalIndent(JsonSnakeCase{Value: data}, "", " ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 转首字母驼峰json
func ObjectToJsonCamel(data any) string {
	bytes, err := jjson.Marshal(JsonCamelCase{Value: data})
	if err != nil {
		return ""
	}

	return string(bytes)
}

func SprintPrivateValue(data any) string {
	str := fmt.Sprintf("%+v", data)
	str = strings.ReplaceAll(str, " ", "\n ")
	str = strings.ReplaceAll(str, "{", "\n{\n ")
	str = strings.ReplaceAll(str, "}", "\n}")
	return str
}
