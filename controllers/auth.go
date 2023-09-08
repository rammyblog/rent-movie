package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rammyblog/rent-movie/database"
	"github.com/rammyblog/rent-movie/models"
	"github.com/rammyblog/rent-movie/package/app"
	"github.com/rammyblog/rent-movie/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	var input AuthRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		appG.Response(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)
}

func RegisterUser(c *gin.Context) {
	appG := app.Gin{C: c}
	input := new(CreateUserRequest)
	if err := c.ShouldBindJSON(&input); err != nil {
		appG.Response(http.StatusBadRequest, err.Error())
		return
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		appG.Response(http.StatusBadRequest, err.Error())
		return
	}
	user := models.User{Email: input.Email, Password: string(encpw), Name: input.Name}

	result := database.DB.Create(&user)
	if result.Error != nil {
		appG.Response(http.StatusBadRequest, result.Error.Error())
		return
	}
	appG.Response(http.StatusCreated, types.MapUserToUser(user))
}
