package result

// 临时票据结果
type CodeResult struct {
	Code int `json:"code"`
}

type Credentials struct {
	OpenId  string `json:"openId"`
	Unionid string `json:"unionid"`
}
