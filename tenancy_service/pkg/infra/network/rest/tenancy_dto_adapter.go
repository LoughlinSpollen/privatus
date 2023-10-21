package rest

import (
	"fmt"

	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	"github.com/google/uuid"

	"encoding/base64"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type tenancyDTO struct {
	ID      string `json:"id,omitempty"`
	MlModel string `json:"ml-model,omitempty"`
}

type tenancyDTOAdapter struct {
}

type TenancyDTOAdapter interface {
	TenancyToBytes(tenancyModel *model.Tenancy) ([]byte, error)
	BytesToTenancy(tenancyBytes []byte) (*model.Tenancy, error)
}

func NewTenancyDTOAdapter() *tenancyDTOAdapter {
	return &tenancyDTOAdapter{}
}

func (a *tenancyDTOAdapter) TenancyToBytes(tenancyModel *model.Tenancy) ([]byte, error) {
	log.Debug("tenancyDTOAdapter TenancyToBytes")

	tenancyDTO := a.toDTO(tenancyModel)
	jsonStr, err := json.Marshal(tenancyDTO)
	if err != nil {
		return nil, err
	}
	b64 := base64.StdEncoding.EncodeToString(jsonStr)
	tenancyDTOBytes := []byte(b64)
	return tenancyDTOBytes, nil
}

func (a *tenancyDTOAdapter) BytesToTenancy(tenancyBytes []byte) (*model.Tenancy, error) {
	log.Debug("tenancyDTOAdapter BytesToTenancy")

	var tenancyModel *model.Tenancy
	var tenancyDTO tenancyDTO
	b64, err := base64.StdEncoding.DecodeString(string(tenancyBytes))
	if err != nil {
		log.Warn(fmt.Printf("tenancyDTOAdapter fromDTO failed to decode tenancy DTO string (%s): %v", string(tenancyBytes), err))
		return nil, err
	}
	err = json.Unmarshal(b64, &tenancyDTO)
	if err != nil {
		log.Warn(fmt.Printf("tenancyDTOAdapter fromDTO failed to unmarshall tenancy json (%s): %v", b64, err))
		return nil, err
	}
	tenancyModel, err = a.fromDTO(&tenancyDTO)
	if err != nil {
		return nil, err
	}
	return tenancyModel, err
}

func (a *tenancyDTOAdapter) fromDTO(tenancyDTO *tenancyDTO) (*model.Tenancy, error) {
	log.Debug("tenancyDTOAdapter fromDTO")

	modelTenancy := model.NewTenancy([]byte(tenancyDTO.MlModel))
	if tenancyDTO.ID != "" {
		var err error
		modelTenancy.ID, err = uuid.Parse(tenancyDTO.ID)
		if err != nil {
			log.Warn(fmt.Printf("tenancyDTOAdapter fromDTO failed to parse tenancy uuid: %v", err))
			return nil, err
		}
	}

	return modelTenancy, nil
}

func (a *tenancyDTOAdapter) toDTO(tenancyModel *model.Tenancy) *tenancyDTO {
	log.Debug("tenancyDTOAdapter toDTO")

	dtoTenancy := tenancyDTO{
		ID:      tenancyModel.ID.String(),
		MlModel: string(tenancyModel.MlModel),
	}

	return &dtoTenancy
}
