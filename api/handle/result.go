package handle

type ApiResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewApiResult create an instance of the ApiResult
func NewApiResult(code int, message string, data any) ApiResult {
	r := ApiResult{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return r
}

// NewErrorApiResult create an instance of the ApiResult that contains error info
func NewErrorApiResult(code int, message string) ApiResult {
	return NewApiResult(code, message, nil)
}
