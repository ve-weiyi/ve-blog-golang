package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/headerconst"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
)

// 未登录token
// 未登录时，token = md5(tm,ts)
func SignToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := c.Request.Header.Get(headerconst.HeaderToken)
		tm := c.Request.Header.Get(headerconst.HeaderTerminal)
		ts := c.Request.Header.Get(headerconst.HeaderTimestamp)

		//glog.Infof("api is no login required. tk:%v, tm:%v,ts:%v", tk, tm, ts)
		// 请求头缺少参数
		if tk == "" || tm == "" || ts == "" {
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserNotPermission, "无效请求"))
			c.Abort()
			return
		}
		// 判断 token = md5(tm,ts)
		if tk != crypto.Md5v(tm, ts) {
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserNotPermission, "无效请求"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// 登录token
// 登录时，token = md5(uid,ts)，从redis中获取token对应的用户信息
func LoginToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := c.Request.Header.Get(headerconst.HeaderToken)
		uid := c.Request.Header.Get(headerconst.HeaderUid)

		glog.Infof("api is login required. tk:%v, uid:%v", tk, uid)
		// 请求头缺少参数
		if tk == "" || uid == "" {
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserNotPermission, "无效请求"))
			c.Abort()
			return
		}
		// 判断 uid = cache.get(token)

		c.Next()
	}
}
