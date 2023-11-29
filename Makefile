postgres:
	docker run --name postgres12 -p 5400:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createDb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

startDb:
	docker start postgres12

dropDb:
	docker exec -it postgres12 dropdb simple_bank

migrateUp:		
	migrate -path db/migration -database "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...		

.PHONY: postgres createDb dropDb migrateUp migrateDown sqlc startDb test