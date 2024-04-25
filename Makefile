# Docker container name
CONTAINER_NAME = hostel_gin_pg
# PostgreSQL database name
DB_NAME = hostel_gin_db

# Create PostgreSQL container
createpg:
	docker run --name $(CONTAINER_NAME) -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=hostelGoGinSecret -d postgres:16-alpine

# Start PostgreSQL container
startpg:
	docker start $(CONTAINER_NAME)

# Stop PostgreSQL container
stoppg:
	docker stop $(CONTAINER_NAME)

# Remove PostgreSQL container
removepg:
	docker rm $(CONTAINER_NAME)

run:
	go run cmd/main.go

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=root --owner=root $(DB_NAME)

# Drop specified database
dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb $(DB_NAME)

.PHONY: createpg startpg stoppg removepg run createdb dropdb