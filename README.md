# identity-go

An identity microservice written in Go; implemented with gRPC.

Auto generate the RESTful API and Swagger docs; to the OpenAPI specification.

## CLI commands

Start the service: `go run cmd/main.go`

Update go dependencies: `go mod tidy`

Generate protobuf code: `buf generate`

Update protobuf dependencies: `buf mod update`

## Other commands

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"email": "x@x", "password": "x"}' \
  http://localhost:8080/login
```

# Testing

| Tests                                                                                        | Description  |
| -------------------------------------------------------------------------------------------- | ------------ |
| `docker exec identity-go go test ./...`                                                      | All packages |
| `docker exec identity-go go test`                                                            | Main only    |
| `docker exec identity-go go test -run ./... github.com/clintvidler/identity-go/service/data` | Data package |

| Test flags | Description   |
| ---------- | ------------- |
| `-v`       | Verbose       |
| `-cover`   | Test coverage |
