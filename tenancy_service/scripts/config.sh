#! /usr/bin/env sh

cd ../database
source ./scripts/db-config.sh
cd ../tenancy_service

export PGPORT=5432
export PGHOST="0.0.0.0"

export TENANCY_API_HTTP_PORT=8080

export ML_SERVICE_HOST="0.0.0.0"
export ML_SERVICE_PORT=1025

export MPC_SERVICE_HOST="0.0.0.0"
export MPC_SERVICE_PORT=1026

export FEDERATION_SERVICE_HOST="0.0.0.0"
export FEDERATION_SERVICE_PORT=1027

export EXCHANGE_SERVICE_HOST="0.0.0.0"
export EXCHANGE_SERVICE_PORT=1028

export SERVICE_BACKOFF=100
export SERVICE_MAX_RETRY=40
export SERVICE_TIMEOUT_RETRY=2
export SERVICE_TIMEOUT=5