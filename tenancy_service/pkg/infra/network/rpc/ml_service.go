package rpc

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	pb "github.com/LoughlinSpollen/tenancy_service/build/protos/ml_service_api/pkg/api/v1"
	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type mlService struct {
	connection *grpc.ClientConn
	mlClient   pb.MLClient
}

func NewMLService() *mlService {
	return &mlService{}
}

func (s *mlService) Connect() error {
	log.Debug("mlService Connect")

	opts := []grpc_retry.CallOption{}
	address := fmt.Sprintf("%s:%d", *mlHost, *mlPort)
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		log.Printf("Could not connect to ML service %s:%d : %v \n", *mlHost, *mlPort, err)
		return err
	}
	s.connection = conn
	s.mlClient = pb.NewMLClient(s.connection)
	log.Printf("ML-service connected %s:%d \n", *mlHost, *mlPort)
	return nil
}

func (s *mlService) Close() {
	log.Debug("mlService Close")

	if s.connection == nil {
		log.Warn("Connection to ML service could not be closed - invalid.")
		return
	}
	s.connection.Close()
}

func (s *mlService) Training(training *model.Training) error {
	log.Debug("mlService Training")

	if s.mlClient == nil {
		log.Fatal("Could not send training, not connected to ML service.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), msgTimeout)
	defer cancel()

	response, err := s.mlClient.Training(ctx,
		&pb.TrainingRequest{States: training.States},
		grpc_retry.WithCodes(retriableErrors...),
		grpc_retry.WithMax(uint(*maxRetry)),
		grpc_retry.WithPerRetryTimeout(connTimeoutRetry),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(connBackoff)))
	if err != nil {
		return err
	}
	if errStr := response.GetError(); errStr != "" {
		return errors.New(errStr)
	}

	return nil
}
