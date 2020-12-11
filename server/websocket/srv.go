package websocket

import (
	"IM/lib/cache"
	"IM/protobuf"
	grpc "IM/server/grpc/client"
	"fmt"
	"time"
)

func SendMessageToUsers(formId string, appId string, targetIds []string, data string) (pushId []string, err error) {
	for _, uid := range targetIds {
		if uid == formId {
			continue
		}
		client := GetUserClient(appId, uid)
		if client == nil {
			fmt.Println("用户不在线")
			// 走推送
			pushId = append(pushId, uid)
			continue
		}
		client.SendMsg([]byte(data))
	}
	return
}

// 在会话中发消息
func SendMessageToChat(appId string, uid string, targetIds []string, seq string, message string) (pushId []string, err error) {
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		if IsLocal(server) {
			data := protobuf.GetMsgData("msg", uid, seq, message)
			pushId, err = SendMessageToUsers(uid, appId, targetIds, data)
		} else {
			// todo  grpc
			pushId, err = grpc.SendMsgAll(server, appId, uid, targetIds, seq, message)
		}
	}

	return
}
