package http

import (
	"github.com/DarkSoul94/metamemorys_backend/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, ucAuth auth.AuthUC) {
	h := NewAuthHandler(ucAuth)

	authEndpoints := router.Group("/auth")
	{
		//http://localhost:8000/metamemorys/auth/signup
		authEndpoints.POST("/signup", h.SignUp)
		//http://localhost:8000/metamemorys/auth/signin
		authEndpoints.POST("/signin", h.SignIn)

	}
}
