#! /usr/bin/env sh
source ./scripts/db-config.sh

FILE=$POSTGRES_DATA/postmaster.pid
if test -f "$FILE"; then
    echo "${PRIVATUS_DB_NAME} stopping database"
    pg_ctl -D $POSTGRES_DATA stop 2>&1 &
else
    echo "${PRIVATUS_DB_NAME} database was not started"
fi
