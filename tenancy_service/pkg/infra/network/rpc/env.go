package rpc

import (
	"flag"
	"time"

	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/env"
	"google.golang.org/grpc/codes"
)

var (
	fs           = flag.NewFlagSet("rpc-services", flag.ExitOnError)
	timeout      = fs.Int("service-timeout", env.WithDefaultInt("SERVICE_TIMEOUT", 2), "service timeout")
	backoff      = fs.Int("service-backoff", env.WithDefaultInt("SERVICE_BACKOFF", 100), "service backoff")
	maxRetry     = fs.Int("service-max-retry", env.WithDefaultInt("SERVICE_MAX_RETRY", 40), "service max retry")
	timeoutRetry = fs.Int("service-timeout-retry", env.WithDefaultInt("SERVICE_TIMEOUT_RETRY", 2), "service timeout retry")
	mlHost       = fs.String("ml-service-host", env.WithDefaultString("ML_SERVICE_HOST", "0.0.0.0"), "ml-service host")
	mlPort       = fs.Int("ml-service-port", env.WithDefaultInt("ML_SERVICE_PORT", 1025), "ml-service port")
)

var (
	msgTimeout       time.Duration = time.Duration(*timeout) * time.Second
	connBackoff      time.Duration = time.Duration(*backoff) * time.Millisecond
	connTimeoutRetry time.Duration = time.Duration(*timeoutRetry) * time.Second
	retriableErrors                = []codes.Code{codes.Unavailable, codes.DataLoss}
)
