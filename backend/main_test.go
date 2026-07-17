package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pirateking92/teamtarot/backend/env"
	"github.com/pirateking92/teamtarot/backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// The suite runs against an isolated in-memory SQLite DB with a small fixture
// deck, so tests need no Postgres, no network, and no API key. Interpretation
// is stubbed via cfg.Testing.
type TestSuiteEnv struct {
	suite.Suite
	app      *gin.Engine
	db       *gorm.DB
	res      *httptest.ResponseRecorder
	deckSize int
}

func (s *TestSuiteEnv) SetupSuite() {
	gin.SetMode(gin.TestMode)

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	s.Require().NoError(err)
	s.Require().NoError(db.AutoMigrate(
		&models.Card{}, &models.User{}, &models.Session{},
		&models.Reading{}, &models.ReadingCard{}, &models.DailyCard{},
	))

	// Seed a small fixture deck (enough for every spread up to celtic-cross).
	suits := []string{"wands", "cups", "swords", "pentacles"}
	var cards []models.Card
	for i := 1; i <= 4; i++ {
		cards = append(cards, models.Card{
			Name: "Major " + string(rune('A'+i)), NameShort: "ar0" + string(rune('0'+i)),
			Arcana: "major", MeaningUp: "up meaning", MeaningReverse: "rev meaning",
			Description: "desc", ImageFileName: "ar0" + string(rune('0'+i)) + ".jpg",
		})
	}
	for si, suit := range suits {
		for r := 2; r <= 4; r++ {
			ns := suit[:2] + "0" + string(rune('0'+r)) + string(rune('0'+si))
			cards = append(cards, models.Card{
				Name: suit + " card", NameShort: ns, Arcana: "minor", Suit: suit,
				MeaningUp: "up", MeaningReverse: "rev", Description: "d", ImageFileName: ns + ".jpg",
			})
		}
	}
	s.Require().NoError(db.Create(&cards).Error)
	s.deckSize = len(cards)

	s.db = db
	cfg := env.Config{Testing: true, FrontendOrigin: "http://localhost:5173"}
	s.app = setupApp(cfg, db)
}

func (s *TestSuiteEnv) deckCount() int { return s.deckSize }

func (s *TestSuiteEnv) SetupTest() { s.res = httptest.NewRecorder() }

func (s *TestSuiteEnv) do(method, path, body string) []byte {
	return s.doTok(method, path, body, "")
}

func (s *TestSuiteEnv) doTok(method, path, body, token string) []byte {
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	req, _ := http.NewRequest(method, path, reader)
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	s.app.ServeHTTP(s.res, req)
	data, _ := io.ReadAll(s.res.Body)
	return data
}

// register creates a user and returns its session token.
func (s *TestSuiteEnv) register(email string) string {
	s.res = httptest.NewRecorder()
	data := s.doTok("POST", "/api/auth/register", `{"email":"`+email+`","password":"supersecret","displayName":"T"}`, "")
	s.Require().Equal(http.StatusCreated, s.res.Code, string(data))
	var out struct {
		Token string `json:"token"`
	}
	s.Require().NoError(json.Unmarshal(data, &out))
	return out.Token
}

func TestSuite(t *testing.T) { suite.Run(t, new(TestSuiteEnv)) }

func (s *TestSuiteEnv) Test_Health() {
	s.do("GET", "/api/health", "")
	assert.Equal(s.T(), http.StatusOK, s.res.Code)
}

func (s *TestSuiteEnv) Test_Spreads() {
	data := s.do("GET", "/api/spreads", "")
	assert.Equal(s.T(), http.StatusOK, s.res.Code)
	var out struct {
		Spreads []struct {
			Type      string `json:"type"`
			Positions []any  `json:"positions"`
		} `json:"spreads"`
	}
	s.Require().NoError(json.Unmarshal(data, &out))
	assert.Len(s.T(), out.Spreads, 4)
}

func (s *TestSuiteEnv) Test_ListCards() {
	data := s.do("GET", "/api/cards", "")
	assert.Equal(s.T(), http.StatusOK, s.res.Code)
	var out struct {
		Cards []models.Card `json:"cards"`
	}
	s.Require().NoError(json.Unmarshal(data, &out))
	assert.Len(s.T(), out.Cards, s.deckCount())
}

func (s *TestSuiteEnv) Test_CreateAndPollReading() {
	data := s.do("POST", "/api/readings", `{"querentName":"Test","story":"a story","spreadType":"past-present-future"}`)
	assert.Equal(s.T(), http.StatusCreated, s.res.Code)

	var created struct {
		Reading models.Reading `json:"reading"`
	}
	s.Require().NoError(json.Unmarshal(data, &created))
	assert.Len(s.T(), created.Reading.Cards, 3)
	assert.NotEqual(s.T(), "", created.Reading.ID.String())

	// Poll for the async (stubbed) interpretation to complete.
	id := created.Reading.ID.String()
	var status string
	for i := 0; i < 20; i++ {
		s.res = httptest.NewRecorder()
		poll := s.do("GET", "/api/readings/"+id, "")
		var got struct {
			Reading models.Reading `json:"reading"`
		}
		s.Require().NoError(json.Unmarshal(poll, &got))
		status = got.Reading.Status
		if status == models.ReadingComplete {
			assert.NotEmpty(s.T(), got.Reading.Interpretation)
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	assert.Equal(s.T(), models.ReadingComplete, status)
}

func (s *TestSuiteEnv) Test_GetReading_NotFound() {
	s.do("GET", "/api/readings/11111111-1111-1111-1111-111111111111", "")
	assert.Equal(s.T(), http.StatusNotFound, s.res.Code)
}

func (s *TestSuiteEnv) Test_ListReadings_RequiresAuth() {
	s.do("GET", "/api/readings", "")
	assert.Equal(s.T(), http.StatusUnauthorized, s.res.Code)
}

func (s *TestSuiteEnv) Test_Register_Me_And_DuplicateEmail() {
	token := s.register("me-test@example.com")

	s.res = httptest.NewRecorder()
	data := s.doTok("GET", "/api/auth/me", "", token)
	assert.Equal(s.T(), http.StatusOK, s.res.Code)
	var me struct {
		User models.User `json:"user"`
	}
	s.Require().NoError(json.Unmarshal(data, &me))
	assert.Equal(s.T(), "me-test@example.com", me.User.Email)

	// Duplicate registration is rejected.
	s.res = httptest.NewRecorder()
	s.doTok("POST", "/api/auth/register", `{"email":"me-test@example.com","password":"supersecret"}`, "")
	assert.Equal(s.T(), http.StatusConflict, s.res.Code)
}

func (s *TestSuiteEnv) Test_Login_WrongPassword() {
	s.register("login-test@example.com")

	s.res = httptest.NewRecorder()
	s.doTok("POST", "/api/auth/login", `{"email":"login-test@example.com","password":"wrongpass"}`, "")
	assert.Equal(s.T(), http.StatusUnauthorized, s.res.Code)
}

func (s *TestSuiteEnv) Test_AuthedReading_ShowsInHistory() {
	token := s.register("history-test@example.com")

	s.res = httptest.NewRecorder()
	data := s.doTok("POST", "/api/readings", `{"querentName":"T","story":"s","spreadType":"single"}`, token)
	s.Require().Equal(http.StatusCreated, s.res.Code)
	var created struct {
		Reading models.Reading `json:"reading"`
	}
	s.Require().NoError(json.Unmarshal(data, &created))

	s.res = httptest.NewRecorder()
	data = s.doTok("GET", "/api/readings", "", token)
	assert.Equal(s.T(), http.StatusOK, s.res.Code)
	var hist struct {
		Readings []models.Reading `json:"readings"`
	}
	s.Require().NoError(json.Unmarshal(data, &hist))
	found := false
	for _, r := range hist.Readings {
		if r.ID == created.Reading.ID {
			found = true
		}
	}
	assert.True(s.T(), found, "created reading should appear in the user's history")
}
