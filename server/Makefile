# ==============================================================================
# Main

run:
	go run ./cmd/mycloud/main.go

build:
	go build ./cmd/mycloud/main.go

# ==============================================================================
# Go migrate postgresql

force:
	migrate -database postgres://postgres:postgres@localhost:5431/mycloud_db?sslmode=disable -path migrations force 1

version:
	migrate -database postgres://postgres:postgres@localhost:5431/mycloud_db?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5431/mycloud_db?sslmode=disable -path migrations up 1
		# migrate -path migrations/ -database "postgres://postgres:postgres@localhost:5431/mycloud_db?sslmode=disable" -verbose up

migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5431/mycloud_db?sslmode=disable -path migrations down 1

# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)