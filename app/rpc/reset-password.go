package rpc

import (
	"context"

	proto "github.com/clintvidler/identity-go/gen/proto/server"
)

func (s IdentityService) ResetPasswordStart(ctx context.Context, req *proto.ResetPasswordStartRequest) (*proto.ResetPasswordStartReponse, error) {

	return &proto.ResetPasswordStartReponse{}, nil
}

func (s IdentityService) ResetPasswordPending(ctx context.Context, req *proto.ResetPasswordPendingRequest) (*proto.ResetPasswordPendingReponse, error) {

	return &proto.ResetPasswordPendingReponse{}, nil
}

func (s IdentityService) ResetPasswordFinish(ctx context.Context, req *proto.ResetPasswordFinishRequest) (*proto.ResetPasswordFinishReponse, error) {

	return &proto.ResetPasswordFinishReponse{}, nil
}
