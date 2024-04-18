package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func InitWebsocket(engine *gin.Engine) {
	engine.GET("/ws", wsPage)
	engine.Run(":8080")
}

func wsPage(ctx *gin.Context) {
	upgrader := &websocket.Upgrader{}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	fmt.Println("connected: ", conn.RemoteAddr().String())
}
