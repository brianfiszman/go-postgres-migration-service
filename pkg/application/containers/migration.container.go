package containers

import (
	"github.com/brianfiszman/migration-service/pkg/domain/interfaces"
	"github.com/brianfiszman/migration-service/pkg/domain/services"
	"github.com/brianfiszman/migration-service/pkg/infrastructure/repositories"
)

type MigrationContainer struct {
	MigrationService    interfaces.MigrationServicer
	MigrationRepository interfaces.MigrationRepositorier
	UserRepository      interfaces.Repositorier
}

func NewMigrationContainer(Database interfaces.SQLDatabaser) (m MigrationContainer) {
	m = MigrationContainer{
		MigrationRepository: repositories.NewMigrationRepository(Database, Database.GetDatabaseName()),
		UserRepository:      repositories.NewUserRepository(Database),
	}

	m.MigrationService = services.NewMigrationService(m.MigrationRepository, m.UserRepository)

	return
}
