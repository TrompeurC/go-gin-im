package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// setMode 设置应用的模式

func SetMode() {
	// 设置框架gin的模式
	mode := viper.GetString("app.mode")
	switch mode {
	case gin.DebugMode:
		fallthrough
	case gin.ReleaseMode:
		gin.SetMode(mode)
	}

	// 设置GORM及其他包的日志模式，通常在初始化具体工具完成
}
