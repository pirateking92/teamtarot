package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"main.go/env"
	"main.go/models"
	// "gorm.io/gorm"
)

type TestSuiteEnv struct {
	suite.Suite
	// db      *gorm.DB
	// token   string
	// token2  string
	// userID  int
	// userID2 int
	app *gin.Engine
	res *httptest.ResponseRecorder
}

// Tests are run before they start
func (suite *TestSuiteEnv) SetupSuite() {
	env.LoadEnv(".test.env")
	// models.OpenDatabaseConnection()
	// models.AutoMigrateModels()
	// suite.db = models.Database
	suite.app = setupApp()
	// suite.userID = 1
	// suite.token, _ = auth.GenerateToken(uint(suite.userID))
	// suite.userID2 = 2
	// suite.token2, _ = auth.GenerateToken(uint(suite.userID2))

}

func (suite *TestSuiteEnv) SetupTest() {
	suite.res = httptest.NewRecorder()
	// suite.TearDownTest()
}

// Running after each test
// func (suite *TestSuiteEnv) TearDownTest() {
// 	suite.db.Exec("DELETE FROM comments")
// 	suite.db.Exec("DELETE FROM posts")
// 	suite.db.Exec("DELETE FROM users")

// }

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(TestSuiteEnv))
}

// Can retrieve 3 cards successfully
func (suite *TestSuiteEnv) Test_GetThreeCards() {
	app := suite.app

	req, _ := http.NewRequest("GET", "/cards", nil)
	app.ServeHTTP(suite.res, req)

	responseData, _ := io.ReadAll(suite.res.Body)
	var jsonCards struct {
		Cards []models.JSONCard
	}

	_ = json.Unmarshal(responseData, &jsonCards)

	assert.Equal(suite.T(), 200, suite.res.Code)
	assert.Len(suite.T(), jsonCards.Cards, 3)
	// assert.Contains(suite.T(), jsonCards.Cards[0], "CardName")
}

// Checks that two different responses given two different sets of cards
func (suite *TestSuiteEnv) Test_GetThreeCardsIsRandom() {
	app := suite.app

	//Response 1
	req, _ := http.NewRequest("GET", "/cards", nil)
	app.ServeHTTP(suite.res, req)

	responseData, _ := io.ReadAll(suite.res.Body)
	var jsonCards struct {
		Cards []models.JSONCard
	}

	_ = json.Unmarshal(responseData, &jsonCards)
	//Response 2
	req2, _ := http.NewRequest("GET", "/cards", nil)
	app.ServeHTTP(suite.res, req2)

	responseData2, _ := io.ReadAll(suite.res.Body)
	var jsonCards2 struct {
		Cards []models.JSONCard
	}

	_ = json.Unmarshal(responseData2, &jsonCards2)
	assert.NotEqual(suite.T(), jsonCards.Cards[0].CardName, jsonCards2.Cards[0].CardName) //1 in 24336 of this failing!!!
}
