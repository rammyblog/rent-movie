package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rammyblog/rent-movie/database"
	"github.com/rammyblog/rent-movie/helpers"
	"github.com/rammyblog/rent-movie/middleware/jwt"
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
}

type MovieRentRequest struct {
	DueDate      time.Time `json:"dueDate" binding:"required"`
	DateBorrowed time.Time `json:"dateBorrowed" binding:"required"`
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

func RentMovie(c *gin.Context) {
	appG := app.Gin{C: c}
	var input MovieRentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		appG.Response(http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	var movie models.Movie
	var rent models.Rent

	movieId, parsingError := helpers.ConvertStringToInt(&c.Params, "id")

	if parsingError != nil {
		log.Fatal(parsingError)
		appG.Response(http.StatusBadRequest, "Error occurred")
		return
	}

	userId, userTokenErr := jwt.GetUserIdFromToken(c)

	if userTokenErr != nil {
		log.Fatal(userTokenErr)
		appG.Response(http.StatusBadRequest, "Error occurred")
		return
	}

	if err := database.DB.First(&movie, "id = ? AND available = ?", movieId, true).Error; err != nil {
		appG.Response(http.StatusBadRequest, "Opps... Movie not found/available!")
		return
	}

	if err := database.DB.First(&user, userId).Error; err != nil {
		appG.Response(http.StatusBadRequest, "User not found!")
		return
	}

	if result := database.DB.First(&rent, "movie_id = ? AND user_id = ?", movieId, userId); result.Error == nil {
		if rent.ID != 0 {
			appG.Response(http.StatusBadRequest, "Opps... You have borrowed this movie earlier!")
			return
		}
	}

	if input.DateBorrowed.After(input.DueDate) {
		appG.Response(http.StatusBadRequest, "Borrowed date is after Due date")
		return
	}

	rentedMovie := models.Rent{
		DueDate:      input.DueDate,
		DateBorrowed: input.DateBorrowed,
		UserID:       int(userId),
		User:         user,
		MovieId:      movieId,
		Movie:        movie,
		Returned:     false,
	}

	err := database.DB.Create(&rentedMovie).Error
	if err != nil {
		log.Fatal(err)
		appG.Response(http.StatusBadRequest, "An Error occurred while creating a movie")
		return
	}
	appG.Response(http.StatusCreated, rentedMovie)

}

func ReturnMovie(c *gin.Context) {
	appG := app.Gin{C: c}

	var rent models.Rent

	movieId, parsingError := helpers.ConvertStringToInt(&c.Params, "id")

	if parsingError != nil {
		log.Fatal(parsingError)
		appG.Response(http.StatusBadRequest, "Error occurred")
		return
	}

	userId, userTokenErr := jwt.GetUserIdFromToken(c)

	if userTokenErr != nil {
		log.Fatal(userTokenErr)
		appG.Response(http.StatusBadRequest, "Error occurred")
		return
	}

	if result := database.DB.First(&rent, "movie_id = ? AND user_id = ?", movieId, userId); result.Error != nil {
		log.Fatal(result.Error)
		appG.Response(http.StatusBadRequest, "An Error occurred while returning a movie")
		return
	}
	if rent.Returned {
		appG.Response(http.StatusOK, "Already returned")
		return

	}

	database.DB.Model(&rent).Updates(models.Rent{Returned: true, ReturnedAt: time.Now()})
	appG.Response(http.StatusOK, rent)

}
