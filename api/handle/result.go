package handle

import "encoding/json"

type ApiResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JsonApiRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

func NewApiJsonResult(code int, message string, data interface{}) JsonApiRes {
	marshal, _ := json.Marshal(data)
	return JsonApiRes{
		Code:    code,
		Message: message,
		Data:    marshal,
	}
}

// NewApiResult create an instance of the ApiResult
func NewApiResult(code int, message string, data interface{}) *ApiResult {
	return &ApiResult{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewErrorApiResult create an instance of the ApiResult that contains error info
func NewErrorApiResult(code int, message string) *ApiResult {
	return NewApiResult(code, message, nil)
}
