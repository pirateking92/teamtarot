package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pirateking92/teamtarot/backend/middleware"
	"github.com/pirateking92/teamtarot/backend/models"
	"github.com/pirateking92/teamtarot/backend/services"
)

const sessionMaxAge = int(services.SessionDuration / time.Second) // in seconds

type registerRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register creates an account and starts a session.
// POST /api/auth/register
func (h *Handler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if req.Email == "" || len(req.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required and password must be at least 8 characters"})
		return
	}

	user, token, err := services.Register(h.DB, req.Email, req.Password, req.DisplayName)
	if errors.Is(err, services.ErrEmailTaken) {
		c.JSON(http.StatusConflict, gin.H{"error": "that email is already registered"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create account"})
		return
	}

	h.setSessionCookie(c, token)
	c.JSON(http.StatusCreated, gin.H{"user": user, "token": token})
}

// Login verifies credentials and starts a session.
// POST /api/auth/login
func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user, token, err := services.Login(h.DB, req.Email, req.Password)
	if errors.Is(err, services.ErrInvalidCredentials) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not sign in"})
		return
	}

	h.setSessionCookie(c, token)
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

// Logout revokes the current session.
// POST /api/auth/logout
func (h *Handler) Logout(c *gin.Context) {
	if token := middleware.TokenFromRequest(c); token != "" {
		_ = services.DeleteSession(h.DB, token)
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(middleware.SessionCookieName, "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "signed out"})
}

// Me returns the authenticated user.
// GET /api/auth/me  (requires auth)
func (h *Handler) Me(c *gin.Context) {
	v, _ := c.Get("user")
	user, ok := v.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) setSessionCookie(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteLaxMode)
	// secure=false for local http dev; the BFF also stores its own first-party
	// cookie, so this cookie is primarily for standalone/API use.
	c.SetCookie(middleware.SessionCookieName, token, sessionMaxAge, "/", "", false, true)
}
