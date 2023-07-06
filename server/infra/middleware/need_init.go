package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/global"
)

// 处理跨域请求,支持options访问
func NeedInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.DB == nil {
			response.OkWithDetailed(c, "前往初始化数据库", gin.H{
				"needInit": true,
			})
			c.Abort()
		} else {
			c.Next()
		}
		// 处理请求
	}
}
