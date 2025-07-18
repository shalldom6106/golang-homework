package routers

import (
	"init_order/task4/controllers"
	"init_order/task4/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRouter(c *gin.Engine) {
	comments := c.Group("/comments")
	//直接访问的接口,获取文章的评论
	comments.GET("/findcomment/:postID", controllers.FindComment)
	//需登录后访问
	comments.Use(middlewares.JWTAuthMiddleware())
	{
		comments.POST("createcomment", controllers.CreateComment)
	}
}
