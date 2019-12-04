#!/bin/sh
set -e

if [ "$APP_ENV" != 'production' ]; then
    echo >&2 "Waiting for Postgres to be ready..."
    until pg_isready --timeout=0 --dbname="${DATABASE_URL}"; do
        sleep 1
    done
fi

if [ "$APP_ENV" = 'production' ]; then
    api
else
    go get github.com/pilu/fresh && fresh
fi
