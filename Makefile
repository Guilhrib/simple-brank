postgres:
	docker run --name local-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -d postgres

createdb:
	docker exec -it local-postgres createdb --username=postgres --password --owner=postgres simplebank

dropdb:
	docker exec -it local-postgres dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	DB_SOURCE=postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable go test -v -cover ./...

server:
	go run main.go

.PHONY:
	postgres
	createdb
	dropdb
	migrateup
	migratedown
	sqlc
	test
	server