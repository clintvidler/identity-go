package rpc

import (
	"context"
	"log"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
)

func (s IdentityService) RegisterStart(ctx context.Context, req *proto.RegisterStartRequest) (*proto.RegisterStartReponse, error) {
	log.Printf("RegisterStart(%s, %s)", req.Email, req.Password)

	return &proto.RegisterStartReponse{}, nil
}

func (s IdentityService) RegisterPending(ctx context.Context, req *proto.RegisterPendingRequest) (*proto.RegisterPendingReponse, error) {
	log.Println("RegisterPending()", req.Key)

	return &proto.RegisterPendingReponse{}, nil
}

func (s IdentityService) RegisterFinish(ctx context.Context, req *proto.RegisterFinishRequest) (*proto.RegisterFinishReponse, error) {
	log.Printf("RegisterFinish(%s) %s", req.Key, req.DisplayName)

	return &proto.RegisterFinishReponse{}, nil
}
