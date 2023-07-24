package console

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mewway/go-laravel/contracts/config"
	"github.com/mewway/go-laravel/contracts/database/orm"
	"github.com/mewway/go-laravel/database/db"
	"github.com/mewway/go-laravel/database/gorm"
	"github.com/mewway/go-laravel/support"
	"github.com/mewway/go-laravel/support/str"
	gormio "gorm.io/gorm"
)

type StructHelperInterface interface {
}

type StructHelper struct {
	StructHelperInterface
	Db        *gormio.DB
	DefaultDb string
}

type Column struct {
	ColumnName    string
	ColumnType    string
	IsNullable    string
	ColumnKey     string
	ColumnComment string
	Extra         string
}

const (
	IndexPrimary = "PRI"
	IndexUnique  = "UNI"
	IndexNormal  = "MUL"
)

var typeMap = map[string]string{
	"int":        "int64",
	"tinyint":    "int8",
	"smallint":   "int16",
	"mediumint":  "int32",
	"bigint":     "int64",
	"float":      "float32",
	"double":     "float64",
	"decimal":    "string", // 假设使用了第三方库处理十进制数据类型
	"char":       "string",
	"varchar":    "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"longtext":   "string",
	"text":       "string",
	"enum":       "MyEnumType", // 自定义枚举类型
	"set":        "[]string",   // 切片，表示集合
	"binary":     "[]byte",
	"varbinary":  "[]byte",
	"blob":       "[]byte",
	"date":       "time.Time",
	"time":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"year":       "int",
	"boolean":    "bool",
	"json":       "interface{}",   // 使用空接口，可以处理任意JSON数据
	"geometry":   "geos.Geometry", // 假设使用了第三方库处理几何类型
	"polygon":    "geos.Geometry", // 假设使用了第三方库处理几何类型
	"uuid":       "string",
	"xml":        "string",
}

var DefaultDb string

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
		helper.DefaultDb = conf.GetString(fmt.Sprintf("database.connections.%s.database", conn))
		return helper, nil
	}
	return helper, nil
}

func (s StructHelper) GetTableStruct(database, table string) []*Column {
	sqlPattern := "SELECT COLUMN_NAME as column_name, COLUMN_TYPE as column_type, IS_NULLABLE as is_nullable, COLUMN_KEY as column_key, COLUMN_COMMENT as column_comment, EXTRA as extra FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?;"
	var cols []*Column
	s.Db.Raw(sqlPattern, database, table).Scan(&cols)
	return cols
}

func (s StructHelper) StringColumns(cols []*Column) string {
	result := ""
	for _, col := range cols {
		result += support.TabSpace + col.String() + "\n"
	}
	return strings.Trim(result, " ")
}

func (c Column) String() string {
	pattern := "%s %s `json:\"%s\" gorm:\"%s\" xorm:\"%s\"` // %s"
	columnName := str.Case2Camel(c.ColumnName)

	columnType := c.parseType()
	columnJsonAlias := str.Camel2Case(c.ColumnName)
	tagGorm := c.parseGorm()
	tagXorm := c.parseXorm()
	comment := strings.ReplaceAll(c.ColumnComment, "\n", " ")
	return fmt.Sprintf(pattern, columnName, columnType, columnJsonAlias, tagGorm, tagXorm, comment)
}

func (c Column) parseType() string {
	reg := regexp.MustCompile(`\w+`)
	columnType := "string"
	realType := reg.FindString(c.ColumnType)
	if _, ok := typeMap[realType]; ok == true {
		columnType = typeMap[realType]
	}
	if strings.Contains(c.ColumnType, "nsigned") && strings.HasPrefix(columnType, "float") == false {
		columnType = "u" + columnType
	}
	return columnType
}

func (c Column) parseGorm() string {
	var gorm []string
	switch c.ColumnKey {
	case IndexPrimary:
		gorm = append(gorm, "primaryKey")
	case IndexUnique:
		gorm = append(gorm, "unique")
	case IndexNormal:
		gorm = append(gorm, "index")
	}

	if c.IsNullable == "NO" {
		gorm = append(gorm, "not null")
	}

	if strings.Contains(c.Extra, "auto_increment") {
		gorm = append(gorm, "autoIncrement")
	}
	return strings.Join(gorm, " ")
}

func (c Column) parseXorm() string {
	return c.ColumnType
}
