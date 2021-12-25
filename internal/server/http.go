package server

import (
	"github.com/102er/apiserver/internal/server/route"
	"github.com/102er/apiserver/internal/service"
	"github.com/gin-gonic/gin"
)

func NewGinHttpServer(user *service.UserService) *gin.Engine {
	router := gin.Default()
	rg := router.Group("/api/v1")
	route.RegisterUserHandler(rg, user)
	return router
}
