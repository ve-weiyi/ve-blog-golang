package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/iputil"
)

func LimitIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		source := c.ClientIP()
		address, err := iputil.GetIpInfoByBaidu(source)
		global.LOG.Println("OperationRecord GetIpInfoByBaidu:", err)
		c.Set("ip_source", source)
		c.Set("ip_address", address)
		c.Next()
	}
}
