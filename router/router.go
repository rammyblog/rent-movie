package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rammyblog/rent-movie/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/register", controllers.RegisterUser)

	// apiV1 := r.Group("/api/v1")

	return r
}
