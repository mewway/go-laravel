package hash

import (
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.hash"

type ServiceProvider struct {
}

func (hash *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeConfig()), nil
	})
}

func (hash *ServiceProvider) Boot(app foundation.Application) {

}
