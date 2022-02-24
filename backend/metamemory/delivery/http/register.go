package http

import (
	"github.com/DarkSoul94/metamemorys_backend/metamemory"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc metamemory.MetaUsecase, middlewares ...gin.HandlerFunc) {
	h := NewMetaHandler(uc)

	memberEndpoints := router.Group("/member")
	memberEndpoints.Use(middlewares...)
	{
		memberEndpoints.GET("/list", h.GetMemberList)
		memberEndpoints.POST("/create", h.CreateMember)
	}

	fileEndpoints := router.Group("/file")
	fileEndpoints.Use(middlewares...)
	{
		fileEndpoints.POST("/create", h.CreateFile)
		fileEndpoints.GET("/list", h.GetMemberFiles)
	}
}
