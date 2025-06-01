include .env

build:
	@go build main.go -o /bin/gobank
	
test:
	@go test -v -coverage ./...
	
run:
	@go run main.go
	
cmp-up:
	@docker compose up -d --build
cmp-down:
	@docker compose down -v
	
createdb:
	@docker exec -it $(DB_CONTAINER) createdb -U $(POSTGRES_USER) -W $(POSTGRES_PASSWORD) 
dropdb:
	@docker exec -it $(DB_CONTAINER) dropdb -U $(POSTGRES_USER) gobank
connectdb: 
	@docker exec -it $(DB_CONTAINER) psql -d $(POSTGRES_DB) -U $(POSTGRES_USER) -W $(POSTGRES_PASSWORD)
	
migrate:
	@migrate create -ext sql -dir database/db/migrations -seq gobank_schema
migrate-up:
	@migrate -database $(DB_PATH) -path database/db/migrations up 
migrate-down:
	@migrate -database $(DB_PATH) -path database/db/migrations down 
fixdirty:
	@docker exec -it $(DB_CONTAINER) psql -c "drop table if exists schema_migrations;" -d $(POSTGRES_DB) -U $(POSTGRES_USER) -W $(POSTGRES_PASSWORD)
	
.PHONY: createdb dropdb connectdb cmp-up cmp-down test run build create-migrate migrate-up migrate-down
