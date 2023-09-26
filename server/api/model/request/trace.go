package request

import "time"

// 链路追踪 https://coder55.com/article/159071
type spanContext struct {
	traceId string // TraceID 表示tracer的全局唯一ID
	spanId  string // SpanId 标示单个trace中某一个span的唯一ID，在trace中唯一
}

// SpanContext 保存了分布式追踪的上下文信息，包括 Trace id，Span id 以及其它需要传递到下游的内容。
type SpanContext interface {
	TraceId() string                     // get TraceId
	SpanId() string                      // get SpanId
	Visit(fn func(key, val string) bool) // 自定义操作TraceId，SpanId
}

// 一个 REST 调用或者数据库操作等，都可以作为一个 span 。 span 是分布式追踪的最小跟踪单位，一个 Trace 由多段 Span 组成。
type Span struct {
	ctx           spanContext // 传递的上下文
	serviceName   string      // 服务名
	operationName string      // 操作
	startTime     time.Time   // 开始时间戳
	flag          string      // 标记开启trace是 server 还是 client
	children      int         // 本 span fork出来的 childsnums
}
