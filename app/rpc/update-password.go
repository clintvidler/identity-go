package rpc

import (
	"context"
	"os"
	"strconv"
	"strings"

	proto "github.com/clintvidler/identity-go/gen/proto/server"
)

func (s IdentityService) UpdatePassword(ctx context.Context, req *proto.UpdatePasswordRequest) (*proto.UpdatePasswordReponse, error) {
	uid, err := strconv.Atoi(ctx.Value(KeyUid{}).(string))
	if err != nil {
		return nil, err
	}

	u, err := s.data.User.ReadOne(uint(uid), "")
	if err != nil {
		return nil, err
	}

	if err := u.ComparePassword(req.Old); err != nil {
		return nil, err
	}

	if err := u.SetPassword(req.New); err != nil {
		return nil, err
	}

	if err := s.data.User.UpdateOne(u); err != nil {
		return nil, err
	}

	// Notify the user

	emailTo := u.Email
	emailFrom := "no-reply@" + strings.Split(os.Getenv("FRONTEND_URL"), ":")[0]
	emailSubject := "Account password changed"
	emailBody := "If this wasn't you, contact admin."
	s.emailClient.Send(emailTo, emailFrom, emailSubject, emailBody)

	return &proto.UpdatePasswordReponse{}, nil
}
