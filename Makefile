include .env

build:
	@go build -o ./bin/gobank ./src 	
test:
	@go test ./... -v -cover 
run:
	@go run ./src/main.go
dev:
	@air
	
start:
	@docker start $(DB_CONTAINER)
stop:
	@docker stop $(DB_CONTAINER)
cmp-up:
	@docker compose up -d --build
cmp-down:
	@docker compose down -v
	
createdb:
	@docker exec -it $(DB_CONTAINER) createdb -U $(POSTGRES_USER) -W $(POSTGRES_PASSWORD) 
dropdb:
	@docker exec -it $(DB_CONTAINER) dropdb -U $(POSTGRES_USER) gobank
connectdb: 
	@docker exec -it $(DB_CONTAINER) psql -d $(POSTGRES_DB) -U $(POSTGRES_USER)
	
migrate:
	@migrate create -ext sql -dir database/db/migrations -seq gobank_schema
migrate-up:
	@migrate -database $(DB_PATH) -path database/db/migrations up 
migrate-down:
	@migrate -database $(DB_PATH) -path database/db/migrations down 
fixdirty:
	@docker exec -it $(DB_CONTAINER) psql -c "drop table if exists schema_migrations;" -d $(POSTGRES_DB) -U $(POSTGRES_USER) -W $(POSTGRES_PASSWORD)
gensqlc:
	@sqlc generate
	
.PHONY: createdb dropdb connectdb cmp-up cmp-down test run build create-migrate migrate-up migrate-down gensqlc
