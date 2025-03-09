package utils

import (
	"lonely-monitor/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

// Response 统一响应结构
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
	TraceId string `json:"traceId,omitempty"`
}

// LoginResponse 登录响应数据
type LoginResponse struct {
	Token string `json:"token"`
}

// Success 成功响应
func Success(c *app.RequestContext, data any) *Response {
	return &Response{
		Code:    errno.SuccessCode,
		Message: "success",
		Data:    data,
		TraceId: GetTraceId(c),
	}
}

// Error 错误响应
func Error(c *app.RequestContext, code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
		TraceId: GetTraceId(c),
	}
}
