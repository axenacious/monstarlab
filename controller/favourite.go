package controller

import (
	"github.com/gin-gonic/gin"
	"monstarlab/model"
	"net/http"
	jwt "github.com/appleboy/gin-jwt/v2"
	"fmt"
	"strconv"
)

// FavouriteController is the Favourite controller
type FavouriteController struct{}

func (ctrl *FavouriteController) GetFavourite(c *gin.Context) {
	var favourite model.Favourite
	var user model.User
	var movie model.Movie

	email := jwt.ExtractClaims(c)["email"]
	if err := user.GetFirstByEmail(fmt.Sprintf("%v", email)); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} 

	if err := favourite.GetFavourites(user.UserID); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	movieListID := getMovieID(favourite)

	if err := movie.GetMovieByID(movieListID); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Movie_list": movie})
	return
}

func (ctrl *FavouriteController) SaveFavourite(c *gin.Context) {
	favourite := model.Favourite{
		{
			UserID: "",
			MovieID: "",
		},
	}
	var user model.User

	movieID := c.Param("id")
	
	email := jwt.ExtractClaims(c)["email"]
	if err := user.GetFirstByEmail(fmt.Sprintf("%v", email)); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	favourite[0].UserID = strconv.Itoa(user.UserID)
	favourite[0].MovieID = movieID

	if err := favourite.SaveFavourite(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Favourite Added"})
	return
}

func getMovieID (favourite model.Favourite) ([]string){
	listID := []string{}

	for _, val := range favourite{
		listID = append(listID, val.MovieID)
	}

	return listID
}
