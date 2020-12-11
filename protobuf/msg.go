package protobuf

import (
	"IM/common"
	"github.com/golang/protobuf/proto"
)

const (
	MessageTypeText = "text"
)

func NewTextMsg(from, message string) *SocketMessage {
	return &SocketMessage{
		Type:    MessageTypeText,
		Message: message,
		From:    from,
	}
}

func getTextMsgRes(cmd, uid, seq, message string) string {
	textMsg := NewTextMsg(uid, message)
	textMsgProto, err := proto.Marshal(textMsg)
	if err != nil {

	}
	head := NewResponseHead(seq, cmd, common.SUCCESS, "", textMsgProto)
	return head.ToProto()
}

func GetMsgData(cmd, uid, seq, message string) string {
	return getTextMsgRes(cmd, uid, seq, message)
}
