package routers

import (
	"taskflow-samrat/controllers"
	"taskflow-samrat/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/task-flow-sam/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	auth.Use(middleware.NoAuthMiddleWare())
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	projects := r.Group("/projects")
	// projects.Use(middleware.AuthMiddleWare())
	{
		projects.POST("/", controllers.CreateProject)
		projects.GET("/", controllers.GetAllProjects)
		projects.GET("/:projectId", controllers.GetProjectById)
		projects.PATCH("/:projectId", controllers.UpdateProjectById)
		projects.DELETE("/:projectId", controllers.DeleteProjectById)
	}

	return r
}
