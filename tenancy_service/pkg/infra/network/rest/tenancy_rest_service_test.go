package rest_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	infra "github.com/LoughlinSpollen/tenancy_service/pkg/infra"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/network/rest"
	"github.com/google/uuid"

	"encoding/base64"
	"encoding/json"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tenancy rest unit test suite")
}

type tenancyUsecaseStub struct{}

func (u *tenancyUsecaseStub) CreateTenancy(tenancy *model.Tenancy) error {
	tenancy.ID = uuid.MustParse("cd27e265-9605-4b4b-a0e5-3003ea9cc4da")
	tenancy.MlModel = []byte("Python code here")
	return nil
}

func (u *tenancyUsecaseStub) CreateFederation(federation *model.Federation) error {
	federation.ID = uuid.MustParse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	federation.TenancyID = uuid.MustParse("cd27e265-9605-4b4b-a0e5-3003ea9cc4da")
	federation.Threshold = 1
	federation.Epochs = 1
	federation.Rate = 1
	federation.Rounds = 1
	federation.Batch = 1
	return nil
}

type tenancyUsecaseFailParamStub struct{}

func (u *tenancyUsecaseFailParamStub) CreateTenancy(*model.Tenancy) error {
	return nil
}

func (u *tenancyUsecaseFailParamStub) CreateFederation(*model.Federation) error {
	return errors.New("not found")
}

var _ = Describe("Tenancy Rest API", func() {
	var (
		server  infra.TenancyRestAdapter
		request *http.Request
	)

	Describe("Success", func() {
		BeforeEach(func() {
			usecaseStub := tenancyUsecaseStub{}
			server = rest.NewTenancyRestService(&usecaseStub)
			go func() {
				server.Connect()
			}()
		})
		AfterEach(func() {
			go func() {
				server.Close()
			}()
		})

		Describe("POST /v1/tenancy - create tenancy for federated learning", func() {
			Context("with valid JSON", func() {
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

				It("returns created, 201-StatusCodeCreated, and an UUID", func() {
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

		Describe("POST /v1/tenancy/{id}/federation - successfully add training federation", func() {
			Context("with valid JSON", func() {
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

				It("returns created, 201-StatusCodeCreated, and an the federation payload plus UUIDs ", func() {
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

	Describe("Failure bad params", func() {
		BeforeEach(func() {
			usecaseFailStub := tenancyUsecaseFailParamStub{}
			server = rest.NewTenancyRestService(&usecaseFailStub)
			go func() {
				server.Connect()
			}()
		})
		AfterEach(func() {
			server.Close()
		})

		Describe("POST /v1/tenancy/null/federation - missing params when add training federation", func() {
			Context("with valid JSON", func() {
				BeforeEach(func() {
					createRequestPayload := map[string]interface{}{}
					bodyStr, err := json.Marshal(createRequestPayload)
					Expect(err).ShouldNot(HaveOccurred())
					b64 := base64.StdEncoding.EncodeToString(bodyStr)
					body := []byte(b64)

					request, _ = http.NewRequest("POST", "http://localhost:8080/v1/tenancy/null/federation", bytes.NewReader(body))
					request.Header.Set("Content-Type", "application/vnd.api+json")
					request.Header.Set("Accept", "application/vnd.api+json")
					request.Header.Set("Content-Length", fmt.Sprint(len(body)))
				})

				It("returns 404-StatusBadRequest, and an empty payload", func() {
					response, err := http.DefaultClient.Do(request)
					Expect(err).ShouldNot(HaveOccurred())
					resBody, err := io.ReadAll(response.Body)
					Expect(err).ShouldNot(HaveOccurred())

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(string(resBody)).Should(Equal("Invalid tenancy ID"))
				})
			})
		})
	})
})
