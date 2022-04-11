package interfaces

type Databaser interface {
	ConnectDatabase()
	Ping() error
	CreateDatabase(databaseName string) error
}
