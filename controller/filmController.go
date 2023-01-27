package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yutt/go-movies-api/initializers"
	"github.com/yutt/go-movies-api/model"
)

// @Summary      Get all films
// @Description  Get all films
// @Tags         films
// @Produce      json
// @Success      200  {object} []model.Film
// @Failure	 	 500
// @Failure		 400
// @Router       /v1/films [get]
// @Security     ApiKeyAuth
// @param Authorization header string true "Authorization"
func GetFilms(c *gin.Context) {
	var films []model.Film
	if outcome := initializers.DB.Model(&model.Film{}).Find(&films); outcome.Error != nil {
		c.AbortWithError(http.StatusNotFound, outcome.Error)
	}
	c.JSON(http.StatusOK, gin.H{"data": films})
}

// @Summary      Get a film
// @Description  Get a film
// @Tags         films
// @Produce      json
// @Success      200  {object} model.Film
// @Failure	 	 500
// @Failure		 400
// @Router       /v1/films/:id [get]
// @Security     ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param		 id path int true "film id"
func GetFilm(c *gin.Context) {
	var film model.Film
	if outcome := initializers.DB.Model(&model.Film{}).Find(&film, c.Param("id")); outcome.Error != nil {
		c.AbortWithError(http.StatusNotFound, outcome.Error)
	}
	c.JSON(http.StatusOK, gin.H{"data": film})
}

// @Summary      Create a film
// @Description  Create a film
// @Tags         films
// @Produce      json
// @Success      200  {object} model.Film
// @Failure	 	 500
// @Failure		 400
// @Router       /v1/films [post]
// @Security     ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param		 data body CreateFilmBody true "define body of the film to create"
func CreateFilm(c *gin.Context) {
	var serializedFilm CreateFilmBody
	if c.BindJSON(&serializedFilm) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	var realFilm model.Film
	realFilm.Title = serializedFilm.Title
	realFilm.ReleaseDate = serializedFilm.ReleaseDate
	realFilm.Director = serializedFilm.Director
	realFilm.Synopsis = serializedFilm.Synopsis
	realFilm.Genres = generateGenresList(serializedFilm.Genres)
	realFilm.CreatedAt = time.Now()
	realFilm.UpdatedAt = time.Now()

	if user, found := c.Get("user"); found {
		realFilm.UserID = user.(model.User).ID
	}

	if outcome := initializers.DB.Model(&model.Film{}).Create(&realFilm); outcome.Error != nil {
		c.AbortWithError(http.StatusBadRequest, outcome.Error)
	}
	c.JSON(http.StatusCreated, gin.H{"data": realFilm})
}

// @Summary      Update a film
// @Description  Update a film
// @Tags         films
// @Produce      json
// @Success      200  {object} model.Film
// @Failure	 	 500
// @Failure		 400
// @Router       /v1/films/:id [put]
// @Security     ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param		 data body CreateFilmBody true "define body of the film to update"
// @Param		 id path int true "film id"
func UpdateFilm(c *gin.Context) {
	var (
		film           model.Film
		createFilmBody CreateFilmBody
	)
	if outcome := initializers.DB.Model(&model.Film{}).Find(&film, c.Param("id")); outcome.Error != nil {
		c.AbortWithError(http.StatusNotFound, outcome.Error)
	}
	if user, found := c.Get("user"); found {
		if user.(model.User).ID != film.UserID {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
	if c.BindJSON(&createFilmBody) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if outcome := initializers.DB.Model(&model.Film{}).Save(&film); outcome.Error != nil {
		c.AbortWithError(http.StatusBadRequest, outcome.Error)
	}
	c.JSON(http.StatusOK, gin.H{"data": film})
}

// @Summary      Delete a film
// @Description  Delete a film
// @Tags         films
// @Produce      json
// @Success      200  {object} model.Film
// @Failure	 	 500
// @Failure		 400
// @Router       /v1/films/:id [delete]
// @Security     ApiKeyAuth
// @Param		 id path int true "film id"
// @param 		Authorization header string true "Authorization"
func DeleteFilm(c *gin.Context) {
	var film model.Film
	if outcome := initializers.DB.Model(&model.Film{}).Find(&film, c.Param("id")); outcome.Error != nil {
		c.AbortWithError(http.StatusNotFound, outcome.Error)
	}
	if user, found := c.Get("user"); found {
		if user.(model.User).ID != film.UserID {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
	if outcome := initializers.DB.Model(&model.Film{}).Delete(&film); outcome.Error != nil {
		c.AbortWithError(http.StatusBadRequest, outcome.Error)
	}
	c.JSON(http.StatusOK, gin.H{"data": film})
}

type CreateFilmBody struct {
	Title       string    `json:"title" binding:"required,min=3,max=300"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Director    string    `json:"director" binding:"required,min=3,max=300"`
	Cast        string    `json:"cast" binding:"required,max=1000"`
	Synopsis    string    `json:"synopsis" binding:"required,max=1000"`
	Genres      []string  `json:"genres" binding:"required"`
}

func generateGenresList(genres []string) []model.Genre {
	var genresList []model.Genre

	for _, genrestr := range genres {
		var genre model.Genre
		initializers.DB.Model(&model.Genre{}).Where("name = ?", genrestr).First(&genre)
		if genre.ID != 0 {
			genresList = append(genresList, genre)
		} else {
			genre.Name = genrestr
			initializers.DB.Model(&model.Genre{}).Create(&genre)
			genresList = append(genresList, genre)
		}
	}
	return genresList
}
