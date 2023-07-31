package apidocs

import (
	"fmt"
	"path"
	"regexp"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate"
)

type Config struct {
	OutRoot   string
	ApiRoot   []string
	ModelRoot []string

	ImportPkgPaths []string
	IgnoredModels  []string
	ReplaceModels  map[string]string

	ApiFuncNameAs  func(api *ApiDeclare) string
	ApiFieldNameAs func(field *ModelField) string
	ApiFieldTypeAs func(field *ModelField) string
}

type AstApiDoc struct {
	Config

	//ApiDocs []*TsApiDoc

	ApiDeclares  []*ApiDeclare
	TypeDeclares []*ModelDeclare
}

func NewAstApiDoc(config Config) *AstApiDoc {
	return &AstApiDoc{
		Config: config,
	}
}

func (s *AstApiDoc) Parse() {
	// 解析api定义
	var apis []*ApiDeclare
	for _, root := range s.ApiRoot {
		doc := ParseApiDocsByRoot(root)
		apis = append(apis, doc...)
	}

	// 解析model定义
	var models []*ModelDeclare
	for _, root := range s.ModelRoot {
		model := ParseApiModelsByRoot(root)
		models = append(models, model...)
	}

	// 根据tag对api分类
	s.ApiDeclares = apis
	s.TypeDeclares = models

	//fmt.Println("ApiDeclares:", jsonconv.ObjectToJsonIndent(apis))
	//fmt.Println("TypeDeclares:", jsonconv.ObjectToJsonIndent(models))
}

// 生成 TypeScript
func (s *AstApiDoc) GenerateTsTypeFile() {
	var tsDeclares []*TsModelDeclare
	for _, model := range s.TypeDeclares {
		item := s.convertTsModelDeclare(model)
		if item != nil {
			tsDeclares = append(tsDeclares, item)
		}
	}
	fmt.Println("tsDeclares:", jsonconv.ObjectToJsonIndent(tsDeclares))
	meta := plate.PlateMeta{
		Key:            "",
		AutoCodePath:   path.Join(s.OutRoot, "types.ts"),
		Replace:        true,
		TemplateString: ModelTypeScript,
		Data:           tsDeclares,
	}

	err := meta.CreateTempFile()
	if err != nil {
		fmt.Println("生成 TypeScript 时发生错误:", err)
	}
}

func (s *AstApiDoc) GenerateTsApiFiles() {
	tsApiDocs := s.GroupTsApiDocs(s.ApiDeclares)
	fmt.Println("tsApiDocs:", jsonconv.ObjectToJsonIndent(tsApiDocs))
	var metas []plate.PlateMeta
	for _, apiDoc := range tsApiDocs {
		meta := plate.PlateMeta{
			Key:            "",
			AutoCodePath:   path.Join(s.OutRoot, fmt.Sprintf("%s.ts", jsonconv.Camel2Case(apiDoc.Tag))),
			Replace:        true,
			TemplateString: ApiTypeScript,
			Data:           apiDoc,
		}
		fmt.Println("apiDocs:", jsonconv.ObjectToJsonIndent(apiDoc))
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

func (s *AstApiDoc) GroupTsApiDocs(docs []*ApiDeclare) []*TsApiDoc {
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
		// 添加导入的model
		for _, param := range params {
			for _, model := range s.TypeDeclares {
				if model.Name == param {
					// 过滤需要忽略的model
					var ignored bool
					for _, ign := range s.IgnoredModels {
						fmt.Println("IgnoredModels:", ign, model.Name)
						if ign == model.Name {
							ignored = true
							break
						}
					}
					if ignored {
						break
					}
					item := s.convertTsModelDeclare(model)
					// 去重，已添加的不再添加
					var has bool
					for _, decl := range apiDoc.ModelDeclares {
						if item.Name == decl.Name {
							has = true
							break
						}
					}
					if !has {

						apiDoc.ModelDeclares = append(apiDoc.ModelDeclares, item)
					}
					break
				}
			}
		}

	}

	return apiDocs
}

func (s *AstApiDoc) convertTsModelDeclare(model *ModelDeclare) *TsModelDeclare {
	if model == nil {
		return nil
	}

	name := getTypeScriptType(model.Name)
	tsFields := make([]*ModelField, 0)

	// 需要替换名称的model
	for k, v := range s.ReplaceModels {
		//fmt.Println("ReplaceModels  k:", k, "v:", name)
		if k == name {
			name = v
			break
		}
	}

	for _, field := range model.Fields {
		tsField := &ModelField{
			Name:    jsonconv.Camel2Case(field.Name),
			Type:    getTypeScriptType(field.Type),
			Comment: field.Comment,
		}

		tsFields = append(tsFields, tsField)
	}

	tsModel := &TsModelDeclare{
		Name:   name,
		Extend: s.convertTsModelDeclare(model.Extend),
		Fields: tsFields,
	}

	return tsModel
}

func (s *AstApiDoc) convertTsApiDeclare(doc *ApiDeclare) *TsApiDeclare {
	re := regexp.MustCompile(`\{(.+?)\}`)
	count := len(doc.Path) + len(doc.Query) + len(doc.Form)
	if doc.Body != nil {
		count++
	}
	var tsDoc = &TsApiDeclare{
		Tag:          doc.Tag,
		FunctionName: s.ApiFuncNameAs(doc),
		Summary:      doc.Summary,
		Url:          re.ReplaceAllString(doc.Url, "${$1}"),
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

func (s *AstApiDoc) convertTsParams(list []*ApiParam) []*ApiParam {
	if list == nil {
		return nil
	}
	var out []*ApiParam
	for _, in := range list {
		out = append(out, s.convertTsParam(in))
	}

	return out
}

func (s *AstApiDoc) convertTsParam(in *ApiParam) *ApiParam {
	if in == nil {
		return nil
	}

	out := &ApiParam{
		Name: in.Name,
		Type: getTypeScriptType(in.Type),
	}

	return out
}

func (s *AstApiDoc) convertRequestStr(doc *ApiDeclare) string {
	params := make([]string, 0)
	types := make([]string, 0)
	if doc.Header != nil {
		for _, param := range doc.Header {
			params = append(params, param.Name)
			types = append(types, getTypeScriptType(param.Type))
		}
	}
	if doc.Path != nil {
		for _, param := range doc.Path {
			params = append(params, param.Name)
			types = append(types, getTypeScriptType(param.Type))
		}
	}
	if doc.Query != nil {
		for _, param := range doc.Query {
			params = append(params, param.Name)
			types = append(types, getTypeScriptType(param.Type))
		}
	}
	if doc.Form != nil {
		for _, param := range doc.Form {
			params = append(params, param.Name)
			types = append(types, getTypeScriptType(param.Type))
		}
	}
	if doc.Body != nil {
		params = append(params, doc.Body.Name)
		types = append(types, getTypeScriptType(doc.Body.Type))
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

// response.Response{data=response.PageResult{list=[]entity.Api}}-->Response<PageResult<Api>>
func (s *AstApiDoc) convertResponseStr(data string) string {
	// 提取参数
	params := extractFieldsByAst(data)

	// 替换ts类型
	for i, param := range params {
		params[i] = getTypeScriptType(param)
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
