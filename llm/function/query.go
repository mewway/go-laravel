// Package function
// @Description: 实现从数据库中读取数据表 数据列 注释等其他功能
package function

import (
	"encoding/json"
	"fmt"

	"github.com/mewway/go-laravel/contracts/config"
	"github.com/mewway/go-laravel/contracts/database/orm"
	"github.com/mewway/go-laravel/database/db"
	"github.com/mewway/go-laravel/database/gorm"
	"github.com/spf13/cast"
	gormio "gorm.io/gorm"
)

type Conn struct {
	DB        *gormio.DB
	DefaultDb string
}

func NewDB(conf config.Config) (database *Conn) {
	database = &Conn{}
	conn := conf.GetString("database.default")
	driver := conf.GetString("database.connections." + conn + ".driver")

	gormConfig := db.NewConfigImpl(conf, conn)
	writeConfigs := gormConfig.Writes()
	if len(writeConfigs) == 0 {
		return nil
	}
	switch orm.Driver(driver) {
	case orm.DriverMysql:
		dail := gorm.NewDialectorImpl(conf, conn)
		dbConf := db.NewConfigImpl(conf, conn)
		instance, err := gorm.NewGormImpl(conf, conn, dbConf, dail).Make()
		if err != nil {
			return nil
		}
		database.DB = instance
		database.DefaultDb = conf.GetString(fmt.Sprintf("database.connections.%s.database", conn))
		return database
	}
	return database
}
func (db *Conn) QueryTableDistinctCountMap(database, table string, cols []string) string {
	sqlPattern := "SELECT count(DISTINCT(`%s`)) as d_count FROM `%s`.`%s`;"
	countMap := make(map[string]int)
	result := &struct {
		DCount int `json:"d_count"`
	}{}
	for _, col := range cols {
		db.DB.Raw(fmt.Sprintf(sqlPattern, col, database, table)).Scan(result)
		countMap[col] = cast.ToInt(result.DCount)
	}
	indent, err := json.Marshal(countMap)
	if err != nil {
		return ""
	}
	return string(indent)
}

func (db *Conn) DDL(database, table string) string {
	sqlPattern := "SHOW CREATE TABLE `%s`.`%s`;"
	result := struct {
		Name        string
		CreateTable string `gorm:"column:Create Table"`
	}{}
	db.DB.Raw(fmt.Sprintf(sqlPattern, database, table)).Scan(&result)
	return result.CreateTable
}
