package main

import (
	"context"
	"testing"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(t *testing.T) {
	client, err := NewClient("host.docker.internal:9900")
	assert.NoError(t, err)

	resp, err := client.Login("x@x", "x")
	assert.NoError(t, err)
	assert.Equal(t, resp.Access, "abc")
	assert.Equal(t, resp.Refresh, "xyz")
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
