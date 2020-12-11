package main

import (
	"IM/lib/mysqllib"
	"IM/lib/redislib"
	"IM/routers"
	"IM/server/websocket"
	"IM/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	initConfig()
	initRedis()
	initMysql()
	initFile()

	// 初始化路由
	router := gin.Default()

	routers.Init(router)
	routers.WebsocketInit()

	task.Init()
	task.ServerInit()

	go websocket.StartWebSocket()

	//go grpc.Init()

	go open()

	httpPort := viper.GetString("app.httpPort")
	http.ListenAndServe(":"+httpPort, router)
}

func initConfig() {
	viper.SetConfigName("config/app")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func initMysql() {
	mysqllib.InitDB()
}

func initRedis() {
	redislib.InitClient()
}

func initFile() {
	gin.DisableConsoleColor()

	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultErrorWriter = io.MultiWriter(f)
}

func open() {
	time.Sleep(1000 * time.Millisecond)

	httpUrl := viper.GetString("app.httpUrl")
	httpUrl = "http://" + httpUrl + "/home/index"

	fmt.Println("访问页面体验：", httpUrl)

	cmd := exec.Command("open", httpUrl)
	cmd.Output()
}
