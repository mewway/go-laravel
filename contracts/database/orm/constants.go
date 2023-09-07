package orm

const (
	DriverMysql      Driver = "mysql"
	DriverPostgresql Driver = "postgresql"
	DriverSqlite     Driver = "sqlite"
	DriverSqlserver  Driver = "sqlserver"
	DriverHttp       Driver = "http"
	DriverGrpc       Driver = "grpc"
)

type Driver string

func (d Driver) String() string {
	return string(d)
}
