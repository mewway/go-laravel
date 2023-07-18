package crypt

import (
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.crypt"

type ServiceProvider struct {
}

func (crypt *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewAES(app.MakeConfig()), nil
	})
}

func (crypt *ServiceProvider) Boot(app foundation.Application) {

}
