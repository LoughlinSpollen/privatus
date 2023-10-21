#! /usr/bin/env sh

echo "creating role in schema ${PRIVATUS_DB_NAME}"


# readwrite role for PRIVATUS_DB_USER
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "create role readwrite;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "grant connect on database ${PRIVATUS_DB_NAME} to readwrite;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "grant usage, create on schema ${PRIVATUS_DB_NAME} to readwrite;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "grant select, insert, update, delete on all tables in schema ${PRIVATUS_DB_NAME} to readwrite;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "alter default privileges in schema ${PRIVATUS_DB_NAME} grant select, insert, update, delete on tables to readwrite;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "grant usage on all sequences in schema ${PRIVATUS_DB_NAME} to readwrite;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "alter default privileges in schema ${PRIVATUS_DB_NAME} grant usage on sequences to readwrite;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "alter default privileges grant all on functions to readwrite;"

PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "grant readwrite to ${PRIVATUS_DB_USER};"
