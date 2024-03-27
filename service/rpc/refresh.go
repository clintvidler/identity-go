package rpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *IdentityService) Refresh(ctx context.Context, req *proto.RefreshRequest) (*proto.RefreshReponse, error) {
	// Validate token
	claims, err := s.ParseClaims(fmt.Sprint(req.RefreshToken))
	if err != nil {
		return nil, err
	}
	uid := claims["sub"].(string)
	aid := claims["aud"].(string)

	// Ensure the token exists in the current family
	if err := s.data.Token.Exists(req.RefreshToken); err != nil {
		return nil, err
	}

	// Read the most recently created token for the user/app
	lt, err := s.data.Token.ReadLatestByIds(uid, aid)
	if err != nil {
		return nil, err
	}

	// Invalidate (remove) all of the users tokens if an old token from the current family was provided
	if req.RefreshToken != lt {
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
