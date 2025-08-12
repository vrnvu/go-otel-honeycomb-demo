# go-otel-honeycomb-demo

Simple exploration demo for OpenTelemetry + Honeycomb. 

Three tiny HTTP services:

- server (port 8080) calls provider and db
- provider (port 8081) returns foo/bar/baz with random outcomes
- db (port 8082) returns 200/500 randomly

## Build

Requires Go.

1) Per-service env files (copy and edit):

```
cp .env.example cmd/server/.env
cp .env.example cmd/provider/.env
cp .env.example cmd/db/.env
```

2) Build locally:

```
go build ./...
```

3) Run locally:

Make sure you source .env or inject it, i.e

```
source .env && go run cmd/server/main.go
source .env && go run cmd/provider/main.go
source .env && go run cmd/db/main.go
```

## Run with Docker Compose

Configure .env first. 

Then build and start all services:

```
docker compose up
```

## Curl examples

User-Id header is required to test baggages and multi-span tracing.

```
curl -H "User-Id: User" http://localhost:8080/foo
curl -H "User-Id: User" http://localhost:8080/bar
curl -H "User-Id: User" http://localhost:8080/baz
```