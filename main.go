package main

import (
	"momo-im/pkg/websocket"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	engine := gin.Default()
	websocket.InitWebsocket(engine)
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
