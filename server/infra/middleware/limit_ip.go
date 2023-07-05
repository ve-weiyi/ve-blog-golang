package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/iputil"
)

func LimitIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		var source string
		switch ip {
		case "":
		case "127.0.0.1":
			source = "localhost"
		case "localhost":
			source = "localhost"
		default:
			location, err := iputil.GetIpInfoByBaidu(ip)
			global.LOG.Println("OperationRecord GetIpInfoByBaidu:", err)
			source = location.Location
		}
		c.Set("ip_address", ip)
		c.Set("ip_source", source)
		c.Next()
	}
}
