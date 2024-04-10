package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-gin-im/tool"
)

// main
func main() {
	// 初始化配置文件
	tool.LocalConfig()

	// 初始化路由
	router := gin.Default()

	// 启动服务
	router.Run(viper.GetString("app.addr"))
}
