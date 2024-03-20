package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"google.golang.org/grpc/metadata"
)

func (s *IdentityService) CurrentUser(ctx context.Context, req *proto.CurrentUserRequest) (*proto.CurrentUserReponse, error) {
	// TODO: Move some of this into an IsAuth middleware/interceptor

	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in the context")
	}
	access := headers.Get("access")[0]
	if len(access) < 1 {
		return nil, errors.New("no token found in the metadata")
	}

	claims, err := s.ParseClaims(access)
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
