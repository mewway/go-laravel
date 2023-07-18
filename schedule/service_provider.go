package schedule

import (
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.schedule"

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeArtisan(), app.MakeCache(), app.MakeLog()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {

}
