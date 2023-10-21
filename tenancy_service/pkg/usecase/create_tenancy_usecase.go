package usecase

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	infra "github.com/LoughlinSpollen/tenancy_service/pkg/infra"
)

type TenancyUsecase interface {
	CreateTenancy(tenancy *model.Tenancy) error
	CreateFederation(federation *model.Federation) error
}

type tenancyUsecase struct {
	trainingRepo infra.TenancyRepositoryAdapter
	mpcService   infra.MPCServiceAdapter
	mlService    infra.MLServiceAdapter
}

func NewTenancyUsecase(trainingRepo infra.TenancyRepositoryAdapter,
	mpcService infra.MPCServiceAdapter,
	mlService infra.MLServiceAdapter) *tenancyUsecase {
	return &tenancyUsecase{
		trainingRepo: trainingRepo,
		mpcService:   mpcService,
		mlService:    mlService,
	}
}

func (u *tenancyUsecase) CreateTenancy(tenancy *model.Tenancy) error {
	log.Debug("tenancyUsecase CreateTenancy")

	if err := u.trainingRepo.Save(tenancy); err != nil {
		return err
	}
	return nil
}

func (u *tenancyUsecase) CreateFederation(federation *model.Federation) error {
	log.Debug("tenancyUsecase createFederation")

	if federation.ID == uuid.Nil {
		if err := u.trainingRepo.AddFederation(federation); err != nil {
			return err
		}
	} else {
		if err := u.trainingRepo.UpdateFederation(federation); err != nil {
			return err
		}
	}
	return nil
}
