package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
)

type BaseController struct {
}

func NewBaseController() BaseController {
	return BaseController{}
}

// 获取请求上下文
func (m *BaseController) GetRequestContext(c *gin.Context) (*request.Context, error) {

	reqCtx := &request.Context{}
	reqCtx.Token = c.GetHeader(constant.HeaderToken)
	reqCtx.Uid = cast.ToInt64(c.GetHeader(constant.HeaderUid))
	reqCtx.IpAddress = c.ClientIP()
	reqCtx.UserAgent = c.Request.UserAgent()
	reqCtx.Context = c.Request.Context()
	return reqCtx, nil
}

// IP限流

type IsValidChecker interface {
	IsValid() error
}

func (m *BaseController) ShouldBindJSON(c *gin.Context, req interface{}) error {
	//value := reflect.ValueOf(req)
	//if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct {
	//	if err := m.BindJSONIgnoreCase(c, req); err != nil {
	//		return apierror.NewApiError(apierror.CodeMissingParameter, "参数错误").WrapMessage(err.Error())
	//	}
	//}

	if err := c.ShouldBindJSON(&req); err != nil {
		return apierr.ErrorInvalidParam.WrapMessage(err.Error())
	}

	isValid, ok := req.(IsValidChecker)
	if !ok {
		return nil
	}

	if err := isValid.IsValid(); err != nil {
		return apierr.ErrorInvalidParam.WrapMessage(err.Error())
	}

	return nil
}

// 把请求参数转换为小写
func (m *BaseController) BindJSONIgnoreCase(c *gin.Context, req interface{}) (err error) {

	var tmp map[string]interface{}
	err = c.ShouldBindJSON(&tmp)
	if err != nil {
		return err
	}
	//如果obj已经是指针，则此处不需要指针
	js := jsonconv.ObjectToJsonSnake(tmp)
	err = json.Unmarshal([]byte(js), req)
	//m.Log.Logger(js)
	//m.Log.JsonIndent(req)
	if err != nil {
		glog.Error(err)
	}
	return err
}

func (m *BaseController) ShouldBindQuery(c *gin.Context, req interface{}) error {
	// ShouldBindQuery使用tag "form"
	if err := c.ShouldBind(req); err != nil {
		return apierr.ErrorInvalidParam.WrapMessage(err.Error())
	}
	isValid, ok := req.(IsValidChecker)
	if !ok {
		return nil
	}
	return isValid.IsValid()
}

func (m *BaseController) ShouldBind(c *gin.Context, req interface{}) error {
	if c.Request.Method == "GET" {
		return m.ShouldBindQuery(c, req)
	}
	return m.ShouldBindJSON(c, req)
}

func (m *BaseController) Response(c *gin.Context, code int64, msg string, data interface{}) {
	obj := response.Body{
		Code:    code,
		Message: msg,
		Data:    data,
		TraceId: c.Request.Context().Value("X-Trace-ID").(string),
	}
	c.JSON(http.StatusOK, obj)

	//全部转下划线json
	//c.Render(http.StatusOK, camelJSONRender{render.JSON{Data: obj}})
}

func (m *BaseController) StreamResponse(c *gin.Context, data string) {
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

func (m *BaseController) ResponseOk(c *gin.Context, data interface{}) {
	m.Response(c, http.StatusOK, "操作成功", data)
}

func (m *BaseController) ResponseError(c *gin.Context, err error) {
	glog.Error("操作失败!", err)
	//debug.PrintStack() // 打印调用栈

	switch e := err.(type) {
	case *apierr.ApiError:
		m.Response(c, e.Code, e.Error(), e.Error())
		return

	case *json.UnmarshalTypeError:
		m.Response(c, apierr.ErrorInternalServerError.Code, "json解析错误", e.Error())
		return

	case *mysql.MySQLError:
		switch e.Number {
		case 1062:
			m.Response(c, apierr.ErrorSqlQueryError.Code, "数据已存在", e.Error())
			return
		default:
			m.Response(c, apierr.ErrorSqlQueryError.Code, "数据库错误", response.SqlErrorI18n(e))
			return
		}
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		m.Response(c, apierr.ErrorSqlQueryError.Code, "数据不存在", err.Error())
		return
	}

	m.Response(c, apierr.ErrorInternalServerError.Code, apierr.ErrorInternalServerError.Error(), err.Error())
}
