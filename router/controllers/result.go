package controllers

type ApiResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewApiResult 创建一个API返回结果对象
func NewApiResult(code int, message string, data interface{}) *ApiResult {
	return &ApiResult{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewErrorApiResult 创建一个API返回错误结果对象
func NewErrorApiResult(code int, message string) *ApiResult {
	return NewApiResult(code, message, nil)
}
