include .env

# Default values (can be overridden in .env)
PORT ?= 8080
APP_ENV ?= local
PROTOCOL ?= "mongodb+srv"
USERNAME ?= "mongo"
PASSWORD ?= "mongo"
HOST ?= "localhost:27017"
APPNAME ?= "rewards_and_recognition"

up:
	@echo "Starting containers..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stoping containers..."
	docker-compose down

build:
	go build -o ${APPNAME} ./main/

start:
	@echo "Starting application..."
	@env PORT=${PORT} APP_ENV=${APP_ENV} MONGO_DB_USERNAME=${USERNAME} MONGO_DB_PASSWORD=${PASSWORD} MONGO_DB_HOST=${HOST} MONGO_DB_DATABASE=${APPNAME} MONGO_DB_PROTOCOL=${PROTOCOL} ./${APPNAME}

restart: build start

format_all_code:
	go fmt ./...
