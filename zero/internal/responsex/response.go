package responsex

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go.opentelemetry.io/otel/trace"
)

type Body struct {
	Code        int         `json:"code"`
	Message     string      `json:"msg"`
	Data        interface{} `json:"data,omitempty"`
	EncryptData interface{} `json:"encrypt_data,omitempty"`
	TraceId     string      `json:"trace_id"`
}

// Response 统一封装成功响应值.
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	// 1. err不为nil的情况，匹配错误码返回
	if err != nil {
		body := Body{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    resp,
			TraceId: GetTraceId(r),
		}
		httpx.OkJsonCtx(r.Context(), w, body)
		return
	}

	// 2. err为nil的情况，返回成功响应
	body := Body{
		Code:    http.StatusOK,
		Message: "success!",
		Data:    resp,
		TraceId: GetTraceId(r),
	}
	httpx.OkJsonCtx(r.Context(), w, body)
}

// GetTraceId 获取TraceId.
func GetTraceId(r *http.Request) string {
	var traceId string
	spanCtx := trace.SpanContextFromContext(r.Context())
	if spanCtx.HasTraceID() {
		traceId = spanCtx.TraceID().String()
	}

	return traceId
}

// GetLanguage 获取app设置的Language，根据Language返回多语言.
func GetLanguage(r *http.Request) string {
	if len(r.Header["Language"]) > 0 {
		language := r.Header["Language"][0]

		return language
	}

	return "en"
}
