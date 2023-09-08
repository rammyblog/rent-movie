package jwt

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/rammyblog/rent-movie/models"
)

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func permissionDenied(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"msg":  http.StatusText(httpCode),
		"data": data,
	})

	c.Abort()
}

func CreateJwt(user *models.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"id":        user.ID,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func GetUserIdFromToken(c *gin.Context) (uint, error) {
	bearerTokenString := c.GetHeader("Authorization")
	tokenString := strings.Split(bearerTokenString, " ")[1]
	token, err := validateJWT(tokenString)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func WithJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerTokenString := c.GetHeader("Authorization")
		if bearerTokenString == "" {
			permissionDenied(c, http.StatusBadRequest, "Permission denied")
			return
		}
		tokenString := strings.Split(bearerTokenString, " ")[1]
		_, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(c, http.StatusBadGateway, "Permission denied")
			return
		}
		c.Next()
	}
}
