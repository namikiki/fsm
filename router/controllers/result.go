package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// SuccessResponse 返回处理成功消息
func SuccessResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// ErrorResponse  返回处理失败消息
func ErrorResponse(c *gin.Context, code int, msg string, err error) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:  code,
		Msg:   msg,
		Error: err.Error(),
	})
}
