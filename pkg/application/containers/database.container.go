package containers

import (
	"github.com/brianfiszman/migration-service/pkg/domain/interfaces"
	"github.com/brianfiszman/migration-service/pkg/infrastructure/adapters"
	"github.com/brianfiszman/migration-service/pkg/infrastructure/config"
)

type DatabaseContainer struct {
	DatabaseConfig config.DatabaseConfig
	Database       interfaces.SQLDatabaser
}

func NewDatabaseContainer() *DatabaseContainer {
	var d *DatabaseContainer = &DatabaseContainer{
		DatabaseConfig: config.GetDatabaseConfig(),
		Database:       adapters.NewPostgreSQLAdapter(),
	}

	d.Database.ConnectDatabase()

	return d
}
