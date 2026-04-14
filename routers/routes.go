package routers

import (
	"taskflow-samrat/controllers"
	"taskflow-samrat/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.Use(middleware.NoAuthMiddleWare())
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	return r
}
