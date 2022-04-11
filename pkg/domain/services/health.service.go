package services

import "github.com/brianfiszman/migration-service/pkg/domain/interfaces"

type HealthService struct {
	Database interfaces.Databaser
}

func (h HealthService) GetHealthiness() error {
	return h.Database.Ping()
}
