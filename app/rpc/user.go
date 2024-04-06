package rpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/clintvidler/identity-go/app/util"
	proto "github.com/clintvidler/identity-go/gen/proto/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type KeyUid struct{}
type KeyAid struct{}

func (s *IdentityService) IsAuth(ctx context.Context, req *proto.IsAuthRequest) (*proto.IsAuthReponse, error) {
	return &proto.IsAuthReponse{Id: fmt.Sprint(ctx.Value(KeyUid{}))}, nil
}

func (s *IdentityService) CurrentUser(ctx context.Context, req *proto.CurrentUserRequest) (*proto.CurrentUserReponse, error) {
	id := fmt.Sprint(ctx.Value(KeyUid{}))

	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Find user by email
	u, err := s.data.User.ReadOne(uint(_id), "")
	if err != nil {
		return nil, err
	}

	return &proto.CurrentUserReponse{
		Id:          fmt.Sprint(u.Id),
		DisplayName: fmt.Sprint(u.DisplayName),
		Email:       fmt.Sprint(u.Email),
		UpdatedAt:   fmt.Sprint(u.UpdatedAt),
		CreatedAt:   fmt.Sprint(u.CreatedAt),
		ExpiredAt:   fmt.Sprint(u.ExpiredAt),
	}, nil
}

func IsAuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// List of methods that require auth
	protected := []string{
		"/proto.IdentityService/IsAuth",
		"/proto.IdentityService/CurrentUser",
	}

	// Read the method
	method, ok := grpc.Method(ctx)
	if !ok {
		return nil, errors.New("an unexpected error has occured")
	}

	// If the method is not protected, skip to next
	if !util.Contains(protected, method) {
		return handler(ctx, req)
	}

	// Read the metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in the context")
	}

	var access string

	access, err = util.GetCookie(md, "access")
	if err != nil {
		if len(md.Get("access")) > 0 {
			access = md.Get("access")[0]
		}
	}

	if access == "" {
		return nil, errors.New("no access token")
	}

	// Access the IdentityService struct
	identityService, ok := info.Server.(*IdentityService)
	if !ok {
		return nil, errors.New("an unexpected error has occured")
	}

	// Parse access token claims
	claims, err := identityService.ParseClaims(access)
	if err != nil {
		// TODO: respond with 401: unauthorized and listen for this on frontend

		return nil, fmt.Errorf("access token: %s", err.Error())
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
