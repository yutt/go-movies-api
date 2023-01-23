package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/yutt/go-movies-api/initializers"
	"github.com/yutt/go-movies-api/logger"
	"github.com/yutt/go-movies-api/model"
)

// Function that checks if the user is logged in
func RequireLogin(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Warning.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["exp"], claims["ID"])

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var (
			user         model.User
			usersFetched int64
		)

		initializers.DB.Limit(1).Find(&user, "id = ?", claims["ID"]).Count(&usersFetched)
		if usersFetched == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
