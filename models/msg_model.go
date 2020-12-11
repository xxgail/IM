package models

import "IM/common"

const (
	MessageTypeText = "text"
)

// 定义消息体
type Message struct {
	Type    string `json:"type"` // text
	Message string `json:"message"`
	From    string `json:"from"`
}

// 文本消息初始化
func NewTextMsg(from, message string) *Message {
	return &Message{
		Type:    MessageTypeText,
		Message: message,
		From:    from,
	}
}

func getTextMsgRes(cmd, uid, seq, message string) string {
	textMsg := NewTextMsg(uid, message)
	head := NewResponseHead(seq, cmd, common.SUCCESS, "", textMsg)
	return head.ToString()
}

func GetMsgData(cmd, uid, seq, message string) string {
	return getTextMsgRes(cmd, uid, seq, message)
}
