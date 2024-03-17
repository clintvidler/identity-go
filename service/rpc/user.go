package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
)

func (s *IdentityService) CurrentUser(ctx context.Context, req *proto.CurrentUserRequest) (*proto.CurrentUserReponse, error) {
	// TODO: Move some of this into an IsAuth middleware/interceptor

	// Access token must be provided
	if req.Access == "" {
		// Unauthorised
		err := errors.New("no access token provided")
		return nil, err
	}

	claims, err := s.ParseClaims(req.Access)
	if err != nil {
		// Unauthorised
		log.Printf("Error: %s", err)
		return nil, err
	}

	tokenSubject := claims["sub"].(string)
	tokenAudience := claims["aud"].(string)
	log.Printf("Audience: %s", tokenAudience)

	// TODO: Verify Audience

	return &proto.CurrentUserReponse{Id: tokenSubject}, nil
}
