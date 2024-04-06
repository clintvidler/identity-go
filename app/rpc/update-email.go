package rpc

import (
	"context"

	proto "github.com/clintvidler/identity-go/gen/proto/server"
)

func (s IdentityService) UpdateEmailStart(ctx context.Context, req *proto.UpdateEmailStartRequest) (*proto.UpdateEmailStartReponse, error) {

	return &proto.UpdateEmailStartReponse{}, nil
}

func (s IdentityService) UpdateEmailFinish(ctx context.Context, req *proto.UpdateEmailFinishRequest) (*proto.UpdateEmailFinishReponse, error) {

	return &proto.UpdateEmailFinishReponse{}, nil
}
