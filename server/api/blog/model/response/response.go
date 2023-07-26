package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
}

type PageInfo struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message" `
	Data    interface{} `json:"data"`
}

const (
	ERROR   = 504
	SUCCESS = 200
)

func Result(c *gin.Context, code int, msg string, data interface{}) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(c, SUCCESS, "操作成功", map[string]interface{}{})
}

func OkWithMessage(c *gin.Context, message string) {
	Result(c, SUCCESS, message, map[string]interface{}{})
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "查询成功", data)
}

func OkWithDetailed(c *gin.Context, message string, data interface{}) {
	Result(c, SUCCESS, message, data)
}

func Fail(c *gin.Context) {
	Result(c, ERROR, "操作失败", map[string]interface{}{})
}

func FailWithMessage(c *gin.Context, message string) {
	Result(c, ERROR, message, map[string]interface{}{})
}

func FailWithDetailed(c *gin.Context, message string, data interface{}) {
	Result(c, ERROR, message, data)
}
