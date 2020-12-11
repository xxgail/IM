package routers

import (
	"IM/controllers/chat"
	"IM/controllers/home"
	"IM/controllers/user"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")

	homeRouter := router.Group("/home")
	{
		homeRouter.GET("/index", home.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("addFriend", user.AddFriend)
		userRouter.POST("addChat", user.AddChat)
	}

	chatRouter := router.Group("/chat")
	{
		chatRouter.GET("list/:chatId", chat.List)
		chatRouter.POST("send", chat.Send)
	}
}
