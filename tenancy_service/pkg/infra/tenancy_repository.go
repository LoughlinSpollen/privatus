package infra

import (
	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	"github.com/google/uuid"
)

type TenancyRepositoryAdapter interface {
	Connect() error
	Close()
	Save(tenancy *model.Tenancy) error
	AddFederation(federation *model.Federation) error
	UpdateFederation(federation *model.Federation) error
	AddTraining(training *model.Training) error
	UpdateTraining(training *model.Training) error
	ReadFederation(federationID uuid.UUID) (*model.Federation, error)
	ReadTenancy(tenancyID uuid.UUID) (*model.Tenancy, error)
}
