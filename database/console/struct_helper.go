package console

import (
	"errors"

	"github.com/mewway/go-laravel/contracts/config"
	"github.com/mewway/go-laravel/contracts/database/orm"
	"github.com/mewway/go-laravel/database/db"
	"github.com/mewway/go-laravel/database/gorm"
	gormio "gorm.io/gorm"
)

type StructHelperInterface interface {
}

type StructHelper struct {
	StructHelperInterface
	Db *gormio.DB
}

type Column struct {
	ColumnName    string
	ColumnType    string
	IsNullable    string
	ColumnKey     string
	ColumnComment string
	Extra         string
}

func NewStructHelper(conf config.Config) (helper *StructHelper, err error) {
	helper = &StructHelper{}
	conn := conf.GetString("database.default")
	driver := conf.GetString("database.connections." + conn + ".driver")

	gormConfig := db.NewConfigImpl(conf, conn)
	writeConfigs := gormConfig.Writes()
	if len(writeConfigs) == 0 {
		return nil, errors.New("can't not found database configuration")
	}
	switch orm.Driver(driver) {
	case orm.DriverMysql:
		dail := gorm.NewDialectorImpl(conf, conn)
		dbConf := db.NewConfigImpl(conf, conn)
		instance, err := gorm.NewGormImpl(conf, conn, dbConf, dail).Make()
		if err != nil {
			return nil, errors.New("can't not connect to database by the given configuration")
		}
		helper.Db = instance

		return helper, nil
	}
	return helper, nil
}

func (s StructHelper) getTableStruct(database, table string) []*Column {
	sqlPattern := "SELECT COLUMN_NAME as column_name, COLUMN_TYPE as column_type, IS_NULLABLE as is_nullable, COLUMN_KEY as column_key, COLUMN_COMMENT as column_comment, EXTRA as extra FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?;"
	var cols []*Column
	s.Db.Raw(sqlPattern, database, table).Scan(cols)
	return cols
}
