# ==============================================================================
# Main

run:
	go run ./cmd/auth/main.go

build:
	go build ./cmd/auth/main.go



# ==============================================================================
# Go migrate postgresql

force:
	migrate -database postgres://postgres:postgres@localhost:5432/file_db?sslmode=disable -path migrations force 1

version:
	migrate -database postgres://postgres:postgres@localhost:5432/file_db?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5432/file_db?sslmode=disable -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5432/file_db?sslmode=disable -path migrations down 1



# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build
