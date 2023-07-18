package console

import (
	"github.com/mewway/go-laravel/console/console"
	consolecontract "github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/foundation"
)

const Binding = "cicada.console"

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	receiver.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app foundation.Application) {
	artisan := app.MakeArtisan()
	config := app.MakeConfig()
	artisan.Register([]consolecontract.Command{
		console.NewListCommand(artisan),
		console.NewKeyGenerateCommand(config),
		console.NewMakeCommand(),
	})
}
