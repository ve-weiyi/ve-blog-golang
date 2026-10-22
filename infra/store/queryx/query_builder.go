package queryx

import (
	"regexp"
	"strings"
)

const (
	// MaxPageSize 是分页大小上限，防止一次查询取走过多数据。
	MaxPageSize = 1000
	// defaultPageSize 与 NewQueryBuilder 的初始值保持一致。
	defaultPageSize = 10
)

// sortItemPattern 只允许「列名 [asc|desc]」，列名可带一个表前缀。
// 排序会原样进入 ORDER BY，因此必须在这里挡住子查询、函数调用与其它表达式。
var sortItemPattern = regexp.MustCompile(`(?i)^[a-z_][a-z0-9_]*(\.[a-z_][a-z0-9_]*)?(\s+(asc|desc))?$`)

// QueryBuilder 查询条件构建器
type QueryBuilder struct {
	page       int
	size       int
	sorts      string
	conditions []string
	params     []any
}

// Option 查询选项函数
type Option func(*QueryBuilder)

// WithPage 设置页码
func WithPage(page int) Option {
	return func(qb *QueryBuilder) {
		qb.page = page
	}
}

// WithSize 设置每页大小
func WithSize(size int) Option {
	return func(qb *QueryBuilder) {
		qb.size = size
	}
}

// WithSorts 设置排序。非法的排序项会被丢弃，避免客户端输入进入 ORDER BY。
func WithSorts(sorts ...string) Option {
	return func(qb *QueryBuilder) {
		if sanitized := sanitizeSorts(sorts); sanitized != "" {
			qb.sorts = sanitized
		}
	}
}

// WithCondition 添加查询条件
func WithCondition(condition string, value ...any) Option {
	return func(qb *QueryBuilder) {
		qb.conditions = append(qb.conditions, condition)
		qb.params = append(qb.params, value...)
	}
}

// NewQueryBuilder 创建查询构建器
func NewQueryBuilder(opts ...Option) *QueryBuilder {
	qb := &QueryBuilder{
		page:       1,
		size:       defaultPageSize,
		sorts:      "id desc",
		conditions: make([]string, 0),
		params:     make([]any, 0),
	}

	for _, opt := range opts {
		opt(qb)
	}

	return qb
}

// Build 构建查询条件和参数
func (qb *QueryBuilder) Build() (page int, size int, sorts string, conditions string, params []any) {
	page = qb.page
	if page < 1 {
		page = 1
	}

	size = qb.size
	if size <= 0 {
		size = defaultPageSize
	}
	if size > MaxPageSize {
		size = MaxPageSize
	}

	return page, size, qb.sorts, strings.Join(qb.conditions, " and "), qb.params
}

// sanitizeSorts 按逗号拆分后逐项校验，只保留「列名 [asc|desc]」。
func sanitizeSorts(sorts []string) string {
	items := make([]string, 0, len(sorts))
	for _, sort := range sorts {
		for _, item := range strings.Split(sort, ",") {
			item = strings.TrimSpace(item)
			if item == "" || !sortItemPattern.MatchString(item) {
				continue
			}
			items = append(items, item)
		}
	}

	return strings.Join(items, ",")
}
