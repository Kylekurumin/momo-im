package main

import (
	"momo-im/app"
	"momo-im/pkg/websocket"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	engine := gin.Default()
	go websocket.InitWebsocket(engine)

	app.Init(engine)

	httpPort := viper.GetString("app.httpPort")
	_ = engine.Run(":" + httpPort)
}

func initConfig() {
	viper.SetConfigName("config/app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
