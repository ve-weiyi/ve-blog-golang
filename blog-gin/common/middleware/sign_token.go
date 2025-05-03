package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
)

// 未登录token
// 未登录时，token = md5(tm,ts)
func SignToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := c.Request.Header.Get(restx.HeaderToken)
		tm := c.Request.Header.Get(restx.HeaderTerminalId)
		ts := c.Request.Header.Get(restx.HeaderTimestamp)

		//glog.Infof("api is no login required. tk:%v, tm:%v,ts:%v", tk, tm, ts)
		// 请求头缺少参数
		if tk == "" || tm == "" || ts == "" {
			response.ResponseError(c, bizerr.NewBizError(bizerr.CodeUserUnLogin, "用户未登录"))
			c.Abort()
			return
		}
		// 判断 token = md5(tm,ts)
		if tk != crypto.Md5v(tm, ts) {
			response.ResponseError(c, bizerr.NewBizError(bizerr.CodeUserLoginExpired, "无效请求"))
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
		tk := c.Request.Header.Get(restx.HeaderToken)
		uid := c.Request.Header.Get(restx.HeaderUid)

		glog.Infof("api is login required. tk:%v, uid:%v", tk, uid)
		// 请求头缺少参数
		if tk == "" || uid == "" {
			response.ResponseError(c, bizerr.NewBizError(bizerr.CodeUserUnLogin, "用户未登录"))
			c.Abort()
			return
		}
		// 判断 uid = cache.get(token)

		c.Next()
	}
}
