package common

type JsonResult struct {
	Code   uint32      `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}

// 定义错误返回
func Response(code uint32, message string, data interface{}) JsonResult {
	message = GetErrorMessage(code, message)
	jsonMap := grantMap(code, message, data)

	return jsonMap
}

func grantMap(code uint32, message string, data interface{}) JsonResult {
	jsonMap := JsonResult{
		Code:   code,
		Msg:    message,
		Result: data,
	}

	return jsonMap
}
