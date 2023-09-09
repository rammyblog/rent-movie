package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rammyblog/rent-movie/database"
	"github.com/rammyblog/rent-movie/helpers"
	"github.com/rammyblog/rent-movie/models"
	"github.com/rammyblog/rent-movie/package/app"
)

type MovieCreateRequest struct {
	Name        string    `json:"name" binding:"required"`
	Year        int       `json:"year" binding:"required"`
	Rating      float64   `json:"rating" binding:"required"`
	ReleaseDate time.Time `json:"releaseDate" binding:"required"`
	Summary     string    `json:"summary" binding:"required"`
	Available   bool      `json:"available" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	Genre       []int     `json:"genre" binding:"required"`
	// Genre       []models.Genre
}

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

func GetSingleMovie(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := helpers.ConvertStringToInt(&c.Params, "id")

	if err != nil {
		log.Fatal(err)
		appG.Response(http.StatusBadRequest, "Error occurred")
		return
	}

	var movie models.Movie

	if err := database.DB.Preload("Genre").First(&movie, id).Error; err != nil {
		appG.Response(http.StatusBadRequest, "Record not found!")
		return
	}
	appG.Response(http.StatusOK, movie)
}

func CreateMovie(c *gin.Context) {
	appG := app.Gin{C: c}
	var input MovieCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		appG.Response(http.StatusBadRequest, err.Error())
		return
	}

	var genres []models.Genre
	if err := database.DB.Find(&genres, input.Genre).Error; err != nil {
		log.Fatal(err)
		appG.Response(http.StatusBadRequest, "Error occurred while search for a genre")
		return
	}

	movie := models.Movie{
		Name:        input.Name,
		Year:        input.Year,
		Rating:      input.Rating,
		ReleaseDate: input.ReleaseDate,
		Summary:     input.Summary,
		Available:   input.Available,
		Image:       input.Image,
		Genre:       genres,
	}

	err := database.DB.Create(&movie).Error
	if err != nil {
		log.Fatal(err)
		appG.Response(http.StatusBadRequest, "An Error occurred while creating a movie")
		return
	}
	appG.Response(http.StatusCreated, movie)
}