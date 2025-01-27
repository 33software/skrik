#!/bin/sh
set -e

echo "Applying swagger docs..."
swag init -g internal/app/app.go

exec "$@"
