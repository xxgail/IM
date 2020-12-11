package models

import (
	"fmt"
	"time"
)

const (
	heartbeatTimeout = 3 * 60 // 用户心跳超时时间
)

// 用户在线状态
type UserOnline struct {
	Ip            string `json:"ip"`            // Ip
	Port          string `json:"port"`          // 端口
	AppId         string `json:"appId"`         // appId
	UId           string `json:"uid"`           // 用户Id
	ClientIp      string `json:"clientIp"`      // 客户端Ip
	ClientPort    string `json:"clientPort"`    // 客户端端口
	LoginTime     int64  `json:"loginTime"`     // 用户上次登录时间
	HeartbeatTime int64  `json:"heartbeatTime"` // 用户上次心跳时间
	LogOutTime    int64  `json:"logOutTime"`    // 用户退出登录的时间
	Qua           string `json:"qua"`           // qua
	DeviceToken   string `json:"deviceToken"`   // 设备信息
	IsLogoff      bool   `json:"isLogoff"`      // 是否下线
}

/**********************  数据处理  *********************************/

// 用户登录
func UserLogin(ip, port string, appId string, uid string, addr string, loginTime int64) (userOnline *UserOnline) {
	userOnline = &UserOnline{
		Ip:            ip,
		Port:          port,
		AppId:         appId,
		UId:           uid,
		ClientIp:      addr,
		LoginTime:     loginTime,
		HeartbeatTime: loginTime,
		IsLogoff:      false,
	}

	return
}

// 用户心跳
func (u *UserOnline) Heartbeat(currentTime int64) {

	u.HeartbeatTime = currentTime
	u.IsLogoff = false

	return
}

// 用户退出登录
func (u *UserOnline) LogOut() {
	currentTime := time.Now().Unix()
	u.LogOutTime = currentTime
	u.IsLogoff = true

	return
}

/**********************  数据操作  *********************************/

// 用户是否在线
func (u *UserOnline) IsOnline() bool {
	if u.IsLogoff {
		fmt.Println("用户是否在线 用户已经下线", u.AppId, u.UId)
		return false
	}

	currentTime := time.Now().Unix()
	if u.HeartbeatTime < (currentTime - heartbeatTimeout) {
		fmt.Println("用户是否在线 心跳超时", u.AppId, u.UId, u.HeartbeatTime)
		return false
	}

	return true
}

// 用户是否在本台机器上
func (u *UserOnline) UserIsLocal(localIp, localPort string) bool {
	return u.Ip == localIp && u.Port == localPort
}
