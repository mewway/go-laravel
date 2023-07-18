//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package database

import (
	"context"

	"github.com/google/wire"

	"github.com/mewway/go-laravel/contracts/config"
	"github.com/mewway/go-laravel/database/db"
	"github.com/mewway/go-laravel/database/gorm"
)

//go:generate wire
func InitializeOrm(ctx context.Context, config config.Config, connection string) (*OrmImpl, error) {
	wire.Build(NewOrmImpl, gorm.QuerySet, gorm.GormSet, db.ConfigSet, gorm.DialectorSet)

	return nil, nil
}
