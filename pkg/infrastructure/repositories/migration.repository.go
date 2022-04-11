package repositories

import (
	"github.com/brianfiszman/migration-service/pkg/domain/interfaces"
	"github.com/sirupsen/logrus"
)

type MigrationRepository struct {
	Database     interfaces.SQLDatabaser
	databaseName string
}

func NewMigrationRepository(d interfaces.SQLDatabaser, dbName string) MigrationRepository {
	return MigrationRepository{
		Database:     d,
		databaseName: dbName,
	}
}

func (u MigrationRepository) CreateDatabase() {
	err := u.Database.CreateDatabase(u.databaseName)

	if err != nil {
		logrus.Fatal(err)
	}
}

func (u MigrationRepository) CreateSchema() {
	err := u.Database.CreateSchema()

	if err != nil {
		logrus.Fatal(err)
	}
}
