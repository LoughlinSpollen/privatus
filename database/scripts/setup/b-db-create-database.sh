#! /usr/bin/env sh

echo "creating database ${PRIVATUS_DB_NAME}"
createdb --template=template0 --locale=en_US.UTF-8 --encoding=UTF8 ${PRIVATUS_DB_NAME}

psql -d ${PRIVATUS_DB_NAME} -c "alter user ${PRIVATUS_DB_ADMIN} with encrypted password '${PRIVATUS_DB_ADMIN_PASSWORD}';"
psql -d ${PRIVATUS_DB_NAME} -c "alter user ${PRIVATUS_DB_ADMIN} with superuser;"
