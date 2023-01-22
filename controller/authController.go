package controller

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/yutt/go-movies-api/initializers"
	"github.com/yutt/go-movies-api/model"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

// @Summary      Create a user
// @Description  Create a new user with the data provided
// @Tags         auth
// @Produce      json
// @Success      200  {object} model.User
// @Failure	 	 500
// @Failure		 400
// @Router       /v1/auth/register [post]
// @Param		 data body RegisterBody true "data for new user"
func Register(c *gin.Context) {
	var body RegisterBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

	}
	var user model.User
	user.Username = body.Username

	//TODO: add salt && pepper
	if hashedPw, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		user.Password = string(hashedPw)
	}
	if outcome := initializers.DB.Create(&user); outcome.Error != nil {
		c.AbortWithError(http.StatusBadRequest, outcome.Error)
	}
	c.JSON(http.StatusCreated, gin.H{"data": &model.User{ID: user.ID, Username: user.Username, CreatedAt: user.CreatedAt}})
}

// @Summary      Logs in a user and returns a JWT token
// @Description  Logs in a user and returns a JWT token
// @Tags         auth
// @Produce      json
// @Success      200
// @Failure	 	 500
// @Failure		 400
// @Router       /v1/auth/login [post]
// @Param		 data body RegisterBody true "username and password to login"
func Login(c *gin.Context) {
	var body RegisterBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var user model.User

	if outcome := initializers.DB.First(&user, "username = ?", body.Username); outcome.Error != nil {
		c.AbortWithError(http.StatusBadRequest, outcome.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect User or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect User or password"})
		return
	}
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtExpirationHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Application misconfiguration"})
		log.Fatalf("JWT_EXPIRATION_HOURS is not a valid integer: %v", err)
		return
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * time.Duration(jwtExpirationHours)).Unix(),
			"UserID":   user.ID,
		})

		tokenString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error signing token"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		}
	}

}
