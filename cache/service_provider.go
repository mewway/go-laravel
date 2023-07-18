package cache

import (
	"github.com/mewway/go-laravel/cache/console"
	contractsconsole "github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.cache"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		config := app.MakeConfig()
		log := app.MakeLog()
		store := config.GetString("cache.default")

		return NewApplication(config, log, store)
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]contractsconsole.Command{
		console.NewClearCommand(app.MakeCache()),
	})
}
