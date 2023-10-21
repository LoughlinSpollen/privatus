#! /usr/bin/env sh

eval go build -v -o ./build/tenancy_service -tags debug
eval ./build/tenancy_service