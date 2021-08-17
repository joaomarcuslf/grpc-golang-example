# auth-blog-service

![Actions Workflow](https://github.com/joaomarcuslf/grpc-golang-example/workflows/go/badge.svg)

## How to Start Development

1. Copy ```sample.env``` to ```.env``` and rename the variables if you need
2. Build the images and run the containers:

```sh
$ cp sample.env .env
$ docker-compose up --build -d
$ docker-compose logs -f api
```

## Generate gRPC files

```sh
$ docker exec -it grpc-golang-example_api_1 bash
$ protoc --proto_path=proto ./proto/*.proto --go_out=plugins=grpc:./
```

## Test with Client

```sh
$ docker exec -it grpc-golang-example_api_1 bash
$ go run client/main.go
```

- API: [http://localhost:5000](http://localhost:5000)
