package mocks

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type MockPostgreSQLAdapter struct {
	isConnected bool
}

func NewMockPostgreSQLAdapter() (adapter *MockPostgreSQLAdapter) {
	adapter = &MockPostgreSQLAdapter{}

	return
}

func (d *MockPostgreSQLAdapter) ConnectDatabase() {
	d.isConnected = true

	logrus.Info("Connected to localhost:5432")
}

func (d *MockPostgreSQLAdapter) Ping() error {
	if !d.isConnected {
		return errors.New("Not connected to the database")
	}

	return nil
}

func (d *MockPostgreSQLAdapter) CreateDatabase(databaseName string) error {
	if !d.isConnected {
		return errors.New("Not connected to the database")
	}

	return nil
}

func (d *MockPostgreSQLAdapter) CreateSchema() error {
	if !d.isConnected {
		return errors.New("Not connected to the database")
	}

	return nil
}

func (d *MockPostgreSQLAdapter) CreateTable(tableName string, tableStructure string) error {
	if !d.isConnected {
		return errors.New("Not connected to the database")
	}

	return nil
}

func (d *MockPostgreSQLAdapter) GetDatabaseName() string {
	return "mockName"
}

func (d *MockPostgreSQLAdapter) GetSchemaName() string {
	return "mockName"
}
