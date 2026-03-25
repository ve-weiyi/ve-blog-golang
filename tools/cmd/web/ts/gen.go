package ts

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ve-weiyi/pkg/kit/quickstart/gotplgen"
	"github.com/ve-weiyi/pkg/utils/jsonconv"
)

type TsApiService struct {
	ServiceName string
	Groups      map[string][]TsApiGroup // [account][a,b,c]
	Types       map[string]TsType
}

type TsApiGroup struct {
	Prefix     string
	Middleware []string
	Routes     []TsApiRoute
}

type TsApiRoute struct {
	Summery  string
	Handler  string
	Path     string
	Method   string
	Request  string
	Response string

	PathFields  []string
	QueryFields []string
	FormFields  []string
	BodyFields  []string
}

type TsType struct {
	Comment string
	Name    string
	Extends []string
	Fields  []TsTypeField
}

type TsTypeField struct {
	Name     string
	Type     string
	Comment  string
	Nullable bool
}

func generateApiTs(sv TsApiService, tplPath string, outPath string, nameAs string) error {
	fmt.Println(jsonconv.AnyToJsonIndent(sv.Groups))

	t := path.Join(tplPath, "api.ts.tpl")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	var metas []gotplgen.TemplateMeta
	for k, v := range sv.Groups {

		// 收集当前组使用的类型
		usedTypes := make(map[string]TsType)
		for _, g := range v {
			for _, r := range g.Routes {
				if r.Request != "" && r.Request != "any" {
					name := strings.Trim(r.Request, "[]")
					collectTypeRecursively(name, sv.Types, usedTypes)
				}
				if r.Response != "" && r.Response != "any" {
					name := strings.Trim(r.Response, "[]")
					collectTypeRecursively(name, sv.Types, usedTypes)
				}
			}
		}

		meta := gotplgen.TemplateMeta{
			Mode:           gotplgen.ModeCreateOrReplace,
			CodeOutPath:    path.Join(outPath, fmt.Sprintf("%s.ts", k)),
			TemplateString: string(tpl),
			FunMap: map[string]any{
				"Join": func(s []string) string {
					return strings.Join(s, ", ")
				},
			},
			Data: map[string]any{
				"Name":        jsonconv.Case2Camel(k),
				"TsApiGroups": v,
				"Types":       usedTypes,
			},
		}

		metas = append(metas, meta)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func generateTypesTs(sv TsApiService, tplPath string, outPath string, nameAs string) error {
	t := path.Join(tplPath, "types.ts.tpl")
	o := path.Join(outPath, "types.ts")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	meta := gotplgen.TemplateMeta{
		Mode:           gotplgen.ModeCreateOrReplace,
		CodeOutPath:    o,
		TemplateString: string(tpl),
		FunMap: map[string]any{
			"Join": func(s []string) string {
				return strings.Join(s, ", ")
			},
		},
		Data: sv.Types,
	}

	return meta.Execute()
}

func generatePermsTs(sv TsApiService, tplPath string, outPath string, nameAs string) error {
	t := path.Join(tplPath, "perms.ts.tpl")
	o := path.Join(outPath, "perms.ts")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	meta := gotplgen.TemplateMeta{
		Mode:           gotplgen.ModeCreateOrReplace,
		CodeOutPath:    o,
		TemplateString: string(tpl),
		FunMap:         nil,
		Data: map[string]any{
			"Name":   jsonconv.Case2Camel(sv.ServiceName),
			"Groups": sv.Groups,
		},
	}

	return meta.Execute()
}

// collectTypeRecursively 递归收集类型及其依赖
func collectTypeRecursively(typeName string, allTypes map[string]TsType, usedTypes map[string]TsType) {
	// 如果已经收集过，直接返回
	if _, exists := usedTypes[typeName]; exists {
		return
	}

	// 获取类型定义
	tsType, exists := allTypes[typeName]
	if !exists {
		return
	}

	// 添加到已使用类型中
	usedTypes[typeName] = tsType

	// 递归收集继承的类型
	for _, extend := range tsType.Extends {
		collectTypeRecursively(extend, allTypes, usedTypes)
	}

	// 递归收集字段中引用的类型
	for _, field := range tsType.Fields {
		fieldType := extractTypeName(field.Type)
		if fieldType != "" && isCustomType(fieldType) {
			collectTypeRecursively(fieldType, allTypes, usedTypes)
		}
	}
}

// extractTypeName 从类型字符串中提取类型名称
func extractTypeName(typeStr string) string {
	// 移除数组标记
	typeStr = strings.TrimSuffix(typeStr, "[]")
	// 移除泛型参数
	if idx := strings.Index(typeStr, "<"); idx != -1 {
		typeStr = typeStr[:idx]
	}
	return typeStr
}

// isCustomType 判断是否为自定义类型
func isCustomType(typeName string) bool {
	builtinTypes := map[string]bool{
		"string":  true,
		"number":  true,
		"boolean": true,
		"any":     true,
		"File":    true,
	}
	return !builtinTypes[typeName]
}
