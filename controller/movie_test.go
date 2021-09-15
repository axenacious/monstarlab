package controller

import (
	// "bytes"
	// "fmt"
	// "net/http"
	"net/http/httptest"
	 "testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"monstarlab/config"
)

type MovieTestSuite struct {
	suite.Suite
	rec     *httptest.ResponseRecorder
	context *gin.Context
	app     *gin.Engine
	ctrl    *MovieController
}

func (suite *MovieTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	config.Server.Mode = gin.ReleaseMode
	suite.rec = httptest.NewRecorder()
	suite.context, suite.app = gin.CreateTestContext(suite.rec)
	suite.ctrl = new(MovieController)
}

func (suite *MovieTestSuite) TestMovies() {
	suite.ctrl.Movies(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func (suite *MovieTestSuite) TestGetMovies() {
	suite.context.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	suite.ctrl.GetMovies(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func TestMovieTestSuite(t *testing.T) {
	suite.Run(t, new(MovieTestSuite))
}
