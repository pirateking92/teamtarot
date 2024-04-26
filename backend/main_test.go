package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"main.go/controllers"
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

func RequestSetup(app *gin.Engine, suite *TestSuiteEnv, reqType string, path string) []byte {
	req, _ := http.NewRequest(reqType, path, nil)
	app.ServeHTTP(suite.res, req)
	responseData, _ := io.ReadAll(suite.res.Body)
	return responseData
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

// Checks that the response code is 200
func (suite *TestSuiteEnv) Test_GetThreeCards_ResponseCode() {
	app := suite.app
	responseData := RequestSetup(app, suite, "GET", "/cards")

	var jsonCards struct {
		Cards []models.JSONCard
	}

	_ = json.Unmarshal(responseData, &jsonCards)

	assert.Equal(suite.T(), 200, suite.res.Code)
}

// Checks that the response is 3 cards
func (suite *TestSuiteEnv) Test_GetThreeCards_ExpectedFormat() {
	app := suite.app
	responseData := RequestSetup(app, suite, "GET", "/cards")

	var jsonCards struct {
		Cards []models.JSONCard
	}

	_ = json.Unmarshal(responseData, &jsonCards)

	assert.Len(suite.T(), jsonCards.Cards, 3)
}

// Checks that two different responses given two different sets of cards
func (suite *TestSuiteEnv) Test_GetThreeCardsIsRandom() {
	app := suite.app

	//Response 1
	responseData := RequestSetup(app, suite, "GET", "/cards")
	var jsonCards struct {
		Cards []models.JSONCard
	}

	_ = json.Unmarshal(responseData, &jsonCards)
	//Response 2
	responseData2 := RequestSetup(app, suite, "GET", "/cards")
	var jsonCards2 struct {
		Cards []models.JSONCard
	}

	_ = json.Unmarshal(responseData2, &jsonCards2)
	assert.NotEqual(suite.T(), jsonCards.Cards[0].CardName, jsonCards2.Cards[0].CardName) //1 in 24336 of this failing!!!
}

// Checks that the response code is 200

func (suite *TestSuiteEnv) Test_GetAndInterpretCards_ResponseCode() {
	app := suite.app

	responseData := RequestSetup(app, suite, "GET", "/cards")
	var jsonCards struct {
		Cards     []models.JSONCard
		RequestID uuid.UUID
	}

	_ = json.Unmarshal(responseData, &jsonCards)

	_ = RequestSetup(app, suite, "GET", "/cards/interpret/"+jsonCards.RequestID.String())

	assert.Equal(suite.T(), 200, suite.res.Code)
}

// Checks that the interpretation response is as expected
func (suite *TestSuiteEnv) Test_GetAndInterpretCards_ExpectedFormat() {
	app := suite.app

	responseData := RequestSetup(app, suite, "GET", "/cards")
	var jsonCards struct {
		Cards     []models.JSONCard
		RequestID uuid.UUID
	}

	_ = json.Unmarshal(responseData, &jsonCards)

	responseData2 := RequestSetup(app, suite, "GET", "/cards/interpret/"+jsonCards.RequestID.String())
	var interpretationResponse struct {
		Interpretation string
	}

	_ = json.Unmarshal(responseData2, &interpretationResponse)

	assert.Equal(suite.T(), controllers.LocalStorage[jsonCards.RequestID.String()], "This is a test interpretation")
}
