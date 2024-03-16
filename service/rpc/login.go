package rpc

import (
	"context"
	"log"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
)

func (IdentityService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginReponse, error) {
	log.Printf("Login(%s, %s)", req.Email, req.Password)

	return &proto.LoginReponse{Access: "abc", Refresh: "xyz"}, nil
}
