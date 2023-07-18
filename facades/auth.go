package facades

import (
	"github.com/mewway/go-laravel/contracts/auth"
	"github.com/mewway/go-laravel/contracts/auth/access"
)

func Auth() auth.Auth {
	return App().MakeAuth()
}

func Gate() access.Gate {
	return App().MakeGate()
}
