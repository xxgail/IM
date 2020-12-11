package chat

import (
	"IM/common"
	"IM/controllers"
	"IM/lib/cache"
	"IM/lib/mysqllib"
	"IM/lib/redislib"
	"IM/server/websocket"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var (
	err         error
	data        map[string]interface{}
	redisClient = redislib.GetClient()
	mysqlClient = mysqllib.GetMysqlConn()
)

// 查询mysql
func List(c *gin.Context) {
	chatId := c.Param("chatId")
	fmt.Println("----------------chatid", chatId)
	data := make(map[string]interface{})
	data["list"] = []string{"47809", "47810", "47818"}
	controllers.Response(c, common.OK, "", data)
}

type SendParam struct {
	MsgId    string `form:"msgId" json:"msgId"`
	ChatId   string `form:"chatId" json:"chatId"`
	Message  string `form:"message" json:"message"`
	ChatType string `form:"chatType" json:"chatType"` // single-单聊、group-群聊
	Uid      string `form:"uid" json:"uid"`
}

func Send(c *gin.Context) {
	// 定义接口返回data
	contentType := c.Request.Header.Get("Content-Type")
	fmt.Println("Content-Type:", contentType)

	var param SendParam

	switch contentType {
	case "application/json":
		err = c.ShouldBindJSON(&param)
	case "application/x-www-form-urlencoded":
		err = c.ShouldBindWith(&param, binding.Form)
	}
	if err != nil {
		fmt.Println(err)
	}

	uid := param.Uid
	appId := "101"

	chatType := param.ChatType
	msgId := param.MsgId
	chatId := param.ChatId
	message := param.Message

	if cache.SeqDuplicates(msgId) {
		fmt.Println("给用户发送消息 重复提交:", msgId)
		controllers.Response(c, common.OK, "", data)

		return
	}

	var targetIds []string
	if chatType == "group" { // 群聊 chatId 为群聊id
		// 1.判断用户是否已加入群
		// 2.判断是否已禁言（？
		// 发送
		// 查询所有群成员 // todo
		// 群成员 []string
		targetIds = []string{"47809", "47810", "47818"}
	} else { // 单聊 chatId 为目标用户id
		// 1.判断目标用户是否为好友（即可发送
		// 发送
		targetIds = append(targetIds, chatId)
	}

	pushId, err := websocket.SendMessageToChat(appId, uid, targetIds, msgId, message)
	if err != nil {
		fmt.Println("发送失败", err.Error())
		// 推送
	}

	// 不在线走推送
	fmt.Println(pushId)
	controllers.Response(c, common.OK, "", data)
}
