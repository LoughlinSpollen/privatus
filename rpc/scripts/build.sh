#! /usr/bin/env sh

echo "Generating GPRC/Protobuf"

echo "> Generateing Python bindings for mpc services & ml service"
mkdir -p tmp/build/protos/ml_service_api
mkdir -p tmp/build/protos/mpc_service_api

cp ml-service-api.proto tmp/build/protos/ml_service_api
cp mpc-service-api.proto tmp/build/protos/mpc_service_api

cd tmp
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. build/protos/ml_service_api/ml-service-api.proto
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. build/protos/mpc_service_api/mpc-service-api.proto
cd ..

rm -rf ../mpc_service/build/protos/mpc_service_api
rm -rf ../federation_service/build/protos/mpc_service_api
rm -rf ../exchange_service/build/protos/mpc_service_api

mkdir -p ../mpc_service/build/protos/mpc_service_api
mkdir -p ../federation_service/build/protos/mpc_service_api
mkdir -p ../exchange_service/build/protos/mpc_service_api

cp tmp/build/protos/mpc_service_api/*.py ../mpc_service/build/protos/mpc_service_api
cp tmp/build/protos/mpc_service_api/*.py ../federation_service/build/protos/mpc_service_api
cp tmp/build/protos/mpc_service_api/*.py ../exchange_service/build/protos/mpc_service_api


rm -rf ../ml_service/build/protos/ml_service_api
mkdir -p ../ml_service/build/protos/ml_service_api
cp tmp/build/protos/ml_service_api/*.py ../ml_service/build/protos/ml_service_api


cd tmp
echo "> Generating Golang bindings for mpc services & ml service"
mkdir -p ../../tenancy_service/build/protos/ml_service_api
mkdir -p ../../tenancy_service/build/protos/mpc_service_api
protoc --proto_path=build/protos/ml_service_api/ --go_out=../../tenancy_service/build/protos/ml_service_api --go-grpc_out=../../tenancy_service/build/protos/ml_service_api build/protos/ml_service_api/*.proto
protoc --proto_path=build/protos/mpc_service_api/ --go_out=../../tenancy_service/build/protos/mpc_service_api --go-grpc_out=../../tenancy_service/build/protos/mpc_service_api build/protos/mpc_service_api/*.proto
cd ..
rm -rf tmp