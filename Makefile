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
	migrate -path service/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_DEV_PORT}/datastore?sslmode=disable" -verbose up;
# Test DB
	migrate -path service/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_TEST_PORT}/datastore?sslmode=disable" -verbose up;
mg-down:
# Dev DB
	migrate -path service/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_DEV_PORT}/datastore?sslmode=disable" -verbose down -all;
# Test DB
	migrate -path service/data/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_TEST_PORT}/datastore?sslmode=disable" -verbose down -all
