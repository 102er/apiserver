//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/102er/apiserver/internal"
	"github.com/102er/apiserver/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initApi() (*gin.Engine, error) {
	panic(wire.Build(internal.ProviderSet, server.NewGinHttpServer))
}
