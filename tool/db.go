package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() error {
	// 获取dns
	dns := viper.GetString("db.dsn")
	//打开连接数据库
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	// 设置gorm日志级别
	switch gin.Mode() {
	case gin.DebugMode:
		db.Config.Logger.LogMode(logger.Info)
	case gin.ReleaseMode:
		db.Config.Logger.LogMode(logger.Warn)
	}

	// 设置全局变量
	DB = db

	return nil
}
