package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"go.opentelemetry.io/otel/trace"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
)

type Body struct {
	Code    int64       `json:"code"`
	Message string      `json:"message" `
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id"`
}

func Response(c *gin.Context, code int64, msg string, data interface{}) {
	traceID := trace.SpanContextFromContext(c.Request.Context())

	obj := Body{
		Code:    code,
		Message: msg,
		Data:    data,
		TraceId: traceID.TraceID().String(),
	}
	c.JSON(bizerr.CodeSuccess, obj)
}

func ResponseOk(c *gin.Context, data interface{}) {
	Response(c, bizerr.CodeSuccess, "Operation successful", data)
}

func ResponseError(c *gin.Context, err error) {
	//debug.PrintStack() // 打印调用栈

	switch e := err.(type) {
	case *bizerr.BizError:
		Response(c, e.Code, e.Error(), e.Error())
		return

	case *mysql.MySQLError:
		Response(c, bizerr.CodeSqlQueryError, "Database error", SqlErrorI18n(e))

	default:
		Response(c, bizerr.CodeInternalServerError, "Server Error", e.Error())
	}

	Response(c, bizerr.CodeCaptchaVerify, "Server error", err.Error())
}
