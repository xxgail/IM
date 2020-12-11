package protobuf

import "github.com/golang/protobuf/proto"

func NewResponse(code uint32, msg string, result []byte) *SocketResponse {
	return &SocketResponse{
		Code:   code,
		Msg:    msg,
		Result: result,
	}
}

func NewResponseHead(seq, cmd string, code uint32, msg string, result []byte) *SocketResponseHead {
	response := NewResponse(code, msg, result)
	return &SocketResponseHead{
		Seq:      seq,
		Cmd:      cmd,
		Response: response,
	}
}

func (m *SocketResponseHead) ToProto() string {
	res, _ := proto.Marshal(m)
	return string(res)
}
