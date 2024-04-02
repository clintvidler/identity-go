package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"github.com/clintvidler/identity-go/service/data"
)

func (s IdentityService) RegisterStart(ctx context.Context, req *proto.RegisterStartRequest) (*proto.RegisterStartReponse, error) {
	key, err := s.data.User.UpsertPendingRegistration(req.Email)
	if err != nil {
		return nil, err
	}

	log.Println("Confirmation key:", key)

	// TODO: Send email

	return &proto.RegisterStartReponse{}, nil
}

func (s IdentityService) RegisterPending(ctx context.Context, req *proto.RegisterPendingRequest) (*proto.RegisterPendingReponse, error) {
	email, err := s.data.User.ReadPendingRegistration(req.Key)
	if err != nil {
		return nil, err
	}

	return &proto.RegisterPendingReponse{Email: email}, nil
}

func (s IdentityService) RegisterFinish(ctx context.Context, req *proto.RegisterFinishRequest) (*proto.RegisterFinishReponse, error) {
	email, err := s.data.User.ReadPendingRegistration(req.Key)
	if err != nil {
		return nil, err
	}

	var user data.User
	user.Email = email
	user.DisplayName = req.DisplayName
	user.SetPassword(req.Password)

	id, err := s.data.User.Register(user)
	if err != nil {
		return nil, err
	}

	// TODO: Send email

	return &proto.RegisterFinishReponse{Id: fmt.Sprint(id)}, nil
}
