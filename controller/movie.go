package controller

import (
	"github.com/gin-gonic/gin"
	"monstarlab/model"
	"net/http"
	"strconv"
)

// MovieController is the movie controller
type MovieController struct{}

func (ctrl *MovieController) Movies(c *gin.Context) {
	var movies model.Movie
	searchQuery := c.Query("search")

	if searchQuery != "" {
		if err := movies.Search(searchQuery); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := movies.GetAllMovies(); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Movie_list": movies})
	return
}

func (ctrl *MovieController) GetMovies(c *gin.Context) {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	
	var movies model.Movie

	ids := []string{c.Param("id")}

	if err := movies.GetMovieByID(ids); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Movie_list": movies})
	return
}
