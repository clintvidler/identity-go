package rpc

import (
	"context"
	"strconv"

	proto "github.com/clintvidler/identity-go/gen/proto/server"
)

func (s IdentityService) UpdateDisplayName(ctx context.Context, req *proto.UpdateDisplayNameRequest) (*proto.UpdateDisplayNameReponse, error) {
	uid, err := strconv.Atoi(ctx.Value(KeyUid{}).(string))
	if err != nil {
		return nil, err
	}

	u, err := s.data.User.ReadOne(uint(uid), "")
	if err != nil {
		return nil, err
	}

	u.DisplayName = req.New

	if err := s.data.User.UpdateOne(u); err != nil {
		return nil, err
	}

	return &proto.UpdateDisplayNameReponse{}, nil
}
