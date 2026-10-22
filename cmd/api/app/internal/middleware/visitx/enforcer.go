package visitx

import "github.com/ve-weiyi/blog-cloud/infra/http/routex"

type Enforcer interface {
	IsVisitPage(path string, method string) (pageName string, ok bool)
}

type visitRule struct {
	method   string
	pattern  string
	pageName string
}

type defaultEnforcer struct {
	rules []visitRule
}

// NewVisitEnforcer 注册需要记录访问日志的页面路由。
func NewVisitEnforcer() Enforcer {
	return &defaultEnforcer{
		rules: []visitRule{
			{method: "GET", pattern: "/api/v1/home", pageName: "首页"},
			{method: "GET", pattern: "/api/v1/about-me", pageName: "关于我"},
			{method: "GET", pattern: "/api/v1/albums/:id", pageName: "相册详情"},
			{method: "GET", pattern: "/api/v1/articles/:id", pageName: "文章详情"},
			{method: "GET", pattern: "/api/v1/talks/:id", pageName: "说说详情"},
			{method: "GET", pattern: "/api/v1/messages", pageName: "留言"},
			{method: "GET", pattern: "/api/v1/categories", pageName: "分类"},
			{method: "GET", pattern: "/api/v1/tags", pageName: "标签"},
		},
	}
}

func (e *defaultEnforcer) IsVisitPage(path string, method string) (string, bool) {
	path = routex.Normalize(path)
	for _, rule := range e.rules {
		if rule.method == method && routex.Match(rule.pattern, path) {
			return rule.pageName, true
		}
	}
	return "", false
}
