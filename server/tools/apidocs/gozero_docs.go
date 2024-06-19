package apidocs

import (
	"fmt"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/server/tools/apidocs/apiparser"
)

type GozeroDoc struct {
	Config

	Parser apiparser.ApiParser

	ApiDeclares  []*apiparser.ApiDeclare
	TypeDeclares []*apiparser.ModelDeclare
}

func NewGozeroDoc(config Config) *GozeroDoc {
	cfg := apiparser.AstParserConfig{
		ApiBase: "/api/v1",
	}
	return &GozeroDoc{
		Parser: apiparser.NewAstParser(cfg),
		Config: config,
	}
}

func (s *GozeroDoc) Parse() (err error) {
	// 解析api定义
	var apis []*apiparser.ApiDeclare
	apis, err = s.Parser.ParseApiDocsByRoots(s.ApiRoot...)
	if err != nil {
		return fmt.Errorf("解析api定义时发生错误:%v", err)
	}

	// 解析model定义
	var models []*apiparser.ModelDeclare
	models, err = s.Parser.ParseModelDocsByRoots(s.ModelRoot...)
	if err != nil {
		return fmt.Errorf("解析model定义时发生错误:%v", err)
	}

	// 根据tag对api分类
	s.ApiDeclares = apis

	for _, m := range models {
		// 过滤需要忽略的model
		_, ok := s.IgnoredModels[m.Type]
		if ok {
			continue
		}
		s.TypeDeclares = append(s.TypeDeclares, m)
	}

	//fmt.Println("ApiDeclares:", jsonconv.ObjectToJsonIndent(apis))
	//fmt.Println("TypeDeclares:", jsonconv.ObjectToJsonIndent(models))
	return nil
}

// 查找接口model定义
func (s *GozeroDoc) findModelDeclare(name string) *apiparser.ModelDeclare {
	name = strings.Trim(name, "[]")

	for _, model := range s.TypeDeclares {
		if model.Type == name {
			return model
		}

		// package name 都相等的情况
		//if model.Pkg != "" {
		//	if fmt.Sprintf("%v.%v", model.Pkg, name) == model.Type {
		//		return model
		//	}
		//}
	}

	return nil
}

func (s *GozeroDoc) GenerateGoZeroApiFiles() {
	apiTpl, _ := files.ReadFile("api.tpl")
	// 根据tag进行分组
	var apiGroups = map[string][]*apiparser.ApiDeclare{}
	for _, api := range s.ApiDeclares {
		apiGroups[api.Tag] = append(apiGroups[api.Tag], api)
	}

	tsApiDocs := s.convertZeroApiDocs(apiGroups)
	//fmt.Println("tsApiDocs:", jsonconv.ObjectToJsonIndent(tsApiDocs))
	var metas []invent.TemplateMeta
	for _, apiDoc := range tsApiDocs {
		meta := invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(s.OutRoot, fmt.Sprintf("api/%s.api", jsonconv.Case2Snake(apiDoc.Tag))),
			TemplateString: string(apiTpl),
			FunMap: map[string]any{
				"joinArray": apiparser.JoinArray,
				"add": func(index int, num int) int {
					return index + num
				},
				"messageType": func(name string) string {
					if name == "int" {
						return "int64"
					}

					name = strings.ReplaceAll(name, "[]int", "[]int64")

					return name
				},
			},
			Data: apiDoc,
		}
		//fmt.Println("apiDocs:", jsonconv.ObjectToJsonIndent(apiDoc))
		metas = append(metas, meta)
	}

	for _, meta := range metas {
		err := meta.Execute()
		if err != nil {
			fmt.Println("生成 TypeScript 时发生错误:", err)
		}
		fmt.Println("TypeScript 文件已生成：", meta.CodeOutPath)
	}
}

func (s *GozeroDoc) GenerateGoZeroTypeFiles() {
	typeTpl, _ := files.ReadFile("api_type.tpl")
	tsModelDeclares := make(map[string]*TsModelDeclare)
	for _, model := range s.TypeDeclares {
		// 过滤需要忽略的model
		_, ok := s.IgnoredModels[model.Type]
		if ok {
			continue
		}

		item := s.convertZeroModelDeclare(model)
		if item != nil {
			tsModelDeclares[item.Name] = item
		}
	}

	//fmt.Println("tsModelDeclares:", jsonconv.ObjectToJsonIndent(tsModelDeclares))
	meta := invent.TemplateMeta{
		Key:            "",
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    path.Join(s.OutRoot, "api/types.api"),
		TemplateString: string(typeTpl),
		FunMap: map[string]any{
			"joinArray": apiparser.JoinArray,
			"add": func(index int, num int) int {
				return index + num
			},
			"messageType": func(name string) string {
				if name == "int" {
					return "int64"
				}

				name = strings.ReplaceAll(name, "[]int", "[]int64")

				return name
			},
		},
		Data: tsModelDeclares,
	}

	err := meta.Execute()
	if err != nil {
		fmt.Println("生成 TypeScript 时发生错误:", err)
	}
}

func (s *GozeroDoc) GenerateGoZeroRpcFiles() {
	apiTpl, _ := files.ReadFile("rpc.tpl")
	// 根据tag进行分组
	var apiGroups = map[string][]*apiparser.ApiDeclare{}
	for _, api := range s.ApiDeclares {
		apiGroups[api.Tag] = append(apiGroups[api.Tag], api)
	}

	tsApiDocs := s.convertZeroApiDocs(apiGroups)
	//fmt.Println("tsApiDocs:", jsonconv.ObjectToJsonIndent(tsApiDocs))
	var metas []invent.TemplateMeta
	for _, apiDoc := range tsApiDocs {
		meta := invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(s.OutRoot, fmt.Sprintf("proto/%s.proto", jsonconv.Case2Snake(apiDoc.Tag))),
			TemplateString: string(apiTpl),
			FunMap: map[string]any{
				"joinArray": apiparser.JoinArray,
				"add": func(index int, num int) int {
					return index + num
				},
				"messageType": func(name string) string {
					if name == "interface{}" {
						return "google.protobuf.Any"
					}

					if name == "int" {
						return "int64"
					}

					if name == "int64" {
						return name
					}

					name = strings.ReplaceAll(name, "[]", "repeated ")
					name = strings.ReplaceAll(name, "*", "")
					name = strings.ReplaceAll(name, "int", "int64")

					return name
				},
			},
			Data: apiDoc,
		}
		//fmt.Println("apiDocs:", jsonconv.ObjectToJsonIndent(apiDoc))
		metas = append(metas, meta)
	}

	for _, meta := range metas {
		err := meta.Execute()
		if err != nil {
			fmt.Println("生成 TypeScript 时发生错误:", err)
		}
		fmt.Println("TypeScript 文件已生成：", meta.CodeOutPath)
	}
}

func (s *GozeroDoc) GenerateGoZeroRpcTypeFiles() {
	typeTpl, _ := files.ReadFile("rpc_message.tpl")
	tsModelDeclares := make(map[string]*TsModelDeclare)
	for _, model := range s.TypeDeclares {
		// 过滤需要忽略的model
		_, ok := s.IgnoredModels[model.Type]
		if ok {
			continue
		}

		item := s.convertZeroModelDeclare(model)
		if item != nil {
			tsModelDeclares[item.Name] = item
		}
	}

	//fmt.Println("tsModelDeclares:", jsonconv.ObjectToJsonIndent(tsModelDeclares))
	meta := invent.TemplateMeta{
		Key:            "",
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    path.Join(s.OutRoot, "message.proto"),
		TemplateString: string(typeTpl),
		FunMap: map[string]any{
			"joinArray": apiparser.JoinArray,
			"add": func(index int, num int) int {
				return index + num
			},
			"messageType": func(name string) string {
				if name == "interface{}" {
					return "google.protobuf.Any"
				}

				if name == "int" {
					return "int64"
				}

				if name == "int64" {
					return name
				}

				name = strings.ReplaceAll(name, "[]", "repeated ")
				name = strings.ReplaceAll(name, "*", "")
				name = strings.ReplaceAll(name, "int", "int64")

				return name
			},
		},
		Data: tsModelDeclares,
	}

	err := meta.Execute()
	if err != nil {
		fmt.Println("生成 TypeScript 时发生错误:", err)
	}
}

// 转换为分组为ts api文件
func (s *GozeroDoc) convertZeroApiDocs(groups map[string][]*apiparser.ApiDeclare) []*TsApiDoc {

	var apiDocs []*TsApiDoc

	for tag, group := range groups {
		var params []*apiparser.ApiParam
		var tsApiDeclare []*TsApiDeclare
		for _, api := range group {
			tsApiDeclare = append(tsApiDeclare, s.convertZeroApiDeclare(api))
			params = append(params, collectRequestParams(api)...)
			params = append(params, collectResponseParams(api)...)
		}

		// 移除重复元素
		//params = removeDuplicates(params)
		//fmt.Println("params:", params)

		tsModelsMap := make(map[string]*TsModelDeclare)
		for _, param := range params {

			//fmt.Println("tsModels:", param, jsonconv.ObjectToJsonIndent(s.findTsModelDeclareByName(param)))
			model := s.findModelDeclare(param.Type)
			tsModel := s.convertZeroModelDeclare(model)
			if tsModel == nil {
				continue
			}

			tsModelsMap[tsModel.Name] = tsModel

			// 继承
			if tsModel.Extends != nil {
				for _, v := range tsModel.Extends {
					m := s.findModelDeclare(v)
					tm := s.convertZeroModelDeclare(m)
					if tm == nil {
						continue
					}

					tsModelsMap[tm.Name] = tm
				}
			}

			// 属性
			if tsModel.Fields != nil {
				for _, v := range tsModel.Fields {
					m := s.findModelDeclare(v.Name)
					tm := s.convertZeroModelDeclare(m)
					if tm == nil {
						continue
					}

					tsModelsMap[tm.Name] = tm
				}
			}
		}

		var tsModels []*TsModelDeclare
		var tsModelsName []string

		for _, item := range tsModelsMap {
			tsModels = append(tsModels, item)
			tsModelsName = append(tsModelsName, item.Name)
		}

		sort.Slice(tsModels, func(i, j int) bool {
			return tsModels[i].Name < tsModels[j].Name
		})

		sort.Slice(tsModelsName, func(i, j int) bool {
			return tsModelsName[i] < tsModelsName[j]
		})

		var apiDoc *TsApiDoc
		apiDoc = &TsApiDoc{
			Tag:              jsonconv.Case2Snake(tag),
			ImportPkgPaths:   s.ImportPkgPaths,
			ImportModelTypes: tsModelsName,
			ModelDeclares:    tsModels,
			ApiDeclares:      tsApiDeclare,
		}

		apiDocs = append(apiDocs, apiDoc)
	}
	return apiDocs
}

func (s *GozeroDoc) convertZeroApiDeclare(doc *apiparser.ApiDeclare) *TsApiDeclare {
	var reqStr string
	if doc.Body != nil {
		reqStr = doc.Body.Type[strings.LastIndex(doc.Body.Type, ".")+1:]
	}

	if reqStr == "any" {
		reqStr = ""
	}
	// 提取参数
	var respStr string
	res := collectResponseParams(doc)
	if doc.Response != nil {
		respStr = res[len(res)-1].Type
		respStr = strings.ReplaceAll(respStr, "response.", "")
		respStr = strings.ReplaceAll(respStr, "entity.", "")
	}

	if respStr == "any" {
		respStr = ""
	}

	// {id}->${id}
	re := regexp.MustCompile(`\{(.+?)\}`)
	var tsDoc = &TsApiDeclare{
		Tag: doc.Tag,
		FunctionName: func(api *apiparser.ApiDeclare) string {
			return fmt.Sprintf("%v", jsonconv.Ucfirst(api.FunctionName))
		}(doc),
		Summary:  doc.Summary,
		Base:     s.ApiBase,
		Route:    re.ReplaceAllString(doc.Router, ":$1"),
		Method:   doc.Method,
		Header:   convertTsParams(doc.Header),
		Path:     convertTsParams(doc.Path),
		Query:    convertTsParams(doc.Query),
		Form:     convertTsParams(doc.Form),
		Body:     convertTsParam(doc.Body),
		Request:  reqStr,
		Response: respStr,
	}

	return tsDoc
}

func (s *GozeroDoc) convertZeroModelDeclare(model *apiparser.ModelDeclare) *TsModelDeclare {
	if model == nil {
		return nil
	}

	tsFields := make([]*TsModelField, 0)
	tsExtends := make([]string, 0)

	// 属性
	for _, field := range model.Fields {
		tsField := &TsModelField{
			Name: field.Name,
			Type: func() string {
				if field.Type == "time.Time" {
					return "int64"
				}
				return field.Type
			}(),
			Json:    jsonconv.Case2Snake(field.Name),
			Comment: field.Comment,
		}

		tsFields = append(tsFields, tsField)
	}
	// 继承
	for _, extend := range model.Extend {
		tsExtends = append(tsExtends, s.ApiFieldTypeAs(extend.Type))
	}

	tsModel := &TsModelDeclare{
		Name:    s.ApiFieldTypeAs(model.Type),
		Extends: tsExtends,
		Fields:  tsFields,
	}

	return tsModel
}
