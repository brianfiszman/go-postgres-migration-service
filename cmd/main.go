package main

import "github.com/brianfiszman/migration-service/pkg/application/containers"

func main() {
	var app containers.AppContainer = containers.NewAppContainer()

	// Server initializes listening
	app.MigrationContainer.MigrationService.GenerateMigration()
}
