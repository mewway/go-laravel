package facades

import (
	"github.com/mewway/go-laravel/contracts/hash"
)

func Hash() hash.Hash {
	return App().MakeHash()
}
