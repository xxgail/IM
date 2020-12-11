package cache

import (
	"IM/common"
	"IM/lib/redislib"
	"IM/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

const (
	userOnlinePrefix    = "im:user:online:" // 用户在线状态
	userOnlineCacheTime = 24 * 60 * 60
)

/*********************  查询用户是否在线  ************************/
func getUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)

	return
}

func GetUserOnlineInfo(userKey string) (userOnline *models.UserOnline, err error) {
	redisClient := redislib.GetClient()

	key := getUserOnlineKey(userKey)

	data, err := redisClient.Get(common.Ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("GetUserOnlineInfo", userKey, err)

			return
		}

		fmt.Println("GetUserOnlineInfo", userKey, err)

		return
	}

	userOnline = &models.UserOnline{}
	err = json.Unmarshal(data, userOnline)
	if err != nil {
		fmt.Println("获取用户在线数据 json Unmarshal", userKey, err)

		return
	}

	fmt.Println("获取用户在线数据", userKey, "time", userOnline.LoginTime, userOnline.HeartbeatTime, "Ip", userOnline.Ip, userOnline.IsLogoff)

	return
}

// 设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *models.UserOnline) (err error) {

	redisClient := redislib.GetClient()
	key := getUserOnlineKey(userKey)

	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		fmt.Println("设置用户在线数据 json Marshal", key, err)

		return
	}

	_, err = redisClient.Do(common.Ctx, "setEx", key, userOnlineCacheTime, string(valueByte)).Result()
	if err != nil {
		fmt.Println("设置用户在线数据 ", key, err)

		return
	}

	return
}
