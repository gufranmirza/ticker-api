# ticker-api
API that provides the functionality to manage the treding into stocks

## Build Project
```
make
```

## Apply Migrations
```
make migrate-up
```
or to create new migration
```
make migrate-new
```

## Running in Production 
```
make start
```

## Development
```
make run
```

## API Documentation
[Swagger API documetation](http://localhost:8001/ticker-api/v1/swagger/index.html)

[Swagger file](https://github.com/gufranmirza/ticker-api/blob/main/src/web/docs/swagger.yaml)

## Testing
Running test
```
make test
```

## Generate Mocks
Generate package mocks
```
make mock
```

## Running Docker image
Build and run docker image
```
docker-compose up -d --build
```
