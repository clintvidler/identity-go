package rpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *IdentityService) Refresh(ctx context.Context, req *proto.RefreshRequest) (*proto.RefreshReponse, error) {
	log.Println("Refreshing token...")

	// /

	// Read the metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in the context")
	}

	var refresh string

	if md["grpcgateway-cookie"] != nil {
		// log.Println(md["grpcgateway-cookie"])
		cookies := make(map[string]string)
		for _, e := range strings.Split(md.Get("grpcgateway-cookie")[0], "; ") {
			parts := strings.Split(e, "=")
			cookies[parts[0]] = parts[1]
		}
		refresh = cookies["rt"]
	} else {
		// Read the token from the metadata
		if len(md.Get("refresh")) < 1 {
			return nil, errors.New("no refresh token")
		}
		refresh = md.Get("refresh")[0]
	}

	log.Println(refresh)

	// /

	// Validate token
	claims, err := s.ParseClaims(fmt.Sprint(refresh))
	if err != nil {
		return nil, fmt.Errorf("Refresh token: %s", err.Error())
	}
	uid := claims["sub"].(string)
	aid := claims["aud"].(string)

	// Ensure the token exists in the current family
	if err := s.data.Token.Exists(refresh); err != nil {
		return nil, err
	}

	// Read the most recently created token for the user/app
	lt, err := s.data.Token.ReadLatestByIds(uid, aid)
	if err != nil {
		return nil, err
	}

	// Invalidate (remove) all of the users tokens if an old token from the current family was provided
	if refresh != lt {
		if err := s.data.Token.DeleteAppFamily(uid, aid); err != nil {
			return nil, err
		}
		return nil, errors.New("an old token was used")
	}

	// Generate new tokens
	at, rt, err := s.Generate(uid, aid)
	if err != nil {
		return nil, err
	}

	// Creating outgoing metadata
	if err = grpc.SendHeader(ctx, metadata.Pairs(
		"access-token", at,
		"refresh-token", rt,
	)); err != nil {
		return nil, err
	}

	return &proto.RefreshReponse{}, nil
}
