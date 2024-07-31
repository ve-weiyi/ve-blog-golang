package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

// IP限流
func LimitIP(svcCtx *svctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()

		v, ok := svcCtx.LocalCache.Get(key)
		if !ok {
			svcCtx.LocalCache.Put(key, 1)
		}

		// 短时间内请求10次
		if cast.ToInt(v) > 10 {
			c.JSON(http.StatusOK, apierr.ErrorFrequentRequest)
			c.Abort()
			return
		}

		c.Next()
	}
}
