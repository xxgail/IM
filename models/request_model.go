package models

/************************  请求数据  **************************/
// 通用请求数据格式
type Request struct {
	Seq  string      `json:"seq"`            // 消息的唯一Id
	Cmd  string      `json:"cmd"`            // 请求命令字
	Data interface{} `json:"data,omitempty"` // 数据 json
}

// 登录请求数据
type Login struct {
	ServiceToken string `json:"serviceToken"` // 验证用户是否登录
	AppId        string `json:"appId,omitempty"`
	Uid          string `json:"uid,omitempty"`
}

// 心跳请求数据
type HeartBeat struct {
	Uid string `json:"uid,omitempty"`
}
