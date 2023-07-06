package request

type ChatRecord struct {
	Type      int    `json:"type"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Content   string `json:"content"`
	UserId    int    `json:"userId"`
	IpAddress string `json:"ipAddress"`
	IpSource  string `json:"ipSource"`
}
