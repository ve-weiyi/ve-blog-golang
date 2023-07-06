package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

// GinLogger 用于替换gin框架的Logger中间件，不传参数，直接这样写
func GinLogger() gin.HandlerFunc {
	logger := global.LOG
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next() // 执行视图函数
		// 视图函数执行完成，统计时间，记录日志
		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.RequestURI),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
