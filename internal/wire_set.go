package internal

import (
	"github.com/102er/apiserver/internal/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(service.NewUserService)
