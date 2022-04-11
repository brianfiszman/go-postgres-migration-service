package containers

import "github.com/sirupsen/logrus"

type AppContainer struct {
	MigrationContainer MigrationContainer
	DatabaseContainer  DatabaseContainer
}

func NewAppContainer() (appContainer AppContainer) {
	appContainer.DatabaseContainer = *NewDatabaseContainer()
	logrus.Info("DatabaseContainer Initialized")

	appContainer.MigrationContainer = NewMigrationContainer(appContainer.DatabaseContainer.Database)
	logrus.Info("MigrationContainer Initialized")

	return
}
