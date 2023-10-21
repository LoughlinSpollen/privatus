package infra

import "github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"

type MLServiceAdapter interface {
	Connect() error
	Close()
	Training(training *model.Training) error
}
