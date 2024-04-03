# identity-go

An identity microservice written in Go; implemented with gRPC.

Auto generate the RESTful API and Swagger docs; to the OpenAPI specification.

## CLI commands

Start the service: `make up`

Update go dependencies: `go mod tidy`

Generate protobuf code: `make protoc`

## Other commands

HTTP login request using curl:

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"email": "x@x", "password": "x"}' \
  http://localhost:9800/login
```

## Testing

| Tests                                                                                        | Description  |
| -------------------------------------------------------------------------------------------- | ------------ |
| `docker exec identity-go go test ./...`                                                      | All packages |
| `docker exec identity-go go test ./cmd`                                                      | Main only    |
| `docker exec identity-go go test -run ./... github.com/clintvidler/identity-go/service/data` | Data package |

| Test flags | Description   |
| ---------- | ------------- |
| `-v`       | Verbose       |
| `-cover`   | Test coverage |

# gRPC Client

`./cmd/main_test.go` provides an example client connection in action.

To connect to this gRPC service from another project:

1. Copy the gen folder from this project over to the other project, specicially:
   - `./gen/proto/go/proto/service.pb.go`
   - `./gen/proto/go/proto/service.pb.gw.go`
   - `./gen/proto/go/proto/service_grpc.pb.go`
2. Create and connect to a service client:

```go
package main

import (
	"client/gen/proto/go/proto"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := NewClient(":9900")
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Login("x@x", "x")
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

type Client struct {
	service proto.IdentityServiceClient
}

func NewClient(target string) (*Client, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewIdentityServiceClient(conn)

	return &Client{service: client}, nil
}

func (c *Client) Login(email, password string) (*proto.LoginReponse, error) {
	return c.service.Login(context.Background(), &proto.LoginRequest{Email: email, Password: password})
}
```
