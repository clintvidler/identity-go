package rpc

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"github.com/clintvidler/identity-go/service/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type KeyUid struct{}
type KeyAid struct{}

func (s *IdentityService) CurrentUser(ctx context.Context, req *proto.CurrentUserRequest) (*proto.CurrentUserReponse, error) {
	return &proto.CurrentUserReponse{Uid: fmt.Sprint(ctx.Value(KeyUid{}))}, nil
}

func IsAuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// List of methods that require auth
	protected := []string{
		"/proto.IdentityService/CurrentUser",
	}

	// Read the method
	method, ok := grpc.Method(ctx)
	if !ok {
		return nil, errors.New("an unexpected error has occured")
	}

	// If rhe method is not protected, skip to next
	if !util.Contains(protected, method) {
		return handler(ctx, req)
	}

	// Read the metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in the context")
	}

	// Read the access token from the metadata
	if len(md.Get("access")) < 1 {
		return nil, errors.New("no access token")
	}
	access := md.Get("access")[0]

	// Access the IdentityService struct
	identityService, ok := info.Server.(*IdentityService)
	if !ok {
		return nil, errors.New("an unexpected error has occured")
	}

	// Parse access token claims
	claims, err := identityService.ParseClaims(access)
	if err != nil {
		return nil, err
	}

	tokenSubject := claims["sub"].(string)
	tokenAudience := claims["aud"].(string)
	log.Printf("Subject: %s", tokenSubject)
	log.Printf("Audience: %s", tokenAudience)

	// TODO: Verify audience is as expected

	// Save uid and aid to context
	_ctx := ctx
	_ctx = context.WithValue(_ctx, KeyUid{}, tokenSubject)
	_ctx = context.WithValue(_ctx, KeyAid{}, tokenAudience)

	return handler(_ctx, req)
}
