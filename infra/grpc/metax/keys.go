package metax

// RPC metadata keys 由 HTTP 网关写入、RPC 服务读取，与 HTTP 头名分开，
// 避免和 gRPC 保留键冲突。
const (
	HeaderRPCRemoteAgent = "rpc-remote-agent"
	HeaderRPCRemoteIP    = "rpc-remote-ip"
)
