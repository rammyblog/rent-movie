package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rammyblog/rent-movie/middleware/jwt"
	"github.com/rammyblog/rent-movie/package/app"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}

	userId, err := jwt.GetUserIdFromToken(c)
	if err != nil {
		appG.Response(http.StatusBadRequest, err.Error())
		return
	}

	appG.Response(http.StatusCreated, userId)
}
