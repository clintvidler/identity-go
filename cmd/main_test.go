package main

import (
	"context"
	"testing"

	proto "github.com/clintvidler/identity-go/gen/proto/server"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func TestMain(t *testing.T) {
	client, err := NewClient("host.docker.internal:9900")
	assert.NoError(t, err)

	ctx := context.Background()

	_, header, _, err := client.Login(ctx, "x@x", "x")
	assert.NoError(t, err)
	assert.NotEmpty(t, header.Get("access-token"))
	assert.NotEmpty(t, header.Get("refresh-token"))

	// for key, value := range header {
	// 	log.Printf("%s => %s", key, value)
	// }
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

func (c *Client) Login(ctx context.Context, email, password string) (resp *proto.LoginReponse, header metadata.MD, trailer metadata.MD, err error) {

	c.service.Login(
		ctx,
		&proto.LoginRequest{Email: email, Password: password},
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)

	return
}
