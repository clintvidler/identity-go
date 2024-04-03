package rpc

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/clintvidler/identity-go/app/data"
	proto "github.com/clintvidler/identity-go/gen/proto/server"
)

func (s IdentityService) RegisterStart(ctx context.Context, req *proto.RegisterStartRequest) (*proto.RegisterStartReponse, error) {
	key, err := s.data.User.UpsertPendingRegistration(req.Email)
	if err != nil {
		return nil, err
	}

	emailTo := req.Email
	emailFrom := "no-reply@" + strings.Split(os.Getenv("FRONTEND_URL"), ":")[0]
	emailSubject := "Confirm your account"
	emailBody := "http://" + os.Getenv("FRONTEND_URL") + "/register/" + key + "\n\nExpires in 24 hours."
	s.emailClient.SendEmail(emailTo, emailFrom, emailSubject, emailBody)

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

	emailTo := email
	emailFrom := "no-reply@" + strings.Split(os.Getenv("FRONTEND_URL"), ":")[0]
	emailSubject := "Welcome, account confirmed"
	emailBody := "http://" + os.Getenv("FRONTEND_URL") + "/login"

	s.emailClient.SendEmail(emailTo, emailFrom, emailSubject, emailBody)

	return &proto.RegisterFinishReponse{Id: fmt.Sprint(id)}, nil
}
