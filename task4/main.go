package main

import (
	"init_order/task4/config"
	"init_order/task4/middlewares"
	"init_order/task4/routers"
	"init_order/task4/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())
	r.Use(gin.Recovery())
	//初始化配置环境
	config.InitEnv()
	//初始化日志
	config.InitLogger()
	//初始化数据库连接
	config.InitDB()
	//初始化jwt
	utils.InitJWT()
	//注册路由
	routers.AuthRouter(r)
	routers.PostRouter(r)
	routers.CommentRouter(r)
	r.Run(":8080")
}
