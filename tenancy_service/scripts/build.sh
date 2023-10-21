#! /usr/bin/env sh

# eval rm -rf ./build/
# eval go mod init github.com/LoughlinSpollen/tenancy_service
# eval go test -c ./test/unit -ldflags="-w -s" -o ./build/unit_suite.test
# eval go test -c ./test/integration -ldflags="-w -s" -o ./build/integration_suite.test

eval go mod tidy
eval go build -v -o build/tenancy_service

echo "done"