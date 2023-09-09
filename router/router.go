package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rammyblog/rent-movie/controllers"
	"github.com/rammyblog/rent-movie/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/register", controllers.RegisterUser)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.WithJWTAuth())
	{
		apiv1.GET("/user", controllers.GetUser)

		// movies
		apiv1.GET("/movies", controllers.GetAllMovies)
		apiv1.POST("/movies", controllers.CreateMovie)
		apiv1.GET("/movies/:id", controllers.GetSingleMovie)
		apiv1.POST("/movies/rent/:id", controllers.RentMovie)
		apiv1.POST("/movies/return/:id", controllers.ReturnMovie)

	}

	return r
}
