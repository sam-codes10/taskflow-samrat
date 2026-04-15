#!/bin/bash
set -e

# Construct Goose compatible connection string based on env variables
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} host=${DB_HOST} port=${DB_PORT} sslmode=disable"

echo "Running Database Migrations with Goose..."
# Explicitly pointing to our migrations folder dynamically wrapped inside container
goose -dir migrations up

echo "Starting TaskFlow Samrat Application..."
# Replace current bash shell exclusively with our go server ensuring PID 1 transfers to go binary gracefully
exec ./taskflow
