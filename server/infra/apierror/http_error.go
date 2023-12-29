package apierror

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/apierror/codes"
)

// 定义ERROR常量
// 1xx Informational（信息性状态码）
var (
	ErrorContinue           = NewApiError(codes.CodeContinue, "继续")
	ErrorSwitchingProtocols = NewApiError(codes.CodeSwitchingProtocols, "切换协议")
	ErrorProcessing         = NewApiError(codes.CodeProcessing, "处理中")
)

// 2xx Success（成功状态码）
var (
	ErrorOK                   = NewApiError(codes.CodeOK, "请求成功")
	ErrorCreated              = NewApiError(codes.CodeCreated, "已创建")
	ErrorAccepted             = NewApiError(codes.CodeAccepted, "已接受")
	ErrorNonAuthoritativeInfo = NewApiError(codes.CodeNonAuthoritativeInfo, "非权威信息")
	ErrorNoContent            = NewApiError(codes.CodeNoContent, "无内容")
	ErrorResetContent         = NewApiError(codes.CodeResetContent, "重置内容")
	ErrorPartialContent       = NewApiError(codes.CodePartialContent, "部分内容")
	ErrorMultiStatus          = NewApiError(codes.CodeMultiStatus, "多状态")
	ErrorAlreadyReported      = NewApiError(codes.CodeAlreadyReported, "已报告")
	ErrorIMUsed               = NewApiError(codes.CodeIMUsed, "使用了IM")
)

// 3xx Redirection（重定向状态码）
var (
	ErrorMultipleChoices   = NewApiError(codes.CodeMultipleChoices, "多种选择")
	ErrorMovedPermanently  = NewApiError(codes.CodeMovedPermanently, "永久移动")
	ErrorFound             = NewApiError(codes.CodeFound, "已找到")
	ErrorSeeOther          = NewApiError(codes.CodeSeeOther, "查看其他位置")
	ErrorNotModified       = NewApiError(codes.CodeNotModified, "未修改")
	ErrorUseProxy          = NewApiError(codes.CodeUseProxy, "使用代理")
	ErrorTemporaryRedirect = NewApiError(codes.CodeTemporaryRedirect, "临时重定向")
	ErrorPermanentRedirect = NewApiError(codes.CodePermanentRedirect, "永久重定向")
)

// 4xx Client Errors（客户端错误状态码）
var (
	ErrorBadRequest                  = NewApiError(codes.CodeBadRequest, "请求错误")
	ErrorUnauthorized                = NewApiError(codes.CodeUnauthorized, "未授权")
	ErrorPaymentRequired             = NewApiError(codes.CodePaymentRequired, "需要付款")
	ErrorForbidden                   = NewApiError(codes.CodeForbidden, "禁止访问")
	ErrorNotFound                    = NewApiError(codes.CodeNotFound, "未找到")
	ErrorMethodNotAllowed            = NewApiError(codes.CodeMethodNotAllowed, "方法不允许")
	ErrorNotAcceptable               = NewApiError(codes.CodeNotAcceptable, "不可接受")
	ErrorProxyAuthRequired           = NewApiError(codes.CodeProxyAuthRequired, "代理认证要求")
	ErrorRequestTimeout              = NewApiError(codes.CodeRequestTimeout, "请求超时")
	ErrorConflict                    = NewApiError(codes.CodeConflict, "冲突")
	ErrorGone                        = NewApiError(codes.CodeGone, "资源不可用")
	ErrorLengthRequired              = NewApiError(codes.CodeLengthRequired, "需要内容长度")
	ErrorPreconditionFailed          = NewApiError(codes.CodePreconditionFailed, "先决条件失败")
	ErrorPayloadTooLarge             = NewApiError(codes.CodePayloadTooLarge, "请求实体过大")
	ErrorURITooLong                  = NewApiError(codes.CodeURITooLong, "请求的URI过长")
	ErrorUnsupportedMediaType        = NewApiError(codes.CodeUnsupportedMediaType, "不支持的媒体类型")
	ErrorRangeNotSatisfiable         = NewApiError(codes.CodeRangeNotSatisfiable, "范围不符合要求")
	ErrorExpectationFailed           = NewApiError(codes.CodeExpectationFailed, "期望失败")
	ErrorMisdirectedRequest          = NewApiError(codes.CodeMisdirectedRequest, "误导的请求")
	ErrorUnprocessableEntity         = NewApiError(codes.CodeUnprocessableEntity, "无法处理的实体")
	ErrorLocked                      = NewApiError(codes.CodeLocked, "已锁定")
	ErrorFailedDependency            = NewApiError(codes.CodeFailedDependency, "依赖失败")
	ErrorTooEarly                    = NewApiError(codes.CodeTooEarly, "过早")
	ErrorUpgradeRequired             = NewApiError(codes.CodeUpgradeRequired, "需要升级")
	ErrorPreconditionRequired        = NewApiError(codes.CodePreconditionRequired, "先决条件要求")
	ErrorTooManyRequests             = NewApiError(codes.CodeTooManyRequests, "请求过多")
	ErrorRequestHeaderFieldsTooLarge = NewApiError(codes.CodeRequestHeaderFieldsTooLarge, "请求头字段太大")
	ErrorUnavailableForLegalReasons  = NewApiError(codes.CodeUnavailableForLegalReasons, "由于法律原因不可用")
)

// 5xx Server Errors（服务器错误状态码）
var (
	ErrorInternalServerError           = NewApiError(codes.CodeInternalServerError, "服务器内部错误")
	ErrorNotImplemented                = NewApiError(codes.CodeNotImplemented, "未实现")
	ErrorBadGateway                    = NewApiError(codes.CodeBadGateway, "错误的网关")
	ErrorServiceUnavailable            = NewApiError(codes.CodeServiceUnavailable, "服务不可用")
	ErrorGatewayTimeout                = NewApiError(codes.CodeGatewayTimeout, "网关超时")
	ErrorHTTPVersionNotSupported       = NewApiError(codes.CodeHTTPVersionNotSupported, "不支持的HTTP版本")
	ErrorVariantAlsoNegotiates         = NewApiError(codes.CodeVariantAlsoNegotiates, "变体也协商")
	ErrorInsufficientStorage           = NewApiError(codes.CodeInsufficientStorage, "存储不足")
	ErrorLoopDetected                  = NewApiError(codes.CodeLoopDetected, "检测到循环")
	ErrorNotExtended                   = NewApiError(codes.CodeNotExtended, "未扩展")
	ErrorNetworkAuthenticationRequired = NewApiError(codes.CodeNetworkAuthenticationRequired, "需要网络认证")
)
