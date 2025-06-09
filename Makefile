include .env

# General application commands
build:
	@go build -o ./bin/gobank ./src 	
test:
	@go test ./... -v -cover 
run:
	@go run ./src/main.go
dev:
	@air
	
# Docker commands mostly abbreviations
startdb:
	@docker start $(DB_CONTAINER)
stopdb:
	@docker stop $(DB_CONTAINER)
# startcache:
#     @docker start $(CACHE_CONTAINER)
# stopcache:
#     @docker stop $(CACHE_CONTAINER)
cmp-up:
	@docker compose up -d --build
cmp-down:
	@docker compose down -v
	
# Database related commands
createdb:
	@docker exec -it $(DB_CONTAINER) createdb -U $(POSTGRES_USER) -W $(POSTGRES_PASSWORD) 
dropdb:
	@docker exec -it $(DB_CONTAINER) dropdb -U $(POSTGRES_USER) gobank
connectdb: 
	@docker exec -it $(DB_CONTAINER) psql -d $(POSTGRES_DB) -U $(POSTGRES_USER)
	
# Only migration related commands
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

# Add to phony list	
.PHONY: createdb dropdb connectdb cmp-up cmp-down test run build create-migrate migrate-up migrate-down gensqlc
