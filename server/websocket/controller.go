package websocket

import (
	"IM/common"
	"IM/lib/cache"
	"IM/models"
	"IM/protobuf"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"time"
)

func Ping(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	fmt.Println("webSocket_request ping接口", client.Addr, seq, message)

	data = "pong"

	return
}

// 用户登录
func Login(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {
	code = common.OK
	currentTime := time.Now().Unix()

	request := &protobuf.SocketLogin{}
	if err := proto.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		fmt.Println("用户登录 解析数据失败", seq, err)

		return
	}

	//request := &models.Login{}
	//if err := json.Unmarshal(message, request); err != nil {
	//	code = common.ParameterIllegal
	//	fmt.Println("用户登录 解析数据失败", seq, err)
	//
	//	return
	//}
	fmt.Println("webSocket_request 用户登录", seq, "ServiceToken", request.ServiceToken)

	if request.Uid == "" {
		code = common.UnauthorizedUserId
		fmt.Println("用户登录 非法的用户", seq, request.Uid)

		return
	}

	if client.IsLogin() {
		fmt.Println("用户登录 用户已经登录", client.AppId, client.Uid, seq)
		code = common.OperationFailure

		return
	}

	client.Login(request.AppId, request.Uid, currentTime)

	// 存储数据
	userOnline := models.UserLogin(serverIp, serverPort, request.AppId, request.Uid, client.Addr, currentTime)
	err := cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		fmt.Println("用户登录 SetUserOnlineInfo", seq, err)

		return
	}

	// 用户登录
	login := &login{
		AppId:  request.AppId,
		Uid:    request.Uid,
		Client: client,
	}
	fmt.Println("login", login)
	clientManager.Login <- login

	fmt.Println("用户登录 成功", seq, client.Addr, request.Uid)

	return
}

func Heartbeat(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {
	code = common.OK
	currentTime := time.Now().Unix()

	request := &protobuf.SocketHeartBeat{}
	if err := proto.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		fmt.Println("心跳接口 解析数据失败", seq, err)

		return
	}
	//request := &models.HeartBeat{}
	//if err := json.Unmarshal(message, request); err != nil {
	//	code = common.ParameterIllegal
	//	fmt.Println("心跳接口 解析数据失败", seq, err)
	//
	//	return
	//}

	fmt.Println("webSocket_request 心跳接口", client.AppId, client.Uid)

	if !client.IsLogin() {
		fmt.Println("心跳接口 用户未登录", client.AppId, client.Uid, seq)
		code = common.NotLoggedIn

		return
	}

	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err != nil {
		if err == redis.Nil {
			code = common.NotLoggedIn
			fmt.Println("心跳接口 用户未登录", seq, client.AppId, client.Uid)

			return
		} else {
			code = common.ServerError
			fmt.Println("心跳接口 GetUserOnlineInfo", seq, client.AppId, client.Uid, err)

			return
		}
	}

	client.HeartBeat(currentTime)
	userOnline.Heartbeat(currentTime)
	err = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		fmt.Println("心跳接口 SetUserOnlineInfo", seq, client.AppId, client.Uid, err)

		return
	}

	return
}
