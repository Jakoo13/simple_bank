createPgContainer:
	docker run --name postgres-12 -p 5432:5432 -e POSTGRES_USER=jakew20 -e POSTGRES_PASSWORD=Fatkid06 -d postgres:12-alpine

startPgContainer:
	docker start postgres-12

createdb:
	docker exec -it postgres-12 createdb --username=jakew20 --owner=jakew20 simple_bank

dropdb:
	docker exec -it postgres-12 dropdb --username=jakew20 simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc test server migrateupone migratedownone createPgContainer startPgContainer