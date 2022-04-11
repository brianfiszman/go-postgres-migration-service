package adapters_test

import (
	"testing"

	"github.com/brianfiszman/migration-service/pkg/infrastructure/adapters/mocks"
)

func Test(t *testing.T) {
	pgxconn := mocks.NewMockPostgreSQLAdapter()

	t.Run("Should ping the database and fail", func(t *testing.T) {
		err := pgxconn.Ping()

		if err == nil {
			t.Error("Shouldn't ping if its disconnected")
		}
	})

	t.Run("Should ping the database", func(t *testing.T) {
		// Connect to database
		pgxconn.ConnectDatabase()

		err := pgxconn.Ping()

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Should create database", func(t *testing.T) {
		// Connect to database
		pgxconn.ConnectDatabase()

		err := pgxconn.CreateDatabase("mockname")

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Should create schema", func(t *testing.T) {
		// Connect to database
		pgxconn.ConnectDatabase()

		err := pgxconn.CreateSchema()

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Should create table", func(t *testing.T) {
		// Connect to database
		pgxconn.ConnectDatabase()

		err := pgxconn.CreateTable("mockname", "mockstructure")

		if err != nil {
			t.Error(err)
		}
	})
}
