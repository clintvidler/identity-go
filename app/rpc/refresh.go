package rpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/clintvidler/identity-go/app/util"
	proto "github.com/clintvidler/identity-go/gen/proto/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *IdentityService) Refresh(ctx context.Context, req *proto.RefreshRequest) (*proto.RefreshReponse, error) {
	// Read the metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in the context")
	}

	var refresh string

	refresh, err := util.GetCookie(md, "refresh")
	if err != nil {
		if len(md.Get("refresh")) > 0 {
			refresh = md.Get("refresh")[0]
		}
	}

	if refresh == "" {
		return nil, errors.New("no refresh token")
	}

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
