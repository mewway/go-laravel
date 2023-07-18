package facades

import (
	foundationcontract "github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/foundation"
)

func App() foundationcontract.Application {
	return foundation.App
}
