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
	// env.LoadEnv(".test.env")
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

// {
//     "cards": [
//         {
//             "name": "Two of Pentacles",
//             "type": "minor",
//             "meaning_up": "On the one hand it is represented as a card of gaiety, recreation and its connexions, which is the subject of the design; but it is read also as news and messages in writing, as obstacles, agitation, trouble, embroilment.",
//             "meaning_rev": "Enforced gaiety, simulated enjoyment, literal sense, handwriting, composition, letters of exchange.",
//             "desc": "A young man, in the act of dancing, has a pentacle in either hand, and they are joined by that endless cord which is like the number 8 reversed.",
//             "image_file_name": "pe02.jpg"
//         },
//         {
//             "name": "The Hanged Man",
//             "type": "major",
//             "meaning_up": "Wisdom, circumspection, discernment, trials, sacrifice, intuition, divination, prophecy.",
//             "meaning_rev": "Selfishness, the crowd, body politic.",
//             "desc": "The gallows from which he is suspended forms a Tau cross, while the figure--from the position of the legs--forms a fylfot cross. There is a nimbus about the head of the seeming martyr. It should be noted (1) that the tree of sacrifice is living wood, with leaves thereon; (2) that the face expresses deep entrancement, not suffering; (3) that the figure, as a whole, suggests life in suspension, but life and not death. It is a card of profound significance, but all the significance is veiled. One of his editors suggests that Éliphas Lévi did not know the meaning, which is unquestionable nor did the editor himself. It has been called falsely a card of martyrdom, a card a of prudence, a card of the Great Work, a card of duty; but we may exhaust all published interpretations and find only vanity. I will say very simply on my own part that it expresses the relation, in one of its aspects, between the Divine and the Universe.\nHe who can understand that the story of his higher nature is imbedded in this symbolism will receive intimations concerning a great awakening that is possible, and will know that after the sacred Mystery of Death there is a glorious Mystery of Resurrection.",
//             "image_file_name": "ar12.jpg"
//         },
//         {
//             "name": "The Last Judgment",
//             "type": "major",
//             "meaning_up": "Change of position, renewal, outcome. Another account specifies total loss though lawsuit.",
//             "meaning_rev": "Weakness, pusillanimity, simplicity; also deliberation, decision, sentence.",
//             "desc": "I have said that this symbol is essentially invariable in all Tarot sets, or at least the variations do not alter its character. The great angel is here encompassed by clouds, but he blows his bannered trumpet, and the cross as usual is displayed on the banner. The dead are rising from their tombs--a woman on the right, a man on the left hand, and between them their child, whose back is turned. But in this card there are more than three who are restored, and it has been thought worth while to make this variation as illustrating the insufficiency of current explanations. It should be noted that all the figures are as one in the wonder, adoration and ecstacy expressed by their attitudes. It is the card which registers the accomplishment of the great work of transformation in answer to the summons of the Supernal--which summons is heard and answered from within.\nHerein is the intimation of a significance which cannot well be carried further in the present place. What is that within us which does sound a trumpet and all that is lower in our nature rises in response--almost in a moment, almost in the twinkling of an eye? Let the card continue to depict, for those who can see no further, the Last judgment and the resurrection in the natural body; but let those who have inward eyes look and discover therewith. They will understand that it has been called truly in the past a card of eternal life, and for this reason it may be compared with that which passes under the name of Temperance.",
//             "image_file_name": "ar20.jpg"
//         }
//     ]
// }

// Can create a user successfully
// func (suite *TestSuiteEnv) Test_PostUsers_CorrectJSON() {
// 	app, token := suite.app, suite.token

// 	res := httptest.NewRecorder()
// 	var jsonStr = []byte(`{"email":"tester@email.com", "password": "testpassword", "username": "tester"}`)
// 	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
// 	app.ServeHTTP(res, req)

// 	assert.Equal(suite.T(), 201, res.Code)
// }

// // Not allowed to create a user with incorrect JSON data
// func (suite *TestSuiteEnv) Test_PostUsers_IncorrectJSON() {
// 	app, token := suite.app, suite.token

// 	var jsonStr = []byte(`{"message":"Test Post"}`)
// 	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
// 	app.ServeHTTP(suite.res, req)

// 	assert.Equal(suite.T(), 400, suite.res.Code)
// }

// // Can retrieve all posts
// func (suite *TestSuiteEnv) Test_GetPosts() {
// 	app, token := suite.app, suite.token
// 	// Create a user
// 	user := models.User{
// 		Email:    "testuser6@example.com",
// 		Password: "password123",
// 	}
// 	user.Save()
// 	newPost := models.Post{
// 		Message: "Test Post",
// 		UserID:  user.ID,
// 	}
// 	newPost.Save()

// 	req, _ := http.NewRequest("GET", "/posts", nil)
// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
// 	app.ServeHTTP(suite.res, req)

// 	responseData, _ := io.ReadAll(suite.res.Body)
// 	var jsonPosts struct {
// 		Posts []controllers.JSONPost
// 	}

// 	_ = json.Unmarshal(responseData, &jsonPosts)

// 	assert.Equal(suite.T(), 200, suite.res.Code)
// 	assert.Equal(suite.T(), "Test Post", jsonPosts.Posts[0].Message)
// }
