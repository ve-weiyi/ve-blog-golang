package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
)

// 未登录的校验
func SignToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		uid := c.Request.Header.Get("uid")

		glog.Println(fmt.Sprintf("api is no login required. token:%v ,uid:%v", token, uid))

		c.Set("token", token)
		c.Set("uid", uid)
		c.Next()
	}
}
