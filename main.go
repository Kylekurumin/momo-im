package main

import "github.com/spf13/viper"

func main() {

}

func initConfig() {
	viper.SetConfigName("config/app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}
