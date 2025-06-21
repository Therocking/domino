APP_NAME=domino
BIN_DIR=./bin

all: build

dev:
	go run ./cmd/main.go

build:
	go build -o $(BIN_DIR)/$(APP_NAME) cmd/main.go

run: build
	$(BIN_DIR)/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)/$(APP_NAME)

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm --env-file .env -p 8080:8080 $(APP_NAME):latest

docker-compose:
	docker compose -f .\docker-componse.yml up -d --build