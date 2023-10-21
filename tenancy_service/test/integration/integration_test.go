package integration_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	infra "github.com/LoughlinSpollen/tenancy_service/pkg/infra"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/network/rest"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/network/rpc"
	"github.com/LoughlinSpollen/tenancy_service/pkg/usecase"
	stub "github.com/LoughlinSpollen/tenancy_service/test/integration/stub"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tenancy server integration test suite")
}

var _ = Describe("Tenancy server entry points", func() {
	var (
		db                 infra.TenancyRepositoryAdapter
		mpcService         infra.MPCServiceAdapter
		mlService          infra.MLServiceAdapter
		tenancyRestService infra.TenancyRestAdapter
		request            *http.Request
	)

	BeforeEach(func() {
		db = stub.NewRepositoryStub()
		err := db.Connect()
		Expect(err).ShouldNot(HaveOccurred())

		mpcService = rpc.NewMPCService()
		mlService = rpc.NewMLService()

		tenancyUsecase := usecase.NewTenancyUsecase(db, mpcService, mlService)
		tenancyRestService = rest.NewTenancyRestService(tenancyUsecase)
		go func() {
			tenancyRestService.Connect()
		}()
	})

	AfterEach(func() {
		db.Close()
		mpcService.Close()
		mlService.Close()
		tenancyRestService.Close()
	})

	Describe("privatus tenancy end-point", func() {
		Context("create a new ML tenancy", func() {
			BeforeEach(func() {
				createRequestPayload := map[string]interface{}{
					"ml-model": "Python code here",
				}
				bodyStr, err := json.Marshal(createRequestPayload)
				Expect(err).ShouldNot(HaveOccurred())
				b64 := base64.StdEncoding.EncodeToString(bodyStr)
				body := []byte(b64)

				request, _ = http.NewRequest("POST", "http://localhost:8080/v1/tenancy", bytes.NewBuffer(body))
				request.Header.Set("Content-Type", "application/vnd.api+json")
				request.Header.Set("Accept", "application/vnd.api+json")
				request.Header.Set("Content-Length", fmt.Sprint(len(body)))
			})
			It("successfully created, returning a 201-StatusCode", func() {
				response, err := http.DefaultClient.Do(request)
				Expect(err).ShouldNot(HaveOccurred())
				resBody, err := io.ReadAll(response.Body)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(response.StatusCode).To(Equal(201))

				b64, err := base64.StdEncoding.DecodeString(string(resBody))
				Expect(err).ShouldNot(HaveOccurred())

				var tenancyDTO map[string]interface{}
				err = json.Unmarshal(b64, &tenancyDTO)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(tenancyDTO["id"]).Should(Equal("cd27e265-9605-4b4b-a0e5-3003ea9cc4da"))
				Expect(tenancyDTO["ml-model"]).Should(Equal("Python code here"))
			})
		})
	})
	Describe("privatus federation end-points", func() {
		Context("add federation settings", func() {
			BeforeEach(func() {
				createRequestPayload := map[string]interface{}{
					"threshold": 1,
					"epochs":    1,
					"rate":      1,
					"rounds":    1,
					"batch":     1,
				}
				bodyStr, err := json.Marshal(createRequestPayload)
				Expect(err).ShouldNot(HaveOccurred())
				b64 := base64.StdEncoding.EncodeToString(bodyStr)
				body := []byte(b64)

				request, _ = http.NewRequest("POST", "http://localhost:8080/v1/tenancy/cd27e265-9605-4b4b-a0e5-3003ea9cc4da/federation", bytes.NewReader(body))
				request.Header.Set("Content-Type", "application/vnd.api+json")
				request.Header.Set("Accept", "application/vnd.api+json")
				request.Header.Set("Content-Length", fmt.Sprint(len(body)))
			})

			It("returns created, 201-StatusCodeCreated, and an the federation payload plus UUIDs", func() {
				response, err := http.DefaultClient.Do(request)
				Expect(err).ShouldNot(HaveOccurred())
				resBody, err := io.ReadAll(response.Body)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(response.StatusCode).To(Equal(201))

				b64, err := base64.StdEncoding.DecodeString(string(resBody))
				Expect(err).ShouldNot(HaveOccurred())

				var federationDTO map[string]interface{}
				err = json.Unmarshal(b64, &federationDTO)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(federationDTO["id"]).Should(Equal("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"))
				Expect(federationDTO["tenancy-id"]).Should(Equal("cd27e265-9605-4b4b-a0e5-3003ea9cc4da"))
				Expect(int32(federationDTO["threshold"].(float64))).Should(Equal(int32(1)))
				Expect(int32(federationDTO["epochs"].(float64))).Should(Equal(int32(1)))
				Expect(int32(federationDTO["rate"].(float64))).Should(Equal(int32(1)))
				Expect(int32(federationDTO["rounds"].(float64))).Should(Equal(int32(1)))
				Expect(int32(federationDTO["batch"].(float64))).Should(Equal(int32(1)))
			})
		})
	})
})
