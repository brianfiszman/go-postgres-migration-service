package repositories

import (
	"fmt"

	"github.com/brianfiszman/migration-service/pkg/domain/interfaces"
	"github.com/sirupsen/logrus"
)

const ENTITY_NAME = "user"
const TABLE_STRUCTURE = `
  id serial PRIMARY KEY,
  name VARCHAR (30) UNIQUE NOT NULL
`

type UserRepository struct {
	Database       interfaces.SQLDatabaser
	EntityName     string
	TableStructure string
}

func NewUserRepository(d interfaces.SQLDatabaser) UserRepository {
	return UserRepository{
		Database:       d,
		EntityName:     fmt.Sprintf("%s.%s", d.GetSchemaName(), ENTITY_NAME),
		TableStructure: TABLE_STRUCTURE,
	}
}

func (u UserRepository) CreateEntityTable() {
	err := u.Database.CreateTable(u.EntityName, u.TableStructure)

	if err != nil {
		logrus.Fatal(err)
	}
}
