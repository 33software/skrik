#!/bin/sh
set -e

echo "Applying swagger docs..."
swag init -g ./routes/routes.go

exec "$@"