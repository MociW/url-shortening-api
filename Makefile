#!make
include *.env


migrateup:
	migrate -path db/migrations -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(DATABASE_PORT))/$(DATABASE)" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(DATABASE_PORT))/$(DATABASE)" -verbose down

test:
	go test -v -cover ./...

server:
	go run cmd/main.go

.PHONY: migrateup migratedown test server