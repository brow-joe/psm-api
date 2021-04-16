## System

   * go              : go1.16.2

## Available Scripts

In the project directory, you can run:

### `go get github.com/gorilla/mux`

Adds the dependency

### `go mod vendor`

Creates the vendor of the premises

### `go run main.go`

Launch the application

### `go test -v`

Run tests

## Docker build

In the project directory, you can run:

### Build image

```sh
docker build -t psm-api .
```

### Run container

```sh
docker run -d -p 8080 -p 8080 --net=host psm-api
```
