#! /usr/bin/env sh


source ./scripts/db-config.sh 

FILE=$POSTGRES_DATA/postmaster.pid
if [ ! -f "$FILE" ]; then
    echo "starting ${PRIVATUS_DB_NAME} database"
    postgres -D $POSTGRES_DATA >logfile 2>&1 &
    sleep 2
else
    echo "${PRIVATUS_DB_NAME} database already started"
fi
