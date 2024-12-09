package restx

import (
	"net/http"
	"strings"
)

/**
HTTP请求头部中常用的信息包括但不限于以下字段：

Host: 请求的目标主机和端口号。
User-Agent: 发出请求的客户端信息，如浏览器类型、版本等。
Accept: 客户端能够接收的内容类型。
Accept-Language: 客户端优先接受的语言。
Accept-Encoding: 客户端能够解码的编码方式，如gzip。
Authorization: 客户端提供的认证信息。
Connection: 控制不同网络连接选项。
Content-Type: 请求体的媒体类型。
Content-Length: 请求体的长度。
Cookie: 服务器发送到客户端的cookie。
Referer: 发起请求的原始URL。
X-Forwarded-For: 通过HTTP代理或负载均衡器发送请求的客户端的原始IP地址。
X-Real-IP: 类似于X-Forwarded-For，通常用于获取客户端的真实IP地址。

RemoteAddr和X-Real-IP在HTTP请求中的区别如下：
RemoteAddr: 是http.Request结构体的一个字段，它直接由Go的HTTP服务器提供，包含了发起请求的客户端的IP地址和端口号。这个地址可能是直接连接到服务器的客户端的地址，或者是最后一个代理服务器的地址，如果请求经过了代理。
X-Real-IP: 是一个非标准的HTTP请求头部字段，通常由反向代理服务器（如Nginx或HAProxy）设置，用于记录原始请求的客户端IP地址。当请求通过一个或多个代理时，X-Real-IP用于传递原始客户端的IP地址
在没有代理的情况下，RemoteAddr通常就是客户端的真实IP地址。但是在使用代理的情况下，RemoteAddr可能是代理服务器的地址，而X-Real-IP则是客户端的真实IP地址。因此，如果你的服务部署在使用了反向代理的环境中，通常需要检查X-Real-IP来获取客户端的真实IP地址。
*/

// 自定义rpc请求头部，防止和grpc的头部冲突
const (
	HeaderRPCUserAgent = "rpc-user-agent"
	HeaderRPCClientIP  = "rpc-client-ip"
)

// 通用请求头部
const (
	HeaderRemoteAddr    = "remote-addr"
	HeaderUserAgent     = "user-agent"
	HeaderReferer       = "referer"
	HeaderXForwardedFor = "x-forwarded-for"
	HeaderXRealIP       = "x-real-ip"
)

// 自定义的HTTP请求头部字段
const (
	// 自定义请求头部
	HeaderAppName   = "app-name"
	HeaderTimezone  = "timezone"
	HeaderCountry   = "country"
	HeaderLanguage  = "language"
	HeaderTimestamp = "timestamp"
	// 游客id
	HeaderTerminal = "terminal"
	// 游客签名 token = md5(terminal,timestamp)
	HeaderXAuthToken = "x-auth-token"

	// 用户id
	HeaderUid = "uid"
	// 用户token认证信息，与uid一起使用
	HeaderToken = "token"
	// 用户auth认证信息,与uid一起使用
	HeaderAuthorization = "authorization"

	// 防重放限制 sign=md5(id+ts+secret)
	HeaderXRequestId   = "x-request-id"
	HeaderXRequestTs   = "x-request-ts"
	HeaderXRequestSign = "x-request-sign"
)

var HeaderFields = []string{
	HeaderRemoteAddr,
	HeaderUserAgent,
	HeaderReferer,
	HeaderXForwardedFor,
	HeaderXRealIP,

	HeaderAppName,
	HeaderTimezone,
	HeaderCountry,
	HeaderLanguage,
	HeaderTimestamp,

	HeaderTerminal,
	HeaderXAuthToken,
	HeaderUid,
	HeaderToken,
	HeaderAuthorization,
}

// RestHeader restful请求头部(Representational State Transfer 表述性状态转移)
type RestHeader struct {
	HeaderAppName       string `json:"app-name" header:"app-name,optional"`
	HeaderTimezone      string `json:"timezone" header:"timezone,optional"`
	HeaderCountry       string `json:"country" header:"country,optional"`
	HeaderLanguage      string `json:"language" header:"language,optional"`
	HeaderTimestamp     string `json:"timestamp" header:"timestamp,optional"`
	HeaderTerminal      string `json:"terminal" header:"terminal,optional"`
	HeaderXAuthToken    string `json:"x-auth-token" header:"x-auth-token,optional"`
	HeaderUid           string `json:"uid" header:"uid,optional"`
	HeaderToken         string `json:"token" header:"token,optional"`
	HeaderAuthorization string `json:"authorization" header:"authorization,optional"`
}

func ParseRestHeader(r *http.Request) *RestHeader {
	header := &RestHeader{}
	header.HeaderAppName = r.Header.Get(HeaderAppName)
	header.HeaderTimezone = r.Header.Get(HeaderTimezone)
	header.HeaderCountry = r.Header.Get(HeaderCountry)
	header.HeaderLanguage = r.Header.Get(HeaderLanguage)
	header.HeaderTimestamp = r.Header.Get(HeaderTimestamp)
	header.HeaderTerminal = r.Header.Get(HeaderTerminal)
	header.HeaderXAuthToken = r.Header.Get(HeaderXAuthToken)
	header.HeaderUid = r.Header.Get(HeaderUid)
	header.HeaderToken = r.Header.Get(HeaderToken)
	header.HeaderAuthorization = r.Header.Get(HeaderAuthorization)
	return header
}

func GetClientIP(r *http.Request) string {
	// 从 X-Forwarded-For 头部获取
	xff := r.Header.Get(HeaderXForwardedFor)
	if xff != "" {
		// X-Forwarded-For 可能包含多个逗号分隔的 IP 地址
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0]) // 取第一个 IP
	}

	// 如果没有 X-Forwarded-For，则尝试从 X-Real-IP 获取
	xRealIP := r.Header.Get(HeaderXRealIP)
	if xRealIP != "" {
		return xRealIP
	}

	// 如果都没有，使用 RemoteAddr，但需要去掉端口
	hostPort := strings.Split(r.RemoteAddr, ":")
	if len(hostPort) > 0 {
		return hostPort[0]
	}

	return ""
}
