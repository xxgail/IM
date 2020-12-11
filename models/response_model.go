package models

import "encoding/json"

type ResponseHead struct {
	Seq      string    `json:"seq"`
	Cmd      string    `json:"cmd"`
	Response *Response `json:"response"`
}

type Response struct {
	Code   uint32      `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}

func NewResponse(code uint32, msg string, result interface{}) *Response {
	return &Response{
		Code:   code,
		Msg:    msg,
		Result: result,
	}
}

func NewResponseHead(seq, cmd string, code uint32, msg string, result interface{}) *ResponseHead {
	response := NewResponse(code, msg, result)

	return &ResponseHead{
		Seq:      seq,
		Cmd:      cmd,
		Response: response,
	}
}

func (r *ResponseHead) ToString() string {
	res, _ := json.Marshal(r)
	return string(res)
}
