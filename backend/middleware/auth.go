// Package middleware holds Gin middleware, currently authentication.
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pirateking92/teamtarot/backend/services"
	"gorm.io/gorm"
)

// SessionCookieName is the cookie the session token is stored under.
const SessionCookieName = "session"

// TokenFromRequest extracts a session token from the session cookie or an
// Authorization: Bearer header (the BFF forwards it as a bearer token).
func TokenFromRequest(c *gin.Context) string {
	if cookie, err := c.Cookie(SessionCookieName); err == nil && cookie != "" {
		return cookie
	}
	auth := c.GetHeader("Authorization")
	if after, ok := strings.CutPrefix(auth, "Bearer "); ok {
		return after
	}
	return ""
}

// OptionalAuth attaches the current user to the context if a valid session is
// present, but never blocks the request. Guests simply have no user set.
func OptionalAuth(gormDB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := TokenFromRequest(c); token != "" {
			if user, err := services.UserForSession(gormDB, token); err == nil {
				c.Set("userID", user.ID)
				c.Set("user", user)
			}
		}
		c.Next()
	}
}

// RequireAuth aborts with 401 unless OptionalAuth has already attached a user.
// It must run after OptionalAuth in the middleware chain.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get("userID"); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}
		c.Next()
	}
}
