package result

type UserResult struct {
	Sub          string `json:"sub"`           // 用户ID
	Name         string `json:"name"`          // 用户名
	Picture      string `json:"picture"`       // 头像URL
	OpenID       string `json:"open_id"`       // 用户在开放平台的唯一标识
	UnionID      string `json:"union_id"`      // 用户在开放平台的统一标识
	EnName       string `json:"en_name"`       // 用户英文名
	TenantKey    string `json:"tenant_key"`    // 租户Key
	AvatarURL    string `json:"avatar_url"`    // 头像URL
	AvatarThumb  string `json:"avatar_thumb"`  // 头像缩略图URL
	AvatarMiddle string `json:"avatar_middle"` // 头像中等尺寸URL
	AvatarBig    string `json:"avatar_big"`    // 头像大尺寸URL
	Email        string `json:"email"`         // 邮箱
	UserID       string `json:"user_id"`       // 用户ID
	EmployeeNo   string `json:"employee_no"`   // 员工工号
	Mobile       string `json:"mobile"`        // 手机号码
}
