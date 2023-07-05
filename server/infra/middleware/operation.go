package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"sync"
)

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.LOG.Println("--->", c.Request.URL)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
