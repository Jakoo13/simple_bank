createPgContainer:
	docker run --name postgres-12 --network bank-network -p 5432:5432 -e POSTGRES_USER=jakew20 -e POSTGRES_PASSWORD=Fatkid06 -d postgres:12-alpine

startPgContainer:
	docker start postgres-12

createdb:
	docker exec -it postgres-12 createdb --username=jakew20 --owner=jakew20 simple_bank

dropdb:
	docker exec -it postgres-12 dropdb --username=jakew20 simple_bank

migrateupaws:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@simple-bank.cxptpm9a77z1.us-east-1.rds.amazonaws.com:5432/simple_bank" -verbose up

migratedownaws:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@simple-bank.cxptpm9a77z1.us-east-1.rds.amazonaws.com:5432/simple_bank" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://jakew20:Fatkid06@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

createApiAndDbContainers:
# Does the same things as pressing play in docker desktop
	docker-compose up

deleteApiAndDbContainers:
	docker-compose down

sqlc:
	sqlc generate

db_docs:
	dbdocs build doc/db.dbml

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: createdb dropdb migrateup migratedown sqlc test server migrateupone migratedownone createPgContainer startPgContainer db_docs createApiAndDbContainers deleteApiAndDbContainers migrateupaws migratedownaws migrateup1 migratedown1 proto