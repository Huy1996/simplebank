#!/bin/sh

set -e # The code exit immediately if it return a non 0 state

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"