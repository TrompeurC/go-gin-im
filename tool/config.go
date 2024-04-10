package tool

import (
	"github.com/spf13/viper"
	"log"
)

func LocalConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}

func setConfigDefault() {
	viper.SetDefault("app.addr", ":8080")
}
