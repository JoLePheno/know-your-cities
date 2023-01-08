#!/bin/bash

psql -c "CREATE DATABASE treezor"

go run cmd/migrations/*.go init

go run cmd/migrations/*.go version

go run cmd/migrations/*.go up