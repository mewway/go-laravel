package grpc

import (
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/contracts/log"
)

const Binding = "cicada.grpc"

var LogFacade log.Log

type ServiceProvider struct {
}

func (route *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeConfig()), nil
	})
}

func (route *ServiceProvider) Boot(app foundation.Application) {
	LogFacade = app.MakeLog()
}
