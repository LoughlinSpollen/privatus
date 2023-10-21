#! /usr/bin/env sh

source ./scripts/db-config.sh

echo "${PRIVATUS_DB_NAME} database starting init - create tables'"

if psql ${PRIVATUS_DB_NAME} -c '\q' 2>&1; then
    ./scripts/setup/f-db-create-tenancy-tables.sh
    echo "database ${PRIVATUS_DB_NAME} database finished init - created tenancy tables'"
else
    echo "${PRIVATUS_DB_NAME} database does not exist. Try running 'make db-setup'"
fi

