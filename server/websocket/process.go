package websocket

import (
	"IM/common"
	"IM/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sync"
)

type DisposeFunc func(client *Client, seq string, message []byte) (code uint32, msg string, data []byte)

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
)

// 注册
func Register(key string, value DisposeFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value

	return
}

func getHandlers(key string) (value DisposeFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()

	value, ok = handlers[key]

	return
}

func ProcessData(client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据 stop", r)
		}
	}()

	request := &protobuf.SocketRequest{}
	err := proto.Unmarshal(message, request)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}

	fmt.Println("request-------------", request)
	//request := &models.Request{}
	//err := json.Unmarshal(message, request)
	//if err != nil {
	//	fmt.Println("处理数据 json Unmarshal", err)
	//	client.SendMsg([]byte("数据不合法"))
	//
	//	return
	//}

	requestData := request.Data

	//requestData, err := json.Marshal(request.Data)
	//if err != nil {
	//	fmt.Println("处理数据 json Marshal", err)
	//	client.SendMsg([]byte("处理数据失败"))
	//
	//	return
	//}

	seq := request.Seq
	cmd := request.Cmd

	var (
		code uint32
		msg  string
		data []byte
	)

	// request
	fmt.Println("request", cmd, client.Addr)

	// 采用 map 注册的方式
	if value, ok := getHandlers(cmd); ok {
		code, msg, data = value(client, seq, []byte(requestData))
	} else {
		code = common.RoutingNotExist
		fmt.Println("处理数据 路由不存在", client.Addr, "cmd", cmd)
	}

	msg = common.GetErrorMessage(code, msg)

	requestHead := protobuf.NewResponseHead(seq, cmd, code, msg, data)
	headByte, err := proto.Marshal(requestHead)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)

		return
	}

	//responseHead := models.NewResponseHead(seq, cmd, code, msg, data)
	//
	//headByte, err := json.Marshal(responseHead)
	//if err != nil {
	//	fmt.Println("处理数据 json Marshal", err)
	//
	//	return
	//}

	client.SendMsg(headByte)

	fmt.Println("response send", client.Addr, client.AppId, client.Uid, "cmd", cmd, "code", code)

	return
}
