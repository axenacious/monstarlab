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

type FavouriteTestSuite struct {
	suite.Suite
	rec     *httptest.ResponseRecorder
	context *gin.Context
	app     *gin.Engine
	ctrl    *FavouriteController
}

func (suite *FavouriteTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	config.Server.Mode = gin.ReleaseMode
	suite.rec = httptest.NewRecorder()
	suite.context, suite.app = gin.CreateTestContext(suite.rec)
	suite.ctrl = new(FavouriteController)
}

func (suite *FavouriteTestSuite) TestGetFavourite() {
	suite.context.Params = gin.Params{gin.Param{Key: "token", Value: "1"}}
	suite.ctrl.GetFavourite(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func (suite *FavouriteTestSuite) TestSaveFavourite() {
	suite.context.Params = gin.Params{gin.Param{Key: "token", Value: "1"}, gin.Param{Key: "id", Value: "1"}}
	suite.ctrl.GetFavourite(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func TestFavouriteTestSuite(t *testing.T) {
	suite.Run(t, new(FavouriteTestSuite))
}
