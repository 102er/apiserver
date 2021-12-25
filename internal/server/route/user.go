package route

import (
	"context"
	"github.com/gin-gonic/gin"
)

type IUserServer interface {
	Create(ctx context.Context)
}

func RegisterUserHandler(rg *gin.RouterGroup, srv IUserServer) *UserHandler {
	uh := &UserHandler{
		srv: srv,
	}
	rg.POST("/", uh.Add)
	rg.DELETE("/:id", uh.Add)
	return uh
}

type UserHandler struct {
	srv IUserServer
}

func (u *UserHandler) Add(ctx *gin.Context) {
	ctx.JSONP(200, gin.H{
		"status": 200,
	})
}
