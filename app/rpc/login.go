package rpc

import (
	"context"
	"fmt"
	"log"

	proto "github.com/clintvidler/identity-go/gen/proto/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func (s IdentityService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginReponse, error) {
	log.Printf("Login(%s, %s)", req.Email, req.Password)

	// Find user by email
	user, err := s.data.User.ReadOne(0, req.Email)
	if err != nil {
		// Bad request: User with that email not found
		log.Printf("Error: %s", err)
		return nil, err
	}

	// Verify users password
	if err = user.ComparePassword(req.Password); err != nil {
		// Bad request: Wrong password
		log.Printf("Error: %s", err)
		return nil, err
	}

	// Read the app_id
	p, _ := peer.FromContext(ctx)
	aid := p.Addr.String()

	// Start a new family of tokens
	if err = s.data.Token.DeleteAppFamily(fmt.Sprint(user.Id), aid); err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	// Generate new tokens
	at, rt, err := s.Generate(fmt.Sprint(user.Id), aid)
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	// Creating outgoing metadata
	if err = grpc.SendHeader(ctx, metadata.Pairs(
		"access-token", at,
		"refresh-token", rt,
	)); err != nil {
		log.Println("err", err)
		return nil, err
	}

	return &proto.LoginReponse{}, nil
}
