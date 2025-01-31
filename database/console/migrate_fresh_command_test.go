package console

import (
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"

	configmock "github.com/mewway/go-laravel/contracts/config/mocks"
	consolemocks "github.com/mewway/go-laravel/contracts/console/mocks"
	ormcontract "github.com/mewway/go-laravel/contracts/database/orm"
	"github.com/mewway/go-laravel/database/gorm"
)

func TestMigrateFreshCommand(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping tests of using docker")
	}

	var (
		mockConfig *configmock.Config
		pool       *dockertest.Pool
		resource   *dockertest.Resource
		query      ormcontract.Query
	)

	beforeEach := func() {
		pool = nil
		mockConfig = &configmock.Config{}
	}

	tests := []struct {
		name  string
		setup func()
	}{
		{
			name: "mysql",
			setup: func() {
				var err error
				docker := gorm.NewMysqlDocker()
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createMysqlMigrations()

			},
		},
		{
			name: "postgresql",
			setup: func() {
				var err error
				docker := gorm.NewPostgresqlDocker()
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createPostgresqlMigrations()
			},
		},
		{
			name: "sqlserver",
			setup: func() {
				var err error
				docker := gorm.NewSqlserverDocker()
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createSqlserverMigrations()
			},
		},
		{
			name: "sqlite",
			setup: func() {
				var err error
				docker := gorm.NewSqliteDocker("goravel")
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createSqliteMigrations()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()

			mockContext := &consolemocks.Context{}
			mockArtisan := &consolemocks.Artisan{}
			migrateCommand := NewMigrateCommand(mockConfig)
			assert.Nil(t, migrateCommand.Handle(mockContext))
			mockContext.On("OptionBool", "seed").Return(false).Once()
			migrateFreshCommand := NewMigrateFreshCommand(mockConfig, mockArtisan)
			assert.Nil(t, migrateFreshCommand.Handle(mockContext))

			// Test MigrateFreshCommand with --seed flag and seeders specified
			mockContext = &consolemocks.Context{}
			mockArtisan = &consolemocks.Artisan{}
			mockContext.On("OptionBool", "seed").Return(true).Once()
			mockContext.On("OptionSlice", "seeder").Return([]string{"MockSeeder"}).Once()
			mockArtisan.On("Call", "db:seed --seeder MockSeeder").Return(nil).Once()
			migrateFreshCommand = NewMigrateFreshCommand(mockConfig, mockArtisan)
			assert.Nil(t, migrateFreshCommand.Handle(mockContext))

			// Test MigrateFreshCommand with --seed flag and no seeders specified
			mockContext = &consolemocks.Context{}
			mockArtisan = &consolemocks.Artisan{}
			mockContext.On("OptionBool", "seed").Return(true).Once()
			mockContext.On("OptionSlice", "seeder").Return([]string{}).Once()
			mockArtisan.On("Call", "db:seed").Return(nil).Once()
			migrateFreshCommand = NewMigrateFreshCommand(mockConfig, mockArtisan)
			assert.Nil(t, migrateFreshCommand.Handle(mockContext))

			var agent Agent
			err := query.Where("name", "goravel").First(&agent)
			assert.Nil(t, err)
			assert.True(t, agent.ID > 0)

			if pool != nil && test.name != "sqlite" {
				assert.Nil(t, pool.Purge(resource))
			}

			removeMigrations()
		})
	}
}
