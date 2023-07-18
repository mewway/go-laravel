package auth

import (
	"context"

	"github.com/mewway/go-laravel/auth/access"
	"github.com/mewway/go-laravel/auth/console"
	contractconsole "github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/foundation"
)

const BindingAuth = "cicada.auth"
const BindingGate = "cicada.gate"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(BindingAuth, func(app foundation.Application) (any, error) {
		config := app.MakeConfig()
		return NewAuth(config.GetString("auth.defaults.guard"),
			app.MakeCache(), config, app.MakeOrm()), nil
	})
	app.Singleton(BindingGate, func(app foundation.Application) (any, error) {
		return access.NewGate(context.Background()), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]contractconsole.Command{
		console.NewJwtSecretCommand(app.MakeConfig()),
		console.NewPolicyMakeCommand(),
	})
}
