package model

import "github.com/google/uuid"

type Tenancy struct {
	ID         uuid.UUID
	MlModel    []byte
	Federation *Federation
}

func NewTenancy(mlModel []byte) *Tenancy {
	return &Tenancy{
		MlModel: mlModel,
	}
}
