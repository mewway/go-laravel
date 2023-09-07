package doc

import (
	consolecontract "github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/doc/console"
)

const Binding = "cicada.doc"

type ServiceProvider struct {
}

func (doc *ServiceProvider) Register(app foundation.Application) {

}

func (doc *ServiceProvider) Boot(app foundation.Application) {
	doc.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app foundation.Application) {
	artisan := app.MakeArtisan()
	artisan.Register([]consolecontract.Command{
		console.NewYapiListCommand(artisan),
		console.NewApifoxListCommand(artisan),
	})
}
