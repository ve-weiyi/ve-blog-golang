package enums

// API相关枚举 (t_api)
const (
	// API追溯操作记录 (traceable)
	APITraceableNo  = 0 // 不需要追溯操作记录
	APITraceableYes = 1 // 需要追溯操作记录

	// API禁用状态 (status)
	APIStatusNormal   = 0 // 正常
	APIStatusDisabled = 1 // 禁用
)

// 角色相关枚举 (t_role)
const (
	// 角色禁用状态 (status)
	RoleStatusNormal   = 0 // 正常
	RoleStatusDisabled = 1 // 禁用

	// 角色默认状态 (is_default)
	RoleIsDefaultNo  = 0 // 非默认角色
	RoleIsDefaultYes = 1 // 默认角色
)

// 菜单相关枚举 (t_menu)
const (
	// 菜单缓存状态 (keep_alive)
	MenuKeepAliveNo  = 0 // 不缓存
	MenuKeepAliveYes = 1 // 缓存

	// 菜单一直显示状态 (always_show)
	MenuAlwaysShowNo  = 0 // 不一直显示
	MenuAlwaysShowYes = 1 // 一直显示

	// 菜单隐藏状态 (is_hidden)
	MenuVisibleNo  = 0 // 隐藏
	MenuVisibleYes = 1 // 显示

	// 菜单禁用状态 (status)
	MenuStatusNormal   = 0 // 正常
	MenuStatusDisabled = 1 // 禁用
)
