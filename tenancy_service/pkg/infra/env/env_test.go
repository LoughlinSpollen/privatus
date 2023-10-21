package env_test

import (
	"os"

	"testing"

	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/env"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInfraEnv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "env config unit test suite")
}

var _ = Describe("infrastructure environment variables", func() {

	Describe("Reading env variables for REST service", func() {

		Context("when the TENANCY_API_HTTP_PORT environment variable is not a number", func() {
			BeforeEach(func() {
				os.Setenv("TENANCY_API_HTTP_PORT", "-")
			})
			AfterEach(func() {
				os.Setenv("TENANCY_API_HTTP_PORT", "8080")
			})
			It("should return the default value", func() {
				result := env.WithDefaultInt("TENANCY_API_HTTP_PORT", 8080)
				Expect(result).Should(Equal(8080))
			})
		})

		Context("when the ML_SERVICE_PORT environment variable is empty", func() {
			BeforeEach(func() {
				os.Setenv("ML_SERVICE_PORT", "")
			})
			AfterEach(func() {
				os.Setenv("ML_SERVICE_PORT", "0.0.0.0")
			})
			It("should return the default value", func() {
				result := env.WithDefaultString("ML_SERVICE_PORT", "0.0.0.0")
				Expect(result).Should(Equal("0.0.0.0"))
			})
		})
	})
})
