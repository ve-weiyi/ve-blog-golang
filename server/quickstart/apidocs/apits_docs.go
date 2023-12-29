package apidocs

import (
	"fmt"
	"path"
	"regexp"

	"github.com/ve-weiyi/ve-blog-golang/server/quickstart/apidocs/apiparser"
	"github.com/ve-weiyi/ve-blog-golang/server/quickstart/plate"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

type Config struct {
	OutRoot   string
	ApiRoot   []string
	ModelRoot []string

	ApiBase        string
	ImportPkgPaths []string
	IgnoredModels  []string
	ReplaceModels  map[string]string

	ApiFuncNameAs  func(api *apiparser.ApiDeclare) string
	ApiFieldNameAs func(field *apiparser.ModelField) string
	ApiFieldTypeAs func(field *apiparser.ModelField) string
}

type AstApiDoc struct {
	Config

	//ApiDocs []*TsApiDoc

	Parser apiparser.ApiParser

	ApiDeclares  []*apiparser.ApiDeclare
	TypeDeclares []*apiparser.ModelDeclare
}

func NewAstApiDoc(config Config) *AstApiDoc {
	cfg := apiparser.AstParserConfig{
		ApiBase: "/api/v1",
	}
	return &AstApiDoc{
		Parser: apiparser.NewAstParser(cfg),
		Config: config,
	}
}

func (s *AstApiDoc) Parse() (err error) {
	// 解析api定义
	var apis []*apiparser.ApiDeclare
	apis, err = s.Parser.ParseApiDocsByRoots(s.ApiRoot...)
	if err != nil {
		return fmt.Errorf("解析api定义时发生错误:%v", err)
	}
	//for _, root := range s.ApiRoot {
	//	doc := ParseApiDocsByRoot(root)
	//	apis = append(apis, doc...)
	//}

	// 解析model定义
	var models []*apiparser.ModelDeclare
	models, err = s.Parser.ParseModelDocsByRoots(s.ModelRoot...)
	if err != nil {
		return fmt.Errorf("解析model定义时发生错误:%v", err)
	}

	//for _, root := range s.ModelRoot {
	//	model := ParseApiModelsByRoot(root)
	//	models = append(models, model...)
	//}

	// 根据tag对api分类
	s.ApiDeclares = apis
	s.TypeDeclares = models

	//fmt.Println("ApiDeclares:", jsonconv.ObjectToJsonIndent(apis))
	//fmt.Println("TypeDeclares:", jsonconv.ObjectToJsonIndent(models))
	return nil
}

// 生成 TypeScript
func (s *AstApiDoc) GenerateTsTypeFile() {
	var tsDeclares []*TsModelDeclare
	for _, model := range s.TypeDeclares {
		// 过滤需要忽略的model
		ignored := false
		for _, ign := range s.IgnoredModels {
			if ign == model.Name {
				fmt.Println("IgnoredModels:", ign, model.Name)
				ignored = true
				break
			}
		}
		if ignored {
			continue
		}

		item := s.convertTsModelDeclare(model)
		if item != nil {
			tsDeclares = append(tsDeclares, item)
		}
	}
	//fmt.Println("tsDeclares:", jsonconv.ObjectToJsonIndent(tsDeclares))
	meta := plate.PlateMeta{
		Key:            "",
		AutoCodePath:   path.Join(s.OutRoot, "types.ts"),
		Replace:        true,
		TemplateString: ModelTypeScript,
		FunMap:         map[string]any{"joinArray": apiparser.JoinArray},
		Data:           tsDeclares,
	}

	err := meta.CreateTempFile()
	if err != nil {
		fmt.Println("生成 TypeScript 时发生错误:", err)
	}
}

func (s *AstApiDoc) GenerateTsApiFiles() {
	tsApiDocs := s.GroupTsApiDocs(s.ApiDeclares)
	//fmt.Println("tsApiDocs:", jsonconv.ObjectToJsonIndent(tsApiDocs))
	var metas []plate.PlateMeta
	for _, apiDoc := range tsApiDocs {
		meta := plate.PlateMeta{
			Key:            "",
			AutoCodePath:   path.Join(s.OutRoot, fmt.Sprintf("%s.ts", jsonconv.Camel2Case(apiDoc.Tag))),
			Replace:        true,
			TemplateString: ApiTypeScript,
			FunMap:         map[string]any{"joinArray": apiparser.JoinArray},
			Data:           apiDoc,
		}
		//fmt.Println("apiDocs:", jsonconv.ObjectToJsonIndent(apiDoc))
		metas = append(metas, meta)
	}

	for _, meta := range metas {
		err := meta.CreateTempFile()
		if err != nil {
			fmt.Println("生成 TypeScript 时发生错误:", err)
		}
		fmt.Println("TypeScript 文件已生成：", meta.AutoCodePath)
	}
}

func (s *AstApiDoc) GroupTsApiDocs(docs []*apiparser.ApiDeclare) []*TsApiDoc {
	var apiDocs []*TsApiDoc
	// 分组
	for _, doc := range docs {
		var apiDoc *TsApiDoc
		// 查找
		for _, item := range apiDocs {
			if item.Tag == doc.Tag {
				apiDoc = item
				break
			}
		}
		// 不存在则创建
		if apiDoc == nil {
			apiDoc = &TsApiDoc{
				Tag:            doc.Tag,
				ImportPkgPaths: s.ImportPkgPaths,
				ApiDeclares:    make([]*TsApiDeclare, 0),
			}
			apiDocs = append(apiDocs, apiDoc)
		}
		//fmt.Println("apiDoc:", jsonconv.ObjectToJsonIndent(apiDoc))
		// 添加
		apiDoc.ApiDeclares = append(apiDoc.ApiDeclares, s.convertTsApiDeclare(doc))
		// 需要导入的model
		params := getModelDeclareName(doc)
		//fmt.Println("params:", doc.Tag, jsonconv.ObjectToJsonIndent(params))
		var tsModels []*TsModelDeclare
		for _, param := range params {
			//fmt.Println("tsModels:", param, jsonconv.ObjectToJsonIndent(s.findTsModelDeclareByName(param)))
			tsModels = append(tsModels, s.findTsModelDeclareByName(param)...)
		}

		for _, item := range tsModels {

			// 去重，已添加的不再添加
			var has bool
			for _, decl := range apiDoc.ModelDeclares {
				if item.Name == decl.Name {
					has = true
					break
				}
			}

			if !has {
				apiDoc.ImportModelTypes = append(apiDoc.ImportModelTypes, item.Name)
				apiDoc.ModelDeclares = append(apiDoc.ModelDeclares, item)
			}
		}

	}

	return apiDocs
}

func (s *AstApiDoc) findTsModelDeclareByName(name string) []*TsModelDeclare {
	var tsModel []*TsModelDeclare
	var model *apiparser.ModelDeclare

	// 过滤需要忽略的model
	for _, ign := range s.IgnoredModels {
		//fmt.Println("IgnoredModels:", ign, param)
		if ign == name {
			return nil
		}
	}
	// []*chatgpt.ChatMessage->chatgpt.ChatMessage
	name = getGoType(name)
	model = s.findModelDeclare(name)

	item := s.convertTsModelDeclare(model)
	if item != nil {
		tsModel = append(tsModel, item)
	}
	return tsModel
}

func (s *AstApiDoc) findModelDeclare(name string) *apiparser.ModelDeclare {

	for _, model := range s.TypeDeclares {
		if model.Name == name {
			return model
		}
		// package name 都相等的情况
		if model.Pkg != "" {
			if fmt.Sprintf("%v.%v", model.Pkg, name) == model.Name {
				return model
			}
		}
	}

	return nil
}

func (s *AstApiDoc) convertTsModelDeclare(model *apiparser.ModelDeclare) *TsModelDeclare {
	if model == nil {
		return nil
	}

	name := GetTypeScriptType(model.Name)
	tsFields := make([]*TsModelField, 0)
	tsExtends := make([]string, 0)

	// 需要替换名称的model
	for k, v := range s.ReplaceModels {
		//fmt.Println("ReplaceModels  k:", k, "v:", name)
		if k == name {
			name = v
			break
		}
	}
	// 属性
	for _, field := range model.Fields {
		tsField := &TsModelField{
			Name:    jsonconv.Camel2Case(field.Name),
			Type:    GetTypeScriptType(field.Type),
			Comment: field.Comment,
		}

		tsFields = append(tsFields, tsField)
	}
	// 继承
	for _, extend := range model.Extend {
		tsExtends = append(tsExtends, extend.Name)
	}

	tsModel := &TsModelDeclare{
		Name:    name,
		Extends: tsExtends,
		Fields:  tsFields,
	}

	return tsModel
}

func (s *AstApiDoc) convertTsApiDeclare(doc *apiparser.ApiDeclare) *TsApiDeclare {
	re := regexp.MustCompile(`\{(.+?)\}`)
	count := len(doc.Path) + len(doc.Query) + len(doc.Form)
	if doc.Body != nil {
		count++
	}
	var tsDoc = &TsApiDeclare{
		Tag:          doc.Tag,
		FunctionName: s.ApiFuncNameAs(doc),
		Summary:      doc.Summary,
		Base:         s.ApiBase,
		Route:        re.ReplaceAllString(doc.Router, "${$1}"),
		Method:       doc.Method,
		Header:       s.convertTsParams(doc.Header),
		Path:         s.convertTsParams(doc.Path),
		Query:        s.convertTsParams(doc.Query),
		Form:         s.convertTsParams(doc.Form),
		Body:         s.convertTsParam(doc.Body),
		Request:      s.convertRequestStr(doc),
		Response:     s.convertResponseStr(doc.Response),
	}

	return tsDoc
}

func (s *AstApiDoc) convertTsParams(list []*apiparser.ApiParam) []*TsApiParam {
	if list == nil {
		return nil
	}
	var out []*TsApiParam
	for _, in := range list {
		out = append(out, s.convertTsParam(in))
	}

	return out
}

func (s *AstApiDoc) convertTsParam(in *apiparser.ApiParam) *TsApiParam {
	if in == nil {
		return nil
	}

	out := &TsApiParam{
		Name: in.Name,
		Type: GetTypeScriptType(in.Type),
	}

	return out
}

func (s *AstApiDoc) convertRequestStr(api *apiparser.ApiDeclare) string {
	params := make([]string, 0)
	types := make([]string, 0)
	//if api.Header != nil {
	//	for _, param := range api.Header {
	//		params = append(params, param.Name)
	//		types = append(types, getTypeScriptType(param.Type))
	//	}
	//}
	if api.Path != nil {
		for _, param := range api.Path {
			params = append(params, param.Name)
			types = append(types, GetTypeScriptType(param.Type))
		}
	}
	if api.Query != nil {
		for _, param := range api.Query {
			params = append(params, param.Name)
			types = append(types, GetTypeScriptType(param.Type))
		}
	}
	if api.Form != nil {
		for _, param := range api.Form {
			params = append(params, param.Name)
			types = append(types, GetTypeScriptType(param.Type))
		}
	}
	if api.Body != nil {
		params = append(params, api.Body.Name)
		types = append(types, GetTypeScriptType(api.Body.Type))
	}

	var result string
	for i, param := range types {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%s: %s", params[i], param)
	}
	return result
}

// response.Response{data=response.PageResult{list=[]entity.User}}-->Response<PageResult<User>>
func (s *AstApiDoc) convertResponseStr(data *apiparser.ApiParam) string {
	if data == nil {
		return ""
	}
	// 提取参数
	params := ExtractFieldsByAst(data.Type)

	// 替换ts类型
	for i, param := range params {
		params[i] = GetTypeScriptType(param)
	}

	// 替换参数
	for i, param := range params {
		for k, v := range s.ReplaceModels {
			if k == param {
				params[i] = v
				break
			}
		}
	}

	if len(params) > 0 {
		// 数组 [a, b, c] 转换为字符串 a<b<c>>
		var result string
		for i, val := range params {
			if i > 0 {
				result += "<"
			}
			result += val
		}
		for i := 0; i < len(params)-1; i++ {
			result += ">"
		}
		return result
	}

	return "any"
}

func getModelDeclareName(method *apiparser.ApiDeclare) []string {
	params := make([]string, 0)
	if method.Body != nil {
		params = append(params, method.Body.Type)
	}

	if len(method.Form) > 0 {
		for _, param := range method.Form {
			params = append(params, param.Type)
		}
	}

	if len(method.Path) > 0 {
		for _, param := range method.Path {
			params = append(params, param.Type)
		}
	}

	if len(method.Query) > 0 {
		for _, param := range method.Query {
			params = append(params, param.Type)
		}
	}

	if method.Response != nil {
		params = append(params, ExtractFieldsByAst(method.Response.Type)...)
	}
	return params
}

func joinArray(arr []string) string {
	var result string
	for i, v := range arr {
		result += v
		if i < len(arr)-1 {
			result += ", "
		}
	}
	return result
}
