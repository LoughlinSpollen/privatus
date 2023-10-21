package rest

import (
	"fmt"

	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"encoding/base64"
	"encoding/json"
)

type federationDTO struct {
	ID        string `json:"id,omitempty"`
	TenancyID string `json:"tenancy-id,omitempty"`
	Threshold int32  `json:"threshold,omitempty"`
	Epochs    int32  `json:"epochs,omitempty"`
	Rate      int32  `json:"rate,omitempty"`
	Rounds    int32  `json:"rounds,omitempty"`
	Batch     int32  `json:"batch,omitempty"`
}

type federationDTOAdapter struct {
}

type FederationDTOAdapter interface {
	FederationToBytes(federationModel *model.Federation) ([]byte, error)
	BytesToFederation(federationBytes []byte) (*model.Federation, error)
}

func NewFederationDTOAdapter() *federationDTOAdapter {
	return &federationDTOAdapter{}
}

func (a *federationDTOAdapter) FederationToBytes(federationModel *model.Federation) ([]byte, error) {
	log.Debug("federationDTOAdapter FederationToBytes")

	federationDTO := a.toDTO(federationModel)
	jsonStr, err := json.Marshal(federationDTO)
	if err != nil {
		return nil, err
	}
	b64 := base64.StdEncoding.EncodeToString(jsonStr)
	federationDTOBytes := []byte(b64)
	return federationDTOBytes, nil
}

func (a *federationDTOAdapter) BytesToFederation(federationBytes []byte) (*model.Federation, error) {
	log.Debug("federationDTOAdapter BytesToFederation")

	var federationModel *model.Federation
	var federationDTO federationDTO
	b64, err := base64.StdEncoding.DecodeString(string(federationBytes))
	if err != nil {
		log.Warn(fmt.Printf("federationDTOAdapter fromDTO failed to decode federation DTO string (%s): %v", string(federationBytes), err))
		return nil, err
	}
	err = json.Unmarshal(b64, &federationDTO)
	if err != nil {
		log.Warn(fmt.Printf("federationDTOAdapter fromDTO failed to unmarshall federation json (%s): %v", b64, err))
		return nil, err
	}
	federationModel, err = a.fromDTO(&federationDTO)
	if err != nil {
		return nil, err
	}
	return federationModel, err
}

func (a *federationDTOAdapter) fromDTO(federationDTO *federationDTO) (*model.Federation, error) {
	log.Debug("federationDTOAdapter fromDTO")

	federation := model.NewFederation()
	var err error
	federation.ID, err = a.toUUID(federationDTO.ID)
	if err != nil {
		return nil, err
	}
	federation.TenancyID, err = a.toUUID(federationDTO.TenancyID)
	if err != nil {
		return nil, err
	}

	federation.Threshold = federationDTO.Threshold
	federation.Epochs = federationDTO.Epochs
	federation.Rate = federationDTO.Rate
	federation.Rounds = federationDTO.Rounds
	federation.Batch = federationDTO.Batch

	return federation, nil
}

func (a *federationDTOAdapter) toUUID(id string) (uuid.UUID, error) {
	log.Debug("federationDTOAdapter toUUID")

	if id != "" {
		uuID, err := uuid.Parse(id)
		if err != nil {
			log.Warn(fmt.Printf("federationDTOAdapter fromDTO failed to parse uuid: %v", err))
			return uuid.Nil, err
		}
		return uuID, nil
	}
	return uuid.Nil, nil
}

func (a *federationDTOAdapter) toDTO(federationModel *model.Federation) *federationDTO {
	log.Debug("federationDTOAdapter toDTO")

	dtoFederation := federationDTO{
		ID:        federationModel.ID.String(),
		TenancyID: federationModel.TenancyID.String(),
		Threshold: federationModel.Threshold,
		Epochs:    federationModel.Epochs,
		Rate:      federationModel.Rate,
		Rounds:    federationModel.Rounds,
		Batch:     federationModel.Batch,
	}

	return &dtoFederation
}
