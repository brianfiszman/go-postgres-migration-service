package services

import "github.com/brianfiszman/migration-service/pkg/domain/interfaces"

type MigrationService struct {
	MigrationRepository interfaces.MigrationRepositorier
	UserRepository      interfaces.Repositorier
}

func NewMigrationService(m interfaces.MigrationRepositorier, u interfaces.Repositorier) MigrationService {
	return MigrationService{
		MigrationRepository: m,
		UserRepository:      u,
	}
}

func (m MigrationService) GenerateMigration() {
	m.MigrationRepository.CreateDatabase()
	m.MigrationRepository.CreateSchema()
	m.UserRepository.CreateEntityTable()
}
