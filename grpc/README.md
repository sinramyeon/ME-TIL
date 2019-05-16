# learn gRPC with Go

### Open salute/salute.proto
1. Define the protocol. Look at the [gRPC documentation on how to protobuf](https://grpc.io/docs/guides/). 
Definition must be precise as we already have a message waiting to be parsed.
2. Generate the Go file, by running `docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=plugins=grpc:. -I. ./salute/salute.proto`

### Open client/client.go
Implement Send

### Open server/server.go
Implement Serve

### Run tests
`go test ./...`

## When all tests have passed
Either run detached or in 2 different terminals:
```bash
go run server/server.go
go run web/web.go
```

Now open localhost:8080
