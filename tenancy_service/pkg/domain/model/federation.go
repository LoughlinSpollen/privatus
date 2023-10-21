package model

import "github.com/google/uuid"

type Federation struct {
	ID        uuid.UUID
	TenancyID uuid.UUID
	Threshold int32
	Epochs    int32
	Rate      int32
	Rounds    int32
	Batch     int32
}

func NewFederation() *Federation {
	return &Federation{}
}
