package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-gin-im/tool"
	"log"
)

// main
func main() {
	// 初始化配置文件
	tool.LocalConfig()

	// 设置应用的模式
	tool.SetMode()

	// 初始化数据看
	if err := tool.InitDB(); err != nil {
		log.Fatalln(err)
	}

	// 初始化路由
	router := gin.Default()

	// 启动服务
	router.Run(viper.GetString("app.addr"))
}
