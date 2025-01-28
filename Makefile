.PHONY: up down build

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build


run_sync:
	go run cmd/sync.go