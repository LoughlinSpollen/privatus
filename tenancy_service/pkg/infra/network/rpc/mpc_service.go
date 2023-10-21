package rpc

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	pb "github.com/LoughlinSpollen/tenancy_service/build/protos/mpc_service_api/pkg/api/v1"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type mpcService struct {
	connection *grpc.ClientConn
	mpcClient  pb.MPCClient
}

func NewMPCService() *mpcService {
	return &mpcService{}
}

func (s *mpcService) Connect(host string, port int) error {
	log.Debug("mpcService Connect")

	opts := []grpc_retry.CallOption{}
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		log.Printf("could not connect to mpc service %s:%d : %v \n", host, port, err)
		return err
	}
	s.connection = conn
	s.mpcClient = pb.NewMPCClient(s.connection)
	log.Printf("mpc-service connected %s:%d \n", host, port)
	return nil
}

func (s *mpcService) Close() {
	log.Debug("mpcService Close")

	if s.connection == nil {
		log.Warn("Connection to mpc service could not be closed - invalid.")
		return
	}
	s.connection.Close()
}

func (s *mpcService) PrimeGen(activeMembers int32) (string, error) {
	log.Debug("mpcService PrimeGen")

	if s.mpcClient == nil {
		log.Fatal("Could not send PrimeGen request, not connected to mpc service.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), msgTimeout)
	defer cancel()

	response, err := s.mpcClient.Prime(ctx,
		&pb.PrimeGenRequest{
			Size: activeMembers,
		},
		grpc_retry.WithCodes(retriableErrors...),
		grpc_retry.WithMax(uint(*maxRetry)),
		grpc_retry.WithPerRetryTimeout(connTimeoutRetry),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(connBackoff)))
	if err != nil {
		return "", err
	}
	if errStr := response.GetError(); errStr != "" {
		return "", errors.New(errStr)
	}

	prime := response.GetGen()
	return prime, nil
}
