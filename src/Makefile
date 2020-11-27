export BINARY_NAME = "ticker-api"
export CONFIG_FILE = "app-config.json"

# global command
.PHONY: all
all: dependencies build

.PHONY: dependencies
dependencies:
	go mod download

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o ../bin/${BINARY_NAME} ./
	cp -r ${CONFIG_FILE} ../bin

.PHONY: run
run:
	go run ./main.go

.PHONY: start
start:
	../bin/${BINARY_NAME}


migrate-up:
	../bin/${BINARY_NAME} migrate up 

migrate-down:
	../bin/${BINARY_NAME} migrate down

migrate-new:
	../bin/${BINARY_NAME} migrate new "new migration script"

api-docs:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag init --output=./web/docs