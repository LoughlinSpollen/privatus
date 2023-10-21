package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/LoughlinSpollen/tenancy_service/debug"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/env"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/network/rest"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/network/rpc"
	"github.com/LoughlinSpollen/tenancy_service/pkg/usecase"

	database "github.com/LoughlinSpollen/tenancy_service/pkg/infra/db"
)

var (
	fs = flag.NewFlagSet("privatus-services", flag.ExitOnError)

	mpcHost = fs.String("mpc-service-host", env.WithDefaultString("MPC_SERVICE_HOST", "0.0.0.0"), "mpc-service host")
	mpcPort = fs.Int("mpc-service-port", env.WithDefaultInt("MPC_SERVICE_PORT", 1026), "mpc-service port")

	mlHost = fs.String("ml-service-host", env.WithDefaultString("ML_SERVICE_HOST", "0.0.0.0"), "ml-service host")
	mlPort = fs.Int("ml-service-port", env.WithDefaultInt("ML_SERVICE_PORT", 1027), "ml-service port")
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	_ = debug.Tags()
}

func main() {
	log.Debug("starting tenancy service")

	dbService := database.NewTenancyDB()
	if err := dbService.Connect(); err != nil {
		log.Fatalf("could not connect to database : %v", err)
	}
	defer dbService.Close()

	mpcService := rpc.NewMPCService()
	// if err := mpcService.Connect(*mpcHost, *mpcPort); err != nil {
	// 	log.Fatalf("could not connect to mpc service : %v", err)
	// }
	// defer mpcService.Close()

	mlService := rpc.NewMLService()
	// if err := mlService.Connect(); err != nil {
	// 	log.Fatalf("could not connect to ML service : %v", err)
	// }
	// defer mlService.Close()

	tenancyUsecase := usecase.NewTenancyUsecase(dbService, mpcService, mlService)
	tenancyRestService := rest.NewTenancyRestService(tenancyUsecase)
	tenancyRestService.Connect()
	defer tenancyRestService.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Debug("stopped tenancy service")
}
