package adapters

import (
	"context"
	"fmt"

	"github.com/brianfiszman/migration-service/pkg/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type PostgreSQLAdapter struct {
	config.DatabaseConfig
	ConnectionPool *pgxpool.Pool
}

func NewPostgreSQLAdapter() (adapter *PostgreSQLAdapter) {
	adapter = &PostgreSQLAdapter{
		DatabaseConfig: config.GetDatabaseConfig(),
	}

	return
}

func (d *PostgreSQLAdapter) ConnectDatabase() {
	dbpool, err := pgxpool.Connect(context.Background(), d.GetConnectionString())

	d.ConnectionPool = dbpool

	if err != nil {
		logrus.Fatal("Unable to connect to database: ", err)
	}

	logrus.Info("Connected to ", d.ConnectionPool.Config().ConnString())
}

func (d *PostgreSQLAdapter) Ping() error {
	return d.ConnectionPool.Ping(context.Background())
}

func (d *PostgreSQLAdapter) CreateDatabase(databaseName string) error {
	var query string = fmt.Sprintf("CREATE DATABASE %s;", databaseName)

	_, err := d.ConnectionPool.Exec(context.Background(), query)

	return err
}

func (d *PostgreSQLAdapter) CreateSchema() error {
	var query string = fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", d.GetSchemaName())

	_, err := d.ConnectionPool.Exec(context.Background(), query)

	return err
}

func (d *PostgreSQLAdapter) CreateTable(tableName string, tableStructure string) error {
	var query string = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(%s);", tableName, tableStructure)

	_, err := d.ConnectionPool.Exec(context.Background(), query)

	return err
}

func (d *PostgreSQLAdapter) GetDatabaseName() string {
	return d.Name
}

func (d *PostgreSQLAdapter) GetSchemaName() string {
	return d.Schema
}
