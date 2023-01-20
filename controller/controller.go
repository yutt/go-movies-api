package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yutt/go-movies-api/model"
)

// @Summary      Get a list of users
// @Description  Retuns a list of users
// @Tags         users
// @Produce      json
// @Success      200  {array} model.User
//
// @Router       /v1/users [get]
func ListUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Users)
}

// @Summary      Get user details
// @Description  Retuns user data for the id provided
// @Tags         users
// @Produce      json
// @Success      200  {object} model.User
// @Failure	 	 400
// @Failure		 404
// @Router       /v1/users/{id} [get]
// @Param		 id path uint64 true "search user with id"
func UserDetails(c *gin.Context) {
	if userId, err := strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		found := false
		for _, val := range model.Users {
			if val.Id == userId {
				c.IndentedJSON(http.StatusOK, val)
				break
			}
		}
		if !found {
			c.IndentedJSON(http.StatusNotFound, nil)
		}
	}
}
