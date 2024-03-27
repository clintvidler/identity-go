package rpc

import (
	"context"
	"fmt"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
)

func (s *IdentityService) Logout(ctx context.Context, req *proto.LogoutRequest) (*proto.LogoutReponse, error) {
	// Validate token
	claims, err := s.ParseClaims(fmt.Sprint(req.RefreshToken))
	if err != nil {
		return nil, err
	}

	// Ensure the token exists in the current family
	if err := s.data.Token.Exists(req.RefreshToken); err != nil {
		return nil, err
	}

	// Delete all tokens by users id
	if err = s.data.Token.DeleteByUser(claims["sub"].(string)); err != nil {
		return nil, err
	}

	return nil, nil
}
