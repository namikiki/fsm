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

// SuccessResponse 创建一个API返回结果对象
func SuccessResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}

// ErrorResponse 创建一个API返回错误结果对象
func ErrorResponse(c *gin.Context, code int, msg string, err error) {
	c.JSON(http.StatusOK, Response{
		Code:  code,
		Msg:   msg,
		Error: err.Error(),
	})
}
