package routers

import (
	"init_order/task4/controllers"
	"init_order/task4/middlewares"

	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.Engine) {
	posts := r.Group("/posts")
	//直接访问的接口，读取所有文章和读取单篇文章
	posts.GET("findposts", controllers.FindPosts)
	posts.GET("findpostbyid/:id", controllers.FindPostById)
	posts.Use(middlewares.JWTAuthMiddleware())
	{
		// 需登录后访问
		posts.POST("createpost", controllers.CreatePost)
		posts.PUT("updatepost/:id", controllers.UpdatePost)
		posts.DELETE("deletepost/:id", controllers.DeletePost)
	}
}
