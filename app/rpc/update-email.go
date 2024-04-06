package rpc

import (
	"context"
	"os"
	"strings"

	proto "github.com/clintvidler/identity-go/gen/proto/server"
)

func (s IdentityService) UpdateEmailStart(ctx context.Context, req *proto.UpdateEmailStartRequest) (*proto.UpdateEmailStartReponse, error) {

	uid := ctx.Value(KeyUid{}).(string)

	emailNew := req.Email

	key, err := s.data.User.UpsertPendingUpdateEmail(uid, emailNew)
	if err != nil {
		return nil, err
	}

	emailTo := emailNew

	emailFrom := "no-reply@" + strings.Split(os.Getenv("FRONTEND_URL"), ":")[0]
	emailSubject := "Confirm change email"
	emailBody := "http://" + os.Getenv("FRONTEND_URL") + "/user/edit/email/" + key + "\n\nExpires in 24 hours."
	s.emailClient.Send(emailTo, emailFrom, emailSubject, emailBody)

	return &proto.UpdateEmailStartReponse{}, nil
}

func (s IdentityService) UpdateEmailFinish(ctx context.Context, req *proto.UpdateEmailFinishRequest) (*proto.UpdateEmailFinishReponse, error) {
	uid, newEmail, err := s.data.User.ReadPendingUpdateEmail(req.Key)
	if err != nil {
		return nil, err
	}

	if err := s.data.User.UpdateEmail(uid, newEmail); err != nil {
		return nil, err
	}

	emailTo := newEmail
	emailFrom := "no-reply@" + strings.Split(os.Getenv("FRONTEND_URL"), ":")[0]
	emailSubject := "Email has been changed"
	emailBody := "Email changed, contact admin if this wan't you."
	s.emailClient.Send(emailTo, emailFrom, emailSubject, emailBody)

	return &proto.UpdateEmailFinishReponse{}, nil
}
