package router

import (
	"github.com/gin-gonic/gin"
	"monstarlab/controller"
	"monstarlab/middleware"
)

// Route makes the routing
func Route(app *gin.Engine) {
	authMiddleware := middleware.Auth()
	
	user := app.Group("/user")
	userController := new(controller.UserController)
	user.GET(
		"/user/:id", userController.GetUser,
	).GET(
		"/signup", userController.SignupForm,
	).POST(
		"/signup", userController.Signup,
	).GET(
		"/login", userController.LoginForm,
	).POST(
		"/login", authMiddleware.LoginHandler,
	)

	app.GET("/refresh_token", authMiddleware.RefreshHandler)
	app.Use(authMiddleware.MiddlewareFunc())
	{
		movieController := new(controller.MovieController)
		favouriteController := new(controller.FavouriteController)
		app.GET(
			"/movies/", movieController.Movies,
		).GET(
			"/movies/:id", movieController.GetMovies,
		).GET(
			"/favourite", favouriteController.GetFavourite,
		).POST(
			"/favourite/:id", favouriteController.SaveFavourite,
		)
	}
}
