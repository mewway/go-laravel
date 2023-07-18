package facades

import "github.com/mewway/go-laravel/contracts/filesystem"

func Storage() filesystem.Storage {
	return App().MakeStorage()
}
