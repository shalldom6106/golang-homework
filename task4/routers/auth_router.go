package routers

import (
	"init_order/task4/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
