package interfaces

type SQLDatabaser interface {
	Databaser
	CreateTable(tableName string, tableStructure string) error
	CreateSchema() error
	GetDatabaseName() string
	GetSchemaName() string
}
