package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DarkSoul94/metamemorys_backend/auth"
	"github.com/DarkSoul94/metamemorys_backend/global_const"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	usecase auth.AuthUC
}

// NewAuthMiddleware ...
func NewAuthMiddleware(usecase auth.AuthUC) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: usecase,
	}).Handle
}

// Handle ...
func (m *AuthMiddleware) Handle(c *gin.Context) {
	
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set(global_const.CtxUserKey, user)
}
