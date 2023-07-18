package log

import (
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.log"

type ServiceProvider struct {
}

func (log *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewLogrusApplication(app.MakeConfig()), nil
	})
}

func (log *ServiceProvider) Boot(app foundation.Application) {

}
