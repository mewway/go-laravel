package console

import (
	"fmt"
	"testing"

	"github.com/mewway/go-laravel/contracts/config"
	configmock "github.com/mewway/go-laravel/contracts/config/mocks"
	databasecontract "github.com/mewway/go-laravel/contracts/database"
	"github.com/stretchr/testify/assert"
	gormio "gorm.io/gorm"
)

func TestColumn_String(t *testing.T) {
	type fields struct {
		ColumnName    string
		ColumnType    string
		IsNullable    string
		ColumnKey     string
		ColumnComment string
		Extra         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Column{
				ColumnName:    tt.fields.ColumnName,
				ColumnType:    tt.fields.ColumnType,
				IsNullable:    tt.fields.IsNullable,
				ColumnKey:     tt.fields.ColumnKey,
				ColumnComment: tt.fields.ColumnComment,
				Extra:         tt.fields.Extra,
			}
			assert.Equalf(t, tt.want, c.String(), "String()")
		})
	}
}

func TestColumn_parseGorm(t *testing.T) {
	type fields struct {
		ColumnName    string
		ColumnType    string
		IsNullable    string
		ColumnKey     string
		ColumnComment string
		Extra         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "case1",
			fields: fields{
				ColumnName:    "status",
				ColumnType:    "tinyint(3) unsigned",
				IsNullable:    "NO",
				ColumnKey:     "",
				ColumnComment: "订单状态 0-未接单 1-已接单 2-超时作废",
				Extra:         "",
			},
			want: "not null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Column{
				ColumnName:    tt.fields.ColumnName,
				ColumnType:    tt.fields.ColumnType,
				IsNullable:    tt.fields.IsNullable,
				ColumnKey:     tt.fields.ColumnKey,
				ColumnComment: tt.fields.ColumnComment,
				Extra:         tt.fields.Extra,
			}
			assert.Equalf(t, tt.want, c.parseGorm(), "parseGorm()")
		})
	}
}

func TestColumn_parseType(t *testing.T) {
	type fields struct {
		ColumnName    string
		ColumnType    string
		IsNullable    string
		ColumnKey     string
		ColumnComment string
		Extra         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "case1",
			fields: fields{
				ColumnName:    "status",
				ColumnType:    "tinyint(3) unsigned",
				IsNullable:    "NO",
				ColumnKey:     "",
				ColumnComment: "订单状态 0-未接单 1-已接单 2-超时作废",
				Extra:         "",
			},
			want: "uint8",
		},
		{
			name: "case2",
			fields: fields{
				ColumnName:    "status",
				ColumnType:    "int(11)",
				IsNullable:    "NO",
				ColumnKey:     "",
				ColumnComment: "订单状态 0-未接单 1-已接单 2-超时作废",
				Extra:         "",
			},
			want: "int64",
		},
		{
			name: "case3",
			fields: fields{
				ColumnName:    "status",
				ColumnType:    "float unsigned",
				IsNullable:    "NO",
				ColumnKey:     "",
				ColumnComment: "订单状态 0-未接单 1-已接单 2-超时作废",
				Extra:         "",
			},
			want: "float32",
		},
		{
			name: "case4",
			fields: fields{
				ColumnName:    "status",
				ColumnType:    "double unsigned",
				IsNullable:    "NO",
				ColumnKey:     "",
				ColumnComment: "订单状态 0-未接单 1-已接单 2-超时作废",
				Extra:         "",
			},
			want: "float64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Column{
				ColumnName:    tt.fields.ColumnName,
				ColumnType:    tt.fields.ColumnType,
				IsNullable:    tt.fields.IsNullable,
				ColumnKey:     tt.fields.ColumnKey,
				ColumnComment: tt.fields.ColumnComment,
				Extra:         tt.fields.Extra,
			}
			assert.Equalf(t, tt.want, c.parseType(), "parseType()")
		})
	}
}

func TestColumn_parseXorm(t *testing.T) {
	type fields struct {
		ColumnName    string
		ColumnType    string
		IsNullable    string
		ColumnKey     string
		ColumnComment string
		Extra         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "case2",
			fields: fields{
				ColumnName:    "status",
				ColumnType:    "int(11)",
				IsNullable:    "NO",
				ColumnKey:     "",
				ColumnComment: "订单状态 0-未接单 1-已接单 2-超时作废",
				Extra:         "",
			},
			want: "int(11)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Column{
				ColumnName:    tt.fields.ColumnName,
				ColumnType:    tt.fields.ColumnType,
				IsNullable:    tt.fields.IsNullable,
				ColumnKey:     tt.fields.ColumnKey,
				ColumnComment: tt.fields.ColumnComment,
				Extra:         tt.fields.Extra,
			}
			assert.Equalf(t, tt.want, c.parseXorm(), "parseXorm()")
		})
	}
}

func TestNewStructHelper(t *testing.T) {
	type args struct {
		conf config.Config
	}

	mockConf := &configmock.Config{}
	mockConf.On("GetString", "database.default").Return("mysql")
	mockConf.On("GetString", "database.connections.mysql.host").Return("127.0.0.1")
	mockConf.On("GetString", "database.connections.mysql.username").Return("root")
	mockConf.On("GetString", "database.connections.mysql.password").Return("root1234")
	mockConf.On("GetString", "database.connections.mysql.database").Return("test")
	mockConf.On("GetString", "database.connections.mysql.charset").Return("utf8mb4")
	mockConf.On("GetString", "database.connections.mysql.loc").Return("general_ci")
	mockConf.On("GetInt", "database.connections.mysql.port").Return(3306)
	mockConf.On("GetString", "database.connections.mysql.driver").Return("mysql")
	var fakeDbConf []databasecontract.Config
	fakeDbConf = append(fakeDbConf, databasecontract.Config{})
	mockConf.On("Get", "database.connections.mysql.write").Return(fakeDbConf)
	mockConf.On("Get", "database.connections.mysql.read").Return(fakeDbConf)

	tests := []struct {
		name       string
		args       args
		wantHelper *StructHelper
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "case1",
			args: args{
				conf: mockConf,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				_, err = NewStructHelper(mockConf)
				assert.Error(t, err, "panic")
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHelper, err := NewStructHelper(tt.args.conf)
			if !tt.wantErr(t, err, fmt.Sprintf("NewStructHelper(%v)", tt.args.conf)) {
				return
			}
			assert.Equalf(t, tt.wantHelper, gotHelper, "NewStructHelper(%v)", tt.args.conf)
		})
	}
}

func TestStructHelper_getTableStruct(t *testing.T) {
	type fields struct {
		StructHelperInterface StructHelperInterface
		Db                    *gormio.DB
	}
	type args struct {
		database string
		table    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Column
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StructHelper{
				StructHelperInterface: tt.fields.StructHelperInterface,
				Db:                    tt.fields.Db,
			}
			assert.Equalf(t, tt.want, s.getTableStruct(tt.args.database, tt.args.table), "getTableStruct(%v, %v)", tt.args.database, tt.args.table)
		})
	}
}
