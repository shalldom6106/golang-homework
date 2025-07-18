package config

import (
	"fmt"
	"init_order/task4/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
	}
	//自动迁移模型
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatalf("数据库迁移失败：%v", err)
	}
	DB = db
}

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic("初始化日志失败: " + err.Error())
	}
}

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("无法加载.env文件: %v", err)
	}
}

// // 初始化项目
// func InitGin() {
// 	r := gin.Default()
// 	r.Use(func(ctx *gin.Context) {
// 		log.WithFields(logrus.Fields{
// 			"method": ctx.Request.Method,
// 			"path":   ctx.Request.URL.Path,
// 		}).Info("Request received")
// 		ctx.Next()
// 	})

// 	r.GET("/", func(ctx *gin.Context) {
// 		log.Info("日志启动了")
// 		ctx.String(200, "hello world")
// 	})
// 	err := r.Run(":8080")
// 	if err != nil {
// 		log.Fatalf("项目启动失败:%v", err)
// 	}
// }
