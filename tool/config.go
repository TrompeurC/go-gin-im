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
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("db.type", "mysql")
	viper.SetDefault("db.dsn", "root:123456@tcp(127.0.0.1:23306)/goim_backend?charset=utf8mb4&parseTime=True&loc=Local)")
}
