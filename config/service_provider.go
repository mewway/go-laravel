package config

import (
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/support"
)

const Binding = "cicada.config"

type ServiceProvider struct {
}

func (config *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(support.EnvPath), nil
	})
}

func (config *ServiceProvider) Boot(app foundation.Application) {

}
