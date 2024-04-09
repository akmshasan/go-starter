dbup:
	docker run -it --name postgres-db -p 5432:5432 -e POSTGRES_USER=test -e POSTGRES_PASSWORD=secret -d postgres:16-alpine3.19

dbdown:
	docker stop postgres-db && docker rm postgres-db

createdb:
	docker exec -it postgres-db createdb --username=test --owner=test fruit_store

dropdb:
	docker exec -it postgres-db dropdb --username=test fruit_store

migrate-init:
	migrate create -ext sql -dir ./db/migration -seq init_schema 

migrate-up:
	migrate -path ./db/migration -database "postgresql://test:secret@localhost:5432/fruit_store?sslmode=disable" -verbose up

migrate-down:
	migrate -path ./db/migration -database "postgresql://test:secret@localhost:5432/fruit_store?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go


.PHONY: dbup dbdown createdb dropdb migrate-up migrate-down sqlc server