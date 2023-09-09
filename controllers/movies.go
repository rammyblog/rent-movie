package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rammyblog/rent-movie/database"
	"github.com/rammyblog/rent-movie/models"
	"github.com/rammyblog/rent-movie/package/app"
)

func GetAllMovies(c *gin.Context) {
	appG := app.Gin{C: c}

	var movies []models.Movie

	err := database.DB.Preload("Genre").Find(&movies).Error

	if err != nil {
		appG.Response(http.StatusBadRequest, "Error occurred")
		return
	}

	appG.Response(http.StatusOK, movies)

}
