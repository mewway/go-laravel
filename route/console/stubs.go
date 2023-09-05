package console

type Stubs struct {
}

func (r Stubs) NewRoute(withApi bool) string {
	if withApi {
		return `package DummyPackage

import (
	"github.com/mewway/go-laravel/contracts/http"
	"github.com/mewway/go-laravel/contracts/route"
	"DummyImportController"
)

func DummyFunction(route route.Route) {
	route.Prefix("DummyApiPrefix").Group(func(routes route.Route) {
		routes.DummyMethod("DummySuffix", DummyApiController)
	})
}
`
	}
	return `package DummyPackage

import (
	"github.com/mewway/go-laravel/contracts/http"
	"github.com/mewway/go-laravel/contracts/route"
	"DummyImportController"
)

func DummyFunction(route route.Route) {
	route.Prefix("DummyApiPrefix").Group(func(routes route.Route) {
		routes.DummyMethod("DummySuffix", func(context http.Context) {
			context.Response().Json(200, map[string]string{"hello": "world"})
		})
	})
}
`
}
