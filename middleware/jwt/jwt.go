package jwt

import (
	"fmt"
	"net/http"
	"os"
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
	return
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

func withJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerTokenString := c.GetHeader("Authorization")
		tokenString := strings.Split(bearerTokenString, " ")[1]
		fmt.Println(tokenString)
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(c, http.StatusBadGateway, "Permission denied")
		}
		claims := token.Claims.(jwt.MapClaims)

		fmt.Println(claims)
		c.Next()

		// if err != nil {
		// 	permissionDenied(w)
		// 	return
		// }

		// if !token.Valid {
		// 	permissionDenied(w)
		// 	return
		// }

		// claims := token.Claims.(jwt.MapClaims)

		// userID, err := getId(r)
		// if err != nil {
		// 	permissionDenied(w)
		// 	return
		// }
		// account, err := s.GetAccountById(userID)

		// if err != nil {
		// 	permissionDenied(w)
		// 	return
		// }
		// if account.Number != int64(claims["accountNumber"].(float64)) {
		// 	permissionDenied(w)
		// 	return
		// }

		// handlerFunc(w, r)
	}
}
