#!make
include .env

rl: 
	make down; sleep 1.0; make up;
up:
	docker compose up -d; sleep 1.0; make mg-up; make logs;
logs:
	docker-compose logs -f main;
down:
	docker-compose down; rm -rf .dbdata;

mg-up:
# Dev DB
	migrate -path app/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_DEV_PORT}/datastore?sslmode=disable" -verbose up;
# Test DB
	migrate -path app/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_TEST_PORT}/datastore?sslmode=disable" -verbose up;
mg-down:
# Dev DB
	migrate -path app/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_DEV_PORT}/datastore?sslmode=disable" -verbose down -all;
# Test DB
	migrate -path app/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_TEST_PORT}/datastore?sslmode=disable" -verbose down -all

protoc:
	rm -rf gen
	mkdir -p gen/docs
# Server gRPC
	protoc -I. --go_out=paths=source_relative:./gen/ --go-grpc_out=paths=source_relative:./gen/ proto/server/*.proto
# Server gRPC-Gateway
	protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:./gen/ proto/server/*.proto
# Server swagger docs
	protoc --openapiv2_out=./gen/docs proto/server/*.proto
# Client gRPC
	protoc -I. --go_out=paths=source_relative:./gen/ --go-grpc_out=paths=source_relative:./gen/ proto/client/*.proto
