#! /usr/bin/env sh

source ./scripts/db-config.sh

rm -rf $POSTGRES_DATA
mkdir -p $POSTGRES_DATA
initdb -D ${POSTGRES_DATA}

sleep 1

echo "database ${PRIVATUS_DB_NAME} deleted"
