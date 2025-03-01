package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"traceId,omitempty"`
}

// LoginResponse 登录响应数据
type LoginResponse struct {
	Token string `json:"token"`
}

// Success 成功响应
func Success(c *app.RequestContext, data interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "success",
		Data:    data,
		TraceId: getTraceId(c),
	}
}

// Error 错误响应
func Error(c *app.RequestContext, code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
		TraceId: getTraceId(c),
	}
}
