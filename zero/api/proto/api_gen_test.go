package proto

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"

	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/invent"
	"github.com/ve-weiyi/ve-blog-golang/server/utils"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/convertx"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

//go:embed blog.api
var testApi string

const typeTemplate = `
{{- range . -}}
{{- range .Docs -}}
{{ . }}
{{ end -}}
export interface {{ .Name }} {
  {{- range .Members }}
  {{ convertJson .Name }}?: {{ convertTsType .Type.RawName }}; {{ .Comment }}
  {{- end }}
}

{{ end -}}
`

const ModelImportTpl = `{{ if .Models }}import { {{ joinArray .Models }} } from "./types"{{ end }}`

const apiTemplate = `
{{- range .Imports -}}
{{.}}
{{ end -}}
` + ModelImportTpl + `
{{- $base := .Base }}
{{ range .Routes }}
/** {{ .AtDoc.Text }} */
export function {{ convertHandler .Handler }}Api(
{{- if .RequestType -}}data ?: {{convertTsType .RequestType.RawName}}{{- end -}}
): Promise<{{convertTsType .ResponseType.RawName}}> {
  return http.request<{{convertTsType .ResponseType.RawName}}>({
    url: ` + "`{{$base}}{{.Path}}`" + `,
    method: "{{ .Method }}",
    {{- if .RequestType }}
    data: data,
    {{- end }}
  })
}
{{ end -}}
`

func TestParseContent(t *testing.T) {
	sp, err := parser.ParseContent(testApi)
	t.Log(err)

	//t.Log(jsonconv.ObjectToJsonIndent(sp))

	//for _, tp := range sp.Types {
	//	t.Logf("%v", jsonconv.ObjectToJsonIndent(tp))
	//}
	//

	for _, g := range sp.Service.Groups {
		t.Logf("%v", jsonconv.ObjectToJsonIndent(g))

		mmp := make(map[string]spec.Type)
		for _, r := range g.Routes {
			if r.RequestType != nil {
				name := convertx.ConvertGoTypeToTsType(r.RequestType.Name())
				name = strings.ReplaceAll(name, "[]", "")
				mmp[name] = r.RequestType
			}
			if r.ResponseType != nil {
				name := convertx.ConvertGoTypeToTsType(r.ResponseType.Name())
				name = strings.ReplaceAll(name, "[]", "")
				mmp[name] = r.ResponseType
			}
		}
		var models []string
		for k := range mmp {

			models = append(models, k)
		}

		meta := invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    fmt.Sprintf("./ts/%s.ts", g.Annotation.Properties["group"]),
			TemplateString: apiTemplate,
			FunMap: map[string]any{
				"joinArray": utils.JoinArray,
				"convertJson": func(name string) string {
					if name == "ID" {
						return "id"
					}
					return jsonconv.Camel2Case(name)
				},
				"convertTsType": convertx.ConvertGoTypeToTsType,
				"convertHandler": func(name string) string {
					return jsonconv.Case2CamelLowerStart(name)
				},
			},
			Data: map[string]any{
				"Imports": []string{`import http from "@/utils/request"`},
				"Models":  models,
				"Base":    g.Annotation.Properties["prefix"],
				"Routes":  g.Routes,
			},
		}
		err = meta.Execute()
		t.Log(err)
	}

}

func CreateTypesTs(sp *spec.ApiSpec) {
	meta := invent.TemplateMeta{
		Key:            "",
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    fmt.Sprintf("./ts/types.ts"),
		TemplateString: typeTemplate,
		FunMap: map[string]any{
			"joinArray": utils.JoinArray,
			"convertJson": func(name string) string {
				if name == "ID" {
					return "id"
				}
				return jsonconv.Camel2Case(name)
			},
			"convertTsType": convertx.ConvertGoTypeToTsType,
		},
		Data: sp.Types,
	}
	err := meta.Execute()
	fmt.Println(err)
}
