#! /usr/bin/env sh

# setup the database
source ../database/scripts/db-config.sh

echo "starting database"
cd ../database
./scripts/db-stop.sh
./scripts/db-delete.sh
./scripts/db-start.sh
./scripts/db-setup.sh
./scripts/db-init.sh
./scripts/db-truncate.sh
cd ../tenancy_service

echo "Unit and integration tests"
ginkgo -timeout=60s -r -v -cover -coverprofile=coverage.out -coverpkg=./... --output-dir=../test-results

# echo "staring services for integration tests"

# echo "starting federation_server"
# cd ../federation_server/
# python main.py &

# echo "starting mpc_service"
# eval cd ../mpc_service/
# python main.py &

# echo "starting exchange_server"
# eval cd ../exchange_server/
# python main.py &

# echo "starting ml_service"
# eval cd ../ml_service/
# python main.py &

# cd ../tenancy_service

# echo "Integration tests"
# go test ./test/integration -v -count=1 

# echo "stopping services after integration tests"

# kill -9 $(lsof -ti tcp:1025) 
# kill -9 $(lsof -ti tcp:1026) 
# kill -9 $(lsof -ti tcp:1027) 
# kill -9 $(lsof -ti tcp:1028) 
cd ../database
./scripts/db-stop.sh

cd ../tenancy_service

echo "Tests complete"

