package api

import (
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/quickstart/tools/parserx/aspec"
)

type Route struct {
	Doc      string // 用户接口
	Handler  string // UserHandler
	Method   string // POST、GET、PUT、DELETE
	Path     string // /api/v1/user
	Request  string
	Response string
}

type GroupRoute struct {
	Name   string // account
	Routes []Route
}

func convertGroups(sp *aspec.ApiSpec) (out []GroupRoute) {
	var groups []GroupRoute
	for _, v := range sp.Service.Groups {
		var routes []Route
		for _, r := range v.Routes {

			var doc string
			for _, d := range r.Doc {
				doc = doc + d
			}

			for _, d := range r.HandlerDoc {
				doc = doc + d
			}

			if r.AtDoc.Text != "" {
				doc = doc + strings.Trim(strings.Trim(r.AtDoc.Text, "\\"), "\"")
			}

			var req, resp string
			if r.RequestType != nil {
				req = r.RequestType.Name()
			}

			if r.ResponseType != nil {
				resp = r.ResponseType.Name()
			}

			rt := Route{
				Method:   strings.ToUpper(r.Method),
				Path:     r.Path,
				Handler:  jsonconv.Case2Camel(r.Handler),
				Doc:      doc,
				Request:  req,
				Response: resp,
			}

			routes = append(routes, rt)
		}

		var name = v.Annotation.Properties["group"]
		if name == "" {
			name = "base"
		}

		g := GroupRoute{
			Name:   name,
			Routes: routes,
		}
		groups = append(groups, g)
	}
	return groups
}
