export BINARY_NAME = "ticker-api"
export CONFIG_FILE = "app-config.json"

# global command
.PHONY: all
all: dependencies build test

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
	../bin/${BINARY_NAME} migrate --action=up 

migrate-down:
	../bin/${BINARY_NAME} migrate --action=down

migrate-new:
	../bin/${BINARY_NAME} migrate --action=new --message="new migration script"

api-docs:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag init --output=./web/docs

.PHONY: mock
mock:	
	mockgen -package servermock github.com/gufranmirza/ticker-api/web/server Server > mocks/servermock/mock.go
	mockgen -package routermock github.com/gufranmirza/ticker-api/web/router Router > mocks/routermock/mock.go
	mockgen -package healthmock github.com/gufranmirza/ticker-api/web/services/health Health > mocks/healthmock/mock.go
	mockgen -package tradesservicemock github.com/gufranmirza/ticker-api/web/services/v1/tradesservice TradesService > mocks/tradesservicemock/mock.go
	mockgen -package tradesdalmock github.com/gufranmirza/ticker-api/dal/tradesdal TradesDal > mocks/tradesdalmock/mock.go

packages = \
	./dal/tradesdal\
	./logging\
	./web/server\
	./web/router\
	./web/services/health\
	./web/services/v1/tradesservice\

.PHONY: test
test:
	@$(foreach package,$(packages), \
    		set -e; \
			go test -coverprofile $(package)/cover.out $(package);)