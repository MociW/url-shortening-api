#!make
include config.env


migrateup:
	migrate -path db/migrations -database "mysql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@tcp($(DATABASE_HOST):$(DATABASE_PORT))/$(DATABASE_NAME)" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@tcp($(DATABASE_HOST):$(DATABASE_PORT))/$(DATABASE_NAME)" -verbose down

test:
	go test -v -cover ./...

server:
	go run cmd/main.go

.PHONY: migrateup migratedown test server