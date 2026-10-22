package queryx

import "testing"

func TestBuildSorts(t *testing.T) {
	tests := []struct {
		name  string
		sorts []string
		want  string
	}{
		{name: "默认排序", want: "id desc"},
		{name: "合法排序", sorts: []string{"created_at desc", "id asc"}, want: "created_at desc,id asc"},
		{name: "带表前缀", sorts: []string{"t.id desc"}, want: "t.id desc"},
		{name: "单个字符串含逗号", sorts: []string{"created_at desc, id asc"}, want: "created_at desc,id asc"},
		{name: "子查询被丢弃", sorts: []string{"id desc,(select 1)"}, want: "id desc"},
		{name: "函数调用被丢弃", sorts: []string{"RAND()"}, want: "id desc"},
		{name: "分号注入被丢弃", sorts: []string{"id desc; drop table t_article"}, want: "id desc"},
		{name: "多级前缀被丢弃", sorts: []string{"a.b.c desc"}, want: "id desc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := NewQueryBuilder(WithSorts(tt.sorts...))
			_, _, sorts, _, _ := qb.Build()
			if sorts != tt.want {
				t.Errorf("sorts = %q, want %q", sorts, tt.want)
			}
		})
	}
}

func TestBuildPagination(t *testing.T) {
	tests := []struct {
		name     string
		page     int
		size     int
		wantPage int
		wantSize int
	}{
		{name: "默认分页", wantPage: 1, wantSize: defaultPageSize},
		{name: "正常值", page: 2, size: 50, wantPage: 2, wantSize: 50},
		{name: "页码为 0", page: 0, size: 20, wantPage: 1, wantSize: 20},
		{name: "页码为负", page: -1, size: 20, wantPage: 1, wantSize: 20},
		{name: "大小为负", page: 1, size: -1, wantPage: 1, wantSize: defaultPageSize},
		{name: "超过上限", page: 1, size: MaxPageSize + 1, wantPage: 1, wantSize: MaxPageSize},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := NewQueryBuilder(WithPage(tt.page), WithSize(tt.size))
			page, size, _, _, _ := qb.Build()
			if page != tt.wantPage || size != tt.wantSize {
				t.Errorf("page, size = %d, %d, want %d, %d", page, size, tt.wantPage, tt.wantSize)
			}
		})
	}
}
