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

func TestTenancyDTOAdapter(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("TenancyDTO Adapter", func() {
	var (
		tenancyDTOAdapter rest.TenancyDTOAdapter
	)
	BeforeEach(func() {
		tenancyDTOAdapter = rest.NewTenancyDTOAdapter()
	})
	Describe("Tenancy Domain model DTO serialization", func() {

		When("converting from an tenancy domain model to serialised tenancy DTO, ", func() {
			tenancyModel := model.NewTenancy([]byte("Python code here"))
			tenancyModel.ID = uuid.New()

			It("should contain only tenancy-id and ml-model attributes with lower (snake) case.", func() {
				tenancyBytes, err := tenancyDTOAdapter.TenancyToBytes(tenancyModel)
				Expect(err).ShouldNot(HaveOccurred())

				b64, err := base64.StdEncoding.DecodeString(string(tenancyBytes))
				Expect(err).ShouldNot(HaveOccurred())

				var tenancyDTO map[string]interface{}
				err = json.Unmarshal(b64, &tenancyDTO)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(tenancyDTO["id"]).Should(Equal(tenancyModel.ID.String()))
				Expect(tenancyDTO["ml-model"]).Should(Equal("Python code here"))
			})
		})

		When("converting a serialised tenancy DTO to an tenancy domain model, ", func() {
			tenancyDTO := map[string]interface{}{
				"id":       "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				"ml-model": "Python code here",
			}
			body, err := json.Marshal(tenancyDTO)
			Expect(err).ShouldNot(HaveOccurred())
			b64 := base64.StdEncoding.EncodeToString(body)
			tenancyDTOBytes := []byte(b64)

			It("should contain all domain model fields from the DTO attributes.", func() {
				tenancyModel, err := tenancyDTOAdapter.BytesToTenancy(tenancyDTOBytes)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tenancyModel.ID).Should(Equal(uuid.MustParse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")))
				Expect(tenancyModel.MlModel).Should(Equal([]byte("Python code here")))
			})
		})
	})
})
