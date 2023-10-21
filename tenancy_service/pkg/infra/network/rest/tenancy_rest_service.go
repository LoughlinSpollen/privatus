package rest

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/env"

	"github.com/LoughlinSpollen/tenancy_service/pkg/usecase"

	"github.com/gorilla/mux"
)

var (
	fs       = flag.NewFlagSet("rest-service", flag.ExitOnError)
	httpPort = fs.Int("tenancy-api-port", env.WithDefaultInt("TENANCY_API_HTTP_PORT", 8080), "tenancy API HTTP port")
)

type tenancyRestService struct {
	tenancyUsecase       usecase.TenancyUsecase
	httpServer           *http.Server
	router               *mux.Router
	tenancyDTOAdapter    *tenancyDTOAdapter
	federationDTOAdapter *federationDTOAdapter
	signatureValidater   *signatureValidater
}

func NewTenancyRestService(tenancyUsecase usecase.TenancyUsecase) *tenancyRestService {
	log.Debug("NewTenancyRestService")

	restService := &tenancyRestService{
		router:               mux.NewRouter().StrictSlash(true),
		tenancyDTOAdapter:    NewTenancyDTOAdapter(),
		federationDTOAdapter: NewFederationDTOAdapter(),
		tenancyUsecase:       tenancyUsecase,
		signatureValidater:   NewSignatureValidater(),
	}
	return restService
}

func (s *tenancyRestService) Connect() {
	log.Debug("tenancyRestService Connect")

	s.router.HandleFunc("/v1/tenancy", s.createTenancy).Methods("POST")
	s.router.HandleFunc("/v1/tenancy/{id}/federation", s.createFederation).Methods("POST")
	address := fmt.Sprintf(":%d", *httpPort)
	log.Info(fmt.Printf("tenancyRestService listening %s", address))

	s.httpServer = &http.Server{Addr: address, Handler: s.router}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "http server error: %v \n", err)
		}
	}()
	log.Debug("HTTP server started")
	<-done
	log.Debug("HTTP server stopped")
}

func (s *tenancyRestService) Close() {
	log.Debug("tenancyRestService Close")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "http server shutdown error: %v \n", err)
	}
}

func (s *tenancyRestService) createTenancy(w http.ResponseWriter, r *http.Request) {
	log.Debug("tenancyRestService createTenancy")

	var tenancy *model.Tenancy
	if tenancy = s.requestToTenancy(w, r); tenancy == nil {
		return //requestToTenancy will have written error response
	}

	if err := s.tenancyUsecase.CreateTenancy(tenancy); err != nil {
		if err.Error() == "Already exists" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Failed to create tenancy")
		return
	}
	if err := s.writeTenancyResponse(tenancy, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Failed to serialise tenancy")
		return
	}
}

func (s *tenancyRestService) createFederation(w http.ResponseWriter, r *http.Request) {
	log.Debug("tenancyRestService createFederation")

	var federation *model.Federation
	if federation = s.requestToFederation(w, r); federation == nil {
		return //requestToFederation will have written error response
	}
	var err error
	federation.TenancyID, err = uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Invalid tenancy ID")
		return
	}

	if err := s.tenancyUsecase.CreateFederation(federation); err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Failed to update tenancy federation")
		return
	}
	if err := s.writeFederationResponse(federation, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Failed to serialise tenancy federation")
		return
	}
}

func (s *tenancyRestService) requestToTenancy(w http.ResponseWriter, r *http.Request) *model.Tenancy {
	log.Debug("tenancyRestService requestToTenancy")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Missing tenancy attributes")
		return nil
	}

	if validated := s.signatureValidater.ValidateJSON(w, r, reqBody); !validated {
		return nil
	}

	tenancy, err := s.tenancyDTOAdapter.BytesToTenancy(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Failed to extract tenancy attributes")
		return nil
	}

	return tenancy
}

func (s *tenancyRestService) requestToFederation(w http.ResponseWriter, r *http.Request) *model.Federation {
	log.Debug("tenancyRestService requestToFederation")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Missing federation attributes")
		return nil
	}

	if validated := s.signatureValidater.ValidateJSON(w, r, reqBody); !validated {
		return nil
	}

	federation, err := s.federationDTOAdapter.BytesToFederation(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "Failed to extract tenancy attributes")
		return nil
	}

	return federation
}

func (s *tenancyRestService) writeTenancyResponse(tenancy *model.Tenancy, w http.ResponseWriter) error {
	log.Debug("tenancyRestService writeTenancyResponse")

	w.WriteHeader(http.StatusCreated)
	tenancyBytes, err := s.tenancyDTOAdapter.TenancyToBytes(tenancy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	w.Header().Add("Content-Type", "application/vnd.api+json")
	_, err = w.Write(tenancyBytes)
	return err
}

func (s *tenancyRestService) writeFederationResponse(federation *model.Federation, w http.ResponseWriter) error {
	log.Debug("tenancyRestService writeFederationResponse")

	federationBytes, err := s.federationDTOAdapter.FederationToBytes(federation)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(federationBytes)
	return err
}
