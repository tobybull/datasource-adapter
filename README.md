# Datasource Adapter

The project name "datasource-adapter" should match the name of the git repo

## Run locally

```bash
go run ./cmd/foobar-adapter/main.go
```

## Build the binary and run
```bash
go build . 
./datasource-adapter
```

## Configuration

This app is configured via environment variables because that is the idiomatic cloud native way to configure a microservice. 

For local development, the app will read from the `.env` file.

## Send a request
```bash
curl -X POST -H "Content-Type: application/json" -d '{"input_data": "hello go", "user_id": 123}' http://localhost:8080/page
```
