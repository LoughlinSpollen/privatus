#! /usr/bin/env sh
echo "creating tables schema ${PRIVATUS_DB_NAME}"

PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public;"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "SET search_path TO ${PRIVATUS_DB_NAME},public;"

# https://www.postgresql.org/docs/9.1/storage-toast.html
# https://dba.stackexchange.com/a/189895
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "CREATE TABLE IF NOT EXISTS ${PRIVATUS_DB_NAME}.tenancy
        (id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
        ml_model TEXT,
        created_at timestamp without time zone DEFAULT current_timestamp,
        updated_at timestamp without time zone);"

PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "CREATE TABLE IF NOT EXISTS ${PRIVATUS_DB_NAME}.federation
        (id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
        tenancy_id uuid REFERENCES ${PRIVATUS_DB_NAME}.tenancy (id),
        threshold INTEGER, 
        epochs INTEGER,
        rate INTEGER,
        rounds INTEGER,
        batch INTEGER,
        created_at timestamp without time zone DEFAULT current_timestamp,
        updated_at timestamp without time zone);"


PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "CREATE TABLE IF NOT EXISTS ${PRIVATUS_DB_NAME}.training
        (id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
        tenancy_id uuid REFERENCES ${PRIVATUS_DB_NAME}.tenancy (id),
        states TEXT NOT NULL,
        created_at timestamp without time zone DEFAULT current_timestamp,
        updated_at timestamp without time zone);"


# updated_at date trigger
read -r -d '' DATE_TRIGGER <<- EOM
create or replace function updated_at_column()
returns trigger as \$$
begin
    new.updated_at = now();
    return new;
end;
\$$ language 'plpgsql';
EOM
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "${DATE_TRIGGER}"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "create trigger updated_at_trigger before insert or update on ${PRIVATUS_DB_NAME}.tenancy for each row execute procedure updated_at_column();"
PGPASSWORD=${PRIVATUS_DB_ADMIN_PASSWORD} psql -d ${PRIVATUS_DB_NAME} -U${PRIVATUS_DB_ADMIN} -c "create trigger updated_at_trigger before insert or update on ${PRIVATUS_DB_NAME}.training for each row execute procedure updated_at_column();"
