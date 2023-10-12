package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const TraceIDKey = "X-Trace-ID"

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试从请求头中获取 Trace ID
		traceID := c.GetHeader(TraceIDKey)

		if traceID == "" {
			// 如果请求头中不存在 Trace ID，则生成一个新的 Trace ID
			traceID = GenerateTraceID()
		}
		// 将 Trace ID 存入 context
		ctx := ContextWithFields(c.Request.Context(), zap.String(TraceIDKey, traceID))
		//ctx := context.WithValue(c.Request.Context(), TraceIDKey, traceID)
		c.Request = c.Request.WithContext(ctx)

		// 将 Trace ID 设置到响应头中，以便后续服务使用
		c.Header(TraceIDKey, traceID)

		c.Next()
	}
}

func GenerateTraceID() string {
	return uuid.NewString()
}

const (
	fieldsContextKey = "contextKey"
)

func ContextWithFields(ctx context.Context, fields ...zap.Field) context.Context {
	if val := ctx.Value(fieldsContextKey); val != nil {
		if arr, ok := val.([]zap.Field); ok {
			allFields := make([]zap.Field, 0, len(arr)+len(fields))
			allFields = append(allFields, arr...)
			allFields = append(allFields, fields...)
			return context.WithValue(ctx, fieldsContextKey, allFields)
		}
	}

	return context.WithValue(ctx, fieldsContextKey, fields)
}
