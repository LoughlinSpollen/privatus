#! /usr/bin/env sh

echo "Generating GPRC/Protobuf"

echo "> Generateing Python bindings for mpc services & ml service"
mkdir -p tmp/build/protos/api_federation
mkdir -p tmp/build/protos/api_tenancy

cp api-federation.proto tmp/build/protos/api_federation
cp api-tenancy.proto tmp/build/protos/api_tenancy

rm -rf ../tenancy_service/build/protos/api_tenancy
rm -rf ../federation_service/build/protos/api_tenancy
rm -rf ../federation_service/build/protos/api_federation
rm -rf ../federation_lib/build/protos/api_federation

mkdir -p ../tenancy_service/build/protos/api_tenancy
mkdir -p ../federation_service/build/protos/api_tenancy
mkdir -p ../federation_service/build/protos/api_federation
mkdir -p ../federation_lib/build/protos/api_federation

cd tmp
echo "> Generating Golang bindings"
protoc --proto_path=build/protos/api_tenancy/ --go_out=../../tenancy_service/build/protos/api_tenancy --go-grpc_out=../../tenancy_service/build/protos/api_tenancy build/protos/api_tenancy/*.proto
protoc --proto_path=build/protos/api_tenancy/ --go_out=../../federation_service/build/protos/api_tenancy --go-grpc_out=../../federation_service/build/protos/api_tenancy build/protos/api_tenancy/*.proto

protoc --proto_path=build/protos/api_federation/ --go_out=../../federation_service/build/protos/api_federation --go-grpc_out=../../federation_service/build/protos/api_federation build/protos/api_federation/*.proto

echo "> Generating C++ bindings"
protoc --proto_path=build/protos/api_federation/ --cpp_out=../../federation_lib/build/protos/api_federation --grpc_out=../../federation_lib/build/protos/api_federation --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` build/protos/api_federation/*.proto

cd ..
rm -rf tmp