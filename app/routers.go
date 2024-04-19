package app

import (
	"momo-im/app/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// users

	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", nil)
		userRouter.GET("/online", nil)
		userRouter.POST("/sendMessage", nil)
		userRouter.POST("/sendMessageAll", controllers.SendMessageAll)
	}
}
