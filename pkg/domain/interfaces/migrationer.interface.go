package interfaces

type MigrationRepositorier interface {
	CreateDatabase()
	CreateSchema()
}
