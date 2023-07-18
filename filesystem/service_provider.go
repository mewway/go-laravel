package filesystem

import (
	configcontract "github.com/mewway/go-laravel/contracts/config"
	filesystemcontract "github.com/mewway/go-laravel/contracts/filesystem"
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.filesystem"

var ConfigFacade configcontract.Config
var StorageFacade filesystemcontract.Storage

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewStorage(app.MakeConfig()), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	StorageFacade = app.MakeStorage()
}
