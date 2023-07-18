package facades

import (
	"github.com/mewway/go-laravel/contracts/database/seeder"
)

func Seeder() seeder.Facade {
	return App().MakeSeeder()
}
