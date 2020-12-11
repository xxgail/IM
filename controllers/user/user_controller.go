package user

import (
	"IM/common"
	"IM/controllers"
	"IM/lib/mysqllib"
	"IM/lib/redislib"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	err         error
	data        map[string]interface{}
	redisClient = redislib.GetClient()
	mysqlClient = mysqllib.GetMysqlConn()
)

type User struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Remark string `json:"remark"`
}

func AddFriend(c *gin.Context) {

}

func JoinChat(c *gin.Context) {
	chatId := c.Param("chatId")

	var uid string
	if uidRes, ok := c.Get("uid"); ok {
		uid = uidRes.(string)
	} else {
		fmt.Println("no uid")
	}

	var appId string
	if appIdRes, ok := c.Get("appId"); ok {
		appId = appIdRes.(string)
	}

	var groupId string
	if groupIdRes, ok := c.Get("groupId"); ok {
		groupId = groupIdRes.(string)
	}

	// 加入群聊
	// 1-入mysql
	// 2-更新redis
	var chatUserExist string
	query := "SELECT id FROM chat_users WHERE chat_id = '" + chatId + "' AND uid = '" + uid + "' AND app_id = '" + appId + "' AND group_id = '" + groupId + "'"
	fmt.Println(query)
	err = mysqlClient.QueryRow(query).Scan(&chatUserExist)
	if err != nil {
		fmt.Println("JoinChat 查询数据库单条用户信息 出错：", err)
	}

	currentTime := time.Now()
	//开启事务
	tx, err := mysqlClient.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}

	if chatUserExist == "" {
		queryBid := "INSERT INTO chat_users (`chat_id`,`uid`,`app_id`,`user_name`, `user_avatar`, `remark`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)"
		stemBid, err := tx.Prepare(queryBid)
		if err != nil {
			fmt.Println("Insert-chat_users----Prepare fail")
		}
		_, err = stemBid.Exec(chatId, uid, appId, uid+"name", uid+"avatar", uid+"name", currentTime, currentTime)
		if err != nil {
			fmt.Println("Insert-chat_users----Exec fail")
		}
	} else {
		controllers.Response(c, common.FAIL, "您已加入该群组", data)
	}

	userInfo := User{
		Uid:    uid,
		Name:   uid + "name",
		Avatar: uid + "avatar",
		Remark: uid + "name",
	}
	userInfoStr, err := json.Marshal(userInfo)

	res, err := redisClient.HSet(common.Ctx, "chatId", "uid", string(userInfoStr)).Result()
	if err != nil {
		fmt.Println("redis 加群成员信息出错", err)
	}
	fmt.Println("加群", res)

	controllers.Response(c, common.SUCCESS, "", data)
}

func ExitChat(c *gin.Context) {
	// 退出群聊

}
