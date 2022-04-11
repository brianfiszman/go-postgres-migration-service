package mocks

import (
	"github.com/brianfiszman/migration-service/pkg/domain/interfaces"
	"github.com/brianfiszman/migration-service/pkg/infrastructure/adapters/mocks"
	"github.com/brianfiszman/migration-service/pkg/infrastructure/config"
)

type MockDatabaseContainer struct {
	DatabaseConfig config.DatabaseConfig
	Database       interfaces.SQLDatabaser
}

func NewMockDatabaseContainer() *MockDatabaseContainer {
	var d *MockDatabaseContainer = &MockDatabaseContainer{
		DatabaseConfig: config.GetDatabaseConfig(),
		Database:       mocks.NewMockPostgreSQLAdapter(),
	}

	d.Database.ConnectDatabase()

	return d
}
