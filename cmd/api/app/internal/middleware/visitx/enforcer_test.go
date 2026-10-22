package visitx

import "testing"

func TestIsVisitPage(t *testing.T) {
	e := NewVisitEnforcer()

	tests := []struct {
		name   string
		method string
		path   string
		want   string
		ok     bool
	}{
		{name: "首页", method: "GET", path: "/api/v1/home", want: "首页", ok: true},
		{name: "文章详情", method: "GET", path: "/api/v1/articles/12", want: "文章详情", ok: true},
		{name: "相册详情", method: "GET", path: "/api/v1/albums/3", want: "相册详情", ok: true},
		{name: "说说详情", method: "GET", path: "/api/v1/talks/9", want: "说说详情", ok: true},
		{name: "带查询串", method: "GET", path: "/api/v1/messages?page=2", want: "留言", ok: true},
		{name: "列表不记录", method: "GET", path: "/api/v1/articles", ok: false},
		{name: "方法不匹配", method: "POST", path: "/api/v1/home", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := e.IsVisitPage(tt.path, tt.method)
			if ok != tt.ok || got != tt.want {
				t.Errorf("IsVisitPage(%q, %q) = %q, %v; want %q, %v", tt.path, tt.method, got, ok, tt.want, tt.ok)
			}
		})
	}
}
