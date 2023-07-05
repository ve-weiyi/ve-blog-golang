package result

type TokenResult struct {
	AccessToken      string `json:"access_token"`       // 访问令牌
	TokenType        string `json:"token_type"`         // 令牌类型
	ExpiresIn        int    `json:"expires_in"`         // 令牌过期时间（秒）
	RefreshToken     string `json:"refresh_token"`      // 刷新令牌
	RefreshExpiresIn int    `json:"refresh_expires_in"` // 刷新令牌过期时间（秒）
	Scope            string `json:"scope"`              // 作用域
}
