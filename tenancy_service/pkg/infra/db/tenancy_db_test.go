package tenancy_db_test

import (
	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	infra "github.com/LoughlinSpollen/tenancy_service/pkg/infra"
	database "github.com/LoughlinSpollen/tenancy_service/pkg/infra/db"
	"github.com/google/uuid"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDatabaseService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tenancy db unit test suite")
}

var _ = Describe("Tenancy Database", func() {
	var (
		db                infra.TenancyRepositoryAdapter
		tenancy           *model.Tenancy
		federationUpdated *model.Federation
		tenancyID         uuid.UUID
		federationID      uuid.UUID
	)

	BeforeEach(func() {
		db = database.NewTenancyDB()
		err := db.Connect()
		Expect(err).ShouldNot(HaveOccurred())
	})
	AfterEach(func() {
		db.Close()
	})
	Describe("persist a new tenancy & federation aggregate", func() {
		Context("with tenancy federation aggregate", func() {
			BeforeEach(func() {
				tenancy = &model.Tenancy{
					MlModel: []byte("Python code here"),
					Federation: &model.Federation{
						Threshold: 1,
						Epochs:    1,
						Rate:      1,
						Rounds:    1,
						Batch:     1,
					},
				}
			})

			It("returns tenancy ID and federation ID, without error", func() {
				err := db.Save(tenancy)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tenancy.ID).ShouldNot(Equal(""))
				Expect(tenancy.MlModel).Should(Equal([]byte("Python code here")))
				Expect(tenancy.Federation.ID).ShouldNot(Equal(""))
				Expect(tenancy.ID).Should(Equal(tenancy.Federation.TenancyID))
				tenancyID = tenancy.ID
				federationID = tenancy.Federation.ID
			})
		})
	})
	Describe("persist a new federation for an existing tenancy", func() {
		Context("with tenancy domain model", func() {
			BeforeEach(func() {
				federationUpdated = &model.Federation{
					ID:        federationID,
					TenancyID: tenancyID,
					Threshold: 2,
					Epochs:    2,
					Rate:      2,
					Rounds:    2,
					Batch:     2,
				}
			})

			It("returns orginal IDs, without error ", func() {
				err := db.UpdateFederation(federationUpdated)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(federationUpdated.TenancyID).Should(Equal(tenancyID))
				Expect(federationUpdated.ID).Should(Equal(federationID))
			})

			It("reads updated tenancy-federation aggregate attributes, without error ", func() {
				tenancySaved, err := db.ReadTenancy(tenancyID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(tenancySaved.ID).Should(Equal(tenancyID))
				Expect(tenancySaved.Federation.ID).Should(Equal(federationID))
				Expect(tenancySaved.Federation.Threshold).Should(Equal(int32(2)))
				Expect(tenancySaved.Federation.Epochs).Should(Equal(int32(2)))
				Expect(tenancySaved.Federation.Rate).Should(Equal(int32(2)))
				Expect(tenancySaved.Federation.Rounds).Should(Equal(int32(2)))
				Expect(tenancySaved.Federation.Batch).Should(Equal(int32(2)))
			})

			It("reads updated federation attributes, without error ", func() {
				federationSaved, err := db.ReadFederation(federationUpdated.ID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(federationSaved.TenancyID).Should(Equal(tenancyID))
				Expect(federationSaved.ID).Should(Equal(federationID))
				Expect(federationSaved.Threshold).Should(Equal(int32(2)))
				Expect(federationSaved.Epochs).Should(Equal(int32(2)))
				Expect(federationSaved.Rate).Should(Equal(int32(2)))
				Expect(federationSaved.Rounds).Should(Equal(int32(2)))
				Expect(federationSaved.Batch).Should(Equal(int32(2)))
			})
		})
	})
})
