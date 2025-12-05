package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
)

type Body struct {
	Code    int64       `json:"code"`
	Message string      `json:"message" `
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id,omitempty"`
}

const (
	ERROR   = 504
	SUCCESS = 200
)

func Response(c *gin.Context, code int64, msg string, data interface{}) {
	traceID := trace.SpanContextFromContext(c.Request.Context())

	obj := Body{
		Code:    code,
		Message: msg,
		Data:    data,
		TraceId: traceID.TraceID().String(),
	}
	c.JSON(http.StatusOK, obj)

	//全部转下划线json
	//c.Render(http.StatusOK, camelJSONRender{render.JSON{Data: obj}})
}

func ResponseOk(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, "Operation successful", data)
}

func ResponseError(c *gin.Context, err error) {
	//debug.PrintStack() // 打印调用栈

	switch e := err.(type) {
	case *bizerr.BizError:
		Response(c, e.Code, e.Error(), e.Error())
		return

	case *json.UnmarshalTypeError:
		Response(c, bizerr.CodeInternalServerError, "JSON parsing error", e.Error())
		return

	case *mysql.MySQLError:
		switch e.Number {
		case 1062:
			Response(c, bizerr.CodeSqlQueryError, "Data already exists", e.Error())
			return
		default:
			Response(c, bizerr.CodeSqlQueryError, "Database error", SqlErrorI18n(e))
			return
		}
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		Response(c, bizerr.CodeSqlQueryError, "Data not found", err.Error())
		return
	}

	Response(c, bizerr.CodeInternalServerError, "Server error", err.Error())
}

func ResponseStream(c *gin.Context, data string) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// 计算要发送的数据的分片数量
	//chunkSize := 1
	intervals := getInternalTime(data)

	go func() {
		for _, char := range data {

			_, err := c.Writer.WriteString(fmt.Sprintf("data: %c\n\n", char))
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Fprintf(c.Writer, "data: %c\n\n", char)
			c.Writer.Flush()
			time.Sleep(intervals)
		}

		// 发送结束标记
		_, err := c.Writer.WriteString("data: \n\n")
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Fprintf(c.Writer, "data: \n\n")
		c.Writer.Flush()
	}()

	// 长连接，等待结束
	<-c.Writer.CloseNotify()
}

func getInternalTime(data string) time.Duration {
	if len(data) < 20 {
		return 200 * time.Millisecond
	}

	if len(data) < 100 {
		return 100 * time.Millisecond
	}

	if len(data) < 500 {
		return 50 * time.Millisecond
	}

	if len(data) < 5000 {
		return 20 * time.Millisecond
	}

	return 10 * time.Millisecond
}
