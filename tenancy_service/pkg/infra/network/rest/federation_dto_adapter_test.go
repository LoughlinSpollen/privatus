package rest_test

import (
	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/network/rest"
	"github.com/google/uuid"

	"encoding/base64"
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFederationDTOAdapter(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("FederationDTO Adapter", func() {
	var federationDTOAdapter rest.FederationDTOAdapter

	BeforeEach(func() {
		federationDTOAdapter = rest.NewFederationDTOAdapter()
	})
	Describe("Federation Domain model DTO serialization", func() {

		When("converting from an federation domain model to serialised federation DTO, ", func() {
			expected := int32(1)
			federationModel := model.NewFederation()
			federationModel.ID = uuid.New()
			federationModel.TenancyID = uuid.New()
			federationModel.Threshold = expected
			federationModel.Epochs = expected
			federationModel.Rate = expected
			federationModel.Rounds = expected
			federationModel.Batch = expected

			It("should contain only federation-id and ml-model attributes with lower (snake) case.", func() {
				federationBytes, err := federationDTOAdapter.FederationToBytes(federationModel)
				Expect(err).ShouldNot(HaveOccurred())

				b64, err := base64.StdEncoding.DecodeString(string(federationBytes))
				Expect(err).ShouldNot(HaveOccurred())

				var federationDTO map[string]interface{}
				err = json.Unmarshal(b64, &federationDTO)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(federationDTO["id"].(string)).Should(Equal(federationModel.ID.String()))
				Expect(federationDTO["tenancy-id"].(string)).Should(Equal(federationModel.TenancyID.String()))
				Expect(int32(federationDTO["threshold"].(float64))).Should(Equal(expected))
				Expect(int32(federationDTO["epochs"].(float64))).Should(Equal(expected))
				Expect(int32(federationDTO["rate"].(float64))).Should(Equal(expected))
				Expect(int32(federationDTO["rounds"].(float64))).Should(Equal(expected))
				Expect(int32(federationDTO["batch"].(float64))).Should(Equal(expected))
			})
		})
		When("converting a serialised federation DTO to an federation domain model, ", func() {
			body := []byte(`{
				"id": "cd27e265-9605-4b4b-a0e5-3003ea9cc4da",
				"tenancy-id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				"threshold": 1,
				"epochs": 1,
				"rate": 1,
				"rounds": 1,
				"batch": 1
			}`)

			b64 := base64.StdEncoding.EncodeToString(body)
			federationDTOBytes := []byte(b64)

			It("should contain all domain model fields from the DTO attributes.", func() {
				federationModel, err := federationDTOAdapter.BytesToFederation(federationDTOBytes)
				expected := int32(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(federationModel.ID).Should(Equal(uuid.MustParse("cd27e265-9605-4b4b-a0e5-3003ea9cc4da")))
				Expect(federationModel.TenancyID).Should(Equal(uuid.MustParse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")))
				Expect(federationModel.Threshold).Should(Equal(expected))
				Expect(federationModel.Epochs).Should(Equal(expected))
				Expect(federationModel.Rate).Should(Equal(expected))
				Expect(federationModel.Rounds).Should(Equal(expected))
				Expect(federationModel.Batch).Should(Equal(expected))
			})
		})
	})
})
