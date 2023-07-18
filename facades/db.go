package facades

import (
	"github.com/mewway/go-laravel/contracts/database/orm"
)

func Orm() orm.Orm {
	return App().MakeOrm()
}
