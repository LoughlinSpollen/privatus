#! /usr/bin/env zsh

if ! [ -x "$(command -v go)" ]; then
  echo 'Error: golang is not installed, Run install-dev in the tenancy_service folder first.' >&2
  exit 1
else
  echo 'found golang'
fi


echo "Installing watchexec, run 'make watch' for auto rebuilds of gRPC"
brew install watchexec

echo
echo "Installing grpc core and protobuffers"

pip install grpcio-tools

brew tap grpc/grpc
brew install grpc

echo
brew install protoc-gen-go
