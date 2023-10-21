package integration_test_stub

import (
	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	"github.com/google/uuid"
)

type tenancyRepositoryStub struct {
	federation *model.Federation
	tenancy    *model.Tenancy
	training   *model.Training
}

func NewRepositoryStub() *tenancyRepositoryStub {
	return &tenancyRepositoryStub{}
}

func (stub *tenancyRepositoryStub) Connect() error {
	return nil
}

func (stub *tenancyRepositoryStub) Close() {
}

func (stub *tenancyRepositoryStub) Save(tenancy *model.Tenancy) error {
	tenancy.ID = uuid.MustParse("cd27e265-9605-4b4b-a0e5-3003ea9cc4da")
	if tenancy.Federation != nil {
		tenancy.Federation.ID = uuid.MustParse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
		tenancy.Federation.TenancyID = tenancy.ID
		stub.federation = tenancy.Federation
	}
	stub.tenancy = tenancy
	return nil
}

func (stub *tenancyRepositoryStub) AddFederation(federation *model.Federation) error {
	federation.ID = uuid.MustParse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	stub.federation = federation
	return nil
}

func (stub *tenancyRepositoryStub) UpdateFederation(federation *model.Federation) error {
	stub.federation = federation
	return nil
}

func (stub *tenancyRepositoryStub) AddTraining(training *model.Training) error {
	training.ID = uuid.New()
	stub.training = training
	return nil
}

func (stub *tenancyRepositoryStub) UpdateTraining(training *model.Training) error {
	stub.training = training
	return nil
}

func (stub *tenancyRepositoryStub) ReadFederation(federationID uuid.UUID) (*model.Federation, error) {
	return stub.federation, nil
}

func (stub *tenancyRepositoryStub) ReadTenancy(tenancyID uuid.UUID) (*model.Tenancy, error) {
	return stub.tenancy, nil
}
