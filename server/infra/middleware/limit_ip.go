package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/iputil"
)

func LimitIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		source := c.ClientIP()
		var address string
		switch source {
		case "":
		case "127.0.0.1":
			address = "localhost"
		case "localhost":
			address = "localhost"
		default:
			location, err := iputil.GetIpInfoByBaidu(source)
			global.LOG.Println("OperationRecord GetIpInfoByBaidu:", err)
			address = location.Location
		}
		c.Set("ip_source", source)
		c.Set("ip_address", address)
		c.Next()
	}
}
