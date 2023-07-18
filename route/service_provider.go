package route

import (
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.route"

type ServiceProvider struct {
}

func (route *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewGin(app.MakeConfig()), nil
	})
}

func (route *ServiceProvider) Boot(app foundation.Application) {

}
