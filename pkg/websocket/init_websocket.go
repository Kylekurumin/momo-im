package websocket

import (
	"fmt"
	websocketClient "momo-im/pkg/client"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

func InitWebsocket(engine *gin.Engine) {
	engine.GET("/ws", wsPage)
	_ = engine.Run(viper.GetString("app.webSocketUrl"))
}

func wsPage(ctx *gin.Context) {
	upgrader := &websocket.Upgrader{}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	fmt.Println("connected: ", conn.RemoteAddr().String())

	c := websocketClient.NewClient(conn.RemoteAddr().String(), conn)

	go c.Write()
}
