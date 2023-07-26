package apidocs

import (
	"fmt"
	"regexp"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate"
)

type Config struct {
	OutRoot   string
	ApiRoot   []string
	ModelRoot []string

	ImportPkgPaths []string
}

type AstApiDoc struct {
	Config

	//ApiDocs []*ApiDoc

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

	//fmt.Println("函数参数:", jsonconv.ObjectToJsonIndent(apis))
	//fmt.Println("函数参数:", jsonconv.ObjectToJsonIndent(models))
}

// 生成 TypeScript
func (s *AstApiDoc) GenerateTypeTsFile() {
	var tsDeclares []*TsModelDeclare
	for _, model := range s.TypeDeclares {
		tsModel := &TsModelDeclare{
			Name:   model.Name,
			Fields: convertTsModelFields(model.Fields),
		}
		tsDeclares = append(tsDeclares, tsModel)
	}

	meta := plate.PlateMeta{
		Key:            "",
		AutoCodePath:   "./preview/type.ts",
		Replace:        true,
		TemplateString: ModelTypeScript,
		Data:           tsDeclares,
	}

	err := meta.CreateTempFile()
	if err != nil {
		fmt.Println("生成 TypeScript 时发生错误:", err)
	}
}

func (s *AstApiDoc) GenerateApiTsFiles() {
	apiDocs := s.GroupApiDocs(s.ApiDeclares)
	fmt.Println("apiDocs:", jsonconv.ObjectToJsonIndent(apiDocs))
	var metas []plate.PlateMeta
	for _, apiDoc := range apiDocs {
		meta := plate.PlateMeta{
			Key:            "",
			AutoCodePath:   fmt.Sprintf("./preview/%s.ts", jsonconv.Camel2Case(apiDoc.Tag)),
			Replace:        true,
			TemplateString: ApiTypeScript,
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

func (s *AstApiDoc) GroupApiDocs(docs []*ApiDeclare) []*ApiDoc {
	var apiDocs []*ApiDoc
	for _, doc := range docs {
		hasAppend := false
		for _, apiDoc := range apiDocs {
			if apiDoc.Tag == doc.Tag {
				apiDoc.ApiDeclares = append(apiDoc.ApiDeclares, convertTsApiDoc(doc))
				hasAppend = true
			}
		}
		if hasAppend {
			continue
		}
		apiDoc := &ApiDoc{
			Tag:            doc.Tag,
			ImportPkgPaths: s.ImportPkgPaths,
			ApiDeclares:    make([]*ApiDeclare, 0),
		}

		apiDocs = append(apiDocs, apiDoc)
	}

	// 收集api需要的导入的model（去重）
	for _, apiDoc := range apiDocs {
		apiDoc.ImportModel = CollectApiModel(apiDoc.ApiDeclares)
	}
	return apiDocs
}

func convertTsApiDoc(doc *ApiDeclare) *ApiDeclare {
	doc.FunctionName = jsonconv.Lcfirst(doc.FunctionName)
	if doc.Body != nil {
		doc.Body.Type = extractEntity(doc.Body.Type)
	}
	doc.Response = extractEntity(doc.Response)
	re := regexp.MustCompile(`\{(.+?)\}`)
	doc.Url = re.ReplaceAllString(doc.Url, "${$1}")
	return doc
}

// response.Response{data=response.PageResult{list=[]entity.Api}}-->Response<PageResult<Api>>
func extractEntity(data string) string {
	model := extractFieldsAfterDot(data)
	if len(model) > 0 {
		// 数组 [a, b, c] 转换为字符串 a<b<c>>
		var result string
		for i, val := range model {
			if i > 0 {
				result += "<"
			}
			result += val
		}
		for i := 0; i < len(model)-1; i++ {
			result += ">"
		}
		return result
	}

	return "any"
}

// response.Response{data=response.PageResult{list=[]entity.Api}} --> Response、PageResult 和 Api
func extractFieldsAfterDot(input string) []string {
	// 定义正则表达式
	re := regexp.MustCompile(`\.(\w+)`)
	// 查找所有匹配的字符串
	matches := re.FindAllStringSubmatch(input, -1)

	// 提取 . 后面的字段并返回切片
	fields := make([]string, len(matches))
	for i, match := range matches {
		fields[i] = match[1]
	}

	return fields
}
