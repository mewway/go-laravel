package route

import (
	"fmt"

	consolecontract "github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/route/console"
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
	route.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app foundation.Application) {
	artisan := app.MakeArtisan()
	gin, err := app.Make(Binding)
	if err != nil {
		return
	}
	callback := func() []string {
		m := make([]string, 0)
		g := gin.(*Gin)
		routes := g.GetInstance().Routes()
		for _, v := range routes {
			m = append(m, fmt.Sprintf("【%s】%s", v.Method, v.Path))
		}
		return m
	}
	artisan.Register([]consolecontract.Command{
		console.NewListCommand(artisan, callback),
		console.NewAddCommand(artisan, callback),
	})
}
