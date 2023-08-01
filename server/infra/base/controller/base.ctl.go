package controller

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/api/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

type BaseController struct {
	Log *glog.Glogger
}

func NewBaseController(svc *svc.ControllerContext) BaseController {
	return BaseController{
		Log: global.LOG,
	}
}

// IP限流
func (m *BaseController) LimitLock(ctx *gin.Context) error {
	key := ctx.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Put(key, 1)
	}
	if cast.ToInt(v) > 10 {
		return codes.NewError(codes.CodeForbiddenOperation, fmt.Sprintf("操作频繁,请在10分钟后再试"))
	}
	return nil
}

func (m *BaseController) ResponseOk(ctx *gin.Context, data interface{}) {
	m.Response(ctx, response.SUCCESS, "操作成功", data)
}

func (m *BaseController) ResponseError(ctx *gin.Context, err error) {
	m.Log.Error("操作失败!", err)
	if e, ok := err.(*codes.ApiError); ok {
		ctx.JSON(http.StatusOK, &response.Response{Code: e.Code(), Message: e.Message()})
		return
	}
	m.Response(ctx, response.ERROR, "操作失败", err.Error())
}

func (m *BaseController) Response(ctx *gin.Context, code int, msg string, data interface{}) {
	obj := response.Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	//ctx.JSON(http.StatusOK, obj)

	//全部转下划线json
	ctx.Render(http.StatusOK, camelJSONRender{render.JSON{Data: obj}})
}

func (m *BaseController) Response500(ctx *gin.Context, res interface{}) {
	ctx.JSON(http.StatusInternalServerError, res)
}

func (m *BaseController) GetRequestContext(ctx *gin.Context) (*request.Context, error) {

	reqCtx := &request.Context{}
	reqCtx.Token = ctx.GetString("token")
	reqCtx.UID = ctx.GetInt("uid")
	reqCtx.Username = ctx.GetString("username")
	reqCtx.IpAddress = ctx.GetString("ip_address")
	reqCtx.IpSource = ctx.GetString("ip_source")
	reqCtx.Context = ctx.Request.Context()
	return reqCtx, nil
}

func (m *BaseController) GetContentUnLogin(ctx *gin.Context) (*request.Context, error) {

	return nil, codes.ErrorUserUnLogin
}

type IsValidChecker interface {
	IsValid() error
}

func (m *BaseController) ShouldBindJSON(ctx *gin.Context, req interface{}) error {
	value := reflect.ValueOf(req)
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		//panic("SetCamelCaseJsonTag only accepts a pointer to a struct")
		if err := ctx.ShouldBindJSON(&req); err != nil {
			return codes.NewError(codes.CodeMissingParameter, "参数错误").Wrap(err)
		}
	} else {
		if err := m.BindJSONIgnoreCase(ctx, req); err != nil {
			return codes.NewError(codes.CodeMissingParameter, "参数错误").Wrap(err)
		}
	}

	isValid, ok := req.(IsValidChecker)
	if !ok {
		return nil
	}
	return isValid.IsValid()
}

// 把请求参数转换为小写
func (m *BaseController) BindJSONIgnoreCase(ctx *gin.Context, req interface{}) (err error) {

	var tmp map[string]interface{}
	err = ctx.ShouldBindJSON(&tmp)
	if err != nil {
		return err
	}
	//如果obj已经是指针，则此处不需要指针
	js := jsonconv.ObjectToJsonSnake(tmp)
	//err = jsonconv.UnmarshalJSONIgnoreCase([]byte(js), req)
	err = jsoniter.Unmarshal([]byte(js), req)
	//m.Log.Println(js)
	//m.Log.JsonIndent(req)
	if err != nil {
		m.Log.Error(err)
	}
	return err
}

func (m *BaseController) ShouldBindQuery(ctx *gin.Context, req interface{}) error {
	// ShouldBindQuery使用tag "form"
	if err := ctx.ShouldBind(req); err != nil {
		return codes.NewError(codes.CodeMissingParameter, "参数错误")
	}
	isValid, ok := req.(IsValidChecker)
	if !ok {
		return nil
	}
	return isValid.IsValid()
}

func (m *BaseController) ShouldBind(ctx *gin.Context, req interface{}) error {
	if ctx.Request.Method == "GET" {
		return m.ShouldBindQuery(ctx, req)
	}
	return m.ShouldBindJSON(ctx, req)
}
