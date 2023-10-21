package model

import "github.com/google/uuid"

type Training struct {
	ID        uuid.UUID
	TenancyID uuid.UUID
	States    []string
}

func NewTraining(tenancyID uuid.UUID, states []string) *Training {
	return &Training{
		TenancyID: tenancyID,
		States:    states,
	}
}
