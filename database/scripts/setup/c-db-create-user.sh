#! /usr/bin/env sh

echo "creating user is ${PRIVATUS_DB_USER}"
createuser ${PRIVATUS_DB_USER}

PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "alter user ${PRIVATUS_DB_USER} with encrypted password '${PRIVATUS_DB_PASSWORD}';"
