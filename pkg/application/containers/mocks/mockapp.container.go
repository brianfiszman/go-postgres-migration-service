package mocks

import (
	"github.com/brianfiszman/migration-service/pkg/application/containers"
	"github.com/sirupsen/logrus"
)

type MockAppContainer struct {
	MockDatabaseContainer MockDatabaseContainer
	MigrationContainer    containers.MigrationContainer
}

func NewMockAppContainer() (appContainer MockAppContainer) {
	appContainer.MockDatabaseContainer = *NewMockDatabaseContainer()
	logrus.Info("MockDatabaseContainer Initialized")

	appContainer.MigrationContainer = containers.NewMigrationContainer(appContainer.MockDatabaseContainer.Database)
	logrus.Info("MigrationContainer Initialized")

	return
}
